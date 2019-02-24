package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	y := []int{0, 1, 2}

	fmt.Println(x[1])

	for i, v := range x[1:3] {
		fmt.Println(i, v)
	}

	x = append(x, y...)

	x = append(x[:2], x[4:]...)

	fmt.Println(x)
}
