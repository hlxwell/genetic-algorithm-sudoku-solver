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

// const SelectionSize = 10

var Dictionary = []string{"W", "O", "R", "D"}

type Chromosome struct {
	CurrentBestSolution *Sudoku
	Generation          int
	Population          []*Sudoku
}

func NewChromosome(generation int) *Chromosome {
	return &Chromosome{
		Generation: generation,
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
	return &Sudoku{Matrix: result}
}

// To evaludate the fitness of a genome.
// sudoku fitness = solved rows + solved columes + solved boxes
func (c Chromosome) Fitness(s *Sudoku) int {
	return s.ValidSolutionCount()
}

// Crossover: only crossover the last CrossoverCount elements.
func (c *Chromosome) Crossover() {
	var newGenome []string
	newPopulation := []*Sudoku{}

	for _, p := range c.Population {
		newGenome = []string{}
		newGenome = append(c.CurrentBestSolution.Matrix[:GenomeSize-CrossoverCount], p.Matrix[GenomeSize-CrossoverCount:]...)
		newPopulation = append(newPopulation, &Sudoku{Matrix: newGenome})
	}

	c.Population = newPopulation
}

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

func (c *Chromosome) Select() {
	bestSolution := c.Population[0]
	bestScore := c.Fitness(bestSolution)

	for _, p := range c.Population {
		score := c.Fitness(p)
		if score > bestScore {
			bestSolution = p
			bestScore = score
		}
	}

	c.CurrentBestSolution = bestSolution
}

func (c *Chromosome) Evolve() {
	for {
		fmt.Println("--- Generation: ", c.Generation)
		c.Select()    // Select the best solution.
		c.Crossover() // Best solution crossover with all elements.
		c.Mutate()    // Random Mutate several elements.
		c.Select()    // Select the best solution.
		c.Generation++

		if c.CurrentBestSolution.ValidSolutionCount() == 12 {
			break
		} else {
			fmt.Println(c.CurrentBestSolution.Matrix, c.CurrentBestSolution.ValidSolutionCount())
		}
	}
}
