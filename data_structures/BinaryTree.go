package main

import (
	"fmt"
)

// Node struct
type Node struct {
	value int
	left  *Node
	right *Node
}

// BinaryTree struct
type BinaryTree struct {
	root *Node
}

// String prints a visual representation of the tree
func (bt *BinaryTree) String() {
	fmt.Println("------------------------------------------------")
	stringify(bt.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.value)
		stringify(n.right, level)
	}
}

func main() {
	a := Node{value: 10}
	b := Node{value: 8}
	c := Node{value: 5}
	d := Node{value: 2}
	e := Node{value: 11}
	f := Node{value: 7, right: &a}
	g := Node{value: 3, left: &b, right: &c}
	h := Node{value: 4, left: &f, right: &g}
	i := Node{value: 9, left: &d, right: &e}
	j := Node{value: 1, left: &h, right: &i}

	tree := BinaryTree{root: &j}

	tree.String()
}
