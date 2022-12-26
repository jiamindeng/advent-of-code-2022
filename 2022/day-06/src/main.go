package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func helper(input string, numUnique int) int {
	chars := strings.Split(input, "")
	left := 0
	seen := map[string]int{}
	for right := 0; right >= left && right < len(chars); right++ {
		count, ok := seen[chars[right]]
		for ok && count > 0 {
			seen[chars[left]]--
			left++
			count = seen[chars[right]]
		}

		if right-left+1 == numUnique {
			return right + 1
		}

		seen[chars[right]] += 1
	}

	return -1
}

func partOne(input string) int {
	return helper(input, 4)
}

func partTwo(input string) int {
	return helper(input, 14)
}

// 012345678012
// zcfzfwzzqfrl
