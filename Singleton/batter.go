package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type databaseConnection struct{}

var (
	singletonInstance *databaseConnection
	once              sync.Once
)

func GetDatabase() *databaseConnection {
	if singletonInstance == nil {
		once.Do(func() {
			fmt.Println("Connecting to database: mongoDb")
			singletonInstance = &databaseConnection{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singletonInstance
}

//Question: sync.Once vs Mutex vs Atomic

var mu sync.Mutex
var counter int = 0

func Increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++

}

var counter3 int = 0

// mutex with WaitGroup
func IncrementWaitGroup(wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	counter3++

	defer wg.Done()
}

// Atomic
var counter2 atomic.Int64

func Increment2(wg *sync.WaitGroup) {
	counter2.Add(1)
	defer wg.Done()
}

func Initialize() {
	once.Do(func() {
		fmt.Println("Initialize only once")
	})
}

func main() {
	// connection1 := GetDatabase()
	// connection2 := GetDatabase()
	// connection3 := GetDatabase()

	// fmt.Println(connection1 == connection2)
	// fmt.Println(connection2 == connection3)

	// Testing sync.Once
	Initialize()
	Initialize()
	Initialize()

	// Testing Mutex
	for i := 0; i < 1000; i++ {
		go Increment()
	}
	fmt.Println("Counter with Mutex:", counter)

	// Testing Mutex using WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go IncrementWaitGroup(&wg)
	}

	wg.Wait()                                                   // Jab tak saari 1000 goroutines khatam na hon, yahan rukega
	fmt.Println("Counter with Mutex  and WaitGroup:", counter3) // Ab hamesha 1000 print hoga

	// Testing Atomic
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Increment2(&wg)

	}
	wg.Wait()
	fmt.Println("Counter with Atomic:", counter2.Load())

}
