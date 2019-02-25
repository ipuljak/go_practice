package main

import (
	"fmt"
	"math"
)

func mainPointers() {
	a := 42
	fmt.Println(a)  // See the item
	fmt.Println(&a) // See the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	//var b *int = &a
	b := &a
	fmt.Println(b)  // See the address
	fmt.Println(*b) // See the item at the address

	// When To Use Pointers
	x := 0
	fmt.Println("x before: ", x)
	fmt.Println("x before: ", &x)
	foo(&x)
	fmt.Println("x after: ", x)
	fmt.Println("x after: ", &x)

	// Method sets - set of methods attached to a type
	c := circle{radius: 5}
	info(c)
	info(&c)
}

func foo(y *int) {
	fmt.Println("y before: ", y)
	fmt.Println("y before: ", *y)
	*y = 43
	fmt.Println("y after: ", y)
	fmt.Println("y after: ", *y)
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// c is a NON POINTER RECEIVER - it will accept c, or *c
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// c is a POINTER RECEIVER - it would only take in a *c
func (c *circle) areaPointer() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println("Area - ", s.area())
}
