package main

import (
	"fmt"
	"parkinglot/internal/allocator"
	"parkinglot/internal/parking"
	"parkinglot/internal/payment"
	"parkinglot/internal/payment/fake"
	"parkinglot/internal/pricing"
	"sync"
)

func main() {
	paymentService := payment.NewPaymentService()

	// Register payment strategies
	paymentService.Register(payment.MethodUPI, &fake.FakeUPI{})
	paymentService.Register(payment.MethodCard, &fake.FakeCard{})
	paymentService.Register(payment.MethodCash, &fake.FakeCash{})

	lot := parking.NewParkingLot(pricing.NewHourlyPricing(), allocator.NewHeapSpotAllocator(allocator.ByNearest))

	// Floor 1: Total 8 parking spots

	// 2 Small spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-S1", Floor: 1, DistanceFromGate: 10, Size: parking.Small,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-S1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-S2", Floor: 1, DistanceFromGate: 20, Size: parking.Small,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-S2")
	}

	// 2 Medium spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-M1", Floor: 1, DistanceFromGate: 5, Size: parking.Medium,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-M1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-M2", Floor: 1, DistanceFromGate: 15, Size: parking.Medium,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-M2")
	}

	// 2 Large spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-L1", Floor: 1, DistanceFromGate: 8, Size: parking.Large,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-L1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-L2", Floor: 1, DistanceFromGate: 25, Size: parking.Large,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-L2")
	}

	// 2 Electric spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-E1", Floor: 1, DistanceFromGate: 30, Size: parking.Electric,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-E1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F1-E2", Floor: 1, DistanceFromGate: 35, Size: parking.Electric,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F1-E2")
	}

	// ================= FLOOR 2 =================

	// 2 Small spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-S1", Floor: 2, DistanceFromGate: 12, Size: parking.Small,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-S1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-S2", Floor: 2, DistanceFromGate: 22, Size: parking.Small,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-S2")
	}

	// 2 Medium spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-M1", Floor: 2, DistanceFromGate: 7, Size: parking.Medium,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-M1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-M2", Floor: 2, DistanceFromGate: 17, Size: parking.Medium,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-M2")
	}

	// 2 Large spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-L1", Floor: 2, DistanceFromGate: 9, Size: parking.Large,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-L1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-L2", Floor: 2, DistanceFromGate: 28, Size: parking.Large,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-L2")
	}

	// 2 Electric spots
	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-E1", Floor: 2, DistanceFromGate: 32, Size: parking.Electric,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-E1")
	}

	if err := lot.AddSpot(&parking.Spot{
		ID: "F2-E2", Floor: 2, DistanceFromGate: 38, Size: parking.Electric,
	}); err != nil {
		fmt.Println("[ERROR] AddSpot failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Spot added: %s\n", "F2-E2")
	}

	fmt.Println("\n========== PARKING LOT SIMULATION ==========")
	fmt.Println("Floors: 2")
	fmt.Println("Floor 1: 2 Small | 2 Medium | 2 Large | 2 Electric")
	fmt.Println("Floor 2: 2 Small | 2 Medium | 2 Large | 2 Electric")
	fmt.Println("Total Floors: 2 | Total Spots: 16")
	fmt.Println("============================================")
	// ==========================================
	// TEST 1: CYCLE PARKING
	// ==========================================

	fmt.Println("\n[TEST 1] Cycle Parking")

	cycle := &parking.Vehicle{
		LicensePlate: "PB-65-BB-2221",
		Type:         parking.Cycle,
	}

	t0, err := lot.ParkVehicle(cycle)

	if err != nil {
		fmt.Println("[ERROR] Cycle Parking Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Cycle Parked | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t0.TicketID, cycle.LicensePlate, t0.SpotID, t0.Floor)
	}

	// ==========================================
	// TEST 1: BIKE PARKING
	// ==========================================

	fmt.Println("\n[TEST 1] Bike Parking")

	bike := &parking.Vehicle{
		LicensePlate: "PB-65-BB-2222",
		Type:         parking.Bike,
	}

	t1, err := lot.ParkVehicle(bike)

	if err != nil {
		fmt.Println("[ERROR] Bike Parking Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Bike Parked | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t1.TicketID, bike.LicensePlate, t1.SpotID, t1.Floor)
	}

	// ==========================================
	// TEST 2: CAR PARKING
	// ==========================================

	fmt.Println("\n[TEST 2] Car Parking")

	car := &parking.Vehicle{
		LicensePlate: "PB-65-AA-1111",
		Type:         parking.Car,
	}

	t2, err := lot.ParkVehicle(car)

	if err != nil {
		fmt.Println("[ERROR] Car Parking Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Car Parked | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t2.TicketID, car.LicensePlate, t2.SpotID, t2.Floor)
	}

	// ==========================================
	// TEST 3: TRUCK PARKING
	// ==========================================

	fmt.Println("\n[TEST 3] Truck Parking")

	truck := &parking.Vehicle{
		LicensePlate: "PB-65-TT-3333",
		Type:         parking.Truck,
	}

	t3, err := lot.ParkVehicle(truck)

	if err != nil {
		fmt.Println("[ERROR] Truck Parking Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Truck Parked | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t3.TicketID, truck.LicensePlate, t3.SpotID, t3.Floor)
	}

	// ==========================================
	// TEST 4: ELECTRIC VEHICLE PARKING
	// ==========================================

	fmt.Println("\n[TEST 4] Electric Vehicle Parking")

	evCar := &parking.Vehicle{
		LicensePlate: "PB-65-AA-1000",
		Type:         parking.EV,
	}

	t4, err := lot.ParkVehicle(evCar)

	if err != nil {
		fmt.Println("[ERROR] EV Parking Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] EV Parked | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t4.TicketID, evCar.LicensePlate, t4.SpotID, t4.Floor)
	}

	// ==========================================
	// TEST 5: DUPLICATE VEHICLE PREVENTION
	// ==========================================

	fmt.Println("\n[TEST 5] Duplicate Vehicle Prevention")

	_, err = lot.ParkVehicle(car)

	if err != nil {
		fmt.Println("[SUCCESS] Duplicate Vehicle Rejected:", err)
	} else {
		fmt.Println("[ERROR] Duplicate Vehicle Was Allowed!")
	}

	// ==========================================
	// TEST 6: SIZE FALLBACK
	// ==========================================

	fmt.Println("\n[TEST 6] Bike Size Fallback")

	// Fill remaining Small spot
	bike3 := &parking.Vehicle{
		LicensePlate: "PB-65-BB-5555",
		Type:         parking.Bike,
	}

	t5, err := lot.ParkVehicle(bike3)

	if err != nil {
		fmt.Println("[ERROR] Second Bike Parking Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Bike Parked | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t5.TicketID, bike3.LicensePlate, t5.SpotID, t5.Floor)
	}

	// Try another Bike after Small spots are occupied
	bike2 := &parking.Vehicle{
		LicensePlate: "PB-65-BB-4444",
		Type:         parking.Bike,
	}

	t6, err := lot.ParkVehicle(bike2)

	if err != nil {
		fmt.Println("[ERROR] Bike Fallback Failed:", err)
	} else {
		fmt.Printf("[SUCCESS] Bike Fallback | Ticket: %s | Vehicle: %s | Spot: %s | Floor: %d\n",
			t6.TicketID, bike2.LicensePlate, t6.SpotID, t6.Floor)
	}

	// Test 8: Concurrent Parking
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			v := &parking.Vehicle{
				LicensePlate: fmt.Sprintf("PB-65-CC-%d", id),
				Type:         parking.Car,
			}

			ticket, parkErr := lot.ParkVehicle(v)

			if parkErr == nil {
				fmt.Printf("Concurrent Car %d parked at %s\n",
					id, ticket.SpotID)
			} else {
				fmt.Printf("Concurrent Car %d rejected: %v\n",
					id, parkErr)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("\n========== UNPARKING SIMULATION ==========")

	// Unpark Bike
	fmt.Println("\n[TEST 9] Unpark Bike")

	fee, err := lot.UnparkVehicle(t1.TicketID)

	if err != nil {
		fmt.Println("[ERROR] Bike Unparking Failed:", err)
	}

	// Process fake payment (variable ka naam 'payResp' rakha hai taaki 'payment' package se clash na ho)
	payResp, err := paymentService.Pay(
		t1.TicketID,
		fee,
		payment.MethodUPI,
	)

	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}

	// 5. Verify payment
	paid, err := paymentService.IsPaid(t1.TicketID)

	if err == nil {
		fmt.Println("Is Paid:", paid)
	}

	fmt.Println("Payment Successful!")
	fmt.Println("Payment ID:", payResp.ID)
	fmt.Println("Payment Status:", payResp.Status, payResp.Method)

	// -------------------------------Unpark Car-----------------------
	fmt.Println("\n[TEST 10] Unpark Car")

	fee, err = lot.UnparkVehicle(t2.TicketID)

	if err != nil {
		fmt.Println("[ERROR] Car Unparking Failed:", err)
	}

	// Process fake payment
	payResp, err = paymentService.Pay(
		t2.TicketID,
		fee,
		payment.MethodCard,
	)

	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}

	// 5. Verify payment
	paid, err = paymentService.IsPaid(t2.TicketID)

	if err == nil {
		fmt.Println("Is Paid:", paid)
	}
	fmt.Println("Payment Successful!")
	fmt.Println("Payment ID:", payResp.ID)
	fmt.Println("Payment Status:", payResp.Status, payResp.Method)

	fmt.Println("\n========== SIMULATION COMPLETED ==========")
}
