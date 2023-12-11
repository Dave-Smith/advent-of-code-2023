package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []string
	Bid   int
	Score int
}

var cardValue map[string]int = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func main() {
	hands := readInput("./input2.txt")
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if diff := b.Score - a.Score; diff != 0 {
			return diff
		}
		for i := 0; i < 5; i++ {
			if diff := cardValue[a.Cards[i]] - cardValue[b.Cards[i]]; diff != 0 {
				return diff
			}
		}
		return 1
	})

	res := 0
	for i := 0; i < len(hands); i++ {
		fmt.Printf("Cards %s, Bid %d, winnings %d\n", hands[i].Cards, hands[i].Bid, hands[i].Bid*(i+1))
		res += hands[i].Bid * (i + 1)
	}
	fmt.Println(hands, res)
}

func scoreHand(cards []string) int {
	counter := make(map[string]int)
	for i := 0; i < len(cards); i++ {
		if val, ok := counter[cards[i]]; ok {
			counter[cards[i]] = val + 1
		} else {
			counter[cards[i]] = 1
		}
	}
	if v, ok := counter["J"]; ok {
		cardsCopy := make([]string, 0)
		copy(cards, cardsCopy)

	}
	if len(counter) == 1 {
		return 1
	}
	if len(counter) == 5 {
		return 7
	}
	twoOfA := false
	threeOfA := false
	for _, v := range counter {
		if v == 4 {
			return 2
		}
		if v == 3 {
			threeOfA = true
		}
		if v == 2 {
			if twoOfA == true {
				return 5
			}
			twoOfA = true
		}
	}
	if twoOfA && threeOfA {
		return 3
	}
	if threeOfA {
		return 4
	}
	return 6 // pair
}

func highestNonJ(hand map[string]int) string {
	cloned := maps.Clone(hand)
	for k, v := range cloned {
		cloned[k] = cardValue[k] * hand[k]
	}
	for k, v := range hand {
		if k == "J" {
			continue
		}

		if v == 4 {
			return k
		}
	}
}

func readInput(filename string) []Hand {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	linescanner := bufio.NewScanner(file)
	linescanner.Split(bufio.ScanLines)

	hands := make([]Hand, 0)
	for linescanner.Scan() {
		line := linescanner.Text()

		h := strings.Split(line, " ")
		if val, err := strconv.Atoi(h[1]); err == nil {
			cards := strings.Split(h[0], "")
			hand := Hand{cards, val, scoreHand(cards)}
			hands = append(hands, hand)
		}
	}
	return hands
}
