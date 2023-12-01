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
