package main

import (
	"fmt"
)

type shape interface {
	area() int
}

type square struct {
	length int
}

type circle struct {
	rad int
}

func (s square) area() int {
	return s.length * s.length
}

func (c circle) area() int {
	return 3 * (c.rad * c.rad)
}

func measure(s shape) {
	fmt.Println("Params of the object", s)
	fmt.Println("area of the object", s.area())
}

func main() {
	squareElem := square{
		length: 12,
	}

	circleElem := circle{
		rad: 7,
	}

	measure(squareElem)
	measure(circleElem)
}
