package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func isNum(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func isSymbol(char byte) bool {
	if !isNum(char) && char != 46 {
		return true
	}
	return false
}

func hasAdjacentRow(row []byte, start, end int) bool {
	// If start is any but the first position, look in the previous grid space
	if start != 0 {
		if isSymbol(row[start-1]) {
			fmt.Printf("Found an adjacent symbol: %s in previous column\n", string(row[start-1]))
			return true
		}
	}
	// Look between start and end positions
	for i := start; i <= end; i++ {
		if isSymbol(row[i]) {
			fmt.Printf("Found an adjacent symbol: %s in current column\n", string(row[i]))
			return true
		}
	}
	// If end is any but the last position, look in the next grid space
	if end != len(row)-1 {
		if isSymbol(row[end+1]) {
			fmt.Printf("Found an adjacent symbol: %s in next column\n", string(row[end+1]))
			return true
		}
	}
	return false
}

func hasAdjacent(grid [][]byte, start, end, row int) bool {
	// If the row is not the first row, check for adjacency in the previous row
	if row != 0 && hasAdjacentRow(grid[row-1], start, end) {
		fmt.Printf("Found an adjacent symbol in previous row\n")
		return true
	}

	// Check if the current row contains adjacenct symbols
	if hasAdjacentRow(grid[row], start, end) {
		fmt.Printf("Found an adjacent symbol in current row\n")
		return true
	}

	// If the row is not the last row, check for adjacency in the next row
	if row != len(grid)-1 && hasAdjacentRow(grid[row+1], start, end) {
		fmt.Printf("Found an adjacent symbol in next row\n")
		return true
	}

	return false
}

func part1() int {
	var ans int
	var start_pos, end_pos int = -1, -1

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make([][]byte, 0)
	lineScanner := bufio.NewScanner(file)

	// Fill out the grid in memory
	for lineScanner.Scan() {
		chars := []byte(lineScanner.Text())
		grid = append(grid, chars)
	}

	for i, row := range grid {
		fmt.Printf("\nLooking at line %s\n", row)

		// fmt.Printf("Length of line is %d\n", len(row))
		for j, col := range grid[i] {
			// fmt.Printf("j = %d\n", j)
			// fmt.Printf("Looking at value %s\n", string(col))
			_ = col
			// We find a new number, set start and end positions as the same
			if isNum(grid[i][j]) && start_pos == -1 {
				start_pos = j
				end_pos = j
				// fmt.Printf("We found the first num: %d and j = %d\n", grid[i][start_pos]-48, j)
			}

			// We are at the final character, and it's a digit (end of number)
			// Or, we found a non-number character (end of number)
			if j == len(row)-1 && isNum(grid[i][j]) {
				end_pos = j
				// fmt.Printf("We found the last num: %d and j = %d\n", grid[i][end_pos]-48, j)

				n_s := ""
				for k := start_pos; k <= end_pos; k++ {
					n_s += string(grid[i][k])
				}
				fmt.Printf("Checking for adjacency of number %s\n", n_s)
				// See if the number we found has an adjacent symbol
				if hasAdjacent(grid, start_pos, end_pos, i) {
					// If it does, calculate the actual numerical value and add to running total
					n_s := ""
					for k := start_pos; k <= end_pos; k++ {
						n_s += string(grid[i][k])
					}
					v, _ := strconv.Atoi(n_s)
					fmt.Printf("Adding number with adjacency: %d\n", v)
					ans += v
				} else {
					fmt.Printf("Did not find an adjacent symbol for number: %s\n", n_s)
				}
				start_pos, end_pos = -1, -1
			}
			// We found our first number. Set the end position, check for adjacency, and reset the positions
			if !isNum(grid[i][j]) && end_pos != -1 {
				end_pos = j - 1
				// fmt.Printf("We found the last num: %d and j = %d\n", grid[i][end_pos]-48, j)

				n_s := ""
				for k := start_pos; k <= end_pos; k++ {
					n_s += string(grid[i][k])
				}
				fmt.Printf("Checking for adjacency of number %s\n", n_s)
				// See if the number we found has an adjacent symbol
				if hasAdjacent(grid, start_pos, end_pos, i) {
					// If it does, calculate the actual numerical value and add to running total
					n_s := ""
					for k := start_pos; k <= end_pos; k++ {
						n_s += string(grid[i][k])
					}
					v, _ := strconv.Atoi(n_s)
					fmt.Printf("Adding number with adjacency: %d\n", v)
					ans += v
				} else {
					fmt.Printf("Did not find an adjacent symbol for number: %s\n", n_s)
				}
				start_pos, end_pos = -1, -1
			}
		}
	}

	return ans
}

func part2() int {
	var ans int
	return ans
}

func main() {
	start := time.Now()

	fmt.Println("Part 1 Solution:", part1())
	p1_time := time.Since(start)
	fmt.Println("Part 1 took", p1_time)

	fmt.Println("Part 2 Solution:", part2())
	fmt.Println("Part 2 took", time.Since(start)-p1_time)
}
