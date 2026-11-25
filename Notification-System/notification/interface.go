package notification

type INotification interface {
	GetContent() string
}

type INotificationDecorator struct {
	notification INotification
}

func (d *INotificationDecorator) GetContent() string {
	return d.notification.GetContent()
}
