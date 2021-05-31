package main

import (
	"math/rand"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

const GenomeSize = 16
const CrossoverCount = 4
const MutateCount = 3
const PopulationSize = 100
const TotalValidSolutionCount = 12 // 12 for 4x4, 27 for 9x9
const SelectionRate = 0.1

// For 4x4
var Dictionary = []string{"W", "O", "R", "D"}

// For 9x9
// var Dictionary = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

type Chromosome struct {
	// CurrentBestSolution *Sudoku
	Generation int
	Population []*Sudoku
}

func NewChromosome() *Chromosome {
	return &Chromosome{}
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
	newPopulation := []*Sudoku{}
	for i := 0; i < PopulationSize; i++ {
		newPopulation = append(newPopulation, tailCrossover(c.SelectParent(), c.SelectParent()))
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
// Pick TOP 1/10 fitness parents
// https://en.wikipedia.org/wiki/Selection_(genetic_algorithm)
func (c *Chromosome) SelectParent() *Sudoku {
	selected1 := rand.Intn(PopulationSize * SelectionRate)
	selected2 := rand.Intn(PopulationSize * SelectionRate)
	if c.Fitness(c.Population[selected1]) >= c.Fitness(c.Population[selected2]) {
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

func (c *Chromosome) SortPopulartionByFitness() {
	sort.Slice(c.Population, func(i, j int) bool {
		return c.Population[i].ValidSolutionCount() > c.Population[j].ValidSolutionCount()
	})
}

func (c *Chromosome) PrintPopulationFitness() {
	arr := []int{}
	for _, p := range c.Population {
		arr = append(arr, p.ValidSolutionCount())
	}
	log.Infoln(arr)
}

// =============== Private Methods ===============

// Crossover Solution 1:
func tailCrossover(father *Sudoku, mother *Sudoku) *Sudoku {
	newGenome := []string{}
	newGenome = append(newGenome, father.Matrix[:GenomeSize-CrossoverCount]...)
	newGenome = append(newGenome, mother.Matrix[GenomeSize-CrossoverCount:]...)
	return NewSudoku(newGenome)
}

// Crossover Solution 2: Zipping rows of parents gene.
func zippingCrossover(father *Sudoku, mother *Sudoku) *Sudoku {
	newGenome := []string{}
	newGenome = append(newGenome, father.Matrix[:4]...)
	newGenome = append(newGenome, mother.Matrix[4:8]...)
	newGenome = append(newGenome, father.Matrix[8:12]...)
	newGenome = append(newGenome, mother.Matrix[12:16]...)
	return NewSudoku(newGenome)
}
