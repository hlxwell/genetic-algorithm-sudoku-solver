package main

import (
	"fmt"
	"testing"
)

var matrix = []string{
	"W", "R", "O", "D",
	"D", "O", "W", "R",
	"O", "D", "R", "W",
	"R", "W", "D", "O",
}

func TestColumns(t *testing.T) {
	s := NewSudoku(matrix)
	fmt.Println(s.Columns())
}

func TestRows(t *testing.T) {
	s := NewSudoku(matrix)
	fmt.Println(s.Rows())
}

func TestBoxes(t *testing.T) {
	s := NewSudoku(matrix)
	fmt.Println(s.Boxes())
}
