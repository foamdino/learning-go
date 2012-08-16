package main

import "fmt"

type node struct {
	data string
	left *node
	right *node
}

func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Printf("data: %s\n", n.data)
		inorder(n.right)
	}
}

func preorder(n *node) {
	if n != nil {
		fmt.Printf("data: %s\n", n.data)
		preorder(n.left)
		preorder(n.right)
	}
}

func postorder(n *node) {
	if n != nil {		
		postorder(n.left)
		postorder(n.right)
		fmt.Printf("data: %s\n", n.data)
	}
}

func add(n *node, data string) {
	if data < n.data {
		if n.left == nil {
			n.left = newNode(data)
		} else {
			add(n.left, data)
		}
	} else if data > n.data {
		if n.right == nil {
			n.right = newNode(data)
		} else {
			add(n.right, data)
		}
	}
}

func remove(n *node, data string) *node {
	if n == nil {
		fmt.Printf("%s not found", data)
		return nil
	}

	if data < n.data {
		n.left = remove(n.left, data)
	} else if data > n.data {
		n.right = remove(n.right, data)
	} else if n.left != nil && n.right != nil {
		if min := findmin(n.right); min != nil {
			n.data = min.data	
		}
		n.right = remove(n.right, data)
	} else {
		if n.left != nil {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

func findmin(n *node) *node {
	if n == nil {
		return nil
	} else if n.left == nil {
		return n
	}

	return findmin(n.left)
}

func newNode(data string) *node {
	n := &node{data, nil, nil}
	return n
}

func main() {
	var root = newNode("kev")

	add(root, "test")
	add(root, "test2")
	add(root, "abc")
	add(root, "def")

	fmt.Printf("inorder\n")
	inorder(root)
	fmt.Printf("\n")

	fmt.Printf("preorder\n")
	preorder(root)
	fmt.Printf("\n")
	
	fmt.Printf("postorder\n")
	postorder(root)
	fmt.Printf("\n")

	fmt.Printf("remove 'test'\n")
	remove(root, "test")
	inorder(root)
	fmt.Printf("\n")	

}