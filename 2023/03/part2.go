package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Star struct {
	Position []int
	Numbers  []int
}

var stars []Star

func isNum(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func isStar(char byte) bool {
	// '*' character is utf-8 character code 42
	if char == 42 {
		return true
	}
	return false
}

func hasAdjacentRow(row []byte, start, end int) int {
	// If start is any but the first position, look in the previous grid space
	if start != 0 {
		if isStar(row[start-1]) {
			// fmt.Printf("Found an adjacent gear in previous column\n")
			return start - 1
		}
	}
	// Look between start and end positions
	for i := start; i <= end; i++ {
		if isStar(row[i]) {
			// fmt.Printf("Found an adjacent gear in current column\n")
			return i
		}
	}
	// If end is any but the last position, look in the next grid space
	if end != len(row)-1 {
		if isStar(row[end+1]) {
			// fmt.Printf("Found an adjacent gear in next column\n")
			return end + 1
		}
	}
	return -1
}

func hasAdjacent(grid [][]byte, start, end, row int) (int, int) {
	// If the row is not the first row, check for adjacency in the previous row
	var y int = -1
	if row != 0 {
		y = hasAdjacentRow(grid[row-1], start, end)
		if y != -1 {
			// fmt.Printf("Found an adjacent gear in previous row\n")
			return row - 1, y
		}
	}

	// If the row is not the last row, check for adjacency in the next row
	if row != len(grid)-1 {
		y = hasAdjacentRow(grid[row+1], start, end)
		if y != -1 {
			// fmt.Printf("Found an adjacent gear in next row\n")
			return row + 1, y
		}
	}

	// Check if the current row contains adjacenct symbols
	y = hasAdjacentRow(grid[row], start, end)
	if y != -1 {
		// fmt.Printf("Found an adjacent gear in current row\n")
		return row, y
	}

	return -1, -1
}

func part2(stars []Star) int {
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
				// fmt.Printf("Checking for adjacency of number %s\n", n_s)
				// See if the number we found has an adjacent gear
				x, y := hasAdjacent(grid, start_pos, end_pos, i)
				if x != -1 && y != -1 {
					n_s := ""
					for k := start_pos; k <= end_pos; k++ {
						n_s += string(grid[i][k])
					}
					v, _ := strconv.Atoi(n_s)
					fmt.Printf("Number found with adjacent gear in position [%d, %d]: %d\n", x, y, v)
					// ans += v

					// There is an adjacent star. If it already exists, add the number to the list
					// If it doesn't exist, create and append to list
					for i, star := range stars {
						if star.Position[0] == x && star.Position[1] == y {
                            stars[i].Numbers = append(stars[i].Numbers, v)
						}
						if i == len(stars)-1 {
							stars = append(stars, Star{
								[]int{x, y},
								[]int{v},
							})
						}
					}
					if len(stars) == 0 {
						stars = append(stars, Star{
							[]int{x, y},
							[]int{v},
						})
					}
					fmt.Printf("Current star list : %v\n", stars)

				} else {
					// fmt.Printf("Did not find an adjacent gear for number: %s\n", n_s)
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
				// fmt.Printf("Checking for adjacency of number %s\n", n_s)
				// See if the number we found has an adjacent gear
				x, y := hasAdjacent(grid, start_pos, end_pos, i)
				if x != -1 && y != -1 {
					n_s := ""
					for k := start_pos; k <= end_pos; k++ {
						n_s += string(grid[i][k])
					}
					v, _ := strconv.Atoi(n_s)
					fmt.Printf("Number found with adjacent gear in position [%d, %d]: %d\n", x, y, v)
					// ans += v

					// There is an adjacent star. If it already exists, add the number to the list
					// If it doesn't exist, create and append to list
					for i, star := range stars {
						fmt.Printf("Star[%d] position[0] = %d\n", i, star.Position[0])
						fmt.Printf("Star[%d] position[1] = %d\n", i, star.Position[1])
						if star.Position[0] == x && star.Position[1] == y {
                            stars[i].Numbers = append(stars[i].Numbers, v)
						} else if i == len(stars)-1 {
							stars = append(stars, Star{
								[]int{x, y},
								[]int{v},
							})
						}
					}
					if len(stars) == 0 {
						stars = append(stars, Star{
							[]int{x, y},
							[]int{v},
						})
					}
					fmt.Printf("Current star list : %v\n", stars)

				} else {
					// fmt.Printf("Did not find an adjacent gear for number: %s\n", n_s)
				}
				start_pos, end_pos = -1, -1
			}
		}
	}

	// Stars list has been populated.
	// Look through all stars and multiply the number list for stars that have only two numbers in the list
	// Add the product to the running total
	var product int = 1
	fmt.Println(stars)
	for i := range stars {
		if len(stars[i].Numbers) == 2 {
			for _, p := range stars[i].Numbers {
			    product *= p
			}
			ans += product
            product = 1
		}
	}

	return ans
}

func main() {
	stars = make([]Star, 0)
	start := time.Now()

	fmt.Println("Part 2 Solution:", part2(stars))
	p1_time := time.Since(start)
	fmt.Println("Part 2 took", p1_time)
}
