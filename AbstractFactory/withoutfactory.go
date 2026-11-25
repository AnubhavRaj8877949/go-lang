package main

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
// 4. BUSINESS LOGIC
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
// 5. MAIN
// ==========================================

func main() {

	// AWS objects directly create kiye

	awsSMS := &AWSSMS{}
	awsEmail := &AWSEmail{}
	awsPush := &AWSPush{}

	SendOrderConfirmation(awsSMS, awsEmail, awsPush)

	fmt.Println("----------------")

	// Firebase objects directly create kiye

	firebaseSMS := &FirebaseSMS{}
	firebaseEmail := &FirebaseEmail{}
	firebasePush := &FirebasePush{}

	SendOrderConfirmation(firebaseSMS, firebaseEmail, firebasePush)
}
