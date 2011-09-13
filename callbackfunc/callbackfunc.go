package main

import "fmt"

func printit(x int) {
	fmt.Printf("x = %v\n", x)
}

func callback(y int, f func(int)) {
	f(y)
}

func main() {
	callback(2, printit)
}