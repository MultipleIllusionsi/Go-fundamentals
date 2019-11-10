package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

////////////////// v2
// func main() {
// 	c := gen()
// 	receive(c)

// 	fmt.Println("about to exit")
// }

// func receive(c <-chan int) {
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

// func gen() <-chan int {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()

// 	return c
// }
