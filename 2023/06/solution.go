package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func fillSlice(line string, p *[]int) {
	l := strings.Fields(strings.Split(line, ":")[1])
	for _, data := range l {
		x, _ := strconv.Atoi(strings.TrimSpace(data))
		*p = append(*p, x)
	}
}

func calculateWins(time, distance, i int) int {
	var ans int
	for t := 0; t <= time; t++ {
		d := (time - t) * t
		// fmt.Printf("Button held for %d seconds, has speed of %d and covers distance %d\n", t, t, distance)
		if d > distance {
			// fmt.Println("This strategy wins the race!!")
			ans++
			// fmt.Printf("There are currently %d ways to win race #%d\n", race_wins, i+1)
		}
	}
	return ans
}

func part1() int {
	var ans int = 1
	var times, distances []int
	times = make([]int, 0)
	distances = make([]int, 0)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			fillSlice(line, &times)
		}
		if strings.Contains(line, "Distance") {
			fillSlice(line, &distances)
		}
	}

	for i := 0; i < len(times); i++ {
		// fmt.Println(strings.Repeat("=", 60))
		// fmt.Printf("Race %d: Best distance is %d meters in %d seconds\n", i, distances[i], times[i])
		race_wins := calculateWins(times[i], distances[i], i)
		// fmt.Printf("There are %d ways to win race #%d\n", race_wins, i+1)
		ans *= race_wins
		// fmt.Println(strings.Repeat("=", 60))
	}

	return ans
}

func part2() int {
	var ans int = 1
	var times, distances []int
	times = make([]int, 0)
	distances = make([]int, 0)

	file, _ := os.Open("input2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			fillSlice(line, &times)
		}
		if strings.Contains(line, "Distance") {
			fillSlice(line, &distances)
		}
	}

	for i := 0; i < len(times); i++ {
		// fmt.Println(strings.Repeat("=", 60))
		// fmt.Printf("Race %d: Best distance is %d meters in %d seconds\n", i, distances[i], times[i])
		race_wins := calculateWins(times[i], distances[i], i)
		// fmt.Printf("There are %d ways to win race #%d\n", race_wins, i+1)
		ans *= race_wins
		// fmt.Println(strings.Repeat("=", 60))
	}

	return ans
}

func main() {
	t0 := time.Now()
	fmt.Printf("Part 1 solution: %d\n", part1())
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
	fmt.Printf("Part 2 solution: %d\n", part2())
	t2 := time.Now()
	fmt.Printf("Part 2 took: %v\n", t2.Sub(t1))
}
