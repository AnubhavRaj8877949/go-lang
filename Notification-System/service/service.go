package service

import (
	"sync"

	"notification-system/notification"
	"notification-system/observer"
)

type NotificationService struct {
	observable    *observer.NotificationObservable
	notifications []notification.INotification
}

var (
	instance *NotificationService
	once     sync.Once
)

func GetInstance() *NotificationService {
	once.Do(func() {
		instance = &NotificationService{
			observable:    observer.NewNotificationObservable(), // creating observable
			notifications: make([]notification.INotification, 0),
		}
	})
	return instance
}

func (ns *NotificationService) GetObservable() *observer.NotificationObservable {
	return ns.observable
}

func (ns *NotificationService) SendNotification(notif notification.INotification) {
	ns.notifications = append(ns.notifications, notif)
	ns.observable.SetNotification(notif)
}

func NewLogger() *observer.Logger {
	obs := GetInstance().GetObservable()
	return observer.NewLoggerWithObservable(obs)
}

func NewNotificationEngine() *observer.NotificationEngine {
	obs := GetInstance().GetObservable()
	engine := observer.NewNotificationEngineWithObservable(obs)
	obs.AddObserver(engine)
	return engine
}
