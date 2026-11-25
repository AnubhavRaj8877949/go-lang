package strategies

import (
	"fmt"
	"time"
)

type PaymentStrategy interface {
	Pay(amount float64) error
	MethodName() string
}

func validateAmount(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %.2f", amount)
	}

	return nil
}

type cardPayment struct {
	CardNumber string
}

func (r *cardPayment) MethodName() string {
	return "Credit Card"
}

func (r *cardPayment) Pay(amount float64) error {

	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement card payment logic here
	fmt.Printf("Paid ₹%.2f  using Credit Card (%s)\n", amount, r.CardNumber)
	time.Sleep(2 * time.Second)

	return nil

}

type cashPayment struct{}

func (r *cashPayment) MethodName() string {
	return "Cash"
}

func (r *cashPayment) Pay(amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement cash payment logic here
	fmt.Println("Processing cash payment...")
	time.Sleep(2 * time.Second)
	return nil
}

type UPIPayment struct {
	UPIID string
}

func (r *UPIPayment) MethodName() string {
	return "UPI"
}

func (r *UPIPayment) Pay(amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement UPI payment logic here
	fmt.Printf("Paid ₹%.2f using UPI (%s)\n", amount, r.UPIID)
	time.Sleep(2 * time.Second)
	return nil
}

type GooglePayment struct{}

func (r *GooglePayment) MethodName() string {
	return "Google Pay"
}

func (r *GooglePayment) Pay(amount float64) error {
	if err := validateAmount(amount); err != nil {
		return err
	}
	// Implement Google payment logic here
	fmt.Println("Processing Google payment...")
	time.Sleep(2 * time.Second)
	return nil
}
