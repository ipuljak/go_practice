package main

import "fmt"

// person struct
type person struct {
	first string
	last  string
}

// agent struct
type agent struct {
	person
	id       int
	licensed bool
}

// Attached this method to agent type
func (a agent) speak() {
	fmt.Printf("My name is %s %s.\n", a.first, a.last)
}

// Attached this method to person type
func (p person) speak() {
	fmt.Printf("My name is %s %s and I am a person.\n", p.first, p.last)
}

// Any other type that has "speak" is also of type human
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Bar - person", h.(person).first)
	case agent:
		fmt.Println("Bar - agent", h.(agent).first)
	}
}

func doStuff() {
	a1 := agent{
		person:   person{"Bob", "Dole"},
		id:       1,
		licensed: true,
	}

	p1 := person{"Ivan", "Smith"}

	a1.speak()
	p1.speak()

	// Polymorphism allows interfaces to accept many different types
	bar(a1)
	bar(p1)

	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function")
	}()

	// Func expression
	f := func() {
		fmt.Println("This is a func expression")
	}

	f()

	g := func(i int) {
		fmt.Println("Func expression - ", i)
	}

	g(100)

	r := printString()

	fmt.Println(r)

	// Return a function
	k := returnFunc()

	l := k()
	fmt.Println(l)

	// Callbacks
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println("Sum of evens - ", s2)

	// Closures
	x1 := incrementor()
	x2 := incrementor()

	fmt.Println(x1())
	fmt.Println(x1())
	fmt.Println(x2())

	fact := factorial(4)
	fmt.Println(fact)
}

func printString() string {
	s := "Hello world!"
	return s
}

func returnFunc() func() int {
	return func() int {
		return 451
	}
}

// Callbacks
func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}

// Passing in a function as a parameter
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

// Closures - close the scope of a variable
var x int // The scope of x is the entire package

// This will increment x without problems
func addToX() {
	x++

	{
		// We can print y here
		y := 1
		fmt.Println(y)
	}

	// but can't print y here
}

func incrementor() func() int {
	var xx int

	return func() int {
		// We can access xx here
		xx++
		return xx
	}
}

// Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
