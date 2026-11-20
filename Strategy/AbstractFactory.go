package main

import (
	"fmt"
	"time"
)

func validateAmount1(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %d", amount)
	}

	return nil
}

type Payment1 interface {
	Pay(amount int) error
}

type Refund interface {
	Refund(amount int) error
}
type Webhook interface {
	Handle() error
}

type RazorpayPayment struct{}

func (r *RazorpayPayment) Pay(amount int) error {
	if err := validateAmount1(amount); err != nil {
		return err
	}
	// Implement Razorpay payment logic here
	println("Processing Razorpay payment...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

type RazorpayRefund struct{}

func (r *RazorpayRefund) Refund(amount int) error {
	if err := validateAmount1(amount); err != nil {
		return err
	}
	// Implement Razorpay refund logic here
	println("Processing Razorpay refund...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

type RazorpayWebhook struct{}

func (r *RazorpayWebhook) Handle() error {

	// Implement Razorpay webhook logic here
	println("Processing Razorpay webhook...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

type PayUPayment struct{}

func (p *PayUPayment) Pay(amount int) error {
	if err := validateAmount1(amount); err != nil {
		return err
	}
	// Implement PayU payment logic here
	println("Processing PayU payment...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

type PayURefund struct{}

func (p *PayURefund) Refund(amount int) error {
	if err := validateAmount1(amount); err != nil {
		return err
	}
	// Implement PayU refund logic here
	println("Processing PayU refund...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

type PayUWebhook struct{}

func (p *PayUWebhook) Handle() error {

	// Implement PayU webhook logic here
	println("Processing PayU webhook...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

// All provider components grouped together.
type PaymentComponents struct {
	Payment Payment1
	Refund  Refund
	Webhook Webhook
}

// this is  messey  code  but  i  will  refactor  it  later  to  make  it  clean
// why this is  messey

func ProcessPayment(provider string) (PaymentComponents, error) {

	var payment Payment1
	var refund Refund
	var webhook Webhook
	var components PaymentComponents
	if provider == "razorpay" {

		payment = &RazorpayPayment{}
		refund = &RazorpayRefund{}
		webhook = &RazorpayWebhook{}

	} else if provider == "payu" {

		payment = &PayUPayment{}
		refund = &PayURefund{}
		webhook = &PayUWebhook{}

	}
	components = PaymentComponents{payment, refund, webhook}

	return components, nil
}

//Abstract Factory  way

type PaymentGatewayFactory interface {
	CreatePayment() Payment1
	CreateRefund() Refund
	CreateWebhook() Webhook
}

type RazorpayFactory struct{}

func (r RazorpayFactory) CreatePayment() Payment1 {
	return &RazorpayPayment{}
}

func (r RazorpayFactory) CreateRefund() Refund {
	return &RazorpayRefund{}
}

func (r RazorpayFactory) CreateWebhook() Webhook {
	return &RazorpayWebhook{}
}

type PayUFactory struct{}

func (p PayUFactory) CreatePayment() Payment1 {
	return &PayUPayment{}
}

func (p PayUFactory) CreateRefund() Refund {
	return &PayURefund{}
}

func (p PayUFactory) CreateWebhook() Webhook {
	return &PayUWebhook{}
}

func ProcessPayment1(factory PaymentGatewayFactory) {

	payment := factory.CreatePayment()
	refund := factory.CreateRefund()
	webhook := factory.CreateWebhook()

	payment.Pay(1000)
	refund.Refund(500)
	webhook.Handle()
}
func main() {
	components, err := ProcessPayment("razorpay")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	components.Payment.Pay(100)
	components.Refund.Refund(50)
	components.Webhook.Handle()

	//batter way to do that

	// razorpayFactory := RazorpayFactory{}

	// ProcessPayment1(razorpayFactory)

	// fmt.Println("---------")

	// payUFactory := PayUFactory{}

	// ProcessPayment1(payUFactory)
}
