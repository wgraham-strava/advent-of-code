package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1() int {
	var nums *[]int
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
		split_line := strings.Split(line, " ")
		nums = string_to_int_slice(&split_line)
		if is_safe(nums, false) {
			ans++
		}
	}

	return ans
}

func part2() int {
	var nums *[]int
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
		split_line := strings.Split(line, " ")
		nums = string_to_int_slice(&split_line)
		if is_safe(nums, true) {
			ans++
		}
	}

	return ans
}

func is_safe(nums *[]int, tolerate bool) bool {
	var diff int
	var dir int = 1
	var safe bool = true

	if (*nums)[1] < (*nums)[0] {
		dir = -1
	}

	for i := 1; i < len(*nums); i++ {
		diff = (*nums)[i] - (*nums)[i-1]
		if diff == 0 || math.Abs(float64(diff)) > 3 || diff*dir < 0 {
			safe = false
		}
	}

	// If repot is not safe, check again, removing one number at a time until safe
	if !safe && tolerate {
		for i := range *nums {

			tol_nums := make([]int, 0)
			tol_nums = append(tol_nums, (*nums)[:i]...)
			tol_nums = append(tol_nums, (*nums)[i+1:]...)

			if is_safe(&tol_nums, false) {
				return true
			}
		}
	}
	return safe
}

func string_to_int_slice(s *[]string) *[]int {
	nums := make([]int, len(*s))
	for i := range *s {
		val, _ := strconv.Atoi((*s)[i])
		nums[i] = val
	}
	return &nums
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
