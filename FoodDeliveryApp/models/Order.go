package models

import (
	"fmt"
	"fooddelivery/strategies"
	"strings"
	"time"
)

type Order interface {
	GetType() string
	ProcessPayment() error
	GetDetails() string
}

type BaseOrder struct {
	OrderId         int
	User            *User
	Restaurant      *Restaurant
	MenuList        []MenuItem
	PaymentStrategy strategies.PaymentStrategy
	Total           float64
	ScheduleTime    string
	CreatedAt       time.Time
	PaymentStatus   string
}

func (b *BaseOrder) processPayment() error {
	if b.PaymentStrategy == nil {
		return fmt.Errorf("payment strategy is not set")
	}

	b.PaymentStatus = "Processing"
	if err := b.PaymentStrategy.Pay(b.Total); err != nil {
		b.PaymentStatus = "Failed"
		return err
	}
	b.PaymentStatus = "Paid"
	return nil
}

func (b BaseOrder) details(orderType, fulfillment string) string {
	var details strings.Builder

	fmt.Fprintf(&details, "Order ID: %d\n", b.OrderId)
	fmt.Fprintf(&details, "Order type: %s\n", orderType)
	if b.CreatedAt.IsZero() {
		fmt.Fprintln(&details, "Created at: Not recorded")
	} else {
		fmt.Fprintf(&details, "Created at: %s\n", b.CreatedAt.Format(time.RFC1123))
	}

	if b.User == nil {
		fmt.Fprintln(&details, "Customer: Not assigned")
	} else {
		fmt.Fprintf(&details, "Customer: %s (ID: %d)\n", b.User.Name, b.User.Id)
		fmt.Fprintf(&details, "Customer address: %s\n", b.User.Address)
		fmt.Fprintf(&details, "Customer email: %s\n", b.User.Email)
		fmt.Fprintf(&details, "Customer mobile: %s\n", b.User.MobileNo)
	}

	if b.Restaurant == nil {
		fmt.Fprintln(&details, "Restaurant: Not assigned")
	} else {
		fmt.Fprintf(&details, "Restaurant: %s (ID: %d)\n", b.Restaurant.Name, b.Restaurant.Id)
		fmt.Fprintf(&details, "Restaurant address: %s\n", b.Restaurant.Address)
	}

	fmt.Fprintf(&details, "%s\n", fulfillment)
	fmt.Fprintln(&details, "Items:")
	for i, item := range b.MenuList {
		fmt.Fprintf(&details, "  %d. [%d] %s - ₹%.2f\n", i+1, item.Code, item.Name, item.Price)
	}
	fmt.Fprintf(&details, "Total: ₹%.2f\n", b.Total)

	paymentMethod := "Not selected"
	if b.PaymentStrategy != nil {
		paymentMethod = b.PaymentStrategy.MethodName()
	}
	paymentStatus := b.PaymentStatus
	if paymentStatus == "" {
		paymentStatus = "Not processed"
	}
	fmt.Fprintf(&details, "Payment method: %s\n", paymentMethod)
	fmt.Fprintf(&details, "Payment status: %s\n", paymentStatus)

	scheduleTime := b.ScheduleTime
	if scheduleTime == "" {
		scheduleTime = "Immediate"
	}
	fmt.Fprintf(&details, "Scheduled for: %s", scheduleTime)

	return details.String()
}

// -----------------------------Delivery Order------------------------
type DeliveryOrder struct {
	BaseOrder
	DeliveryAddress string
}

func (d *DeliveryOrder) GetType() string {
	return "Delivery Order"
}

func (d *DeliveryOrder) ProcessPayment() error {
	return d.BaseOrder.processPayment()
}

func (d *DeliveryOrder) GetDetails() string {
	return d.BaseOrder.details(
		d.GetType(),
		fmt.Sprintf("Delivery address: %s", d.DeliveryAddress),
	)
}

// -------------------------Pickup Order----------------------------
type PickupOrder struct {
	BaseOrder
}

func (d *PickupOrder) GetType() string {
	return "Pickup Order"
}

func (d *PickupOrder) ProcessPayment() error {
	return d.BaseOrder.processPayment()
}

func (d *PickupOrder) GetDetails() string {
	pickupLocation := "Pickup location: Not assigned"
	if d.Restaurant != nil {
		pickupLocation = fmt.Sprintf("Pickup location: %s", d.Restaurant.Address)
	}
	return d.BaseOrder.details(d.GetType(), pickupLocation)
}
