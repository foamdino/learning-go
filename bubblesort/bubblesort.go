package main

import "fmt"

func bubblesort(vals []int) []int {
	swapped := true
	for swapped == true {
		//reset swapped
		swapped = false
		for i,j := 0, 1; j<len(vals); i, j = i+1, j+1 {
			if vals[i] > vals[j] {
				vals[i], vals[j] = vals[j], vals[i]
				swapped = true
			}
		} 
	}
	return vals
}

func main() {
	vals := []int{1,7,2,4,9,0,34,22,12,45,9}
	fmt.Printf("original: %v\n", vals)
	fmt.Printf("sorted: %v\n", bubblesort(vals))
}