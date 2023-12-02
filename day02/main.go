package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := int64(0)

	for scanner.Scan() {
		game := scanner.Text()
		pulls := strings.Split(game, ":")[1]
		var red, green, blue int
		for _, pull := range strings.Split(pulls, ";") {
			for _, cube := range strings.Split(pull, ",") {
				var color string
				var val int
				if _, err := fmt.Sscanf(cube, "%d %s", &val, &color); err != nil {
					log.Fatal(err)
				}
				if color == "red" {
					red = int(math.Max(float64(red), float64(val)))
				}
				if color == "green" {
					green = int(math.Max(float64(green), float64(val)))
				}
				if color == "blue" {
					blue = int(math.Max(float64(blue), float64(val)))
				}
			}
		}
		power := red * green * blue
		res += int64(power)
	}
	fmt.Println(res)
}
func partOne() {
	//12 red cubes, 13 green cubes, and 14 blue cubes
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)

		gameParts := strings.Split(line, ":")
		var id int
		_, err := fmt.Sscanf(gameParts[0], "Game %d", &id)
		if err != nil {
			panic(err)
		}

		fail := false
		for _, pull := range strings.Split(gameParts[1], ";") {
			//fmt.Println(pull)

			for _, cube := range strings.Split(pull, ",") {
				var color string
				var val int
				if _, err := fmt.Sscanf(cube, "%d %s", &val, &color); err != nil {
					panic(err)
				}
				//fmt.Printf("Cube: %s %d", color, val)
				if (color == "red" && val > 12) || (color == "green" && val > 13) || (color == "blue" && val > 14) {
					fmt.Printf("Number to high: %s: %d ", color, val)
					fail = true
				}
			}
			if fail {
				fmt.Printf("Game %v is not good. input: %s\n", id, line)
				break
			}
		}
		if !fail {
			res += id
		}
	}

	fmt.Printf("\n Result: %d", res)
}
