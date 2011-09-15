package main

import "fmt"

func plusTwo() (func(v int) (int)) {
	return func(v int) (int) {
		return v+2
	}
}

func plusX(x int) (func(v int) (int)) {
	return func(v int) (int) {
		return v+x
	}
}

func main() {
	p := plusTwo()
	fmt.Printf("3+2: %d\n", p(3))

	px := plusX(3)
	fmt.Printf("3+3: %d\n", px(3))
}