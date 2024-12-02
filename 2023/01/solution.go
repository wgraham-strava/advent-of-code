package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	var firstNum, lastNum, ans, s int

	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// Read and print each line
	for lineScanner.Scan() {
		line := lineScanner.Text()

		var lastWordPos, firstWordPos = 0, len(line) + 1
		var lastWordValue, firstWordValue = 0, 0

		for i := range numbers {
			firstIndex := strings.Index(line, i)
			if firstIndex != -1 && firstIndex < firstWordPos {
				firstWordPos = firstIndex
				firstWordValue = numbers[i]
			}
			lastIndex := strings.LastIndex(line, i)
			if lastIndex != -1 && lastIndex >= lastWordPos {
				lastWordPos = lastIndex
				lastWordValue = numbers[i]
			}
		}

		firstNum, lastNum, s = 0, 0, 0
		for i, char := range line {
			if char > 47 && char < 65 && firstNum == 0 {
				if i < firstWordPos {
					firstNum = int(char)
				} else {
					firstNum = firstWordValue + 48
				}
			}

			if char > 47 && char < 65 {
				if i >= lastWordPos {
					lastNum = int(char)
				} else {
					lastNum = lastWordValue + 48
				}
			}
		}
		if firstNum == 0 && firstWordValue != 0 {
			firstNum = firstWordValue + 48
		}
		if lastNum == 0 && lastWordValue != 0 {
			lastNum = lastWordValue + 48
		}

		fmt.Println("input:", line)
		fmt.Printf("first: %d\n", firstNum-48)
		fmt.Printf("last: %d\n", lastNum-48)

		s, _ = strconv.Atoi(fmt.Sprint(firstNum) + fmt.Sprint(lastNum))
		fmt.Printf("to add: %d\n\n", s)
		ans += s
	}
	fmt.Println("Solution:", ans)

	stop := time.Since(start)
	fmt.Println("Took", stop)
}
