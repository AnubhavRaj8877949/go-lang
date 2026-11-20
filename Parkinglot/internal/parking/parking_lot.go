package parking

import (
	"fmt"
	"sync"
	"time"
)

// SpotAllocator decide karta hai ki gaadi ko kaunsa spot milega.
type SpotAllocator interface {
	Add(spot *Spot) error
	Allocate(v Vehicle) *Spot
	Release(spot *Spot)
}
type PricingStrategy interface {
	CalculateFee(ticket *Ticket, exitTime time.Time) (int64, error)
}

type ParkingLot struct {
	mu             sync.Mutex
	allocator      SpotAllocator
	allSpotsByID   map[string]*Spot
	activeTickets  map[string]*Ticket
	activeVehicles map[string]string
	pricing        PricingStrategy
	nextTicketNo   int64
	nowFunc        func() time.Time // Testability ke liye clock injectio
}

func NewParkingLot(pricing PricingStrategy, allocator SpotAllocator) *ParkingLot {

	return &ParkingLot{
		allocator:      allocator,
		allSpotsByID:   make(map[string]*Spot), // empty map
		activeTickets:  make(map[string]*Ticket),
		activeVehicles: make(map[string]string),
		pricing:        pricing,
		nowFunc:        time.Now,
	}

}

func (pl *ParkingLot) AddSpot(spot *Spot) error {

	if spot == nil || spot.ID == "" {
		return ErrInvalidInput
	}

	pl.mu.Lock()
	defer pl.mu.Unlock()

	// 1. Same ID pehle se toh nahi hai
	if _, exists := pl.allSpotsByID[spot.ID]; exists {
		return ErrDuplicateSpot
	}

	// 2. Allocator mein daalo (size galat hua toh yahi error dega)
	if err := pl.allocator.Add(spot); err != nil {
		return err
	}
	// 3. Sab sahi hai, tabhi ID map mein save karo
	pl.allSpotsByID[spot.ID] = spot
	return nil
}

func (pl *ParkingLot) ParkVehicle(v *Vehicle) (*Ticket, error) {

	if v == nil || v.LicensePlate == "" {
		return nil, ErrInvalidInput
	}
	pl.mu.Lock()
	defer pl.mu.Unlock()

	// 1. Duplicate Vehicle Check
	if _, isParked := pl.activeVehicles[v.LicensePlate]; isParked {
		return nil, ErrVehicleAlreadyParked
	}

	// 2. Allocator se spot maango (size aur rules allocator decide karega)
	spot := pl.allocator.Allocate(*v)

	if spot == nil {
		return nil, ErrParkingFull
	}

	// 3. Ticket banao (gaadi ki copy rakho)
	now := pl.nowFunc()
	pl.nextTicketNo++
	ticketID := fmt.Sprintf("TKT-%d", pl.nextTicketNo)
	ticket := &Ticket{
		TicketID:  ticketID,
		SpotID:    spot.ID,
		Floor:     spot.Floor,
		Vehicle:   *v, // pointer se value copy
		EntryTime: now,
	}

	// 4. State Update
	spot.IsOccupied = true
	spot.Vehicle = &ticket.Vehicle // spot bhi copy ko point kare, caller ke pointer ko nahi
	pl.activeTickets[ticketID] = ticket
	pl.activeVehicles[v.LicensePlate] = ticketID
	return ticket, nil

}

func (pl *ParkingLot) UnparkVehicle(ticketID string) (int64, error) {
	if ticketID == "" {
		return 0, ErrInvalidInput
	}
	pl.mu.Lock()
	defer pl.mu.Unlock()

	ticket, exists := pl.activeTickets[ticketID]
	if !exists {
		return 0, ErrInvalidTicket
	}
	exitTime := pl.nowFunc()
	fee, err := pl.pricing.CalculateFee(ticket, exitTime)
	if err != nil {
		return 0, err
	}

	spot := pl.allSpotsByID[ticket.SpotID]
	spot.IsOccupied = false
	spot.Vehicle = nil
	pl.allocator.Release(spot)

	delete(pl.activeTickets, ticketID)
	delete(pl.activeVehicles, ticket.Vehicle.LicensePlate)
	return fee, nil
}
