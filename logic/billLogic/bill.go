package billLogic

import (
	"bill/models"
)

func InsertBill(b *models.BillCreatePayload, u *models.User) (*models.Bill, error) {
	bill := b.ToBill()
	bill.UId = 1
	if err := bill.Insert(); err != nil {
		return nil, err
	}
	return bill, nil
}
