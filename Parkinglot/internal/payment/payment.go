package payment

import "time"

type PaymentMethod string

const (
	MethodCash PaymentMethod = "CASH"
	MethodUPI  PaymentMethod = "UPI"
	MethodCard PaymentMethod = "CARD"
)

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "PENDING"
	PaymentSuccess PaymentStatus = "SUCCESS"
	PaymentFailed  PaymentStatus = "FAILED"
)

type Payment struct {
	ID        string
	TicketID  string        // Exit isse dhoondhega
	Amount    int64         // paise
	Method    PaymentMethod // strategy nahi, sirf naam
	Status    PaymentStatus // Exit isse check karega
	CreatedAt time.Time
}

type PaymentStrategy interface {
	Pay(ticketID string, amount int64) error
}
