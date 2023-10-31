package user

import (
	"bill/logic/userLogic"
	"bill/models"
	"bill/modules/constant"
	"bill/modules/jwt_auth"
	"bill/modules/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(g *gin.RouterGroup) {
	g.POST("/login", login)
	g.POST("/regist", regist)
}

func login(c *gin.Context) {
	loginReq := &models.UserLoginPayload{}
	if err := c.ShouldBindJSON(loginReq); err != nil {
		utils.SendErr(c, constant.ErrInvalidParams)
		return
	}

	user, err := userLogic.Login(loginReq)

	if err != nil {
		utils.SendErr(c, err)
		return
	}

	// 发放jwttoken
	token, err := jwt_auth.NewToken(jwt.MapClaims{"id": user.Id})

	if err != nil {
		utils.SendErr(c, constant.ErrServerInternalError)
		return
	}

	//缓存用户到sessions
	if err := user.StoreToRedis(); err != nil {
		utils.SendErr(c, constant.ErrServerInternalError)
		return
	}

	// 种cookie  httpOnly表示客户端js无法操作,secure表示无法通过浏览器查看（如果用postman,是否需要关闭secure?）
	c.SetCookie("token", token, 0,
		"/", "", true, true)

	utils.SendSucc(c, gin.H{"token": token, "user": user})
}

// regist 用户注册
func regist(c *gin.Context) {
	payload := new(models.UserRegistPayload)
	if err := c.ShouldBind(payload); err != nil {
		utils.SendErr(c, constant.ErrInvalidParams)
		return
	}

}
