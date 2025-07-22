package main

import (
	"fmt"
	"sync"
	"time"
)

// Create our worker function for processing work
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}

}

func main() {
	ports := make(chan int, 2)
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)
	end := time.Now()
	elsapsed := end.Sub(start)
	fmt.Printf("Time taken: %d ns", elsapsed.Nanoseconds())
}
