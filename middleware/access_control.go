package middleware

import (
	"bill/models"
	"bill/modules/constant"
	"bill/modules/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionRequired(resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户ID
		u, isExit := c.Get("user")
		if !isExit {
			utils.SendHttpErr(c, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}
		_, ok := u.(*models.User)

		if !ok {
			utils.SendHttpErr(c, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		//if !user.HasPermission(resource, action) {
		//	utils.SendHttpErr(c, http.StatusForbidden, constant.ErrPermissionDenied)
		//	return
		//}
		c.Next()
	}
}
