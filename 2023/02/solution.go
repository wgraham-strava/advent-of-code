package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Find all pulls of a given color, and check to see if any are greater than allowed
func isPossible(line, color string, max_possible int) bool {
    // Use regular expression to find all digits matching the given color
    r := regexp.MustCompile("[0-9]+ " + color)
    matches := r.FindAllString(line, -1)

    // Loop through all matches found. If any match is greater than the max given, return false. This game is not possible
    for _, n := range matches {
        found, _ := strconv.Atoi(strings.Split(n, " ")[0]) 
        if found > max_possible {
            // fmt.Printf("This game cannot happen. Max %ss is %d, and we found %d %ss in this bag pull\n\n", color, max_possible, found, color)
            return false
        }
    }
    return true;
}

// Find the max number of cubes drawn for a given color in a given game
func findMax(line, color string) int {
    // Use regular expression to find all digits matching the given color
    r := regexp.MustCompile("[0-9]+ " + color)
    matches := r.FindAllString(line, -1)

    var max_found int = 0;

    // Loop through all matches found, and return only the highest value
    for _, n := range matches {
        found, _ := strconv.Atoi(strings.Split(n, " ")[0]) 
        if found > max_found { max_found = found }
    }
    return max_found;
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	var ans int

    var max_red int = 12;
    var max_green int = 13;
    var max_blue int = 14;
    var possible bool;

	// Read and print each line
	for lineScanner.Scan() {
        possible = true
		line := lineScanner.Text()

        // GameID is the first found digit in the line
        gameId, _ := strconv.Atoi((regexp.MustCompile("[0-9]+").FindString(line)))
		// fmt.Println("Game ID:", gameId)
		// fmt.Println("input:", line)

        // Check if the game is possible by looking for offending green pulls, then reds, then blues, if still possible
        possible = isPossible(line, "green", max_green)
        if possible { possible = isPossible(line, "red", max_red) };
        if possible { possible = isPossible(line, "blue", max_blue) };

        // If the game is still possible, add game id to running total
        if possible {
            // fmt.Printf("This game is possible! Adding Game ID of %d to current rolling answer of %d\n\n", gameId, ans)
            ans += gameId
        }

	}
    return ans
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	var ans int

    var red, green, blue int

	// Find the max amount of reds, greens, blues for a given game, multiply them
	for lineScanner.Scan() {
        line := lineScanner.Text()

        red = findMax(line, "red")
        green = findMax(line, "green")
        blue = findMax(line, "blue")

        ans += (red * blue * green)
	}
    return ans
}

func main() {
	start := time.Now()

	fmt.Println("Part 1 Solution:", part1())
    p1_time := time.Since(start)
	fmt.Println("Part 1 took", p1_time)

    fmt.Println("Part 2 Solution:", part2())
	fmt.Println("Part 2 took", time.Since(start) - p1_time)
}
