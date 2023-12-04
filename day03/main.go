package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Part struct {
	Number int
	Start  int
	End    int
	Line   int
}

type Point struct {
	x int
	y int
}

var specialChars = map[int]interface{}{33: nil, 34: nil, 35: nil, 36: nil, 37: nil, 38: nil, 39: nil, 40: nil, 41: nil, 42: nil, 43: nil, 44: nil, 45: nil, 47: nil, 58: nil, 59: nil, 60: nil, 61: nil, 62: nil, 63: nil, 64: nil, 91: nil, 92: nil, 93: nil, 94: nil, 95: nil, 96: nil, 123: nil, 124: nil, 125: nil, 126: nil}

func main() {
	partTwo()
}
func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	gears := make([]Point, 0)
	parts := make([]Part, 0)

	var curr string
	lineNum := 0
	for scanner.Scan() {
		curr = scanner.Text()

		start := 0
		for {
			part, ok := nextNumber(curr, start)
			if !ok {
				break
			}
			part.Line = lineNum
			parts = append(parts, part)
			start = part.End + 1
		}

		gears = append(gears, nextGears(curr, lineNum)...)
		lineNum++
	}

	fmt.Printf("\n Result: %d", gearRatios(gears, parts))
}

func gearRatios(gears []Point, parts []Part) int64 {
	res := int64(0)
	for _, g := range gears {
		gearParts := adjacentParts(g, parts)
		if len(gearParts) == 2 {
			res += int64(gearParts[0].Number) * int64(gearParts[1].Number)
		}
	}
	return res
}

func adjacentParts(g Point, parts []Part) []Part {
	res := make([]Part, 0)
	for _, p := range parts {
		for _, a := range adjacent(p) {
			if g == a {
				res = append(res, p)
				break
			}
		}
	}
	return res
}

func adjacent(part Part) []Point {
	p := make([]Point, 0)
	for i := part.Start - 1; i <= part.End+1; i++ {
		p = append(p, Point{i, part.Line - 1})
		p = append(p, Point{i, part.Line})
		p = append(p, Point{i, part.Line + 1})
	}
	return p
}

func nextGears(line string, lineNum int) []Point {
	res := make([]Point, 0)
	for i := 0; i < len(line); i++ {
		if line[i] == 42 {
			res = append(res, Point{i, lineNum})
		}
	}
	return res
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := int64(0)
	parts := make([]Part, 0)

	var curr, prev string
	scanner.Scan()
	next := scanner.Text()
	for scanner.Scan() || (len(next) > 0 && len(prev) > 0) {
		prev = curr
		curr = next
		next = scanner.Text()

		prevMap := mapLine(prev)
		currMap := mapLine(curr)
		nextMap := mapLine(next)

		start := 0
		for {
			part, ok := nextNumber(curr, start)
			if !ok {
				break
			}
			if isPartValid(part, prevMap, currMap, nextMap) {
				parts = append(parts, part)
				res += int64(part.Number)
			}
			// fmt.Printf("Part: %v\n", part)
			start = part.End + 1
		}
		//fmt.Printf("Numbers: %v in %s\n", parts, curr)
	}

	fmt.Printf("\n Result: %d", res)
}

func isPartValid(part Part, prev, curr, next []int) bool {
	pos := make(map[int]interface{})
	for i := part.Start; i <= part.End; i++ {
		pos[i] = nil
	}
	if part.Start > 0 {
		pos[part.Start-1] = nil
	}
	pos[part.End+1] = nil

	for _, v := range prev {
		if _, ok := pos[v]; ok {
			return true
		}
	}
	for _, v := range curr {
		if _, ok := pos[v]; ok {
			return true
		}
	}
	for _, v := range next {
		if _, ok := pos[v]; ok {
			return true
		}
	}

	return false
}

func nextNumber(line string, start int) (Part, bool) {
	if len(line)-1 == start {
		return Part{}, false
	}

	end := start
	number := make([]byte, 0)
	for i := start; i < len(line); i++ {
		if line[i] >= 48 && line[i] <= 57 {
			number = append(number, line[i])
			end = i + 1
		} else if len(number) > 0 {
			break
		}
	}
	if len(number) == 0 {
		return Part{}, false
	}
	partNum, _ := strconv.Atoi(string(number))
	return Part{partNum, end - len(number), end - 1, 0}, true
}

func mapLine(line string) []int {
	lineMap := make([]int, 0)
	for i, v := range line {
		c := int(v)
		if _, ok := specialChars[c]; ok {
			lineMap = append(lineMap, i)
		}
	}
	return lineMap
}
