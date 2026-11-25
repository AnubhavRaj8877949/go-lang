package abstractfactory

import "fmt"

// ==========================================
// 1. PRODUCT INTERFACES
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
// 2. AWS IMPLEMENTATION
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

// ==========================================
// 3. FIREBASE IMPLEMENTATION
// ==========================================

type FirebaseSMS struct{}

func (f *FirebaseSMS) SendSMS(msg string) {
	fmt.Println("[Firebase] SMS:", msg)
}

type FirebaseEmail struct{}

func (f *FirebaseEmail) SendEmail(msg string) {
	fmt.Println("[Firebase] Email:", msg)
}

type FirebasePush struct{}

func (f *FirebasePush) SendPush(msg string) {
	fmt.Println("[Firebase FCM] Push:", msg)
}

// ==========================================
// 4. NORMAL FACTORY FUNCTIONS
// ==========================================

// SMS Factory
func GetSMSSender(provider string) SMSSender {
	if provider == "AWS" {
		return &AWSSMS{}
	}

	return &FirebaseSMS{}
}

// Email Factory
func GetEmailSender(provider string) EmailSender {
	if provider == "AWS" {
		return &AWSEmail{}
	}

	return &FirebaseEmail{}
}

// Push Factory
func GetPushSender(provider string) PushSender {
	if provider == "AWS" {
		return &AWSPush{}
	}

	return &FirebasePush{}
}

// ==========================================
// 5. BUSINESS LOGIC
// ==========================================

func SendOrderConfirmation(
	sms SMSSender,
	email EmailSender,
	push PushSender,
) {

	sms.SendSMS("Aapka order #101 confirm ho gaya hai.")

	email.SendEmail("Order #101 ki receipt.")

	push.SendPush("Order #101 pack ho raha hai!")
}

// ==========================================
// 6. MAIN
// ==========================================

func main() {

	// AWS objects individually create honge

	awsSMS := GetSMSSender("AWS")
	awsEmail := GetEmailSender("AWS")
	awsPush := GetPushSender("AWS")

	SendOrderConfirmation(awsSMS, awsEmail, awsPush)

	fmt.Println("----------------")

	// Firebase objects individually create honge

	firebaseSMS := GetSMSSender("Firebase")
	firebaseEmail := GetEmailSender("Firebase")
	firebasePush := GetPushSender("Firebase")

	SendOrderConfirmation(firebaseSMS, firebaseEmail, firebasePush)
}
