package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "first was changed"
	p.last = "last was changed"

	// (*p).first = ""
	// (*p).last = ""
}

func main() {
	somePerson := person{
		first: "vasya",
		last:  "ivanov",
	}
	fmt.Println(somePerson)

	changeMe(&somePerson)
	fmt.Println(somePerson)
}
