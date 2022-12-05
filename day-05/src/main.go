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

type Move struct {
	count  int
	source int
	target int
}

func partOne(input string) string {
	return helper(input, true)
}

func partTwo(input string) string {
	return helper(input, false)
}

func helper(input string, reverse bool) string {
	boxLine := 9
	numStacks := 9
	boxes := [][]string{}

	lines := strings.Split(input, "\n")

	for i := 0; i < numStacks; i++ {
		boxes = append(boxes, []string{})
	}

	for _, line := range lines[:boxLine-1] {
		for i := 0; i < len(line)-2; i += 4 {
			if line[i:i+3] != "   " {
				boxes[i/4] = append([]string{string(line[i+1])}, boxes[i/4]...)
			}
		}
	}

	for _, line := range lines[boxLine+1:] {
		var move Move
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &move.count, &move.source, &move.target)
		if err != nil {
			fmt.Println(err)
		}

		move.source -= 1
		move.target -= 1

		stackSize := len(boxes[move.source])
		toMove := boxes[move.source][stackSize-move.count : stackSize]
		if reverse {
			for i, j := 0, len(toMove)-1; i < j; i, j = i+1, j-1 {
				toMove[i], toMove[j] = toMove[j], toMove[i]
			}
		}
		boxes[move.source] = boxes[move.source][:stackSize-move.count]
		boxes[move.target] = append(boxes[move.target], toMove...)
	}

	var res string
	for _, stack := range boxes {
		res += stack[len(stack)-1]
	}

	return res
}
