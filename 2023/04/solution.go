package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "math"
	"regexp"
	"time"
)

// Given a number "win", return true if this number exists in "nums"
func isWinner(win int, nums []int) bool {
    // fmt.Printf("Checking number %d\n", win)
    for _, n := range nums {
        if win == n { 
            fmt.Printf("Found a match. Number %d is a winner!\n", win)
            return true 
        }
    }
    return false
}

// Given a list of numbers "winners", return the total sum of all winners in "nums"
func findWinners(winners, nums []int) int {
    fmt.Printf("Checking numbers to find winners...\n")
    var ans int
    for _, w := range winners {
        if isWinner(w, nums) { ans += 1 }
    }
    return ans
}

func part1() int {
	var ans int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	for lineScanner.Scan() {
        var nums, winners []int
        fmt.Println(strings.Repeat("-", 100))

        line := lineScanner.Text()
        
        cardId, _ := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(line))

        s := strings.Split(strings.Split(line, ":")[1], "|")
        nums_string, winners_string := s[0], s[1]
        nums_slice, winners_slice := strings.Fields(nums_string), strings.Fields(winners_string)

        for _, data := range nums_slice {
            n, _ := strconv.Atoi(data)
            nums = append(nums, n)
        }

        for _, data := range winners_slice {
            n, _ := strconv.Atoi(data)
            winners = append(winners, n)
        }

        fmt.Printf("Looking at card number: %d\n", cardId)
        fmt.Printf("Card contains numbers: %v\n", nums)
        // fmt.Printf("Have winning numbers: %v\n", winners)

        // fmt.Printf("Looking to see which winning numbers matches the card\n")
        wins := findWinners(winners, nums)
        value := int(math.Floor(math.Pow(2, float64(wins-1))))
        ans += value

        if wins != 0 {
            fmt.Printf("This card is a winner! Found %d wins with a total value of %d\n", wins, value)
        } else {
            fmt.Printf("This card is a loser! Did not find any wins\n")
        }

        fmt.Println(strings.Repeat("-", 100))
	}

    return ans
}

func part2() int {
	var ans int
    copies := make([]int, 200)

    for i := range copies {
        copies[i] = 1
    }

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	for lineScanner.Scan() {
        var nums, winners []int
        fmt.Println(strings.Repeat("-", 100))

        line := lineScanner.Text()
        
        cardId, _ := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(line))

        s := strings.Split(strings.Split(line, ":")[1], "|")
        nums_string, winners_string := s[0], s[1]
        nums_slice, winners_slice := strings.Fields(nums_string), strings.Fields(winners_string)

        for _, data := range nums_slice {
            n, _ := strconv.Atoi(data)
            nums = append(nums, n)
        }

        for _, data := range winners_slice {
            n, _ := strconv.Atoi(data)
            winners = append(winners, n)
        }

        wins := findWinners(winners, nums)

        if wins != 0 {
            fmt.Printf("This card had winners, so you get more cards yay!! The next %d cards have been copied!\n", wins)
        } else {
            fmt.Printf("This card is a loser! Did not find any more cards\n")
        }

        var cards int
        for i := range copies {
            cards += copies[i]
        }

        for i := 1; i <= wins; i++ {
            copies[cardId+i-1] += 1*copies[cardId-1]
        }

        fmt.Printf("Card %d has %d copies\n", cardId, copies[cardId-1])
        fmt.Printf("You currently have %d cards\n", cards)
        ans += copies[cardId-1]
        fmt.Println(strings.Repeat("-", 100))
	}

    return ans
}

func main() {
	start := time.Now()

	// fmt.Println("Part 1 Solution:", part1())
    p1_time := time.Since(start)
	fmt.Println("Part 1 took", p1_time)

    fmt.Println("Part 2 Solution:", part2())
	fmt.Println("Part 2 took", time.Since(start) - p1_time)
}
