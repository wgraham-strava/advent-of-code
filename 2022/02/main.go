// main package
package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

var totals []int

// A = 65 = rock
// B = 66 = paper
// C = 67 = scissors

// X = 88 = rock
// Y = 89 = paper
// Z = 90 = scissors

func part1(f string) int {
	var ans int

	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := []byte(scanner.Text())
		p1 := int(chars[0]) - 64
		p2 := int(chars[2]) - 87

		var points, hand int
		points = p2 // utf-8 magic
		if p1 == p2 {
			// tie
			hand = 3
		} else if ((p1 + 1) % 3) == (p2 % 3) {
			// p2 win
			hand = 6
		}
		hand += points
		ans += hand
		// log.Printf("%d + %d = %d\nRunning Total: %d\n\n", points, hand, points+hand, ans)

	}
	return ans
}

func part2(f string) int {
	var ans int
	var points, hand int

	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := []byte(scanner.Text())
		p1 := int(chars[0]) - 64
		p2 := int(chars[2]) - 87

		// log.Printf("p1 = %d, p2 = %d\n", p1, p2)
		// X = lose = 88
		// Y = tie = 89
		// Z = win = 90

		points = p2  // utf-8 magic
		if p2 == 2 { // tie
			points = p1
			// log.Printf("Pick %d to force a tie (+3)\n", points)
			hand = 3
		} else if p2 == 3 { // p2 win
			points = (p1 % 3) + 1
			hand = 6
			// log.Printf("Pick %d to force a win (+6)\n", points)
		} else { // p2 lose
			p2 = ((p1 + 1) % 3) + 1
			points = p2
			hand = 0
			// log.Printf("Pick %d to force a loss (+0)\n", points)
		}
		ans += hand
		ans += points
		// log.Printf("%d + %d = %d\n", points, hand, points+hand)
		// log.Printf("Running Total: %d\n\n", ans)

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
