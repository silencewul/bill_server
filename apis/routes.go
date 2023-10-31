package apis

import (
	"bill/apis/bill"
	"bill/apis/category"
	"bill/apis/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	user.RegisterUserRoutes(g)
	bill.RegisterUserRoutes(g)
	category.RegisterUserRoutes(g)
}
