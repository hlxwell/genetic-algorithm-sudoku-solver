package main

import (
	"fmt"
	"math/rand"
	"time"
)

const GenomeSize = 16
const CrossoverCount = 8
const MutateCount = 3
const PopulationSize = 100
const TotalValidSolutionCount = 12

// const SelectionSize = 10

var Dictionary = []string{"W", "O", "R", "D"}

type Chromosome struct {
	// CurrentBestSolution *Sudoku
	Generation    int
	MaxGeneration int
	Population    []*Sudoku
}

func NewChromosome(generation int) *Chromosome {
	return &Chromosome{
		MaxGeneration: generation,
	}
}

// Generate Population
func (c *Chromosome) GeneratePopulation() {
	result := []*Sudoku{}
	for i := 0; i < PopulationSize; i++ {
		result = append(result, c.GenerateGenome())
	}
	c.Population = result
}

// To generate one random solution
func (c *Chromosome) GenerateGenome() *Sudoku {
	result := make([]string, GenomeSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < GenomeSize; i++ {
		n := rand.Intn(len(Dictionary))
		result[i] = Dictionary[n]
	}
	return NewSudoku(result)
}

// To evaludate the fitness of a genome.
// sudoku fitness = solved rows + solved columes + solved boxes
func (c Chromosome) Fitness(s *Sudoku) int {
	return s.ValidSolutionCount()
}

// Crossover: only crossover the last CrossoverCount elements.
// https://en.wikipedia.org/wiki/Crossover_(genetic_algorithm)
func (c *Chromosome) Crossover() {
	var newGenome []string
	newPopulation := []*Sudoku{}

	for i := 0; i < PopulationSize; i++ {
		father := c.SelectParent()
		mather := c.SelectParent()
		newGenome = []string{}
		newGenome = append(father.Matrix[:GenomeSize-CrossoverCount], mather.Matrix[GenomeSize-CrossoverCount:]...)
		newPopulation = append(newPopulation, NewSudoku(newGenome))
	}

	c.Population = newPopulation
}

// https://en.wikipedia.org/wiki/Mutation_(genetic_algorithm)
func (c *Chromosome) Mutate() {
	rand.Seed(time.Now().UnixNano())

	for _, p := range c.Population {
		// mutate MutateCount elements.
		for i := 0; i < MutateCount; i++ {
			randInMatrix := rand.Intn(len(p.Matrix))
			randInDict := rand.Intn(len(Dictionary))
			// Randomly replace 1 element.
			p.Matrix[randInMatrix] = Dictionary[randInDict]
		}
	}
}

// Use Tournament method to choose parent.
// https://en.wikipedia.org/wiki/Selection_(genetic_algorithm)
func (c *Chromosome) SelectParent() *Sudoku {
	selected1 := rand.Intn(PopulationSize)
	selected2 := rand.Intn(PopulationSize)
	if c.Fitness(c.Population[selected1]) > c.Fitness(c.Population[selected2]) {
		return c.Population[selected1]
	}
	return c.Population[selected2]
}

func (c *Chromosome) Elitism() *Sudoku {
	bestSolution := c.Population[0]
	bestScore := c.Fitness(bestSolution)

	for _, p := range c.Population {
		score := c.Fitness(p)
		if score > bestScore {
			bestSolution = p
			bestScore = score
		}
	}

	return bestSolution
}

func (c *Chromosome) Evolve() {
	for i := 0; i < c.MaxGeneration; i++ {
		fmt.Printf("%dth generation:\n", i)

		c.Crossover()       // Best solution crossover with all elements.
		c.Mutate()          // Random Mutate several elements.
		best := c.Elitism() // Select the best solution.

		if best.ValidSolutionCount() == TotalValidSolutionCount {
			break
		} else {
			fmt.Println(best.Matrix, best.ValidSolutionCount())
		}
	}
}
