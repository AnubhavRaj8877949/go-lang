package observer

import "notification-system/notification"

type IObserver interface {
	Update()
}

type IObservable interface {
	AddObserver(observer IObserver)
	RemoveObserver(observer IObserver)
	NotifyObservers()
}

type NotificationObservable struct {
	observers           []IObserver
	currentNotification notification.INotification
}

func NewNotificationObservable() *NotificationObservable {
	return &NotificationObservable{
		observers: make([]IObserver, 0),
	}
}

func (no *NotificationObservable) AddObserver(obs IObserver) {
	no.observers = append(no.observers, obs)
}

func (no *NotificationObservable) RemoveObserver(obs IObserver) {
	for i, o := range no.observers {
		if o == obs {
			no.observers = append(no.observers[:i], no.observers[i+1:]...)
			break
		}
	}
}

func (no *NotificationObservable) NotifyObservers() {
	for _, obs := range no.observers {
		obs.Update()
	}
}

func (no *NotificationObservable) SetNotification(notif notification.INotification) {
	no.currentNotification = notif
	no.NotifyObservers()
}

func (no *NotificationObservable) GetNotification() notification.INotification {
	return no.currentNotification
}

func (no *NotificationObservable) GetNotificationContent() string {
	return no.currentNotification.GetContent()
}
