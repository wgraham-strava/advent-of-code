// main package
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

type Move struct {
	Qty, Src, Dst int
}

type Stack struct {
	ID     int
	Crates []string
}

func (s *Stack) push(v string) *Stack {
	s.Crates = append(s.Crates, v)
	return s
}

func (s *Stack) pop() string {
	v := s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[0 : len(s.Crates)-1]
	return v
}

func (s *Stack) add(v []string) *Stack {
	s.Crates = append(s.Crates, v...)
	return s
}

func (s *Stack) remove(n int) []string {
	v := s.Crates[len(s.Crates)-n:]
	s.Crates = s.Crates[0 : len(s.Crates)-n]
	return v
}

func parseInput(f string) ([]Stack, []Move) {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stacks := make([]Stack, 10)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for i := 1; i < len(line)-1; i += 4 {
			if line[i] >= 'A' && line[i] <= 'Z' {
				crates := []string{}
				crates = append(crates, string(line[i]))
				stackID := (i / 4) + 1
				stacks[stackID-1].ID = stackID
				stacks[stackID-1].Crates = append(stacks[stackID-1].Crates, crates...)
			}
		}
	}

	for i := range stacks {
		slices.Reverse(stacks[i].Crates)
	}

	moves := []Move{}
	for scanner.Scan() {
		line := scanner.Text()
		var qty, src, dst int
		fmt.Sscanf(line, "move %d from %d to %d", &qty, &src, &dst)
		moves = append(moves, Move{qty, src, dst})
	}

	return stacks, moves
}

func part1(f string) string {
	var ans string

	stacks, moves := parseInput(f)

	for _, el := range moves {
		for j := 0; j < el.Qty; j++ {
			s := stacks[el.Src-1].pop()
			stacks[el.Dst-1].push(s)
		}
	}

	for _, el := range stacks {
		if len(el.Crates) != 0 {
			ans += el.Crates[len(el.Crates)-1]
		}
	}
	return ans
}

func part2(f string) string {
	var ans string
	stacks, moves := parseInput(f)

	for _, el := range moves {
		s := stacks[el.Src-1].remove(el.Qty)
		stacks[el.Dst-1].add(s)
	}

	for _, el := range stacks {
		if len(el.Crates) != 0 {
			ans += el.Crates[len(el.Crates)-1]
		}
	}
	return ans
}

func main() {
	ta := time.Now()
	log.Printf("Test input solution: %s\n", part2("test.txt"))
	tb := time.Now()
	log.Printf("Test input took: %v\n", tb.Sub(ta))
	log.Printf("Part 1 solution: %s\n", part1("input.txt"))
	ta = time.Now()
	log.Printf("Part 1 took: %v\n", ta.Sub(tb))
	log.Printf("Part 2 solution: %s\n", part2("input.txt"))
	tb = time.Now()
	log.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
