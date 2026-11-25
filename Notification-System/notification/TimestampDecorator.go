package notification

import (
	"fmt"
	"time"
)

type TimestampDecorator struct {
	INotificationDecorator
}

func NewTimestampDecorator(n INotification) *TimestampDecorator {
	return &TimestampDecorator{
		INotificationDecorator: INotificationDecorator{notification: n},
	}
}

func (t *TimestampDecorator) GetContent() string {
	ts := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s]%s", ts, t.notification.GetContent())
}
