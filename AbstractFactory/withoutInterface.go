package main

import "fmt"

// ==========================================
// 1. AWS IMPLEMENTATION
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
// 2. FIREBASE IMPLEMENTATION
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
// 3. BUSINESS LOGIC
// ==========================================

// Yahan directly concrete struct pointers accept kar rahe hain.

func SendAWSOrderConfirmation(
	sms *AWSSMS,
	email *AWSEmail,
	push *AWSPush,
) {
	sms.SendSMS("Order #101 confirmed")
	email.SendEmail("Order #101 receipt")
	push.SendPush("Order #101 processing")
}

func SendFirebaseOrderConfirmation(
	sms *FirebaseSMS,
	email *FirebaseEmail,
	push *FirebasePush,
) {
	sms.SendSMS("Order #101 confirmed")
	email.SendEmail("Order #101 receipt")
	push.SendPush("Order #101 processing")
}

// ==========================================
// 4. MAIN
// ==========================================

func main() {

	// AWS objects directly create kiye

	awsSMS := &AWSSMS{}
	awsEmail := &AWSEmail{}
	awsPush := &AWSPush{}

	SendAWSOrderConfirmation(awsSMS, awsEmail, awsPush)

	fmt.Println("----------------")

	// Firebase objects directly create kiye

	firebaseSMS := &FirebaseSMS{}
	firebaseEmail := &FirebaseEmail{}
	firebasePush := &FirebasePush{}

	SendFirebaseOrderConfirmation(
		firebaseSMS,
		firebaseEmail,
		firebasePush,
	)
}
