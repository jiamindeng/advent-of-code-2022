package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	input := parseInput(string(f))
	resOne := partOne(input)
	fmt.Println(resOne)
	resTwo := partTwo(input)
	fmt.Println(resTwo)
}

func partOne(nums []int) int {
	count := 0
	prev := nums[0]
	for _, num := range nums[1:] {
		if num > prev {
			count++
		}
		prev = num
	}
	return count
}

func partTwo(nums []int) int {
	count := 0
	prev := nums[0] + nums[1] + nums[2]
	fmt.Println(prev)
	for i := 1; i < len(nums)-2; i++ {
		tripletSum := nums[i] + nums[i+1] + nums[i+2]
		fmt.Println(tripletSum)
		if tripletSum > prev {
			count++
		}
		prev = tripletSum
	}
	return count
}

func parseInput(f string) []int {
	input := []int{}
	lines := strings.Split(f, "\n")
	for _, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		input = append(input, int(num))
	}

	return input
}
