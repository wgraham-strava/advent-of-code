package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"time"
)

var rankMap = map[int]string{
	1: "Five of a kind",
	2: "Four of a kind",
	3: "Full house",
	4: "Three of a kind",
	5: "Two pairs",
	6: "One pair",
	7: "High card",
}

var fiveKind, fourKind, fullHouse, threeKind, twoPair, onePair, highCards []Hand = make([]Hand, 0), make([]Hand, 0), make([]Hand, 0), make([]Hand, 0), make([]Hand, 0), make([]Hand, 0), make([]Hand, 0)

var highCardMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	Cards    []string
	Bet      int
	HighCard int
	// Rank
	// 1 - Five of a kind
	// 2 - Four of a kind
	// 3 - Full house (3 of kind + 2 of kind)
	// 4 - Three of kind
	// 5 - Two pairs
	// 6 - One pair
	// 7 - High card
	Rank int
}

func countMatches(s string, c rune) int {
	var ans int
	for _, x := range s {
		if c == x {
			ans++
		}
	}
	return ans
}

func isRankedLower(a, b Hand) bool {
	for i := range a.Cards {
		a_value := highCardMap[a.Cards[i]]
		b_value := highCardMap[b.Cards[i]]
		if a_value < b_value {
			return true
		} else if a_value > b_value {
			return false
		}
	}
	return false
}

func HandSort(hands []Hand) []Hand {
	//Start the loop in reverse order, so the loop will start with length
	//which is equal to the length of input array and then loop untill
	//reaches 1
	for i := len(hands); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if !isRankedLower(hands[j-1], hands[j]) {
				intermediate := hands[j]
				hands[j] = hands[j-1]
				hands[j-1] = intermediate
			}
		}
	}
	return hands
}

func NewHand(cards []string, bet, highCard, rank int) Hand {
	return Hand{
		Cards:    cards,
		Bet:      bet,
		HighCard: highCard,
		Rank:     rank,
	}
}

func makeHand(line string) Hand {
	var h Hand
	var cards []string

	var possibleTwoPair, possibleFullHouseThree, possibleFullHouseTwo bool = false, false, false
	var highCard int = -1
	var twoPairCard rune = -1

	card_string := strings.Split(line, " ")[0]
	bet, _ := strconv.Atoi(strings.TrimSpace(strings.Split(line, " ")[1]))

	for _, char := range card_string {
		cards = append(cards, string(char))
	}

	for _, char := range card_string {
		highCard = int(math.Max(float64(highCard), float64(highCardMap[string(char)])))
		matches := countMatches(card_string, char)

		if matches == 5 {
			// fmt.Printf("Found a five of a kind: %v\n", line)
			h = NewHand(cards, bet, highCard, 1)
			fiveKind = append(fiveKind, h)
			return h
		} else if matches == 4 {
			// fmt.Printf("Found a four of a kind: %v\n", line)
			h = NewHand(cards, bet, highCard, 2)
			fourKind = append(fourKind, h)
			return h
		} else if matches == 3 {
			possibleFullHouseThree = true
			// fmt.Printf("Found a two or three of a kind, possible full house. Continuing to look.. %v\n", line)
		} else if matches == 2 {
			possibleFullHouseTwo = true
		}

		if matches == 3 && possibleFullHouseTwo {
			// fmt.Printf("Found a full house: %v\n", line)
			h = NewHand(cards, bet, highCard, 3)
			fullHouse = append(fullHouse, h)
			return h
		} else if matches == 2 && possibleFullHouseThree {
			// fmt.Printf("Found a full house: %v\n", line)
			h = NewHand(cards, bet, highCard, 3)
			fullHouse = append(fullHouse, h)
			return h
		} else if matches == 2 && twoPairCard == -1 {
			possibleTwoPair = true
			twoPairCard = char
		} else if matches == 2 && twoPairCard != -1 && twoPairCard != char {
			// fmt.Printf("Found a two pair: %v\n", line)
			h = NewHand(cards, bet, highCard, 5)
			twoPair = append(twoPair, h)
			return h
		}
	}

	if possibleTwoPair {
		// fmt.Printf("Found a one pair: %v\n", line)
		h = NewHand(cards, bet, highCard, 6)
		onePair = append(onePair, h)
		return h
	} else if possibleFullHouseThree {
		// fmt.Printf("Found a three of a kind: %v\n", line)
		h = NewHand(cards, bet, highCard, 4)
		threeKind = append(threeKind, h)
		return h
	} else if !possibleFullHouseThree && !possibleTwoPair {
		// fmt.Printf("Found a high card: %d in line %v\n", highCard, line)
		h = NewHand(cards, bet, highCard, 7)
		highCards = append(highCards, h)
		return h
	}

	h.Cards = cards
	return h
}

func part1() int {
	var ans int = 0
	var hands []Hand
	hands = make([]Hand, 0)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		makeHand(line)
	}

	highCards = HandSort(highCards)
	onePair = HandSort(onePair)
	twoPair = HandSort(twoPair)
	threeKind = HandSort(threeKind)
	fullHouse = HandSort(fullHouse)
	fiveKind = HandSort(fiveKind)
	fourKind = HandSort(fourKind)

	hands = append(hands, highCards...)
	hands = append(hands, onePair...)
	hands = append(hands, twoPair...)
	hands = append(hands, threeKind...)
	hands = append(hands, fullHouse...)
	hands = append(hands, fourKind...)
	hands = append(hands, fiveKind...)

	for i := range hands {
		winnings := (i + 1) * hands[i].Bet
		ans += winnings
	}

	return ans
}

func main() {
	t0 := time.Now()
	fmt.Printf("Part 1 solution: %d\n", part1())
	t1 := time.Now()
	fmt.Printf("Part 1 took: %v\n", t1.Sub(t0))
}
