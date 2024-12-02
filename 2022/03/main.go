// main package
package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func part1(f string) int {
	var ans int

	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := scanner.Text()
		s1 := chars[0 : len(chars)/2]
		s2 := chars[len(chars)/2:]
		// log.Printf("First bit: %s\n", s1)
		// log.Printf("Second bit: %s\n", s2)

		for _, c := range s1 {
			if strings.Contains(s2, string(c)) {
				d := c
				// convert to int and subtract UTF-8 offset
				d -= 38

				// lowercase numbers are higher value in UTF-8, so we need to subtract more
				if d > 52 {
					d -= 58
				}
				log.Printf("%s (%d) is present in both!\n", string(c), d)
				ans += int(d)
				break
			}
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
		// find the one item that is common within a set of 3 lines
		l1 := scanner.Text()
		scanner.Scan()
		l2 := scanner.Text()
		scanner.Scan()
		l3 := scanner.Text()

		for _, c := range l1 {
			if strings.Contains(l2, string(c)) && strings.Contains(l3, string(c)) {
				d := c
				// convert to int and subtract UTF-8 offset
				d -= 38

				// lowercase numbers are higher value in UTF-8, so we need to subtract more
				if d > 52 {
					d -= 58
				}
				log.Printf("%s (%d) is present in all three!\n", string(c), d)
				ans += int(d)
				break
			}
		}
	}
	return ans
}

func main() {
	ta := time.Now()
	log.Printf("Test input solution: %d\n", part2("test.txt"))
	tb := time.Now()
	log.Printf("Test input took: %v\n", tb.Sub(ta))
	log.Printf("Part 1 solution: %d\n", part1("input.txt"))
	ta = time.Now()
	log.Printf("Part 1 took: %v\n", ta.Sub(tb))
	log.Printf("Part 2 solution: %d\n", part2("input.txt"))
	tb = time.Now()
	log.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
