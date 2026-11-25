package strategy

import "fmt"

type INotificationStrategy interface {
	SendNotification(content string)
}

type EmailStrategy struct {
	emailId string
}

func NewEmailStrategy(emailId string) *EmailStrategy {
	return &EmailStrategy{emailId: emailId}
}

func (e *EmailStrategy) SendNotification(content string) {
	fmt.Printf("Sending email Notification to: %s\n%s\n", e.emailId, content)
}

type SMSStrategy struct {
	mobileNumber string
}

func NewSMSStrategy(mobileNumber string) *SMSStrategy {
	return &SMSStrategy{mobileNumber: mobileNumber}
}

func (s *SMSStrategy) SendNotification(content string) {
	fmt.Printf("Sending SMS Notification to: %s\n%s\n", s.mobileNumber, content)
}

type PopUpStrategy struct{}

func NewPopUpStrategy() *PopUpStrategy {
	return &PopUpStrategy{}
}

func (p *PopUpStrategy) SendNotification(content string) {
	fmt.Printf("Sending Popup Notification: \n%s\n", content)
}
