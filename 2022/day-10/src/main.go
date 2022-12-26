package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	fmt.Println(string(f))
	resOne := partOne(string(f))
	fmt.Println(resOne)
	partTwo(string(f))
}

func partOne(input string) int {
	cycle := 1
	signal := 1
	history := map[int]int{}
	ops := strings.Split(input, "\n")
	for _, op := range ops {
		var instruction string
		var dSignal int
		fmt.Sscanf(op, "%s %d", &instruction, &dSignal)
		if instruction == "addx" {
			if cycle%40 == 20 {
				history[cycle] = signal
			}
			cycle += 1
			if cycle%40 == 20 {
				history[cycle] = signal
			}
			cycle += 1
			signal += dSignal
		} else if instruction == "noop" {
			if cycle%40 == 20 {
				history[cycle] = signal
			}
			cycle += 1
		}

	}

	signalStrength := 0
	for cycleNumber, signal := range history {
		signalStrength += cycleNumber * signal
	}

	return signalStrength
}

func partTwo(input string) {
	res := []string{}
	cycle := 0
	position := 1
	ops := strings.Split(input, "\n")
	for _, op := range ops {
		var instruction string
		var dPosition int
		fmt.Sscanf(op, "%s %d", &instruction, &dPosition)
		if instruction == "addx" {
			drawPixel(position, cycle, &res)
			cycle += 1
			drawPixel(position, cycle, &res)
			cycle += 1
			position += dPosition
		} else if instruction == "noop" {
			drawPixel(position, cycle, &res)
			cycle += 1
		}
	}

	for i := range res {
		if i%40 == 0 {
			fmt.Println(res[i : i+40])
		}
	}
}

func incrementSignal(cycle int, signal int, dSignal int, history map[int]int) (int, int) {
	cycle += 1
	signal += dSignal
	return cycle, signal
}

func overlaps(position int, pixel int) bool {
	ds := []int{-1, 0, 1}
	for _, d := range ds {
		if position+d == pixel%40 {
			return true
		}
	}
	return false
}

func drawPixel(position int, pixel int, res *[]string) {
	if overlaps(position, pixel) {
		*res = append(*res, "#")
	} else {
		*res = append(*res, ".")
	}
}
