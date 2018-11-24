package main

import (
	"fmt"
)

// Node struct
type Node struct {
	value int
	next  *Node
}

// LinkedList struct
type LinkedList struct {
	head *Node
}

// Add - add a new item to the LinkedList
func (l *LinkedList) Add(i int) {
	new := Node{value: i}

	if l.head == nil {
		l.head = &new
		return
	}

	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &new
}

// PrintList - print the Linked List
func (l *LinkedList) PrintList() {
	cur := l.head

	fmt.Printf("%d -> ", cur.value)

	for cur.next != nil {
		cur = cur.next
		fmt.Printf("%d -> ", cur.value)
	}

	fmt.Printf("NIL\n")
}

// runExample - run a sample example demonstrating the LinkedList
func runExample() {
	ll := LinkedList{}
	ll.Add(10)
	ll.Add(91)
	ll.Add(22)
	ll.PrintList()
}

func main() {
	runExample()
}
