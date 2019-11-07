package main

import "fmt"

type human struct {
	first string
	last  string
}

type science struct {
	human
	fields []string
}

func main() {
	p1 := science{
		human: human{
			first: "petya",
			last:  "petrov",
		},
		fields: []string{
			"math",
			"phys",
			"music",
		},
	}

	for _, v := range p1.fields {
		fmt.Println(v)
	}
}
