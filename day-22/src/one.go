package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var dirs = []string{">", "v", "<", "^"}
var dirIndex = map[string]int{">": 0, "v": 1, "<": 2, "^": 3}

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
}

func partOne(input string) int {
	steps, turns, map_, rowBounds, colBounds := parseInput(input)
	start := getStartPos(map_)
	curr := start
	for i := 0; i < len(steps)-1; i++ {
		curr = move(curr, steps[i], map_, rowBounds[curr.r], colBounds[curr.c])
		curr = turn(curr, turns[i])
	}
	curr = move(curr, steps[len(steps)-1], map_, rowBounds[curr.r], colBounds[curr.c])
	return (curr.r+1)*1000 + (curr.c+1)*4 + dirIndex[curr.dir]
}

type bound struct {
	min int
	max int
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func parseInput(input string) ([]int, []string, []string, []bound, []bound) {
	lines := strings.Split(input, "\n\n")
	rows := strings.Split(lines[0], "\n")
	isNum := map[rune]bool{}

	for _, num := range "0123456789" {
		isNum[num] = true
	}

	steps := []int{}
	steps_ := strings.FieldsFunc(lines[1], func(r rune) bool {
		return r == 'L' || r == 'R'
	})
	for _, step := range steps_ {
		stepCount, err := strconv.ParseInt(step, 10, 64)
		if err != nil {
			panic(err)
		}

		steps = append(steps, int(stepCount))
	}

	turns := strings.FieldsFunc(lines[1], func(r rune) bool {
		return isNum[r]
	})

	rowBounds := []bound{}
	colBounds := []bound{}
	maxRowLen := 0

	for _, row := range rows {
		maxRowLen = max(maxRowLen, len(row))
		var min, max int
		for c := 0; c < len(row); c++ {
			if string(row[c]) != " " {
				min = c
				break
			}
		}
		for c := len(row) - 1; c >= 0; c-- {
			if string(row[c]) != " " {
				max = c
				break
			}
		}
		rowBounds = append(rowBounds, bound{min, max})
	}

	for r := 0; r < len(rows); r++ {
		diff := maxRowLen - len(rows[r])
		for diff > 0 {
			rows[r] += " "
			diff--
		}
	}

	for c := 0; c < maxRowLen; c++ {
		var min, max int
		for r := 0; r < len(rows); r++ {
			row := rows[r]
			if string(row[c]) != " " {
				min = r
				break
			}
		}
		for r := len(rows) - 1; r >= 0; r-- {
			row := rows[r]
			if string(row[c]) != " " {
				max = r
				break
			}
		}
		colBounds = append(colBounds, bound{min, max})
	}

	return steps, turns, rows, rowBounds, colBounds
}

type Position struct {
	r   int
	c   int
	dir string
}

func turn(pos Position, turn string) Position {
	index := 0
	for i, dir := range dirs {
		if dir == pos.dir {
			index = i
			break
		}
	}

	if turn == "R" {
		pos.dir = dirs[mod(index+1, len(dirs))]
		return pos
	}

	if turn == "L" {
		pos.dir = dirs[mod(index-1, len(dirs))]
		return pos
	}

	return pos
}

func move(pos Position, step int, map_ []string, rowBound bound, colBound bound) Position {
	height := colBound.max - colBound.min + 1
	width := rowBound.max - rowBound.min + 1

	for step > 0 {
		var next string
		var r, c int
		if pos.dir == "<" || pos.dir == ">" {
			if pos.dir == ">" {
				c = mod(pos.c+1-rowBound.min, width) + rowBound.min
			} else if pos.dir == "<" {
				c = mod(pos.c-1-rowBound.min, width) + rowBound.min
			}

			next = string(map_[pos.r][c])
			if next == "." {
				pos.c = c
				step--
			} else if next == "#" {
				break
			}
		}

		if pos.dir == "v" || pos.dir == "^" {
			if pos.dir == "v" {
				r = mod(pos.r+1-colBound.min, height) + colBound.min
			} else if pos.dir == "^" {
				r = mod(pos.r-1-colBound.min, height) + colBound.min

			}

			next = string(map_[r][pos.c])
			if next == "." {
				pos.r = r
				step--
			} else if next == "#" {
				break
			}
		}
	}

	return pos
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}

func getStartPos(rows []string) Position {
	var start Position
	start.dir = ">"
	for r, row := range rows {
		for c, col := range row {
			if string(col) == "." {
				start.r, start.c = r, c
				return start
			}
		}
	}
	return Position{-1, -1, ""}
}
