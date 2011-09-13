package main

import "fmt"

func main() {
	fmt.Println("Using for loop")
	for i := 0; i <= 9; i++ {
		fmt.Printf("i: %d ", i)
	}
	fmt.Println()

	fmt.Println("Using goto")

	i := 0
loop:
	fmt.Printf("i: %d ", i)
	i++
	if i <= 9 {
		goto loop
	}
	fmt.Println()

	fmt.Println("Filling an array")
	var a [10]int
	for i := 0; i <= 9; i++ {
		a[i] = i
	}

	fmt.Println(a)

	fmt.Println("Strings now...")
	for i := 1; i<=100; i++ {
		for j := 0; j <=i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}