// main package
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	"time"
)

var totals []int

func part1(f string) int {
	var ans, c, t int
	var i int

	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		c, _ = strconv.Atoi(line)
		if len(line) == 0 {
			// log.Printf("Found end of bag. Bag contains a total of %d cals\n", t)
			totals = append(totals, t)
			i++
			t = 0
		}
		// log.Printf("Adding snack to bag %d with %d calories\n", i, c)
		t += c
	}
	// add the final one
	totals = append(totals, t)

	slices.Sort(totals)
	ans = totals[len(totals)-1]
	return ans
}

func part2() int {
	ans := totals[len(totals)-1]
	ans += totals[len(totals)-2]
	ans += totals[len(totals)-3]
	return ans
}

func main() {
	ta := time.Now()
	test := part1("test1.txt")
	fmt.Printf("Test input solution: %d\n", test)
	tb := time.Now()
	fmt.Printf("Test input took: %v\n", tb.Sub(ta))
	fmt.Printf("Part 1 solution: %d\n", part1("input.txt"))
	ta = time.Now()
	fmt.Printf("Part 1 took: %v\n", ta.Sub(tb))
	fmt.Printf("Part 2 solution: %d\n", part2())
	tb = time.Now()
	fmt.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
