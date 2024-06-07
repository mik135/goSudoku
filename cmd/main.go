package main

import (
	"sudoku"
	"fmt"
)

func main() {
	grid := sudoku.NewSudoku()
	err := grid.Set(1, 1, 5)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	for _, row := range grid {
		fmt.Println(grid)
	}
}