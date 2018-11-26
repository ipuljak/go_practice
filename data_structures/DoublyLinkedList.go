package main

import (
	"fmt"
)

// Node struct
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// DoublyLinkedList struct
type DoublyLinkedList struct {
	head *Node
}

// Add - add a new item to the DoublyLinkedList
func (l *DoublyLinkedList) Add(i int) {
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
	new.prev = cur
}

// Remove - remove a node from the DoublyLinkedList
func (l *DoublyLinkedList) Remove(i int) {
	cur := l.head

	if cur == nil {
		panic("DoublyLinkedList is empty!")
	}

	if cur.value == i {
		l.head = &Node{}
		return
	}

	for cur.next != nil {
		cur = cur.next

		if cur.value == i {
			cur.prev.next = cur.next
			return
		}
	}
}

// PrintList - print the Linked List
func (l *DoublyLinkedList) PrintList() {
	cur := l.head

	fmt.Printf("%d = ", cur.value)

	for cur.next != nil {
		cur = cur.next
		fmt.Printf("%d = ", cur.value)
	}

	fmt.Printf("NIL\n")
}

// PrintListReverse - print the list backwards
func (l *DoublyLinkedList) PrintListReverse() {
	fmt.Printf("NIL = ")

	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	fmt.Printf("%d = ", cur.value)

	for cur.prev != nil {
		cur = cur.prev

		if cur.prev != nil {
			fmt.Printf("%d = ", cur.value)
		} else {
			fmt.Printf("%d", cur.value)
		}
	}

	fmt.Println()
}

// runExample - run a sample example demonstrating the DoublyLinkedList
func runExample() {
	ll := DoublyLinkedList{}
	ll.Add(10)
	ll.Add(91)
	ll.Add(22)
	ll.Add(183)
	ll.Add(4)
	ll.Add(23)
	ll.PrintList()
	ll.PrintListReverse()

	ll.Remove(183)
	ll.PrintList()

	ll.Remove(23)
	ll.PrintList()

	ll.Remove(10)
	ll.PrintList()
}

func main() {
	runExample()
}
