package main

import "fmt"

// Income struct
type Income struct {
	salary int
}

// Person struct
type Person struct {
	first  string
	last   string
	age    int
	income Income
}

func main() {
	John := Person{
		first:  "John",
		last:   "Smith",
		age:    30,
		income: Income{salary: 50000},
	}

	Bob := Person{
		first:  "Bob",
		last:   "Dole",
		age:    30,
		income: Income{salary: 50000},
	}

	fmt.Println(John, Bob)
	fmt.Println(John.first)
}
