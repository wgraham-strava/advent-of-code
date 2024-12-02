// main package
package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type Forest [][]int

func parseInput(f string) Forest {
	file, _ := os.Open(f)
	defer file.Close()
	var trees Forest

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanBytes)
	var row []int
	for scanner.Scan() {
		r := scanner.Text()
		if r == "\n" {
			trees = append(trees, row)
			row = nil
			continue
		}
		tree, _ := strconv.Atoi(r)
		row = append(row, tree)
	}
	return trees
}

func part1(f string) int {
	trees := parseInput(f)
	var ans int

	// Only check visibility for "inner" trees
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			if trees.isVisible(i, j) {
				ans++
			}
		}
	}

	// Add double the length and width, and subtract 4 for overlapping corners
	ans += len(trees) * 2
	ans += len(trees[0]) * 2
	ans -= 4
	return ans
}

func part2(f string) int {
	trees := parseInput(f)
	var ans int

	// Only check "inner" trees
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			score := trees.calculateScenicScore(i, j)
			if score > ans {
				ans = score
			}
		}
	}

	return ans
}

func (f Forest) isVisible(i, j int) bool {
	// Check "west"-side visibility
	if !slices.Contains(f[i][:j], f[i][j]) && slices.Max(f[i][:j+1]) == f[i][j] {
		return true
	}
	// Check "east"-side visibility
	if !slices.Contains(f[i][j+1:], f[i][j]) && slices.Max(f[i][j:]) == f[i][j] {
		return true
	}
	// Check "north"-side visibility
	for k := i - 1; k >= 0; k-- {
		if f[k][j] >= f[i][j] {
			break
		} else if k == 0 {
			return true
		}
	}
	// Check "south"-side visibility
	for k := i + 1; k < len(f); k++ {
		if f[k][j] >= f[i][j] {
			break
		} else if k == len(f)-1 {
			return true
		}
	}
	return false
}

func (f Forest) calculateScenicScore(i, j int) int {
	var score int
	var up, left, down, right int

	// Check "north"-side visibility
	for k := i - 1; k >= 0; k-- {
		if f[k][j] >= f[i][j] || k == 0 {
			up++
			break
		} else {
			up++
		}
	}

	// Check "west"-side visibility
	for k := j - 1; k >= 0; k-- {
		if f[i][k] >= f[i][j] || k == 0 {
			left++
			break
		} else {
			left++
		}
	}

	// Check "south"-side visibility
	for k := i + 1; k < len(f); k++ {
		if f[k][j] >= f[i][j] || k == len(f)-1 {
			down++
			break
		} else {
			down++
		}
	}

	// Check "east"-side visibility
	for k := j + 1; k < len(f[i]); k++ {
		if f[i][k] >= f[i][j] || k == len(f[i])-1 {
			right++
			break
		} else {
			right++
		}
	}
	score = up * left * down * right
	return score
}

func main() {
	ta := time.Now()
	log.Printf("Test input solution: %d\n", part2("input.txt"))
	tb := time.Now()
	log.Printf("Test input took: %v\n", tb.Sub(ta))
	log.Printf("Part 1 solution: %d\n", part1("input.txt"))
	ta = time.Now()
	log.Printf("Part 1 took: %v\n", ta.Sub(tb))
	log.Printf("Part 2 solution: %d\n", part2("input.txt"))
	tb = time.Now()
	log.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
