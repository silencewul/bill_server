package billLogic

import (
	"bill/models"
)

func InsertBill(b *models.BillCreatePayload, u *models.User) (*models.Bill, error) {
	bill := b.ToBill()
	bill.UId = u.Id
	if err := bill.Insert(); err != nil {
		return nil, err
	}
	return bill, nil
}

func GetUserBills(u *models.User) ([]*models.Bill,error) {
	bills,err := models.GetBills(u)
	return bills,err
}
