package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Hardcode everything
var side = 50
var dirs = []string{">", "v", "<", "^"}
var dirIndex = map[string]int{">": 0, "v": 1, "<": 2, "^": 3}
var orientation = map[int]map[string]Orient{
	0: {">": {1, ">"},
		"v": {2, "v"},
		"<": {3, ">"},
		"^": {5, ">"}},

	1: {">": {4, "<"},
		"v": {2, "<"},
		"<": {0, "<"},
		"^": {5, "^"}},

	2: {">": {1, "^"},
		"v": {4, "v"},
		"<": {3, "v"},
		"^": {0, "^"}},

	3: {">": {4, ">"},
		"v": {5, "v"},
		"<": {0, ">"},
		"^": {2, ">"}},

	4: {">": {1, "<"},
		"v": {5, "<"},
		"<": {3, "<"},
		"^": {2, "^"}},

	5: {">": {4, "^"},
		"v": {1, "v"},
		"<": {0, "v"},
		"^": {3, "^"}},
}

var offset = map[int]Coord{
	0: {0, side},
	1: {0, 2 * side},
	2: {side, side},
	3: {2 * side, 0},
	4: {2 * side, side},
	5: {3 * side, 0},
}

func main() {
	f, _ := os.ReadFile("./input_cube.txt")
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func partTwo(input string) int {
	steps, turns, faces := parseInput(input)
	start := Position{Coord{0, 0}, Orient{0, ">"}}
	curr := start
	for i := 0; i < len(turns); i++ {
		curr = move(curr, steps[i], faces)
		curr = turn(curr, turns[i])
	}
	curr = move(curr, steps[len(steps)-1], faces)

	offsetCoord := offset[curr.o.face]
	return (curr.c.r+1+offsetCoord.r)*1000 + (curr.c.c+1+offsetCoord.c)*4 + dirIndex[curr.o.dir]
}

type Position struct {
	c Coord
	o Orient
}

type Orient struct {
	face int
	dir  string
}

type Coord struct {
	r int
	c int
}

func rotate(oldO, newO Orient, pos Position) Position {
	rotations := mod(dirIndex[newO.dir]-dirIndex[oldO.dir], len(dirs))

	for r := 0; r < rotations; r++ {
		pos = rotateClockwise(pos)
	}

	return pos
}

func rotateClockwise(pos Position) Position {
	side := 50
	r, c := pos.c.r, pos.c.c
	pos.c.r = c
	pos.c.c = side - r - 1
	return pos
}

func orient(pos Position, face int) Position {
	oldO := pos.o
	newO := orientation[pos.o.face][pos.o.dir]
	pos = rotate(oldO, newO, pos)
	pos.o = newO
	return pos
}

func move(pos Position, step int, faces [][]string) Position {
	side := len(faces[0])

	for step > 0 {
		face := faces[pos.o.face]
		var next string
		var r, c int
		if pos.o.dir == ">" || pos.o.dir == "<" {
			if pos.o.dir == ">" {
				c = pos.c.c + 1
			} else if pos.o.dir == "<" {
				c = pos.c.c - 1
			}

			if c >= 0 && c < side {
				next = string(face[pos.c.r][c])
				if next == "." {
					pos.c.c = c
					step--
				} else if next == "#" {
					break
				}
			} else {
				newPos := Position{Coord{pos.c.r, mod(c, side)}, Orient{pos.o.face, pos.o.dir}}
				newPos = orient(newPos, pos.o.face)
				newFace := faces[newPos.o.face]
				newNext := string(newFace[newPos.c.r][newPos.c.c])
				if newNext == "." {
					pos = newPos
					step--
				} else if newNext == "#" {
					break
				}
			}
		} else if pos.o.dir == "v" || pos.o.dir == "^" {
			if pos.o.dir == "v" {
				r = pos.c.r + 1
			} else if pos.o.dir == "^" {
				r = pos.c.r - 1
			}

			if r >= 0 && r < side {
				next = string(face[r][pos.c.c])
				if next == "." {
					pos.c.r = r
					step--
				} else if next == "#" {
					break
				}
			} else {
				newPos := Position{Coord{mod(r, side), pos.c.c}, Orient{pos.o.face, pos.o.dir}}
				newPos = orient(newPos, pos.o.face)
				newFace := faces[newPos.o.face]
				newNext := string(newFace[newPos.c.r][newPos.c.c])
				if newNext == "." {
					pos = newPos
					step--
				} else if newNext == "#" {
					break
				}
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

func turn(pos Position, turn string) Position {
	index := 0
	for i, dir := range dirs {
		if dir == pos.o.dir {
			index = i
			break
		}
	}

	if turn == "R" {
		pos.o.dir = dirs[mod(index+1, len(dirs))]
		return pos
	}

	if turn == "L" {
		pos.o.dir = dirs[mod(index-1, len(dirs))]
		return pos
	}

	return pos
}

func parseInput(input string) ([]int, []string, [][]string) {
	faces := [][]string{}
	isNum := map[rune]bool{}
	for _, num := range "0123456789" {
		isNum[num] = true
	}

	groups := strings.Split(input, "\n\n")
	for _, group := range groups[:len(groups)-1] {
		lines := strings.Split(group, "\n")
		faces = append(faces, lines)
	}

	steps := []int{}
	steps_ := strings.FieldsFunc(groups[len(groups)-1], func(r rune) bool {
		return r == 'L' || r == 'R'
	})
	for _, step := range steps_ {
		stepCount, err := strconv.ParseInt(step, 10, 64)
		if err != nil {
			panic(err)
		}

		steps = append(steps, int(stepCount))
	}

	turns := strings.FieldsFunc(groups[len(groups)-1], func(r rune) bool {
		return isNum[r]
	})

	return steps, turns, faces
}
