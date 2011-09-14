package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibterms(n int) []int {
	result := make([]int, n)
	for i := 1; i<=n; i++ {
		result[i-1] = fib(i)
	}
	return result
}

func main() {
	fmt.Printf("fib(1): %d\n", fib(1))
	fmt.Printf("fib(2): %d\n", fib(2))
	fmt.Printf("fib(3): %d\n", fib(3))
	fmt.Printf("fib(4): %d\n", fib(4))
	fmt.Printf("fibterms(6): %v\n", fibterms(6))
}