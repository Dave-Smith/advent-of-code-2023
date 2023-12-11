package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	report := readInput("./input.txt")
	res := 0
	for _, items := range report {
		res += getPrev(items)
	}
	fmt.Println(res)
}

func getPrev(items []int) int {
	if allZeros(items) {
		return 0
	}

	if len(items) == 1 {
		return 0
	}

	reduced := make([]int, 0)
	for i := 0; i < len(items)-1; i++ {
		reduced = append(reduced, items[i+1]-items[i])
	}

	return items[0] - getPrev(reduced)
}

func getNext(items []int) int {
	if allZeros(items) {
		return 0
	}

	if len(items) == 1 {
		return 0
	}

	reduced := make([]int, 0)
	for i := 0; i < len(items)-1; i++ {
		reduced = append(reduced, items[i+1]-items[i])
	}

	return items[len(items)-1] + getNext(reduced)
}

func allZeros(items []int) bool {
	for _, v := range items {
		if v != 0 {
			return false
		}
	}
	return true
}

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	linescanner := bufio.NewScanner(file)
	linescanner.Split(bufio.ScanLines)

	envReport := make([][]int, 0)
	for linescanner.Scan() {
		line := linescanner.Text()

		lineItems := make([]int, 0)
		for _, token := range strings.Fields(line) {
			if v, err := strconv.Atoi(token); err == nil {
				lineItems = append(lineItems, v)
			}
		}
		envReport = append(envReport, lineItems)
	}

	fmt.Println(envReport)
	return envReport
}
