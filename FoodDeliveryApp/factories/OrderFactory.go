package factories

import (
	"fooddelivery/models"
	"fooddelivery/strategies"
)

type OrderFactory interface {
	CreateOrder(orderType string, id int, user *models.User, ps strategies.PaymentStrategy) models.Order
}
