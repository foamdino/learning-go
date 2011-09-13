package main

//import "fmt"

func main() {
	println("Should print 1 to 5")
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		println(i)
	}

	println("Now with nested loops, break out of outermost loop")
J: for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J
			}
			println(i)
		}
	}

	println("Using continue to skip a value")
	println("Should print 0 to 9, skipping 5")
	for i:=0; i< 10; i++ {
		if i == 5 {
			continue
		}
		println(i)
	}
}