package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fizz := false
		buzz := false
		
		if i % 3 == 0 {
			fizz = true
		} 

		if i % 5 == 0 {
			buzz = true
		}

		if fizz && !buzz{
			fmt.Println("Fizz")
		} else if buzz && !fizz {
			fmt.Println("Buzz")
		} else if fizz && buzz {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}