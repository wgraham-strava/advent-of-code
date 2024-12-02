// main package
package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func anyDupes(w string) bool {
	for i := range w {
		if strings.Contains(w[i+1:], string(w[i])) {
			return true
		}
	}
	return false
}

func distinctWindowPosition(line string, s int) int {
	winSize := s
	for i := 0; i < len(line)-winSize; i++ {
		window := line[i : i+winSize]
		if anyDupes(window) {
			continue
		}
		return i + winSize
	}
	return 0
}

func part1(f string) int {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return distinctWindowPosition(line, 4)
}

func part2(f string) int {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return distinctWindowPosition(line, 14)
}

func main() {
	ta := time.Now()
	log.Printf("Test input solution: %d\n", part1("test.txt"))
	tb := time.Now()
	log.Printf("Test input took: %v\n", tb.Sub(ta))
	log.Printf("Part 1 solution: %d\n", part1("input.txt"))
	ta = time.Now()
	log.Printf("Part 1 took: %v\n", ta.Sub(tb))
	log.Printf("Part 2 solution: %d\n", part2("input.txt"))
	tb = time.Now()
	log.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
