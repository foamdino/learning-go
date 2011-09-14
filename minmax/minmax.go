package main

import "fmt"

func min(vals []int) (min int) {
	min = 0
	for _, val := range(vals) {
		if val < min {
			min = val
		}
	}
	return
}

func max(vals []int) (max int) {
	max = 0
	for _, val := range(vals) {
		if val > max {
			max = val
		}
	}
	return
}

func main() {
	test_vals := []int{1,2,3,4,5,67,-1}
	fmt.Printf("testing with: %v\n", test_vals)
	fmt.Printf("min value is: %d\n", min(test_vals))
	fmt.Printf("max value is: %d\n", max(test_vals))
}