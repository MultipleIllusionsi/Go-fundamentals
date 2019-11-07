package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	msg := fmt.Sprintf("My name is %v %v", p.first, p.last)
	return msg
}

func main() {
	sergey := person{
		first: "Sergey",
		last:  "Zakharov",
		age:   22,
	}

	fmt.Println(sergey.speak())
}
