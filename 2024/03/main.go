package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func part1() int {
	var ans int = 0

	var pattern string = `mul\(\d{1,3},\d{1,3}\)`
	var rg *regexp.Regexp = regexp.MustCompile(pattern)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	// Read and print each line
	for lineScanner.Scan() {
		line := lineScanner.Text()
		matches := rg.FindAllString(line, -1)

		for _, match := range matches {
			// fmt.Println(match)
			ans += do_multiply(match)
		}
	}

	return ans
}

func do_multiply(s string) int {
	nums := strings.Split(strings.Split(strings.Split(s, "(")[1], ")")[0], ",")
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	return a * b
}

func part2() int {
	var pattern string = `mul\(\d{1,3},\d{1,3}\)`
	var do_pattern string = `do\(\)`
	var dont_pattern string = `don't\(\)`
	var rg *regexp.Regexp = regexp.MustCompile(pattern)
	var rg_do *regexp.Regexp = regexp.MustCompile(do_pattern)
	var rg_dont *regexp.Regexp = regexp.MustCompile(dont_pattern)
	var ans int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	// Read and print each line
	for lineScanner.Scan() {
		line := lineScanner.Text()
		fmt.Println("Line = ", line)
		matches := rg.FindAllStringIndex(line, -1)
		_ = matches
		do := rg_do.FindAllIndex([]byte(line), -1)
		dont := rg_dont.FindAllIndex([]byte(line), -1)

		fmt.Printf("matches = %v\n", matches)
		fmt.Printf("dos = %v\n", do)
		fmt.Printf("donts = %v\n", dont)

		fmt.Printf("First match = %s\n", line[matches[0][0]:matches[0][1]])
	}

	// build list of "good ranges"
	// for each match, check to see if it is in a good range
	// if good, mult and add

	return ans
}

func main() {
	t0 := time.Now()
	p1 := part1()
	fmt.Printf("Part 1 solution: %d\n", p1)
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
	p2 := part2()
	fmt.Printf("Part 2 solution: %v\n", p2)
	t0 = time.Now()
	fmt.Printf("Part 2 took: %v\n", t0.Sub(t1))
}