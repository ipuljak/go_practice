package main

import (
	"fmt"
	"sort"
)

func mainSorting() {
	defaultSort()
	customSort()
}

func defaultSort() {
	xi := []int{5, 1, 7, 1, 3, 91, 22}
	xs := []string{"James", "Bond", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}

// Actor struct
type Actor struct {
	Name string
	Age  int
}

// Overriding string functionality
func (a Actor) String() string {
	return fmt.Sprintf("%s: %d", a.Name, a.Age)
}

// ByAge type
type ByAge []Actor

// Sort interface contains Len, Less, Swap
func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func customSort() {
	a1 := Actor{"James", 32}
	a2 := Actor{"Moneypenny", 27}
	a3 := Actor{"Q", 64}
	a4 := Actor{"M", 56}

	actors := []Actor{a1, a2, a3, a4}

	fmt.Println("Actors - ", actors)

	sort.Sort(ByAge(actors))

	fmt.Println("Actors sorted by age - ", actors)
}
