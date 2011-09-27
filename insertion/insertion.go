package main

import "fmt"

func main() {

	unsorted := []int{89, 4, 1, 0, 1, 23, 6, 67, 12, 33}
	fmt.Printf("unsorted: %v\n", unsorted)

	for i := 0; i < len(unsorted); i++ {
		j := i
		for j > 0 && unsorted[j] < unsorted[j-1] {
			unsorted[j], unsorted[j-1] = unsorted[j-1], unsorted[j]
			j--
		}

	}
	fmt.Printf("unsorted: %v\n", unsorted)
}
	