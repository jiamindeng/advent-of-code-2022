package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	fmt.Println(string(f))
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func partOne(input string) int {
	count := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		p := [][]int{}
		for _, pair := range pairs {
			pair_ := strings.Split(pair, "-")
			parsedPair := []int{}
			for _, p_ := range pair_ {
				parsedInt, _ := strconv.ParseInt(p_, 10, 64)
				parsedPair = append(parsedPair, int(parsedInt))
			}
			p = append(p, parsedPair)
		}

		if isInInterval(p[0], p[1]) || isInInterval(p[1], p[0]) {
			count++
		}
	}
	return count
}

func partTwo(input string) int {
	count := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		p := [][]int{}
		for _, pair := range pairs {
			pair_ := strings.Split(pair, "-")
			parsedPair := []int{}
			for _, p_ := range pair_ {
				parsedInt, _ := strconv.ParseInt(p_, 10, 64)
				parsedPair = append(parsedPair, int(parsedInt))
			}
			p = append(p, parsedPair)
		}

		if overlap(p[0], p[1]) || overlap(p[1], p[0]) {
			count++
		}
	}
	return count
}

func isInInterval(a []int, b []int) bool {
	return a[0] <= b[0] && b[1] <= a[1]
}

func overlap(a []int, b []int) bool {
	return isInInterval(a, b) || a[0] <= b[0] && a[1] >= b[0]
}
