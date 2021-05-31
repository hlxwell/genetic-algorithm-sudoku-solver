package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const MaxGeneration = 100000 // max generations
const GenomeSize = 16        // 16/81
const CrossoverCount = 4     // Only for tail crossover method.
const MutateCount = 2        //  mutation times for each generation
const PopulationSize = 100   // Populartion size
const SelectionRate = 0.1    // selection rate for the next gen.

var TotalValidSolutionCount int // 12 for 4x4, 27 for 9x9
var Dictionary []string

func main() {
	log.SetLevel(log.DebugLevel)

	if GenomeSize == 81 {
		// For 9x9
		Dictionary = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		TotalValidSolutionCount = 27
	} else {
		// For 4x4
		Dictionary = []string{"W", "O", "R", "D"}
		TotalValidSolutionCount = 12
	}

	c := NewChromosome()
	c.GeneratePopulation()

	for i := 0; i < MaxGeneration; i++ {
		fmt.Printf("%dth generation:\n", i)

		c.SortPopulartionByFitness()
		c.Crossover()       // Best solution crossover with all elements.
		c.SwapMutate()      // Random Mutate several elements.
		best := c.Elitism() // Select the best solution.

		// return if reach the valid result.
		if best.ValidSolutionCount() == TotalValidSolutionCount {
			break
		}

		fmt.Println("Matrix: ", best.Matrix, "Valid Solution: ", best.ValidSolutionCount())
	}

	c.Elitism().PrettyPrint()
}
