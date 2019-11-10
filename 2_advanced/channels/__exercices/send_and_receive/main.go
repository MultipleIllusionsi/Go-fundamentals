package main

import (
	"fmt"
)

// chan<- int | ONLY SEND (but not receive)
// <-cs      | ONLY RECEIVE from channel
// cr <- 42 | SEND TO CHANNEL
func main() {
	////// doesn't work
	// cs := make(chan<- int)
	////// work
	cs := make(chan int)
	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
