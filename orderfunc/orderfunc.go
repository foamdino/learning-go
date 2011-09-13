package main

import "fmt"

func increasing(x, y int) (low, high int) {
	if x <= y {
		low, high = x, y
	} else {
		low, high = y, x
	}
	return
}

func main() {
	low, high := increasing(3,2)
	fmt.Printf("low: %d, high: %d\n", low, high)

	low, high = increasing(72, 83)
	fmt.Printf("low: %d, high: %d\n", low, high)
}