package main

import (
	"fmt"
	"math"
	"rand"
)

func randomPlusMinus() (val int) {
	switch rand.Int31n(3) {
	case 0:
		val = 0
	case 1:
		val = 1
	case 2:
		val = -1
	}
	return
}


func mutate(src string) (mutated string) {
	charpos := rand.Intn(len(src))
	temp := []byte(src)
	temp[charpos] = uint8(int(src[charpos]) + (rand.Intn(3) -1))
	mutated = string(temp)
	return
}

func fitness(src string, target string) (fitval float64) {
	fitval = 0
	for pos, _ := range(src) {
		fitval += math.Pow((float64(target[pos]) - float64(src[pos])), float64(2))
	}
	
	return
}

func main() {
	var source = "jiKnp4bqpmAbp"
	var target = "Hello, World!"

	//fmt.Printf("fitness: %f\n", fitness("Hello", "Hfllo"))
	//fmt.Printf("original: %s\tmutated: %s\n", "hello", mutate("hello"))

	fitval := int(fitness(source, target))
	i := 0
	for fitval > 0 {
		i += 1
		m := mutate(source)
		var fitval_m = int(fitness(m, target))
		if fitval_m < fitval {
			fitval = fitval_m
			source = m
			fmt.Printf("generation: %d %s[%d]\n", i, m, fitval)
		}
	}
}