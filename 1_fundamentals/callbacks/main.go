package main

import "fmt"

// or (nums []int)
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

// or (f func(xi []int]) int, nums []int])
func isEvenWithFn(f func(xi ...int) int, nums ...int) int {
	newArr := make([]int, 10)
	for _, v := range nums {
		if v%2 == 0 {
			newArr = append(newArr, v)
		}
	}
	//or f(newArr)
	n := f(newArr...)
	return n
}

func main() {
	arr := []int{1, 3, 4, 5, 8}

	//or isEvenWithFn(sum, arr)
	a := isEvenWithFn(sum, arr...)

	fmt.Println(a)
}
