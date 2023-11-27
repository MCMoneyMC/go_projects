package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if areArgumentsValid(args) {
		input := formatArgumentsTo2DSlice(args[1:])
		findSolutions(input)
	} else {
		giveErrorMessage()
	}
}

/*
Sudoku Solver Body
*/

func findSolutions(grid [][]int) {
	spots := findEmptySpots(grid)
	solutions := getSolutions(grid, spots, nil)
	if len(solutions) != 1 {
		giveErrorMessage()
	} else {
		giveSolution(solutions[0])
	}
}

func findEmptySpots(grid [][]int) [][]int {
	var output [][]int
	for i, row := range grid {
		for j, spot := range row {
			if spot == 0 {
				output = append(output, []int{i, j})
			}
		}
	}
	return output
}

func getSolutions(grid [][]int, emptySpots [][]int, currentSolutions [][][]int) [][][]int {
	if len(emptySpots) == 0 {
		dummy := make([][]int, 9)
		copy(dummy, grid)
		return append(currentSolutions, dummy)
	}
	row, col := emptySpots[0][0], emptySpots[0][1]
	for i := 1; i < 10; i++ {
		if isValidPlacement(grid, i, row, col) {
			dummyGrid := copyGrid(grid)
			dummyGrid[row][col] = i
			dummySlice := emptySpots[1:]
			currentSolutions = getSolutions(dummyGrid, dummySlice, currentSolutions)
		}
	}
	return currentSolutions
}

func copyGrid(grid [][]int) [][]int {
	dummy := make([][]int, 9)
	for row := range grid {
		dummy[row] = make([]int, 9)
		copy(dummy[row], grid[row])
	}
	return dummy
}

/*
Printing
*/

func giveErrorMessage() {
	for _, char := range "Error" {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func giveSolution(grid [][]int) {
	for _, row := range grid {
		for i, square := range row {
			z01.PrintRune(rune(square + 48))
			if i < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

/*
Formatting
*/

func formatArgumentsTo2DSlice(args []string) [][]int {
	var grid [][]int
	for _, arg := range args {
		var row []int
		for _, char := range arg {
			if char == '.' {
				row = append(row, 0)
			} else {
				row = append(row, int(char-'0'))
			}
		}
		grid = append(grid, row)
	}
	return grid
}

/*
Validation & Requirements
*/

func areArgumentsValid(args []string) bool {
	if len(args) != 10 {
		return false
	}
	for _, arg := range args[1:] {
		if len(arg) != 9 {
			return false
		}
		for _, char := range arg {
			if !(char == '.' || ('1' <= char && char <= '9')) {
				return false
			}
		}
	}
	return true
}

func getBox(square []int) int {
	row, col := square[0], square[1]
	return (row/3)*3 + col/3
}

func isValidPlacement(grid [][]int, num, row, col int) bool {
	box := getBox([]int{row, col})
	return checkRow(grid, row, num) && checkColumn(grid, col, num) && checkBox(grid, box, num)
}

func checkBox(grid [][]int, box, num int) bool {
	startRow, startCol := (box/3)*3, (box%3)*3
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+startRow][col+startCol] == num {
				return false
			}
		}
	}
	return true
}

func checkRow(grid [][]int, row, num int) bool {
	for col := 0; col < 9; col++ {
		if grid[row][col] == num {
			return false
		}
	}
	return true
}

func checkColumn(grid [][]int, col, num int) bool {
	for row := 0; row < 9; row++ {
		if grid[row][col] == num {
			return false
		}
	}
	return true
}
