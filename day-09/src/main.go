package main

import (
	"fmt"
	"math"
	"os"
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
	rows := strings.Split(input, "\n")
	head := coord{r: 0, c: 0}
	tail := coord{r: 0, c: 0}
	visited := map[string]bool{"0,0": true}

	for _, move := range rows {
		distance := 0
		direction := ""
		fmt.Sscanf(move, "%s %d", &direction, &distance)

		for i := 0; i < distance; i++ {
			head = moveHead[direction](head)
			delta := diff(tail, head)
			if !isAdjacent(delta) {
				tail = moveTail(tail, delta)
				visited[fmt.Sprintf("%d,%d", tail.r, tail.c)] = true
			}
		}
	}

	return len(visited)
}

func partTwo(input string) int {
	rows := strings.Split(input, "\n")
	visited := map[string]bool{"0,0": true}
	rope := map[int]coord{}
	head := rope[0]
	tail := rope[1]

	for _, move := range rows {
		distance := 0
		direction := ""
		fmt.Sscanf(move, "%s %d", &direction, &distance)

		for i := 0; i < distance; i++ {
			head = moveHead[direction](head)
			delta := diff(tail, head)
			rope[0] = head
			if !isAdjacent(delta) {
				tail = moveTail(tail, delta)
				rope[1] = tail
				recurse(rope, rope[2], rope[1], 1, visited)
			}
		}

	}

	return len(visited)
}

type coord struct {
	r int
	c int
}

func isAdjacent(diff coord) bool {
	return math.Abs(float64(diff.r)) <= 1 && math.Abs(float64(diff.c)) <= 1
}

func diff(tail coord, head coord) coord {
	return coord{r: head.r - tail.r, c: head.c - tail.c}
}

func tailDiff(diff coord) coord {
	if diff.c == 0 {
		return coord{r: diff.r / 2, c: 0}
	}

	if diff.r == 0 {
		return coord{r: 0, c: diff.c / 2}
	}

	return coord{r: diff.r / int(math.Abs(float64(diff.r))), c: diff.c / int(math.Abs(float64(diff.c)))}

}

var moveHead = map[string]func(position coord) coord{
	"U": func(position coord) coord {
		return coord{r: position.r + 1, c: position.c}
	},

	"R": func(position coord) coord {
		return coord{r: position.r, c: position.c + 1}
	},

	"D": func(position coord) coord {
		return coord{r: position.r - 1, c: position.c}
	},

	"L": func(position coord) coord {
		return coord{r: position.r, c: position.c - 1}
	},
}

func moveTail(tail coord, delta coord) coord {
	diff := tailDiff(delta)
	return coord{r: tail.r + diff.r, c: tail.c + diff.c}
}

func recurse(rope map[int]coord, tail coord, head coord, depth int, visited map[string]bool) {
	if depth == 9 {
		visited[fmt.Sprintf("%d,%d", head.r, head.c)] = true
		return
	}
	delta := diff(tail, head)
	if !isAdjacent(delta) {
		tail = moveTail(tail, delta)
		depth++
		rope[depth] = tail
		recurse(rope, rope[depth+1], rope[depth], depth, visited)
	}
}
