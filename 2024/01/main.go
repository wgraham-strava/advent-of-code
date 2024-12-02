package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func part1() int {
	var ans int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	// Read and print each line
	for lineScanner.Scan() {
		line := lineScanner.Text()
		// fmt.Println(line)
		split_line := strings.Split(line, "   ")
		p1, _ := strconv.Atoi(split_line[0])
		p2, _ := strconv.Atoi(split_line[1])

		a1 = append(a1, p1)
		a2 = append(a2, p2)
	}

	// fmt.Printf("Unsorted first array is %v\n", a1)
	// fmt.Printf("Unsorted second array is %v\n", a2)
	slices.Sort(a1)
	slices.Sort(a2)
	// fmt.Printf("Sorted first array is %v\n", a1)
	// fmt.Printf("Sorted second array is %v\n", a2)

	for i := range a1 {
		ans += int(math.Abs(float64(a1[i] - a2[i])))
	}

	return ans
}

func part2() int {
	var ans int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	a1 := make([]int, 0)
	occ := make(map[int]int)

	// Read and print each line
	for lineScanner.Scan() {
		line := lineScanner.Text()
		// fmt.Println(line)
		split_line := strings.Split(line, "   ")
		p1, _ := strconv.Atoi(split_line[0])
		p2, _ := strconv.Atoi(split_line[1])

		a1 = append(a1, p1)

		occ[p2] = occ[p2] + 1
		// fmt.Printf("%d occurs %d times\n", p2, occ[p2])
	}

	for _, val := range a1 {
		// fmt.Printf("Looking up occurences of %d\n", val)
		// fmt.Printf("%d occurs %d times\n", val, occ[val])
		ans += val * occ[val]
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
