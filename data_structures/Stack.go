package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack struct definition
type Stack struct {
	top  int
	data []int
}

// New - Initializes a new stack
func (s *Stack) New() {
	s.top = 0
	s.data = []int{}
}

// Check - Prints the top item in the stack
func (s *Stack) Check() {
	fmt.Println(s.top)
}

// Length - Print the length of the data stack
func (s *Stack) Length() {
	fmt.Println(len(s.data))
}

// Print - Print out the stack
func (s *Stack) Print() {
	fmt.Println(s.data)
}

// Add - Adds a new item to the stack
func (s *Stack) Add(i int) {
	s.top = i
	s.data = append(s.data, i)
}

// Pop - Remove and return the top item in the stack
func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		panic("Cannot pop an empty stack!")
	}

	result := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	s.top = s.data[len(s.data)-1]
	return result
}

// printHelp - Prints suggestions for the command line interface for the Stack
func printHelp() {
	fmt.Println()
	fmt.Println("Stack command line arguments...")
	fmt.Println("	NEW - Reset the stack to a new state")
	fmt.Println("	CHECK - Check the top element of the stack")
	fmt.Println("	LENGTH - Check the length of the stack")
	fmt.Println("	PRINT - Print out the elements of the stack")
	fmt.Println("	ADD [int] - Add a new integer value to the stack")
	fmt.Println("	POP - Pop out the top element of the stack")
	fmt.Println("	QUIT - Exit the command line interface")
	fmt.Println()
}

// Execute - Perform an action by inputting it in the command line
// NEW, CHECK, LENGTH, PRINT, ADD [int], POP, QUIT, HELP
func (s *Stack) Execute() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")

		switch strings.ToLower(strs[0]) {
		case "new":
			s.New()
		case "check":
			s.Check()
		case "length":
			s.Length()
		case "print":
			s.Print()
		case "add":
			if len(strs[1]) > 0 {
				integer, _ := strconv.Atoi(strs[1])
				s.Add(integer)
			} else {
				fmt.Println("Please enter an integer!")
			}
		case "pop":
			s.Pop()
		case "quit":
			os.Exit(3)
		case "help":
			printHelp()
		default:
			fmt.Println("Invalid entry. Please type 'help' for valid entries.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

func main() {
	stack := Stack{}
	stack.Execute()
}
