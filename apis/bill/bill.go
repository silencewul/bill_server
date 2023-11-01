package bill

import (
	"bill/logic/billLogic"
	"bill/logic/userLogic"
	"bill/models"
	"bill/modules/constant"
	"bill/modules/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(g *gin.RouterGroup) {
	g.POST("/bill/add", add)
	g.GET("/bills", getBill)
}

func add(c *gin.Context) {
	payload := new(models.BillCreatePayload)
	if err := c.ShouldBindJSON(payload); err != nil {
		utils.SendErr(c, constant.ErrInvalidParams)
		return
	}

	userInfo := userLogic.GetMe(c)
	bill, err := billLogic.InsertBill(payload, userInfo)
	if err != nil {
		utils.SendErr(c, err)
		return
	}
	utils.SendSucc(c, bill)
}

func getBill(c *gin.Context) {
	userInfo := userLogic.GetMe(c)
	bills,err := billLogic.GetUserBills(userInfo)
	if err != nil {
		utils.SendErr(c, err)
		return
	}
	utils.SendSucc(c, bills)
}
