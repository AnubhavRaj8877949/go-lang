package main

import "fmt"

// ==========================================
// 1. PRODUCT INTERFACES (Teeno cheezon ke rules)
// ==========================================
type SMSSender interface {
	SendSMS(msg string)
}
type EmailSender interface {
	SendEmail(msg string)
}
type PushSender interface {
	SendPush(msg string)
}

// ==========================================
// 2. ABSTRACT FACTORY INTERFACE (Ek Family ka Combo)
// ==========================================
type NotificationsFactory interface {
	CreateSMS() SMSSender
	CreateEmail() EmailSender
	CreatePush() PushSender
}

// ==========================================
// 3. FAMILY 1: AWS IMPLEMENTATION
// ==========================================
type AWSSMS struct{}

func (a *AWSSMS) SendSMS(msg string) {
	fmt.Println("[AWS SNS] SMS:", msg)
}

type AWSEmail struct{}

func (a *AWSEmail) SendEmail(msg string) {
	fmt.Println("[AWS SES] Email:", msg)
}

type AWSPush struct{}

func (a *AWSPush) SendPush(msg string) {
	fmt.Println("[AWS Pinpoint] Push:", msg)
}

// AWS Factory (Yeh teeno AWS ke hi degi)

type AWSFactory struct{}

func (f *AWSFactory) CreateSMS() SMSSender {
	return &AWSSMS{}
}

func (f *AWSFactory) CreateEmail() EmailSender {
	return &AWSEmail{}
}
func (f *AWSFactory) CreatePush() PushSender {
	return &AWSPush{}
}

// ==========================================
// 4. FAMILY 2: FIREBASE IMPLEMENTATION
// ==========================================
type FirebaseSMS struct{}

func (fb *FirebaseSMS) SendSMS(msg string) { fmt.Println("[Firebase] SMS:", msg) }

type FirebaseEmail struct{}

func (fb *FirebaseEmail) SendEmail(msg string) { fmt.Println("[Firebase] Email:", msg) }

type FirebasePush struct{}

func (fb *FirebasePush) SendPush(msg string) { fmt.Println("[Firebase FCM] Push:", msg) }

// Firebase Factory (Yeh teeno Firebase ke hi degi)
type FirebaseFactory struct{}

func (f *FirebaseFactory) CreateSMS() SMSSender     { return &FirebaseSMS{} }
func (f *FirebaseFactory) CreateEmail() EmailSender { return &FirebaseEmail{} }
func (f *FirebaseFactory) CreatePush() PushSender   { return &FirebasePush{} }

// Add hear  if u want to add   some   new  Provider

func GetNotificationFactory(provider string) NotificationsFactory {
	if provider == "AWS" {

		return &AWSFactory{}
	}
	return &FirebaseFactory{}
}

func SendOrderConfirmation(factory NotificationsFactory) {

	sms := factory.CreateSMS()
	email := factory.CreateEmail()
	push := factory.CreatePush()

	sms.SendSMS("Aapka order #101 confirm ho gaya hai.")
	email.SendEmail("Order #101 ki receipt.")
	push.SendPush("Order #101 pack ho raha hai!")
}

func main() {
	// App start hote waqt config file se bas EK baar factory banayi
	activeFactory := GetNotificationFactory("AWS")
	activeFactory2 := GetNotificationFactory("Firebase")

	// Ab poore project mein sirf 'activeFactory' pass karni hai
	SendOrderConfirmation(activeFactory)

	SendOrderConfirmation(activeFactory2)

}
