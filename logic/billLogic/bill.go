package billLogic

import (
	"bill/models"
)

func InsertBill(b *models.Bill,u *models.User)(*models.Bill,error) {
	b.UId = u.Id
	if err := b.Insert(); err != nil {
		return nil, err
	}
	return b, nil
}
