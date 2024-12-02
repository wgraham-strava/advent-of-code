package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"time"
)

func showMeTheMoney(start_slice []int, depth int) ([]int, int) {
	allZeros := true
	new_ints := make([]int, 0)

	for i := 0; i < len(start_slice)-1; i++ {
		n := start_slice[i+1] - start_slice[i]
		if allZeros && n != 0 {
			allZeros = false
		}
		new_ints = append(new_ints, n)
	}

	if !allZeros {
		new_ints, _ = showMeTheMoney(new_ints, depth+1)

		new_big_num := start_slice[len(start_slice)-1] + new_ints[len(new_ints)-1]
		new_smol_num := start_slice[0] - new_ints[0]

		start_slice = append(start_slice, new_big_num)
		start_slice = slices.Insert(start_slice, 0, new_smol_num)
	}
	return start_slice, depth
}

func createIntSlice(line string) []int {
	split_line := strings.Split(line, " ")
	var start_slice []int

	for _, x := range split_line {
		d, _ := strconv.Atoi(x)
		start_slice = append(start_slice, d)
	}
	return start_slice
}

func part1() (int, int) {
	var ans_p1, ans_p2 int = 0, 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		int_slice := createIntSlice(line)
		new_line, _ := showMeTheMoney(int_slice, 0)

		ans_p1 += new_line[len(new_line)-1]
		ans_p2 += new_line[0]
	}
	return ans_p1, ans_p2
}

func main() {
	t0 := time.Now()
	p1, p2 := part1()
	fmt.Printf("Part 1 solution: %d\n", p1)
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
	fmt.Printf("Part 2 solution: %v\n", p2)
}
