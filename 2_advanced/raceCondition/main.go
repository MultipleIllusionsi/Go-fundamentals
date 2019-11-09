package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			// Gosched yields the processor, allowing other
			// goroutines to run (but doesnt suspend the curr gorts)
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
