package main

import (
	"fmt"
	)

func hanoi(n, src, dst, tmp int) {
	if n > 0 {
		hanoi(n-1, src, dst, tmp)
		move_disk(n, src, dst, tmp)
		hanoi(n-1, src, dst, tmp)
	}
}

func main() {
	
}