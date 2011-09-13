package main

import "fmt"

func main() {
	nums := []float64{1.0,2.0,3.0,4.0,5.0}

	sum := 0.0
	for _, n := range nums {
		sum += n
	}

	fmt.Println(sum/float64(len(nums)))
	
}