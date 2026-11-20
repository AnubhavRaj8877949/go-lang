package main

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup, oddch, evench chan bool) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-oddch // wait for single
		fmt.Println(i)
		if i < 10 {
			evench <- true // sand   single
		}
	}
}
func printEven(wg *sync.WaitGroup, oddch, evench chan bool) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-evench // wait for single
		fmt.Println(i)
		if i < 10 {
			oddch <- true // sand   single
		}
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	oddch := make(chan bool)
	evench := make(chan bool)
	go printOdd(&wg, oddch, evench)
	go printEven(&wg, oddch, evench)
	oddch <- true
	wg.Wait() // wait  for  go routing  to finish

}
