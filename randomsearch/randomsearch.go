package main

import (
	"fmt"
	"math"
	"math/rand"
	)

type vector struct {
	x float32
	y float32
}

func objective(x vector) (result float64) {
	result = float64((x.x * x.x) + (x.y * x.y))

	return 

}

func randomVector(space []vector) (vector) {
	
	for i, _ := range space {
		space[i].x = (space[i].y - space[i].x) * rand.Float32()
		space[i].y = (space[i].x - space[i].y) * rand.Float32()
	}

	sample := space[rand.Intn(len(space))]

	return sample
}

func search(searchSpace []vector, maxIter int) (vector, float64){

	var bestVector = vector{0.0,0.0}
	var bestCost = math.MaxFloat32

	for i := 0; i<maxIter; i++ {
	
		var candidate = randomVector(searchSpace)
		var candidateCost = objective(candidate)

		if candidateCost < bestCost {
			bestVector = candidate
			bestCost = candidateCost
		}

		fmt.Printf("> iteration %d, best: %v cost: %.2f\n", i, bestVector, bestCost)

	}

	return bestVector, bestCost
}

func makeSpace(size int) ([]vector) {

	var space = make([]vector, size)

	for i := 0; i<size; i++ {
		space[i] = vector{-5.0, 5.0}
	}

	return space
}

func main() {
	fmt.Printf("searching...\n")
	problemSize := 2
	searchSpace := makeSpace(problemSize)
	maxIterations := 100

	var bestVector, bestCost = search(searchSpace, maxIterations)

	fmt.Printf("Best solution: {x:%.2f, y:%.2f}, cost: %.2f\n", bestVector.x, bestVector.y, bestCost)
}