package parking

import "errors"

var (
	ErrInvalidInput           = errors.New("invalid vehicle or spot input")
	ErrVehicleAlreadyParked   = errors.New("vehicle is already parked inside")
	ErrParkingFull            = errors.New("no available spot for this vehicle")
	ErrInvalidTicket          = errors.New("invalid or already used ticket ID")
	ErrDuplicateSpot          = errors.New("spot ID already exists")
	ErrRateNotFound           = errors.New("pricing rate not found for vehicle type")
	ErrPaymentRequired        = errors.New("payment is required before exit")
	ErrInvalidAmount          = errors.New("Invalid amount")
	ErrDuplicatePaymentMethod = errors.New("Duplicate Payment Method")
	ErrAlreadyPaid            = errors.New("Already paid")
	ErrUnknownPaymentMethod   = errors.New(" Umknow Payment Method")
	ErrPaymentInProgress      = errors.New("Error While Processing Payment")
)
