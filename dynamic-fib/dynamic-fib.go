package main

import "fmt"

func fib(i int, known map[int]int) int {
	
	if known[i] > 0 {
		return known[i]
	}

	if i == 0 {
		known[i] = 0
		return 0
	}

	if i == 1 {
		known[i] = 1
		return 1
	}

	if i > 1 {
		var t = fib(i-1, known) + fib(i-2, known)
		known[i] = t
		return t
	}
	return 0

}

func main() {
	known := make(map[int]int)
	fmt.Printf("fib(1): %d\n", fib(1, known))
	fmt.Printf("fib(2): %d\n", fib(2, known))
	fmt.Printf("fib(3): %d\n", fib(3, known))
	fmt.Printf("fib(4): %d\n", fib(4, known))
	fmt.Printf("fib(8): %d\n", fib(8, known))
}