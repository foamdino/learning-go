package main

import (
	"fmt"
	"utf8"
)

func main() {
	fmt.Println("Strings now...")
	for i := 1; i<=100; i++ {
		for j := 0; j <=i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}

	fmt.Println("Counting chars")
	input := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("Number of chars in %s: %d\n", input, len(input))
	bytes := []byte(input)
	fmt.Printf("Number of bytes in %s: %d\n", input, len(bytes))

	fmt.Println("Now with unicode...")
	input = "aΦx"
	for pos, char := range input {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}

	var num_bytes = 0
	for _, char := range input {
		num_bytes += utf8.RuneLen(char)
	}
	fmt.Printf("%s contains %d bytes\n", input, num_bytes)

	fmt.Println("Swapping chars...")
	input = "asSASA ddd dsjkdsjs dk"

	var output = ""
	for pos, char := range input {
		switch pos {
		case 3:
			output = output + "a"
		case 4:
			output = output + "b"
		case 5:
			output = output + "c"
		default:
			output = output + string(char)
		}
	}

	fmt.Println(output)

	fmt.Println("Swapping chars...")
	input = "foobar"
	temp := []byte(input)
	
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	fmt.Println(string(temp))
}