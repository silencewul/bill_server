package apis

import (
	"bill/apis/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	user.RegisterUserRoutes(g)
}
