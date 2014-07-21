package main

import "fmt"

func main() {
	// START OMIT
	// Make an unbuffered channel
	c := make(chan bool)
	go func() {

		// ... do some stuff
		close(c)
	}()

	// do some other stuff

	// have we finished doing the first stuff?
	<-c
	fmt.Println("Finished!")
	// END OMIT
}
