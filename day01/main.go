package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}

var numLookup = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func partTwo() {
	keys := make([]string, len(numLookup))
	for k, _ := range numLookup {
		keys = append(keys, k)
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		nums := ScanLine(line)
		if len(nums) == 0 {
			continue
		}
		tot := int64(nums[0]*10) + int64(nums[len(nums)-1])
		fmt.Printf("Value: %v in line %v\n", tot, line)
		res = res + tot
	}

	fmt.Println(res)
}

func ScanLine(line string) []int {
	res := make([]int, 0)
	const maxWindow int = 6

	for w := 0; w < len(line); w++ {
		for i := 1; i < maxWindow; i++ {
			end := w + i
			if end > len(line) {
				end = len(line)
			}
			token := line[w:end]
			if v, ok := numLookup[token]; ok {
				res = append(res, v)
				break
			}
		}
	}
	fmt.Printf("Numbers found %v in line %v", res, line)
	return res
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		lineScan := bufio.NewScanner(strings.NewReader(line))
		lineScan.Split(bufio.ScanBytes)

		nums := make([]int, 0)
		for lineScan.Scan() {
			c := lineScan.Text()
			if i, e := strconv.Atoi(c); e == nil {
				nums = append(nums, i)
			}
		}
		if len(nums) == 0 {
			continue
		}
		tot := int64(nums[0]*10) + int64(nums[len(nums)-1])
		res = res + tot
	}

	fmt.Println(res)
}
