package middleware

import (
	"bill/modules/constant"
	"bill/modules/utils"
	"github.com/gin-gonic/gin"
)

// SigninRequired 必须是登录用户
func SigninRequired(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		utils.SendErr(c, constant.ErrUnauthorized)
		return
	}

	c.Next()
}
