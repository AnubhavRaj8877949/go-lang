package factories

import (
	"fooddelivery/models"
	"fooddelivery/strategies"
	"time"
)

type NowOrderFactory struct{}

func (nf *NowOrderFactory) CreateOrder(orderType string, id int, user *models.User, ps strategies.PaymentStrategy) models.Order {
	base := models.BaseOrder{
		OrderId:         id,
		Restaurant:      user.Cart.Res,
		MenuList:        user.Cart.MenuList,
		User:            user,
		PaymentStrategy: ps,
		Total:           user.Cart.Total,
		CreatedAt:       time.Now(),
		PaymentStatus:   "Pending",
	}
	if orderType == "Delivery" {
		return &models.DeliveryOrder{BaseOrder: base, DeliveryAddress: user.Address}
	}
	return &models.PickupOrder{BaseOrder: base}
}
