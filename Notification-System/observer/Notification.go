package observer

import "notification-system/strategy"

type NotificationEngine struct {
	notificationObservable *NotificationObservable
	notificationStrategies []strategy.INotificationStrategy
}

func NewNotificationEngineWithObservable(observable *NotificationObservable) *NotificationEngine {
	return &NotificationEngine{
		notificationObservable: observable,
		notificationStrategies: make([]strategy.INotificationStrategy, 0),
	}
}

func (ne *NotificationEngine) AddNotificationStrategy(ns strategy.INotificationStrategy) {
	ne.notificationStrategies = append(ne.notificationStrategies, ns)
}

func (ne *NotificationEngine) Update() {
	notificationContent := ne.notificationObservable.GetNotificationContent()
	for _, s := range ne.notificationStrategies {
		s.SendNotification(notificationContent)
	}
}
