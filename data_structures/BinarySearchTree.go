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

// BinarySearchTree struct
type BinarySearchTree struct {
	root *Node
}

// Add - add a new value to the Binary Search Tree
func (bst *BinarySearchTree) Add(v int) {
	n := Node{value: v}

	if bst.root == nil {
		bst.root = &n
		return
	}

	cur := bst.root

	for cur != nil {
		if v == cur.value {
			fmt.Println("Value already exists in Binary Search Tree!")
			return
		}

		// If value is smaller than the current node, place it to the left
		if v < cur.value {
			if cur.left != nil {
				cur = cur.left
			} else {
				cur.left = &n
				return
			}
			// If the value is greater than the current node, place it to the right
		} else if v > cur.value {
			if cur.right != nil {
				cur = cur.right
			} else {
				cur.right = &n
				return
			}
		}
	}
}

// PREORDER BST TRAVERSAL

// Preorder - Return the slice of node values given a Binary Search Tree
func (bst *BinarySearchTree) Preorder() []int {
	return PreorderTraversal(bst.root, []int{})
}

// PreorderTraversal - Traverse the given node and running slice of values
func PreorderTraversal(root *Node, order []int) []int {
	return order
}

// INORDER BST TRAVERSAL

// Inorder - Print the inorder traversal of the Binary Search Tree
func (bst *BinarySearchTree) Inorder() []int {
	return InorderTraversal(bst.root, []int{})
}

// InorderTraversal - Traverse the given node and running slice of values
func InorderTraversal(root *Node, order []int) []int {
	if root.left != nil {
		return InorderTraversal(root.left, order)
	}

	order = append(order, root.value)

	if root.right != nil {
		return InorderTraversal(root.right, order)
	}

	return order
}

// POSTORDER BST TRAVERSAL

// Postorder - Print the postorder traversal of the Binary Search Tree
func (bst *BinarySearchTree) Postorder() []int {
	return PostorderTraversal(bst.root, []int{})
}

// PostorderTraversal - Traverse the given node and running slice of values
func PostorderTraversal(root *Node, order []int) []int {
	return order
}

// String prints a visual representation of the tree
func (bst *BinarySearchTree) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
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
	root := Node{value: 10}
	tree := BinarySearchTree{root: &root}

	tree.Add(7)
	tree.Add(14)
	tree.Add(5)
	tree.Add(3)
	tree.Add(18)
	tree.Add(15)
	tree.Add(17)
	tree.Add(15)
	tree.Add(4)
	tree.Add(12)

	tree.String()

	preorder := tree.Preorder()
	fmt.Println("Preorder Traversal - ", preorder)
}