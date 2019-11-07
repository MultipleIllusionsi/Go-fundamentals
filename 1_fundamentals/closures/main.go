package main

import "fmt"

func multiply(num1 int) func(num2 int) int {
	return func(num2 int) int {
		return num1 * num2
	}
}

func main() {
	result := multiply(4)(15)

	fmt.Println(result)
}
