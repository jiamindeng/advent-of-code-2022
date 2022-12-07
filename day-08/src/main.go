package main

import (
	"fmt"
	"os"
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
	return 0
}

func partTwo(input string) int {
	return 0
}
