package pricing

import (
	"math"
	"parkinglot/internal/parking"
	"time"
)

// ------------------------Hourly Rate ----------------------------

type HourlyPricing struct {
	Rates map[parking.VehicleType]int64
}

func NewHourlyPricing() *HourlyPricing {
	return &HourlyPricing{
		Rates: map[parking.VehicleType]int64{
			parking.Bike:  10,
			parking.Car:   20,
			parking.Truck: 50,
			parking.EV:    50,
		},
	}
}

func (hp *HourlyPricing) CalculateFee(t *parking.Ticket, exitTime time.Time) (int64, error) {
	rate, ok := hp.Rates[t.Vehicle.Type]
	if !ok {
		return 0, parking.ErrRateNotFound
	}
	duration := exitTime.Sub(t.EntryTime)
	hours := math.Ceil(duration.Hours())
	if hours < 1 {
		hours = 1
	}

	return int64(hours) * rate, nil

}

// -----------------------Flat Rate ---------------------------------
type FlatRatePricing struct {
	Rates map[parking.VehicleType]int64
}

func NewFlatRatePricing() *FlatRatePricing {
	return &FlatRatePricing{
		Rates: map[parking.VehicleType]int64{
			parking.Bike:  30,
			parking.Car:   60,
			parking.Truck: 100,
			parking.EV:    50,
		},
	}
}

func (fp *FlatRatePricing) CalculateFee(t *parking.Ticket, exitTime time.Time) (int64, error) {
	rate, ok := fp.Rates[t.Vehicle.Type]
	if !ok {
		return 0, parking.ErrRateNotFound
	}
	return rate, nil
}

// ----------------------Weekend Rate----------------------------------------
type WeekendPricing struct {
	BasePricing     parking.PricingStrategy
	DiscountPercent int64
}

func NewWeekendPricing(basePricing parking.PricingStrategy, discount int64) *WeekendPricing {
	return &WeekendPricing{
		BasePricing:     basePricing,
		DiscountPercent: discount,
	}
}

func (wp *WeekendPricing) CalculateFee(t *parking.Ticket, exitTime time.Time) (int64, error) {
	// 1. Base strategy se normal fee calculate karo.
	baseFee, err := wp.BasePricing.CalculateFee(t, exitTime)
	if err != nil {
		return 0, err // base fail hua toh hum bhi fail, error aage bhejo
	}
	// 2. Discount calculate karo.
	var discount int64 = 0
	if exitTime.Weekday() == time.Saturday || exitTime.Weekday() == time.Sunday {
		discount = baseFee * wp.DiscountPercent / 100
	}
	// 3. Final fee return karo.
	finalFee := baseFee - discount

	return finalFee, nil
}
