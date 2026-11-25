package main

import "fmt"

// ====================================================
// 1. TARGET INTERFACE (Jo Hamara App Samajhta Hai)
// ====================================================
// Hamara poora project sirf is 'Pay' function ko jaanta hai
type PaymentProcessor interface {
	Pay(amountInRupees int)
}

// ====================================================
// 2. EXISTING WORKING SYSTEM (Purana Razorpay Gateway)
// ====================================================
type Razorpay struct{}

func (r *Razorpay) Pay(amountInRupees int) {
	fmt.Printf("[Razorpay] ₹%d ka payment safaltapoorvak ho gaya.\n", amountInRupees)
}

// ====================================================
// 3. ADAPTEE (Naya 3rd-Party SDK - Jiska Code Hum Badal Nahi Sakte)
// ====================================================
// Dhyan dein: Na iska function naam 'Pay' hai, aur na hi yeh 'int' rupees leta hai!
type StripeUSSDK struct{}

func (s *StripeUSSDK) MakePaymentInCents(cents float64) {
	fmt.Printf("[Stripe US SDK] %.2f cents ka international payment process ho gaya.\n", cents)
}

// ====================================================
// 4. ADAPTER (Bech ka Translator)
// ====================================================
// Yeh struct hamare 'PaymentProcessor' interface ko implement karega,
// aur andar 'StripeUSSDK' ko chupake rakhega.
type StripeAdapter struct {
	stripeSDK *StripeUSSDK // Purane/3rd-party object ko andar rakha (Composition)
}

// Adapter ne hamare app ka rule ('Pay') poora kiya
func (adapter *StripeAdapter) Pay(amountInRupees int) {
	fmt.Println("Adapter: Rupees ko Cents mein convert kar raha hoon...")

	// Translation Logic: 1 Rupee ≈ 1.2 Cents (Example rate)
	cents := float64(amountInRupees) * 1.2

	// Ab andar hi andar 3rd-party SDK ka asli function call kar diya
	adapter.stripeSDK.MakePaymentInCents(cents)
}

// ====================================================
// 5. CLIENT CODE (Hamara Checkout System)
// ====================================================
// Dekhiye! Is function ko bilkul nahi pata ki peeche Razorpay chal raha hai ya Stripe!
func CheckoutOrder(processor PaymentProcessor, amount int) {
	fmt.Println("\n--- Order Checkout Shuru ---")
	processor.Pay(amount)
}

func main() {
	// Case 1: Normal Razorpay use karna (Direct kaam karta hai)
	razorpay := &Razorpay{}
	CheckoutOrder(razorpay, 500)

	// Case 2: Stripe SDK use karna
	// Agar hum seedha 'CheckoutOrder(stripe, 500)' likhenge toh Go Compiler ERROR dega!
	// Isliye humne Stripe ko pehle 'StripeAdapter' ke andar daala:
	rawStripe := &StripeUSSDK{}
	stripeAdapter := &StripeAdapter{stripeSDK: rawStripe}

	// Ab CheckoutOrder bina kisi dikkat ke Stripe ko bhi chala lega!
	CheckoutOrder(stripeAdapter, 500)
}
