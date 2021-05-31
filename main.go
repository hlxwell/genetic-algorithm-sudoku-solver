package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const MaxGeneration = 100000

func main() {
	log.SetLevel(log.DebugLevel)

	c := NewChromosome()
	c.GeneratePopulation()

	for i := 0; i < MaxGeneration; i++ {
		fmt.Printf("%dth generation:\n", i)

		c.SortPopulartionByFitness()
		c.Crossover()       // Best solution crossover with all elements.
		c.RandomMutate()    // Random Mutate several elements.
		best := c.Elitism() // Select the best solution.

		// return if reach the valid result.
		if best.ValidSolutionCount() == TotalValidSolutionCount {
			break
		}

		fmt.Println("Matrix: ", best.Matrix, "Valid Solution: ", best.ValidSolutionCount())
	}

	c.Elitism().PrettyPrint()
}
