package fake

import (
	"fmt"
	"parkinglot/internal/payment"
)

type FakeUPI struct{}

func (f *FakeUPI) Pay(ticketID string, amount int64) error {
	fmt.Println("Processing fake UPI payment...")
	fmt.Printf("Ticket: %s, Amount: %d paise\n", ticketID, amount)
	return nil
}

type FakeCard struct{}

func (f *FakeCard) Pay(ticketID string, amount int64) error {

	fmt.Println("Processing fake Card payment...")
	fmt.Printf("Ticket: %s, Amount: %d paise\n", ticketID, amount)
	return nil
}

type FakeCash struct{}

func (f *FakeCash) Pay(ticketID string, amount int64) error {
	fmt.Println("Processing fake Cash payment...")
	fmt.Printf("Ticket: %s, Amount: %d paise\n", ticketID, amount)
	return nil
}

var (
	_ payment.PaymentStrategy = (*FakeUPI)(nil)
	_ payment.PaymentStrategy = (*FakeCard)(nil)
	_ payment.PaymentStrategy = (*FakeCash)(nil)
)
