package main

import (
	"fmt"
	"log"

	"fooddelivery/strategies"
)

func main() {
	tomato := NewTomatoApp()

	// Load the restaurant catalog before the client starts ordering.
	tomato.InitializeRestaurants()

	// The client creates an account; the app assigns the user ID.
	user := tomato.CreateUser(
		"Anubhav",
		"Sector 17, Chandigarh",
		"8877949888",
		"rajputanubhav65@gmail.com",
	)
	fmt.Printf("Welcome, %s (user #%d)\n", user.Name, user.Id)

	// The client searches for restaurants and chooses one.
	restaurants := tomato.SearchRestaurants("Chandigarh")
	if len(restaurants) == 0 {
		log.Fatal("no restaurants found")
	}
	fmt.Println("Found restaurant:", restaurants[0].Name)
	if err := tomato.SelectRestaurant(user, restaurants[0]); err != nil {
		log.Fatal(err)
	}

	// The client adds menu items to the cart and reviews it.
	if err := tomato.AddToCart(user, 101); err != nil {
		log.Fatal(err)
	}
	if err := tomato.AddToCart(user, 102); err != nil {
		log.Fatal(err)
	}
	tomato.PrintUserCart(user)

	// The client chooses a payment strategy and schedules a pickup order.
	payment := &strategies.UPIPayment{UPIID: "rahul@upi"}
	order, err := tomato.CheckoutScheduled(user, "Pickup", payment, "Today 3PM")
	if err != nil {
		log.Fatal(err)
	}

	// Payment, order storage, notification, and cart clearing happen in TomatoApp.
	if err := tomato.PayForOrder(user, order); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Checkout complete; confirmation notification sent.")
}
