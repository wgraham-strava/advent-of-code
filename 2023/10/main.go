package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"time"
)

func findNextPiece(grid []string, prevX, prevY, x, y int) (int, int) {
	if string(grid[x][y]) == "S" && prevX != -1 && prevY != -1 {
		return steps, steps / 2
	}
	steps++
	var ss string

	var x1, x2, y1, y2 int = -1, -1, -1, -1
	if string(grid[x][y]) == "S" {
		ss = startSymbol(grid, x, y)
	}
	if string(grid[x][y]) == "|" || ss == "|" {
		x1 = x + 1
		x2 = x - 1
		y1 = y
		y2 = y
	} else if string(grid[x][y]) == "-" || ss == "-" {
		x1 = x
		x2 = x
		y1 = y + 1
		y2 = y - 1
	} else if string(grid[x][y]) == "F" || ss == "F" {
		x1 = x + 1
		x2 = x
		y1 = y
		y2 = y + 1
	} else if string(grid[x][y]) == "L" || ss == "L" {
		x1 = x
		x2 = x - 1
		y1 = y + 1
		y2 = y
	} else if string(grid[x][y]) == "J" || ss == "J" {
		x1 = x
		x2 = x - 1
		y1 = y - 1
		y2 = y
	} else if string(grid[x][y]) == "7" || ss == "7" {
		x1 = x + 1
		x2 = x
		y1 = y
		y2 = y - 1
	}

	if prevX == -1 && prevY == -1 {
		return findNextPiece(grid, startX, startY, x1, y1)
	}

	var p, v int
	if x1 == prevX && y1 == prevY {
		prevX, prevY = x, y
		p, v = findNextPiece(grid, prevX, prevY, x2, y2)
	} else {
		prevX, prevY = x, y
		p, v = findNextPiece(grid, prevX, prevY, x1, y1)
	}
	return p, v
}

func startSymbol(grid []string, x, y int) string {
	var startSymbol string = ""

	if strings.ContainsAny(string(grid[x-1][y]), "|7F") && strings.ContainsAny(string(grid[x+1][y]), "|JL") {
		startSymbol = "|"
	} else if strings.ContainsAny(string(grid[x-1][y]), "|7F") && strings.ContainsAny(string(grid[x][y-1]), "-LF") {
		startSymbol = "J"
	} else if strings.ContainsAny(string(grid[x-1][y]), "|7F") && strings.ContainsAny(string(grid[x][y+1]), "-7J") {
		startSymbol = "L"
	} else if strings.ContainsAny(string(grid[x+1][y]), "|JL") && strings.ContainsAny(string(grid[x][y-1]), "-LF") {
		startSymbol = "7"
	} else if strings.ContainsAny(string(grid[x+1][y]), "|JL") && strings.ContainsAny(string(grid[x][y+1]), "-7J") {
		startSymbol = "F"
	} else if strings.ContainsAny(string(grid[x][y-1]), "-FL") && strings.ContainsAny(string(grid[x][y-1]), "-7J") {
		startSymbol = "-"
	} else if strings.ContainsAny(string(grid[x][y-1]), "-FL") && strings.ContainsAny(string(grid[x-1][y]), "|7F") {
		startSymbol = "J"
	} else if strings.ContainsAny(string(grid[x][y-1]), "-FL") && strings.ContainsAny(string(grid[x+1][y]), "|JL") {
		startSymbol = "7"
	} else if strings.ContainsAny(string(grid[x][y+1]), "-7J") && strings.ContainsAny(string(grid[x-1][y]), "|7F") {
		startSymbol = "L"
	} else if strings.ContainsAny(string(grid[x][y+1]), "-7J") && strings.ContainsAny(string(grid[x-1][y]), "|JL") {
		startSymbol = "F"
	} else if strings.ContainsAny(string(grid[x][y+1]), "-7J") && strings.ContainsAny(string(grid[x][y-1]), "-FL") {
		startSymbol = "-"
	} else if strings.ContainsAny(string(grid[x-1][y]), "|7F") && strings.ContainsAny(string(grid[x+1][y]), "|JL") {
		startSymbol = "|"
	}
	return startSymbol
}

func findStartPosition(grid []string) (int, int) {
	var x, y int = 0, 0
	for i := range grid {
		for j := range grid[i] {
			if string(grid[i][j]) == "S" {
				x = i
				y = j
			}
		}
	}
	return x, y
}

var startX, startY, steps int

func part1() int {
	var looper int = 0
	var grid []string = make([]string, 0)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		looper++
	}
	startX, startY = findStartPosition(grid)
	startSymbol := startSymbol(grid, startX, startY)
	fmt.Printf("Puzzle starts at position grid[%d][%d] with pipe piece %v\n", startX, startY, startSymbol)

	_, halfway := findNextPiece(grid, -1, -1, startX, startY)

	return halfway
}

func main() {
	t0 := time.Now()
	p1 := part1()
	fmt.Printf("Part 1 solution: %d\n", p1)
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
}
