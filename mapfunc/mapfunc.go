package main

import "fmt"

func mapper(f func(int) int, l []int) []int {
	for i:=0; i<len(l); i++ {
		l[i] = f(l[i])
	}
	return l
}

func inc(x int) int {
	return x+1
}

func main() {
	test_vals := []int{1,2,3,4,5}
	fmt.Printf("before map: %v\n", test_vals)
	mapper(inc, test_vals)
	fmt.Printf("after map: %v\n", test_vals)
}