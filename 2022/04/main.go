// main package
package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func createSet(in string) []int {
	var set []int
	nums := strings.Split(in, "-")

	s, _ := strconv.Atoi(nums[0])
	// log.Printf("First el of set is %d\n", s)
	e, _ := strconv.Atoi(nums[len(nums)-1])
	// log.Printf("Last el of set is %d\n", e)

	for i := s; i <= e; i++ {
		set = append(set, i)
	}

	return set
}

func containsAny(a, b []int) bool {
	for _, el := range a {
		if slices.Contains(b, el) {
			return true
		}
	}
	return false
}

func isSubset(a, b []int) bool {
	for _, el := range a {
		if !slices.Contains(b, el) {
			return false
		}
	}
	// log.Printf("Set %v is a subset of set %v\n", a, b)
	return true
}

func part1(f string) int {
	var ans int

	file, _ := os.Open(f)
	defer file.Close()

	// read lines, split into two parts
	// split on the '-'s and parse as integers
	// for each part, expand the rep into a set of elements
	// if part a is subset of b or is a superset of b, add to "fully contains" counter
	// return counter
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		assigns := strings.Split(line, ",")
		a := createSet(assigns[0])
		b := createSet(assigns[1])
		// log.Printf("Set a contains: %v\n", a)
		// log.Printf("Set b contains: %v\n", b)
		if len(a) <= len(b) && isSubset(a, b) {
			ans++
		} else if isSubset(b, a) {
			ans++
		}
	}
	return ans
}

func part2(f string) int {
	var ans int

	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		assigns := strings.Split(line, ",")
		a := createSet(assigns[0])
		b := createSet(assigns[1])
		// log.Printf("Set a contains: %v\n", a)
		// log.Printf("Set b contains: %v\n", b)
		if len(a) <= len(b) && containsAny(a, b) {
			ans++
		} else if containsAny(b, a) {
			ans++
		}
	}
	return ans
}

func main() {
	ta := time.Now()
	log.Printf("Test input solution: %d\n", part1("test2.txt"))
	tb := time.Now()
	log.Printf("Test input took: %v\n", tb.Sub(ta))
	log.Printf("Part 1 solution: %d\n", part1("input.txt"))
	ta = time.Now()
	log.Printf("Part 1 took: %v\n", ta.Sub(tb))
	log.Printf("Part 2 solution: %d\n", part2("input.txt"))
	tb = time.Now()
	log.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
