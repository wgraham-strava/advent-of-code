package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

type Point struct {
    x float64
    y float64
}

func checkCol(grid []string, x, y int) bool {
    for i := 0; i < len(grid)-1; i++ {
        if grid[i][y] != 46 { return false }
    }
    return true
}

func shortestPath(a, b Point) int {
    return int(math.Abs(b.x - a.x)+math.Abs(b.y - a.y))
}
func part1() int {
    var ans int
    var emptyRow, emptyCol = true, true
    var rows, cols = make([]int, 0), make([]int, 0)
    var Stars = make([]Point, 0)
    
	contents, _ := os.ReadFile("1.txt")
    grid := strings.Split(string(contents), "\n")

    fmt.Printf("Grid before expansion looks like\n")
    for i := 0; i < len(grid); i++ {
        fmt.Println(grid[i])
    }

    for x := 0; x < len(grid)-1; x++ {
        emptyRow = true

        for y := 0; y < len(grid[x]); y++ {
            emptyCol = true

            if grid[x][y] == 35 { emptyRow = false; emptyCol = false }

            // If column is empty, add this y position to a list of columns to be duped
            if emptyCol && x == 0 && checkCol(grid, x, y) {
                cols = append(cols, y)
            }
        }
        // If row is empty, add this x position to a list of rows to be duped
        if emptyRow {
            rows = append(rows, x)
        }
	}
    fmt.Printf("Rows %v need to be duped\n", rows)
    // fmt.Printf("Columns %v need to be duped\n", cols)

    // Row expansion is working correctly (for test input, at least)
    for i := 0; i < len(rows); i++ {
        grid = slices.Insert(grid, rows[i]+1, fmt.Sprintf("%s", strings.Repeat(".", len(grid[i]))))
    }

    // reverse sort cols, then iterate "forward"
    slices.SortFunc(cols, func(a, b int) int {
        return cmp.Compare(a, b)
    })
    slices.Reverse(cols)

    // Col expansion is working correctly (for test input, at least)
    for i := 0; i < len(grid)-1; i++ {
        // fmt.Printf("Line before insertion\n%v\n", grid[i])
        for j := 0; j < len(cols); j++ {
            grid[i] = grid[i][:cols[j]] + "." + grid[i][cols[j]:]
        }
        // fmt.Printf("Line after insertion\n%v\n", grid[i])
    }

    fmt.Printf("Grid after expansion looks like\n")
    for i := 0; i < len(grid); i++ {
        fmt.Println(grid[i])
    }
    for x := 0; x < len(grid)-1; x++ {
        for y := 0; y < len(grid[x]); y++ {
            if grid[x][y] == 35 { 
                Stars = append(Stars, Point{
                    x: float64(x),
                    y: float64(y),
                })
            }
        }
    }

    combos := combin.Binomial(len(Stars), 2)
    fmt.Printf("Number of possible combinations: %d\n", combos)

    for i := 0; i < len(Stars)-1; i++{
        for j := i+1; j < len(Stars); j++{
        // calculate length from Stars[i] with Stars[j]
        // fmt.Printf("Shortest path between %v and %v is %d\n", Stars[i], Stars[j], shortestPath(Stars[i], Stars[j]))
        ans += shortestPath(Stars[i], Stars[j])
        }
    }

	return ans
}

func main() {
	t0 := time.Now()
	p1 := part1()
	fmt.Printf("Part 1 solution: %d\n", p1)
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
}
