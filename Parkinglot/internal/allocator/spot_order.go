package allocator

import "parkinglot/internal/parking"

// ByNearest: pehle chhota floor, phir gate ke paas wala
func ByNearest(a, b *parking.Spot) bool {
	// Rule 1: Pehle Floor check karo (Floor 1 chhota hai Floor 2 se)
	if a.Floor != b.Floor {
		return a.Floor < b.Floor
	}

	// Rule 2: Agar dono spots same floor par hain, toh Gate se distance check karo
	return a.DistanceFromGate < b.DistanceFromGate

}

// ByCheapest: sabse sasta pehle. Price same ho toh nearest wala.
func ByCheapest(a, b *parking.Spot) bool {
	// agar a.Cost != b.Cost → jiska Cost kam, wo pehle
	if a.Cost != b.Cost {
		return a.Cost < b.Cost
	}
	// warna → return ByNearest(a, b)
	return ByNearest(a, b)
}
