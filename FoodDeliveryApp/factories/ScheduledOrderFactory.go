package factories

import (
	"fooddelivery/models"
	"fooddelivery/strategies"
	"time"
)

type ScheduleOrderFactory struct {
	ScheduleTime string
}

func (sf *ScheduleOrderFactory) CreateOrder(orderType string, id int, user *models.User, ps strategies.PaymentStrategy) models.Order {
	base := models.BaseOrder{
		OrderId:         id,
		Restaurant:      user.Cart.Res,
		MenuList:        user.Cart.MenuList,
		User:            user,
		PaymentStrategy: ps,
		Total:           user.Cart.Total,
		ScheduleTime:    sf.ScheduleTime,
		CreatedAt:       time.Now(),
		PaymentStatus:   "Pending",
	}
	if orderType == "Delivery" {
		return &models.DeliveryOrder{BaseOrder: base, DeliveryAddress: user.Address}
	}
	return &models.PickupOrder{BaseOrder: base}
}
