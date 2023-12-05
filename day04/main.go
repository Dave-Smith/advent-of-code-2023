package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}
func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linescan := bufio.NewScanner(file)
	linescan.Split(bufio.ScanLines)
	games := make([]int, 211)

	res := int64(0)
	gameCounter := 0
	for linescan.Scan() {
		game := strings.Split(linescan.Text(), ":")[1]

		gameNums := strings.Split(game, "|")
		fmt.Printf("%v\n", gameNums[1])
		if len(gameNums) > 2 {
			panic("More than two cards played")
		}

		winning := make(map[int]bool, 0)
		mine := make([]int, 0)
		for _, v := range strings.Fields(gameNums[0]) {
			if n, err := strconv.Atoi(v); err == nil {
				winning[n] = true
			}
		}
		for _, v := range strings.Fields(gameNums[1]) {
			if n, err := strconv.Atoi(v); err == nil {
				if _, ok := winning[n]; ok {
					mine = append(mine, n)
				}
			}
		}

		games[gameCounter]++
		fmt.Printf("Game %v has %v cards\n", gameCounter, games[gameCounter])

		for i := 0; i < len(mine); i++ {
			cardCount := games[gameCounter]
			if gameCounter+i+1 < len(games) {
				games[gameCounter+i+1] = games[gameCounter+i+1] + cardCount
			}
		}

		res += int64(games[gameCounter])
		gameCounter++
		// fmt.Printf("Games: %v\n", games)
	}

	fmt.Printf("Score: %v\n", res)
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linescan := bufio.NewScanner(file)
	linescan.Split(bufio.ScanLines)

	res := int64(0)
	for linescan.Scan() {
		game := strings.Split(linescan.Text(), ":")[1]

		gameNums := strings.Split(game, "|")
		fmt.Printf("%v\n", gameNums[1])
		if len(gameNums) > 2 {
			panic("More than two cards played")
		}

		winning := make(map[int]bool, 0)
		mine := make([]int, 0)
		for _, v := range strings.Fields(gameNums[0]) {
			if n, err := strconv.Atoi(v); err == nil {
				winning[n] = true
			}
		}
		for _, v := range strings.Fields(gameNums[1]) {
			if n, err := strconv.Atoi(v); err == nil {
				if _, ok := winning[n]; ok {
					mine = append(mine, n)
				}
			}
		}
		fmt.Print(len(mine))
		res += int64(math.Pow(2, float64(len(mine)-1)))

		fmt.Println(winning, mine)
	}

	fmt.Printf("Score: %v\n", res)
}
