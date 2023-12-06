package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	Head *PlantingMap
}

type FarmRange struct {
	Dest   int
	Source int
	Range  int
}

type PlantingMap struct {
	Name   string
	Ranges []FarmRange
	Next   *PlantingMap
}

func (p *PlantingMap) FindNext(input int) int {
	fmt.Printf("%d -> %s ", input, p.Name)
	for _, r := range p.Ranges {
		if input >= r.Source && input < r.Source+r.Range {
			return r.Dest + (input - r.Source)
		}
	}

	return input
}

func (a *Almanac) SeedToLocation(source int) int {
	curr := a.Head
	in := source
	fmt.Println("")
	for curr != nil {
		in = curr.FindNext(in)
		curr = curr.Next
	}
	fmt.Printf("-> %d", in)
	return in
}

func NewAlmanac() Almanac {
	return Almanac{
		NewPlantingMap(""),
	}
}

func NewPlantingMap(name string) *PlantingMap {
	return &PlantingMap{name, make([]FarmRange, 0), nil}
}

func (a *Almanac) NextPlantingMap(name string) {
	var curr *PlantingMap = a.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	next := NewPlantingMap(name)
	curr.Next = next
}

func (a *Almanac) AddRange(rng FarmRange) {
	m := a.Head
	for m.Next != nil {
		m = m.Next
	}
	m.Ranges = append(m.Ranges, rng)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	al := NewAlmanac()
	linescanner := bufio.NewScanner(file)
	linescanner.Split(bufio.ScanLines)

	res := math.MaxInt

	seeds := make([]int, 0)
	for linescanner.Scan() {
		line := linescanner.Text()
		fmt.Println(line)

		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")
		if fields[0] == "seeds" {
			for _, token := range strings.Fields(fields[1]) {
				if v, err := strconv.Atoi(token); err == nil {
					seeds = append(seeds, v)
				}
			}
			continue
		}

		if len(fields) > 1 && strings.HasSuffix(fields[0], "map") {
			al.NextPlantingMap(strings.Split(fields[0], " ")[0])
			continue
		}

		rng := make([]int, 0)
		for _, token := range strings.Fields(line) {
			if v, err := strconv.Atoi(token); err == nil {
				rng = append(rng, v)
			}
		}
		if len(rng) == 3 {
			al.AddRange(FarmRange{rng[0], rng[1], rng[2]})
		}
	}

	for _, s := range seeds {
		if loc := al.SeedToLocation(s); loc < res {
			res = loc
		}
	}
	fmt.Printf("\nLocation %d", res)
}
