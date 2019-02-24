package main

import "fmt"

func main() {
	// Create slices
	x := []int{4, 5, 7, 8, 42}
	y := []int{0, 1, 2}

	fmt.Println(x[1])

	// Loop over slice
	for i, v := range x[1:3] {
		fmt.Println(i, v)
	}

	// Merge slices
	x = append(x, y...)

	// Delete from slice
	x = append(x[:2], x[4:]...)

	fmt.Println(x)

	// "make" slices
	a := make([]int, 10, 100)

	fmt.Println(a)

	// Maps

	m := map[string]int{
		"Ivan":  27,
		"Holly": 25,
	}

	// Checking if something exists in a map
	if _, ok := m["Steve"]; ok {
		fmt.Println("Steve exists")
	} else {
		fmt.Println("Steve does not exist")
	}

	// Add elements + range in map
	m["Steve"] = 58

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Delete map
	delete(m, "Steve")
	fmt.Println(m)
}
