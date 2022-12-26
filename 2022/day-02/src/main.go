package main

import (
	"fmt"
	"os"
	"strings"
)

// Map letter move to English
var moves = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var matchPoints = map[string]int{
	"win":  6,
	"draw": 3,
	"loss": 0,
}

var outcome = map[string]string{
	"X": "loss",
	"Y": "draw",
	"Z": "win",
}

var movePoints = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

// From POV of player whose move is represented by value, not key
var statusToMove = map[string]map[string]string{
	"rock": {
		"win":  "paper",
		"draw": "rock",
		"loss": "scissors",
	},
	"paper": {
		"win":  "scissors",
		"draw": "paper",
		"loss": "rock",
	},
	"scissors": {
		"win":  "rock",
		"draw": "scissors",
		"loss": "paper",
	},
}

var moveToStatus = map[string]map[string]string{
	"rock": {
		"paper":    "win",
		"rock":     "draw",
		"scissors": "loss",
	},
	"paper": {
		"scissors": "win",
		"paper":    "draw",
		"rock":     "loss",
	},
	"scissors": {
		"rock":     "win",
		"scissors": "draw",
		"paper":    "loss",
	},
}

func main() {
	f, _ := os.ReadFile("./input.txt")
	fmt.Println(string(f))
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func partOne(input string) int {
	totalScore := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		match := strings.Split(line, " ")
		self := moves[match[1]]
		opponent := moves[match[0]]
		status := moveToStatus[opponent][self]
		totalScore += matchPoints[status] + movePoints[self]
	}
	return totalScore
}

func partTwo(input string) int {
	totalScore := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		match := strings.Split(line, " ")
		gameResult := outcome[match[1]]
		opponent := moves[match[0]]
		self := statusToMove[opponent][gameResult]
		totalScore += matchPoints[gameResult] + movePoints[self]
	}
	return totalScore
}
