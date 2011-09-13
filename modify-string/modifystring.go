package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s) //convert string to byte array
	c[0] = 'c'
	s2 := string(c) //convert byte array to string
	fmt.Printf("%s\n", s2)
}