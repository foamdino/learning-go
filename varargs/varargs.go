package main

import "fmt"

func anyints(arg ...int) {
	for _, n := range(arg) {
		fmt.Printf("n: %d\n", n)
	}
}

func main() {
	anyints(1,2,3)
	anyints(4)
	anyints(5,6,7,8,9)
}