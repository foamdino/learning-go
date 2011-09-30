package main

import (
	"fmt"
	"math"
	"rand"
	"sort"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 abcdefghijklmnopqrstuvwxyz~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"

type creature struct {
	dna string
	fitness int
}

type creatures []creature

func (c creatures) Len() int { return len(c) }
func (c creatures) Less(i, j int) bool {return c[i].fitness < c[j].fitness }
func (c creatures) Swap(i,j int) { c[i], c[j] = c[j], c[i] }

func fitness(src string, target string) (fitval float64) {
	fitval = 0
	for pos, _ := range(src) {
		fitval += math.Pow((float64(target[pos]) - float64(src[pos])), float64(2))
	}
	
	return
}

func createDNA(target string) (dna string) {
	// creates a random string of the same length as the target
	// create our dna
	t := make([]byte, len(target))
	
	for i := 0; i < len(target); i++ {
		//pick a random char from chars and stuff it in the current pos in dna
		t[i] = uint8(chars[int(rand.Intn(len(chars)))])
	}
	dna = string(t)
	return
}

func mutate2(parent1 creature, parent2 creature, target string) (mutated creature) {
	child_dna := []byte(parent1.dna)

	start := rand.Intn(len(parent2.dna))
	stop := rand.Intn(len(parent2.dna))
	if start > stop { //swap them around
		stop, start = start, stop
	}

	// mix parent2 dna into parent1 dna
	for i := start; i <= stop; i++ {
		child_dna[i] = parent2.dna[i]
	}

	// now mutate a position
	charpos := rand.Intn(int(len(child_dna)))
	child_dna[charpos] = uint8(int(child_dna[charpos]) + (rand.Intn(3) -1))
	mutated = creature{string(child_dna), int(fitness(string(child_dna), target))}
	return
}

func randomParent(genepool creatures) (parent creature) {
	wRndNr := int(rand.Float64() * rand.Float64() * float64(len(genepool)))
	parent = genepool[wRndNr]
	return
}

func main() {
	var target = "Hello, World!"

	var gensize = 20
	genepool := make(creatures, gensize)
	
	for i := 0; i < gensize; i ++ {
		var dna = createDNA(target)
		critter := creature{dna, int(fitness(dna, target))}
		genepool[i] = critter
	}

	fmt.Printf("\nunsorted critters\n\n")
	for i:=0; i<len(genepool); i++ {
		fmt.Printf("dna: %s\t fitness: %d\n", genepool[i].dna, genepool[i].fitness)
	}

	generation := 0
	//loop while the fittest individual is not perfect
	for genepool[0].fitness > 0 {
		// sort the pool of individuals
		sort.Sort(genepool)

		// select 2 parents
		parent1 := randomParent(genepool)
		parent2 := randomParent(genepool)

		child := mutate2(parent1, parent2, target)
		// replace the worst individual
		if child.fitness < genepool[len(genepool)-1].fitness { 
			genepool[len(genepool)-1] = child
		}

		generation++

		fmt.Printf("\nsorted with fittest individuals first\n\n")
		for i:=0; i<len(genepool); i++ {
			fmt.Printf("%d\t dna: %s\t fitness: %d\n", generation, genepool[i].dna, genepool[i].fitness)
		}
		
	}

}