package notification

import (
	"fmt"
)

type SignatureDecorator struct {
	INotificationDecorator
	signature string
}

func NewSignatureDecorator(n INotification, signature string) *SignatureDecorator {
	return &SignatureDecorator{
		INotificationDecorator: INotificationDecorator{notification: n},
		signature:              signature,
	}
}

func (s *SignatureDecorator) GetContent() string {
	return fmt.Sprintf("%s\n-- %s", s.notification.GetContent(), s.signature)
}
