package main

import (
	"math"
)

type Sudoku struct {
	Matrix        []string
	EdgeLength    int
	BoxEdgeLength int
}

func NewSudoku(matrix []string) *Sudoku {
	s := &Sudoku{Matrix: matrix}
	s.SetEdgeLength()
	return s
}

func (s *Sudoku) SetEdgeLength() {
	size := len(s.Matrix)
	s.EdgeLength = int(math.Sqrt(float64(size)))
	// It will only support size of 4*4, 9*9, 16*16 ...
	s.BoxEdgeLength = int(math.Sqrt(float64(s.EdgeLength)))
}

func (s *Sudoku) Columns() [][]string {
	result := [][]string{}
	for i := 0; i < s.EdgeLength; i++ {
		tmp := []string{}
		for j := 0; j < s.EdgeLength; j++ {
			tmp = append(tmp, s.Matrix[i+j*4])
		}
		result = append(result, tmp)
	}
	return result
}

func (s *Sudoku) Rows() [][]string {
	result := [][]string{}
	for i := 0; i < s.EdgeLength; i++ {
		tmp := []string{}
		for j := 0; j < s.EdgeLength; j++ {
			tmp = append(tmp, s.Matrix[j+i*4])
		}
		result = append(result, tmp)
	}
	return result
}

func (s *Sudoku) Boxes() [][]string {
	result := [][]string{}
	for i := 0; i < s.EdgeLength; i += s.BoxEdgeLength {
		for j := 0; j < s.EdgeLength; j += s.BoxEdgeLength {
			tmp := []string{}
			for x := 0; x < s.BoxEdgeLength; x++ {
				for y := 0; y < s.BoxEdgeLength; y++ {
					tmp = append(tmp, s.Matrix[(i+x)*4+(j+y)])
				}
			}
			result = append(result, tmp)
		}
	}

	return result
}
