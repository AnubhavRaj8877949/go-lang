package observer

import (
	"fmt"
)

type Logger struct {
	notificationObservable *NotificationObservable
}

func NewLoggerWithObservable(observable *NotificationObservable) *Logger {
	l := &Logger{
		notificationObservable: observable,
	}
	observable.AddObserver(l)
	return l
}

func (l *Logger) Update() {
	fmt.Println("Logging New Notification : \n" + l.notificationObservable.GetNotificationContent())
}
