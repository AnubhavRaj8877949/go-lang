package managers

import (
	"fooddelivery/models"
	"sync"
)

type OrderManager struct {
	Orders []models.Order
}

var (
	orderManagerInstance *OrderManager
	orderOnes            sync.Once
)

func GetOrderManager() *OrderManager {

	orderOnes.Do(func() {
		orderManagerInstance = &OrderManager{
			Orders: []models.Order{},
		}
	})
	return orderManagerInstance

}

func (om *OrderManager) AddOrder(o models.Order) {
	om.Orders = append(om.Orders, o)
}

func (om *OrderManager) GetAllOrders() []models.Order {
	return om.Orders
}
