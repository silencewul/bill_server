package category

import (
	"bill/logic/categoryLogic"
	"bill/modules/constant"
	"bill/modules/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RegisterUserRoutes(g *gin.RouterGroup) {
	g.GET("/category",getCategory)
}

func getCategory(c *gin.Context) {
	str := c.Param("k")
	v, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		utils.SendErr(c, constant.ErrInvalidParams)
	}
	k := int(v)
	if k==0 {
		k=1
	}
	cate,err := categoryLogic.GetCategory(k)
	if err != nil {
		utils.SendErr(c, err)
		return
	}
	utils.SendSucc(c, cate)
}
