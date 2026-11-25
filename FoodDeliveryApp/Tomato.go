package main

import (
	"fmt"

	"fooddelivery/factories"
	"fooddelivery/managers"
	"fooddelivery/models"
	"fooddelivery/strategies"
	"notification-system/notification"
	"notification-system/observer"
	notificationservice "notification-system/service"
	notificationstrategy "notification-system/strategy"
)

type TomatoApp struct {
	restaurantManager   *managers.RestaurantManager
	userManager         *managers.UserManager
	orderManager        *managers.OrderManager
	nextOrderID         int
	notificationService *notificationservice.NotificationService
	notificationEngine  *observer.NotificationEngine
}

func NewTomatoApp() *TomatoApp {
	orderManager := managers.GetOrderManager()
	return &TomatoApp{
		restaurantManager:   managers.GetRestaurantManager(),
		userManager:         managers.GetUserManager(),
		orderManager:        orderManager,
		nextOrderID:         len(orderManager.GetAllOrders()) + 1,
		notificationService: notificationservice.GetInstance(),
	}
}

// InitializeRestaurants loads the sample restaurants and menus for this demo.
func (app *TomatoApp) InitializeRestaurants() {
	if len(app.restaurantManager.RestaurantList) > 0 {
		return
	}

	app.addRestaurant(
		"Spice Garden",
		"Chandigarh",
		models.MenuItem{Code: 101, Name: "Paneer Tikka", Price: 250},
		models.MenuItem{Code: 102, Name: "Butter Naan", Price: 50},
	)
	app.addRestaurant(
		"Bikaner",
		"Delhi",
		models.MenuItem{Code: 201, Name: "Chole Bhature", Price: 120},
	)
}

func (app *TomatoApp) addRestaurant(name, address string, menuItems ...models.MenuItem) *models.Restaurant {
	restaurant := models.NewRestaurant(name, address)
	for _, item := range menuItems {
		restaurant.AddMenuItem(item)
	}
	app.restaurantManager.Add(restaurant)
	return restaurant
}

func (app *TomatoApp) CreateUser(name, address, mobileNo, email string) *models.User {
	user := models.NewUser(name, address, mobileNo, email)
	app.userManager.Add(user)
	return user
}

func (app *TomatoApp) SearchRestaurants(location string) []*models.Restaurant {
	return app.restaurantManager.SearchByLocation(location)
}

func (app *TomatoApp) SelectRestaurant(user *models.User, restaurant *models.Restaurant) error {
	if user == nil {
		return fmt.Errorf("cannot select a restaurant for a nil user")
	}
	if restaurant == nil {
		return fmt.Errorf("cannot select a nil restaurant")
	}
	if user.Cart == nil {
		user.Cart = &models.Cart{}
	}
	if user.Cart.Res != nil && user.Cart.Res.Id != restaurant.Id {
		user.Cart.Clear()
	}
	user.Cart.Res = restaurant
	return nil
}

func (app *TomatoApp) AddToCart(user *models.User, itemCode int) error {
	if user == nil || user.Cart == nil {
		return fmt.Errorf("user or user cart is not initialized")
	}
	if user.Cart.Res == nil {
		return fmt.Errorf("select a restaurant before adding menu items")
	}
	item := user.Cart.Res.GetMenuItem(itemCode)
	if item == nil {
		return fmt.Errorf("menu item %d was not found", itemCode)
	}
	user.Cart.AddToCart(user.Cart.Res, *item)
	return nil
}

func (app *TomatoApp) PrintUserCart(user *models.User) {
	fmt.Println("Items in cart:")
	for _, item := range user.Cart.MenuList {
		fmt.Printf(" - %s: ₹%.2f\n", item.Name, item.Price)
	}
	fmt.Printf("Total: ₹%.2f\n", user.Cart.Total)
}

func (app *TomatoApp) CheckoutNow(
	user *models.User,
	orderType string,
	payment strategies.PaymentStrategy,
) (models.Order, error) {
	factory := &factories.NowOrderFactory{}
	return app.checkout(user, orderType, payment, factory)
}

func (app *TomatoApp) CheckoutScheduled(
	user *models.User,
	orderType string,
	payment strategies.PaymentStrategy,
	scheduleTime string,
) (models.Order, error) {
	factory := &factories.ScheduleOrderFactory{ScheduleTime: scheduleTime}
	return app.checkout(user, orderType, payment, factory)
}

func (app *TomatoApp) checkout(
	user *models.User,
	orderType string,
	payment strategies.PaymentStrategy,
	factory factories.OrderFactory,
) (models.Order, error) {
	if user == nil || user.Cart == nil {
		return nil, fmt.Errorf("user or user cart is not initialized")
	}
	if user.Cart.IsEmpty() {
		return nil, fmt.Errorf("cannot checkout an empty cart")
	}
	if user.Cart.Res == nil {
		return nil, fmt.Errorf("select a restaurant before checkout")
	}
	if payment == nil {
		return nil, fmt.Errorf("select a payment method before checkout")
	}

	order := factory.CreateOrder(orderType, app.nextOrderID, user, payment)
	app.nextOrderID++
	return order, nil
}

func (app *TomatoApp) PayForOrder(user *models.User, order models.Order) error {
	if user == nil || order == nil {
		return fmt.Errorf("user and order are required")
	}
	if err := order.ProcessPayment(); err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	app.orderManager.AddOrder(order)
	app.sendOrderNotification(user, order)
	user.Cart.Clear()
	return nil
}

func (app *TomatoApp) sendOrderNotification(user *models.User, order models.Order) {
	if app.notificationEngine != nil {
		app.notificationService.GetObservable().RemoveObserver(app.notificationEngine)
	}
	app.notificationEngine = notificationservice.NewNotificationEngine()
	app.notificationEngine.AddNotificationStrategy(
		notificationstrategy.NewEmailStrategy(user.Email),
	)

	var message notification.INotification = notification.NewSimpleNotification(
		"Your order has been confirmed!\n" + order.GetDetails(),
	)
	message = notification.NewTimestampDecorator(message)
	app.notificationService.SendNotification(message)
}
