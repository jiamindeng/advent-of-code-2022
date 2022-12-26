package main

import (
	"fmt"
	"os"
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

func partOne(input []string) int {
	return 0
}

func partTwo(input []string) int {
	return 0
}

func parseInput(f string) []string {
	lines := strings.Split(f, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	return lines
}
