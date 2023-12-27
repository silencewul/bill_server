package categoryLogic

import "bill/models"

func GetCategory(k int)  (models.CateList,error) {
	bills,err := models.GetCategory(k)
	return bills,err
}
