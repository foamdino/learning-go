package main

import "fmt"

func gcd(a int, b int) (int) {
	if b == 0 {
		return a
	} 
	return gcd(b, a % b)
	
	
}

func gcdSubtraction(a int, b int) (int) {
	if a == 0 {
		return b
	}

	for b != 0 {
		if a > b {
			a = a - b
		} else {
			b = b -a
		}
	}
	return a
}

func main() {
	fmt.Printf("gcd(1989, 867): %d\n", gcd(1989,867))
	fmt.Printf("gcdSubtraction(1989, 867): %d\n", gcdSubtraction(1989,867))
}