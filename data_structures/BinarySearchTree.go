package main

import (
	"bytes"
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

// PrintPreorder - Print the preorder traversal of the tree
func (bst *BinarySearchTree) PrintPreorder() {
	preorderChannel := make(chan string)
	preorderBuffer := new(bytes.Buffer)

	go bst.Preorder(preorderChannel)

	for {
		val, i := <-preorderChannel
		if !i {
			break
		}

		preorderBuffer.WriteString(val)
		preorderBuffer.WriteString(" ")
	}

	fmt.Println("Preorder Traversal - ", preorderBuffer.String())
}

// Preorder - Return the slice of node values given a Binary Search Tree
func (bst *BinarySearchTree) Preorder(ch chan string) {
	PreorderTraversal(bst.root, ch)
	close(ch)
}

// PreorderTraversal - Traverse the given node and running slice of values
func PreorderTraversal(root *Node, ch chan string) {
	if root == nil {
		return
	}

	ch <- fmt.Sprintf("%v", root.value)
	PreorderTraversal(root.left, ch)
	PreorderTraversal(root.right, ch)
}

// INORDER BST TRAVERSAL

// PrintInorder - Print the inorder traversal of the tree
func (bst *BinarySearchTree) PrintInorder() {
	inorderChannel := make(chan string)
	inorderBuffer := new(bytes.Buffer)

	go bst.Inorder(inorderChannel)

	for {
		val, i := <-inorderChannel
		if !i {
			break
		}

		inorderBuffer.WriteString(val)
		inorderBuffer.WriteString(" ")
	}

	fmt.Println("Inorder Traversal - ", inorderBuffer.String())
}

// Inorder - Print the inorder traversal of the Binary Search Tree
func (bst *BinarySearchTree) Inorder(ch chan string) {
	InorderTraversal(bst.root, ch)
	close(ch)
}

// InorderTraversal - Traverse the given node and running slice of values
func InorderTraversal(root *Node, ch chan string) {
	if root == nil {
		return
	}

	InorderTraversal(root.left, ch)
	ch <- fmt.Sprintf("%v", root.value)
	InorderTraversal(root.right, ch)
}

// POSTORDER BST TRAVERSAL

// PrintPostorder - Print the postorder traversal of the tree
func (bst *BinarySearchTree) PrintPostorder() {
	postorderChannel := make(chan string)
	postorderBuffer := new(bytes.Buffer)

	go bst.Postorder(postorderChannel)

	for {
		val, i := <-postorderChannel
		if !i {
			break
		}

		postorderBuffer.WriteString(val)
		postorderBuffer.WriteString(" ")
	}

	fmt.Println("Postorder Traversal - ", postorderBuffer.String())
}

// Postorder - Print the postorder traversal of the Binary Search Tree
func (bst *BinarySearchTree) Postorder(ch chan string) {
	PostorderTraversal(bst.root, ch)
	close(ch)
}

// PostorderTraversal - Traverse the given node and running slice of values
func PostorderTraversal(root *Node, ch chan string) {
	if root == nil {
		return
	}

	PostorderTraversal(root.left, ch)
	PostorderTraversal(root.right, ch)
	ch <- fmt.Sprintf("%v", root.value)
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

	tree.PrintPreorder()
	tree.PrintInorder()
	tree.PrintPostorder()
}
