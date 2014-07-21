package main

import "fmt"

func main() {
	// START OMIT
	// Make an unbuffered channel
	c := make(chan int)
	go func() {
		// Send a value to the channel
		c <- 1
		c <- 2
		c <- 3
	}()
	// Receive a value from the channel
	fmt.Println(<-c)
	x := <-c
	fmt.Println(x)
	x, ok := <-c
	fmt.Printf("%v %v\n", x, ok)
	close(c)
	x, ok = <-c
	fmt.Printf("%v %v\n", x, ok)
	// END OMIT
}
