package backend

func SolveSudoku(puzzle [][]int) [][]int {
	solveBacktrack(puzzle)
	return puzzle
}

func solveBacktrack(grid [][]int) bool {
	emptyRow, emptyCol := findEmptyCell(grid)
	if emptyRow == -1 && emptyCol == -1 {
		// no empty cells left
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValidMove(grid, emptyRow, emptyCol, num) {
			//num in empty cell
			grid[emptyRow][emptyCol] = num

			// Recur to solve the remaining cell
			if solveBacktrack(grid) {
				return true
			}

			// If placing num not giving solution then backtrack
			grid[emptyRow][emptyCol] = 0
		}
	}

	// If no valid number placed then backtrack
	return false
}

func findEmptyCell(grid [][]int) (int, int) {
	// Find and return the row and column of the first empty cell (with value 0)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // empty cell
}

func isValidMove(grid [][]int, row, col, num int) bool {
	// Check if the number is not already present in the same row, column, or 3x3 subgrid
	return !usedInRow(grid, row, num) && !usedInColumn(grid, col, num) && !usedInBox(grid, row-row%3, col-col%3, num)
}

func usedInRow(grid [][]int, row, num int) bool {
	// Check if the number is already present in the given row
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

func usedInColumn(grid [][]int, col, num int) bool {
	// Check if the number is already present in the given column
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(grid [][]int, startRow, startCol, num int) bool {
	// Check if the number is already present in the 3x3 subgrid starting from startRow, startCol
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return true
			}
		}
	}
	return false
}
