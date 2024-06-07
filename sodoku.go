package sudoku


import (
	"errors"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

// Errors that could occur.
var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in this row")
	ErrInColumn   = errors.New("digit already present in this column")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			d := digits[i][j]
			if d != empty {
				grid[i][j].digit = d
				grid[i][j].fixed = true
			}
		}
	}
	return &grid
}

func (g *Grid) Set(row, column, digit) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDigit
	case g.InRow(row, digit):
		return ErrInRow
	case g.InColumn(column, digit):
		return ErrInColumn
	case g.InRegion(row, column, digit):
		return ErrInRegion
	}

	g[row][column].digit = digit
	return nil

}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}
