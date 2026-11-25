package main

import (
	"notification-system/notification"
	"notification-system/service"
	"notification-system/strategy"
)

func main() {
	notificationService := service.GetInstance()

	// Default Constructors (auto-attach to Singleton Observable)
	_ = service.NewLogger()
	notificationEngine := service.NewNotificationEngine()

	notificationEngine.AddNotificationStrategy(strategy.NewEmailStrategy("random.person@gmail.com"))
	notificationEngine.AddNotificationStrategy(strategy.NewSMSStrategy("+91 9876543210"))
	notificationEngine.AddNotificationStrategy(strategy.NewPopUpStrategy())

	var notif notification.INotification = notification.NewSimpleNotification("Your order has been shipped!")
	notif = notification.NewTimestampDecorator(notif)
	notif = notification.NewSignatureDecorator(notif, "Customer Care")
	notif = notification.NewSignatureDecorator(notif, "High")

	notificationService.SendNotification(notif)
}
