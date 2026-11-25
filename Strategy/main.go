// This is   bad  way to write code.  It is not readable and not maintainable.  Please refactor this code to make it more readable and maintainable
package main

import (
	"fmt"
	"time"
)

func CardPayment(amount int) error {

	if amount <= 0 {
		fmt.Print("Invalid amount")
		return fmt.Errorf("invalid amount")
	}
	// Implement card payment logic here
	fmt.Println("Processing card payment...")
	// fack processing  delay
	time.Sleep(2 * time.Second)

	return nil

}
func CashPayment(amount int) error {
	if amount <= 0 {
		fmt.Print("Invalid amount")
		return fmt.Errorf("invalid amount")
	}
	// Implement cash payment logic here
	fmt.Println("Processing cash payment...")
	// fack processing  delay
	time.Sleep(2 * time.Second)
	return nil
}
func UPIpayment(amount int) error {
	if amount <= 0 {
		fmt.Print("Invalid amount")
		return fmt.Errorf("invalid amount")
	}
	// Implement UPI payment logic here
	fmt.Println("Processing UPI payment...")

	time.Sleep(2 * time.Second)

	return nil
}

func Payment(strategy string, amount int) error {
	switch strategy {
	case "card":
		return CardPayment(amount)
	case "cash":
		return CashPayment(amount)
	case "upi":
		return UPIpayment(amount)
	default:
		return fmt.Errorf("invalid payment strategy")
	}
}

func main4() {

	// result := Payment("card", 100)
	// if result != nil {
	// 	fmt.Println(result)
	// } else {
	// 	fmt.Println("Payment processed successfully")
	// }
	// main2()

}
