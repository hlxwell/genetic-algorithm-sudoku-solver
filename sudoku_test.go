package main

import (
	"testing"
)

var matrix = []string{
	"W", "R", "O", "D",
	"D", "O", "W", "R",
	"O", "D", "R", "W",
	"R", "W", "D", "D",
}

func TestColumns(t *testing.T) {
	s := NewSudoku(matrix)
	if len(s.Columns()) != s.EdgeLength {
		t.Errorf("Columns size should be %d, but was %d", s.EdgeLength, len(s.Columns()))
	}
}

func TestRows(t *testing.T) {
	s := NewSudoku(matrix)
	if len(s.Rows()) != s.EdgeLength {
		t.Errorf("Rows size should be %d, but was %d", s.EdgeLength, len(s.Rows()))
	}
}

func TestBoxes(t *testing.T) {
	s := NewSudoku(matrix)
	if len(s.Rows()) != s.EdgeLength {
		t.Errorf("Boxes count should be %d, but was %d", s.EdgeLength, len(s.Boxes()))
	}
}

func TestValidSolutionCount(t *testing.T) {
	s := NewSudoku(matrix)
	if s.ValidSolutionCount() != 9 {
		t.Errorf("ValidSolutionCount should be %d, but was %d", 9, s.ValidSolutionCount())
	}
}
