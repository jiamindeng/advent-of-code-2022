package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println("1:", resOne)
	resTwo := partTwo(string(f))
	fmt.Println("2:", resTwo)
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	res := []string{}

	for _, line := range lines {
		firstHalf := []string{}
		secondHalf := []string{}
		middle := -1
		chars := strings.Split(line, "")
		if len(chars)%2 != 0 {
			middle = len(chars)/2 + 1
		} else {
			middle = len(chars) / 2
		}

		for _, charString := range chars[:middle] {
			firstHalf = append(firstHalf, charString)
		}

		for _, charString := range chars[middle:] {
			secondHalf = append(secondHalf, charString)
		}

		res = append(res, findDupe(firstHalf, secondHalf))

	}

	sum := 0
	for _, val := range res {
		char, _ := utf8.DecodeRuneInString(val)
		value := runeToValue(char)
		sum += value
	}
	return sum
}

func runeToValue(r rune) int {
	res := int(r) - 96
	if res < 0 {
		return res + 58
	}
	return res
}

func findDupe(a []string, b []string) string {
	seen := map[string]bool{}

	for _, item := range a {
		seen[item] = true
	}

	for _, item := range b {
		if seen[item] {
			return item
		}
	}

	return ""
}

func findTrip(groups [][]string) string {
	threshold := len(groups)
	seen := map[string]int{}
	fmt.Println(groups)
	for i, group := range groups {
		for _, char := range group {
			_, ok := seen[char]
			if !ok {
				seen[char] = 1
			} else {
				seen[char]++
				if seen[char] != i+1 && seen[char] == threshold {
					fmt.Println(seen)
					return char
				}
			}
		}
	}
	return ""
}

func partTwo(input string) int {
	res := []string{}
	lines := strings.Split(input, "\n")
	groups := [][]string{}
	for i, line := range lines {
		group := strings.Split(line, "")
		groups = append(groups, group)
		if (i+1)%3 == 0 {
			res = append(res, findTrip(groups))
			groups = [][]string{}
		}
	}

	sum := 0
	for _, val := range res {
		char, _ := utf8.DecodeRuneInString(val)
		value := runeToValue(char)
		sum += value
	}
	return sum
}
