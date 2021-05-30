package main

import "fmt"

func main() {
	// sudoku := []string{
	// 	"W", "R", "O", "D",
	// 	"D", "O", "W", "R",
	// 	"O", "D", "R", "W",
	// 	"R", "W", "D", "O",
	// }

	c := NewChromosome(10)
	for i := 0; i < 100; i++ {
		fmt.Println(c.Generate())
	}
}
