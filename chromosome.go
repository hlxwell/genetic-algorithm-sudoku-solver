package main

import (
	"math/rand"
	"time"
)

const GenomeSize = 16
const CrossoverRate = 0.1
const MutateRate = 0.02
const PopulationSize = 100
const SelectionSize = 100

type Chromosome struct {
	Generation int
	Population [][]string
}

func NewChromosome(generation int) *Chromosome {
	return &Chromosome{
		Generation: generation,
	}
}

// To generate random solutions
func (c *Chromosome) Generate() []string {
	result := make([]string, 16)
	words := []string{"W", "O", "R", "D"}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 16; i++ {
		n := rand.Intn(len(words))
		result[i] = words[n]
	}

	return result
}

// To evaludate the fitness of a genome.
// sudoku fitness = solved rows + solved columes + solved boxes
func (c *Chromosome) Fitness(sudoku Sudoku) int {
	sudoku

	return 0
}

// Crossover
func (c *Chromosome) Crossover() {

}

func (c *Chromosome) Mutate() {

}

func (c *Chromosome) Select() {

}

type Sudoku struct {
}

func NewSudoku(matrix []string) *Sudoku {

}

func (s *Sudoku) Columns() {

}

func (s *Sudoku) Rows() {

}

func (s *Sudoku) Boxes() {

}
