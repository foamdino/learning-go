package main

import "fmt"

type node struct {
	data string
	left node
	right node
}

type root node

func inorder(root node) {
	if root != nil {
		inorder(root.left)
		fmt.Printf(root.data)
		inorder(root.right)
	}
}

func main() {
	
}