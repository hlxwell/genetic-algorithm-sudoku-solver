package main

import (
	"math/rand"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

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
	rand.Seed(time.Now().UnixNano())
	result := []string{}
	for len(result) < GenomeSize {
		result = append(result, Dictionary...)
	}
	rand.Shuffle(GenomeSize, func(i, j int) { result[i], result[j] = result[j], result[i] })
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
		newPopulation = append(newPopulation, zippingCrossover(c.SelectParent(), c.SelectParent()))
	}
	c.Population = newPopulation
}

// https://en.wikipedia.org/wiki/Mutation_(genetic_algorithm)
func (c *Chromosome) SwapMutate() {
	rand.Seed(time.Now().UnixNano())

	for _, p := range c.Population {
		// mutate MutateCount elements.
		for i := 0; i < MutateCount; i++ {
			rand1 := rand.Intn(len(p.Matrix))
			rand2 := rand.Intn(len(p.Matrix))

			// Swap 2 random elements.
			if rand1 != rand2 {
				tmp := p.Matrix[rand1]
				p.Matrix[rand1] = p.Matrix[rand2]
				p.Matrix[rand2] = tmp
			}
		}
	}
}

func (c *Chromosome) RandomMutate() {
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
	rowSize := father.EdgeLength
	newGenome := []string{}
	for i := 0; i < len(father.Matrix); i += rowSize {
		s, e := i, i+rowSize
		if (i/rowSize)%2 == 0 {
			newGenome = append(newGenome, father.Matrix[s:e]...)
		} else {
			newGenome = append(newGenome, mother.Matrix[s:e]...)
		}
	}
	return NewSudoku(newGenome)
}
