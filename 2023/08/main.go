package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"time"
)

func part1() int {
	var rules map[string][]string = make(map[string][]string)
	var directions string
	var path string = "AAA"

	var ans int = 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read first line into directions
	scanner.Scan()
	directions = scanner.Text()
	// fmt.Printf("Directions for input are: %v\n", directions)

	// Throwaway the empty line
	scanner.Scan()

	// Start building the rules map
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, "=")
		junction := strings.TrimSpace(splits[0])
		splits = strings.Split(strings.Trim(strings.TrimSpace(splits[1]), "()"), ",")
		left, right := strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])

		d := make([]string, 2)
		d[0] = left
		d[1] = right
		rules[junction] = d
	}

get_me_out:
	for true {
		for _, d := range directions {
			if path == "ZZZ" {
				// fmt.Printf("Found the exit!\n")
				break get_me_out
			}
			if string(d) == "L" {
				// fmt.Println("Taking the left path")
				path = rules[path][0]
			} else if string(d) == "R" {
				// fmt.Println("Taking the right path")
				path = rules[path][1]
			}
			ans++
		}
	}

	return ans
}

func main() {
	t0 := time.Now()
	fmt.Printf("Part 1 solution: %d\n", part1())
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
}
