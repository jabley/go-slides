package main

import "fmt"

// START OMIT
func worker(start chan bool) {
	// wait for the starting gun...
	<-start
	// ... do stuff
}

func main() {
	start := make(chan bool)

	for i := 0; i < 100; i++ {
		go worker(start)
	}

	close(start)

	fmt.Println("all workers running now")
}

// END OMIT
