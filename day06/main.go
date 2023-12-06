package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Time int
type Distance int

func main() {
	times, dist := ReadPartTwoInput("./input.txt")
	if len(times) != len(dist) {
		panic("The games list for time and distance record do not match")
	}
	races := make([]int, 0)
	res := 1
	for i := 0; i < len(times); i++ {
		winning := WinningStrategies(times[i], dist[i])
		races = append(races, winning...)
		res = res * len(winning)
	}

	fmt.Printf("Winning strategies %d\n", res)
}

func WinningStrategies(time Time, record Distance) []int {
	winning := make([]int, 0)
	for i := 1; i < int(time)-1; i++ {
		if dist := i * (int(time) - i); dist > int(record) {
			winning = append(winning, dist)
		}
	}
	return winning
}

func ReadPartTwoInput(filename string) (times []Time, dist []Distance) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linescanner := bufio.NewScanner(file)
	linescanner.Split(bufio.ScanLines)

	times = make([]Time, 0)
	dist = make([]Distance, 0)
	for linescanner.Scan() {
		line := linescanner.Text()

		if strings.HasPrefix(line, "Time") {
			token := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
			if val, err := strconv.Atoi(token); err == nil {
				times = append(times, Time(val))
			}
		}

		if strings.HasPrefix(line, "Distance") {
			token := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
			if val, err := strconv.Atoi(token); err == nil {
				dist = append(dist, Distance(val))
			}
		}
	}
	return times, dist
}
func ReadInput(filename string) (times []Time, dist []Distance) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	linescanner := bufio.NewScanner(file)
	linescanner.Split(bufio.ScanLines)

	times = make([]Time, 0)
	dist = make([]Distance, 0)
	for linescanner.Scan() {
		line := linescanner.Text()

		if strings.HasPrefix(line, "Time") {
			for _, token := range strings.Fields(strings.Split(line, ":")[1]) {
				if val, err := strconv.Atoi(token); err == nil {
					times = append(times, Time(val))
				}
			}
		}

		if strings.HasPrefix(line, "Distance") {
			for _, token := range strings.Fields(strings.Split(line, ":")[1]) {
				if val, err := strconv.Atoi(token); err == nil {
					dist = append(dist, Distance(val))
				}
			}
		}
	}
	return times, dist
}
