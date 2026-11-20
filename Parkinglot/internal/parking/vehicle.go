package parking

type VehicleType string

const (
	Bike  VehicleType = "BIKE"
	Car   VehicleType = "CAR"
	Truck VehicleType = "TRUCK"
	EV    VehicleType = "EV"
	Cycle VehicleType = "CYCLE"
)

// Single Source of Truth: VehicleType khud apna required SpotSize batayega
func (t VehicleType) RequiredSpotSize() SpotSize {
	switch t {
	case Bike, Cycle:
		return Small
	case Car:
		return Medium
	case Truck:
		return Large
	case EV:
		return Electric

	default:
		return Large // Safe fallback agar koi unknown type aaye
	}
}

type Vehicle struct {
	ID           string
	LicensePlate string
	Type         VehicleType
	OwnerName    string
	OwnerContact string
}
