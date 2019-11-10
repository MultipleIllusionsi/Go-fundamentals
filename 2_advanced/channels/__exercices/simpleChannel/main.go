package main

import (
	"fmt"
)

/////// doesn't work
func main() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}

//////// 1. anonumys func
// func main() {
// 	c := make(chan int)
// 	go func() {
// 		c <- 42
// 	}()
// 	fmt.Println(<-c)
// }

//////// 2. buffer
// func main() {
// 	c := make(chan int, 1)
// 	c <- 42
// 	fmt.Println(<-c)
// }
