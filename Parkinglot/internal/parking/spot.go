package parking

type SpotSize string

const (
	Small    SpotSize = "SMALL"
	Medium   SpotSize = "MEDIUM"
	Large    SpotSize = "LARGE"
	Electric SpotSize = "ELECTRIC"
)

type Spot struct {
	ID               string
	Floor            int
	DistanceFromGate int   // Min-Heap sorting ke liye (kam distance = gate ke paas)
	Cost             int64 // is spot ka rate, ₹ per hour (VIP spot mehnga, basement sasta)
	Size             SpotSize
	IsOccupied       bool
	Vehicle          *Vehicle
}

// Har spot size ke liye: kis order mein jagah dhoondhni hai
var FallbackOrder = map[SpotSize][]SpotSize{
	Small:    {Small, Medium, Large},
	Medium:   {Medium, Large},
	Large:    {Large},
	Electric: {Electric, Medium, Large},
}
