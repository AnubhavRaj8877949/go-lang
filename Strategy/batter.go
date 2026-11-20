package main

import (
	"fmt"
	"strings"
	"time"
)

type PaymentStrategy interface {
	pay(amount int) error
}

type cardPayment struct{}

func (r *cardPayment) pay(amount int) error {

	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement card payment logic here
	fmt.Println("Processing card payment...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}

type cashPayment struct{}

func (r *cashPayment) pay(amount int) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement cash payment logic here
	fmt.Println("Processing cash payment...")
	// fack processing  delay
	time.Sleep(2 * time.Second)
	return nil
}

type UPIPayment struct{}

func (r *UPIPayment) pay(amount int) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement UPI payment logic here
	fmt.Println("Processing UPI payment...")

	time.Sleep(2 * time.Second)

	return nil
}

type GooglePayment struct{}

func (r *GooglePayment) pay(amount int) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement Google payment logic here
	fmt.Println("Processing Google payment...")

	time.Sleep(2 * time.Second)

	return nil
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

// this  method is constructor for PaymentProcessor (Constructor Injection)
func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{
		strategy: strategy,
	}
}

// this  member function is used to process the payment using the strategy set in the PaymentProcessor struct
func (p *PaymentProcessor) ProcessPayment(amount int) error {

	// Step 1: Check karo strategy nil toh nahi hai
	if p == nil || p.strategy == nil {
		return fmt.Errorf("payment strategy not configured")
	}

	// Step 2: Strategy ka Pay method call karo
	err := p.strategy.pay(amount)
	if err != nil {
		return err
	}

	// Step 3: Error return karo
	return nil
}

func validateAmount(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %d", amount)
	}

	return nil
}

// Factory method to create payment strategy based on the input
func PaymentFactory(paymentType string) (PaymentStrategy, error) {
	switch strings.ToLower(strings.TrimSpace(paymentType)) {

	case "card":
		return &cardPayment{}, nil
	case "cash":
		return &cashPayment{}, nil
	case "upi":
		return &UPIPayment{}, nil
	case "google":
		return &GooglePayment{}, nil
	default:
		return nil, fmt.Errorf("unsupported payment type: %q", paymentType)
	}
}

func main2() {
	fmt.Println("Welcome to the Payment System")
	paymentType := "card" // Change this to "cash" or "upi" to test different strategies
	amount := 100         // Change this to test different amounts
	paymentStrategy, err := PaymentFactory(paymentType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	processor := NewPaymentProcessor(paymentStrategy)

	err = processor.ProcessPayment(amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
