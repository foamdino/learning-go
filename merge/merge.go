package main

import "fmt"

func mergeSort(m []int) (result []int) {
	
	if len(m) <= 1 {
		return m
	}

	var middle = int(len(m) / 2)
	left := []int{}
	right := []int{}

	for pos, v := range(m) {
		if pos < middle {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = mergeSort(left)
	right = mergeSort(right)
	result = merge(left, right)
	return
}


func merge(left []int, right []int) (result []int) {
	var temp = []int{}
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				t := left[0:1]
				temp = append(temp, t...)
				left = left[1:]
			} else {
				t := right[0:1]
				temp = append(temp, t...)
				right = right[1:]
			}
		} else if len(left) > 0 {
			t := left[0:1]
			temp = append(temp, t...)
			left = left[1:]
		} else if len(right) > 0 {
			t := right[0:1]
			temp = append(temp, t...)
			right = right[1:]
		}
	}
	result = temp
	return
}


func main() {
	unsorted :=[]int{6,1,24,90,2,4,3,7,88,199}
	fmt.Printf("%v\n", unsorted)
	sorted := mergeSort(unsorted)
	fmt.Printf("%v\n", sorted)
}