package middleware

import (
	"bill/models"
	"bill/modules/constant"
	"bill/modules/jwt_auth"
	"bill/modules/log"
	"bill/modules/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func CtxData(c *gin.Context) {
	tokenString, cookieErr := c.Cookie("token")

	if cookieErr != nil {
		c.Next()
		return
	}

	claims, ok := jwt_auth.VerifyTokenAsClaims(tokenString)

	if !ok {
		c.Next()
		return
	}

	userID := int(claims["id"].(float64))
	user, err := models.UserFromRedis(userID)

	if err != nil {
		c.Next()
		return
	}

	//TODO 自动延长Token过期时间,重新设置cookie
	if val, ok := claims["exp"]; ok {
		expf := val.(float64)
		exp := int64(expf)
		if time.Now().Unix()-exp < 600 {
			// 发放jwttoken
			token, err := jwt_auth.NewToken(jwt.MapClaims{"id": user.Id})

			if err != nil {
				utils.SendErr(c, constant.ErrServerInternalError)
				return
			}
			//缓存用户到sessions
			if err := models.UserToRedis(user); err != nil {
				log.GetSugar().Errorf("刷新用户token失败!")
				utils.SendErr(c, constant.ErrServerInternalError)
				return
			}

			// 种cookie
			c.SetCookie("token", token, 0,
				"/", "", true, true)
		}
	}

	c.Set("user", user)
	c.Next()
}
