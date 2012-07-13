package main

import (
	"fmt"
)

func next() <- chan int {
	c := make(chan int)
	go func() {
		for i := 0; ; i++ {
			c <- i
		}
	}()
	return c
}

func main() {
	c := next()

	for i := 0; i < 100; i++ {
		fmt.Println("val: ", <-c)
	}
}
