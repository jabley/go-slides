package main

import (
	"fmt"
	"math/rand"
)

// START OMIT
func worker(quit <-chan struct{}, result chan<- int) {
	for {
		select {
		case result <- rand.Intn(10000000):
		case <-quit:
			return
		}
	}
}

func main() {
	quit, result := make(chan struct{}), make(chan int)
	for i := 0; i < 100; i++ {
		go worker(quit, result)
	}
	// Wait for a worker to return a good result
	for {
		if <-result > 9999998 {
			break
		}
	}
	close(quit) // terminate all the workers
	fmt.Println("All done!")
}

// END OMIT
