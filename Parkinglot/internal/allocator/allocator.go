package allocator

import (
	"parkinglot/internal/parking"

	"github.com/zyedidia/generic/heap"
)

// HeapSpotAllocator: har SpotSize ka ek min-heap. Heap ka order `less` decide karta hai
// (ByNearest / ByCheapest), isliye order badalna = sirf naya less function (Strategy).
type HeapSpotAllocator struct {
	available map[parking.SpotSize]*heap.Heap[*parking.Spot]
}

func NewHeapSpotAllocator(less func(a, b *parking.Spot) bool) *HeapSpotAllocator {
	available := make(map[parking.SpotSize]*heap.Heap[*parking.Spot])
	for size := range parking.FallbackOrder {
		available[size] = heap.New(less)
	}
	return &HeapSpotAllocator{available: available}
}

func (a *HeapSpotAllocator) Add(spot *parking.Spot) error {
	if spot == nil || spot.ID == "" {
		return parking.ErrInvalidInput
	}
	h, ok := a.available[spot.Size]
	if !ok {
		return parking.ErrInvalidInput
	}
	h.Push(spot)
	return nil
}

func (a *HeapSpotAllocator) Allocate(v parking.Vehicle) *parking.Spot {
	for _, size := range parking.FallbackOrder[v.Type.RequiredSpotSize()] {
		if spot, ok := a.available[size].Pop(); ok { // ok=false matlab heap khaali, agla size try karo
			return spot
		}
	}
	return nil
}

func (a *HeapSpotAllocator) Release(spot *parking.Spot) {
	if h, ok := a.available[spot.Size]; ok {
		h.Push(spot)
	}
}

var _ parking.SpotAllocator = (*HeapSpotAllocator)(nil)
