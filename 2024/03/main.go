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

	for lineScanner.Scan() {
		line := lineScanner.Text()
		matches := rg.FindAllString(line, -1)

		for _, match := range matches {
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
	var pattern string = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	var rg *regexp.Regexp = regexp.MustCompile(pattern)
	var ans int = 0
	var noop = false

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	var text string
	for lineScanner.Scan() {
		text += lineScanner.Text()
	}
	matches := rg.FindAllString(text, -1)
	_ = matches

	for _, match := range matches {
		if match == "don't()" {
			noop = true
		} else if match == "do()" {
			noop = false
		} else if strings.Contains(match, "mul") && !noop {
			ans += do_multiply(match)
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
	p2 := part2()
	fmt.Printf("Part 2 solution: %v\n", p2)
	t0 = time.Now()
	fmt.Printf("Part 2 took: %v\n", t0.Sub(t1))
}
