package payment

import (
	"fmt"
	"parkinglot/internal/parking"
	"sync"
	"time"
)

type PaymentService struct {
	mu       sync.Mutex
	methods  map[PaymentMethod]PaymentStrategy
	payments map[string]*Payment // ticketID → payment
	nextID   int64
	nowFunc  func() time.Time
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		methods:  make(map[PaymentMethod]PaymentStrategy),
		payments: make(map[string]*Payment),
		nowFunc:  time.Now,
	}

}
func validateAmount(amount int64) error {
	if amount <= 0 {
		return parking.ErrInvalidAmount
	}

	return nil
}

func (s *PaymentService) Register(m PaymentMethod, st PaymentStrategy) error {

	if m == "" || s == nil {
		return parking.ErrInvalidInput
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.methods[m]; exists {

		return parking.ErrDuplicatePaymentMethod
	}
	s.methods[m] = st
	return nil

}
func (s *PaymentService) Pay(
	ticketID string,
	amount int64,
	m PaymentMethod,
) (Payment, error) {

	if s == nil || ticketID == "" {
		return Payment{}, parking.ErrInvalidInput
	}

	if err := validateAmount(amount); err != nil {
		return Payment{}, err
	}

	s.mu.Lock()

	// Find payment strategy.
	st, ok := s.methods[m]
	if !ok {
		s.mu.Unlock()
		return Payment{}, parking.ErrUnknownPaymentMethod
	}

	// Check existing payment.
	old, exists := s.payments[ticketID]

	if exists && old.Status == PaymentSuccess {
		s.mu.Unlock()
		return Payment{}, parking.ErrAlreadyPaid
	}

	if exists && old.Status == PaymentPending {
		s.mu.Unlock()
		return Payment{}, parking.ErrPaymentInProgress
	}

	// Create new pending payment.
	s.nextID++

	p := &Payment{
		ID:        fmt.Sprintf("%d", s.nextID),
		TicketID:  ticketID,
		Amount:    amount,
		Method:    m,
		Status:    PaymentPending,
		CreatedAt: s.nowFunc(),
	}

	s.payments[ticketID] = p

	s.mu.Unlock()

	// Execute strategy outside mutex.
	err := st.Pay(ticketID, amount)

	s.mu.Lock()
	defer s.mu.Unlock()

	if err != nil {
		p.Status = PaymentFailed
		return *p, err
	}

	p.Status = PaymentSuccess

	return *p, nil
}
func (s *PaymentService) IsPaid(ticketID string) (bool, error) {
	if s == nil || ticketID == "" {
		return false, parking.ErrInvalidInput
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.payments[ticketID]

	if !ok {
		return false, nil
	}
	return p.Status == PaymentSuccess, nil

}
