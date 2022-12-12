package main

import (
	"fmt"
	"math"
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

type Coordinate struct {
	r int
	c int
}

func charValue(target string) int {
	index := "abcdefghijklmnopqrstuvwxyzE"
	for i, char := range index {
		if string(char) == target {
			return i
		}
	}

	return -1
}

func partOne(input string) int {
	var start Coordinate
	lines := strings.Split(input, "\n")
	for r, line := range lines {
		for c, char := range line {
			if string(char) == "S" {
				start = Coordinate{r, c}
			}
		}
	}
	return helper(start, lines, false)
}

func partTwo(input string) int {
	end := Coordinate{}
	lines := strings.Split(input, "\n")
	for r, line := range lines {
		for c, char := range line {
			if string(char) == "E" {
				end = Coordinate{r, c}
			}
		}
	}

	return helper(end, lines, true)
}

func helper(start Coordinate, lines []string, fromEnd bool) int {
	neighboring := [4]Coordinate{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	count := 0
	queue := []Coordinate{start}
	visited := map[string]bool{fmt.Sprintf("%d,%d", start.r, start.c): true}

	for len(queue) > 0 {
		for _, node := range queue {
			curr := node
			queue = queue[1:]

			if fromEnd {
				if string(lines[curr.r][curr.c]) == "a" || string(lines[curr.r][curr.c]) == "S" {
					return count
				}
			} else {
				if string(lines[curr.r][curr.c]) == "E" {
					return count
				}
			}

			for _, delta := range neighboring {
				neighbor := Coordinate{curr.r + delta.r, curr.c + delta.c}
				neighborKey := fmt.Sprintf("%d,%d", neighbor.r, neighbor.c)
				inRange := neighbor.r < len(lines) && neighbor.r >= 0 && neighbor.c < len(lines[0]) && neighbor.c >= 0
				canJump := false
				if inRange {
					if fromEnd {
						canJump = charValue(string(lines[curr.r][curr.c]))-charValue(string(lines[neighbor.r][neighbor.c])) <= 1
					} else {
						canJump = charValue(string(lines[neighbor.r][neighbor.c]))-charValue(string(lines[curr.r][curr.c])) <= 1
					}
				}
				if canJump && !visited[neighborKey] {
					visited[fmt.Sprintf("%d,%d", neighbor.r, neighbor.c)] = true
					queue = append(queue, neighbor)
				}

			}
		}
		count++
	}

	return math.MaxInt32
}
