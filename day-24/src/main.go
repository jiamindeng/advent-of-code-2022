package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	plot, blizzards := parseInput(string(f))
	cache := precomp(plot, blizzards)
	resOne := partOne(plot, blizzards, cache)
	fmt.Println(resOne)
	resTwo := partTwo(plot, blizzards, cache)
	fmt.Println(resTwo)
}

func partOne(plot [][]map[string]int, blizzards []blizzard, cache map[int]state) int {
	endCoord := coord{21, 150}
	count, blizzards := simulate(0, coord{0, 1}, endCoord, plot, blizzards, cache)
	return count - 1
}

func partTwo(plot [][]map[string]int, blizzards []blizzard, cache map[int]state) int {
	endCoord := coord{21, 150}
	count, blizzards := simulate(0, coord{0, 1}, endCoord, plot, blizzards, cache)
	count, blizzards = simulate(count, endCoord, coord{0, 1}, plot, blizzards, cache)
	count, blizzards = simulate(count, coord{0, 1}, endCoord, plot, blizzards, cache)
	return count - 1
}

type coord struct {
	r int
	c int
}

type blizzard struct {
	pos coord
	dir string
}

type state struct {
	plot      [][]map[string]int
	blizzards []blizzard
}

func canMove(pos coord, plot [][]map[string]int) bool {
	check := []string{"<", ">", "^", "v", "#"}
	for _, item := range check {
		if plot[pos.r][pos.c][item] > 0 {
			return false
		}
	}
	return true
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}

func moveBliz(bliz blizzard, plot [][]map[string]int) ([][]map[string]int, blizzard) {
	row := len(plot)
	col := len(plot[0])
	move := map[string]func(coord) coord{
		">": func(pos coord) coord {
			return coord{pos.r, mod(pos.c+1, col)}
		},
		"v": func(pos coord) coord {
			return coord{mod(pos.r+1, row), pos.c}
		},
		"<": func(pos coord) coord {
			return coord{pos.r, mod(pos.c-1, col)}
		},
		"^": func(pos coord) coord {
			return coord{mod(pos.r-1, row), pos.c}
		},
	}
	newPos := move[bliz.dir](bliz.pos)
	next := plot[newPos.r][newPos.c]
	if next["#"] < 0 {
		plot[newPos.r][newPos.c][bliz.dir] += 1
		plot[bliz.pos.r][bliz.pos.c][bliz.dir] -= 1
		if plot[bliz.pos.r][bliz.pos.c][bliz.dir] == 0 {
			delete(plot[bliz.pos.r][bliz.pos.c], bliz.dir)
		}
		return plot, blizzard{newPos, bliz.dir}
	} else {
		for next["#"] > 0 {
			newPos = move[bliz.dir](newPos)
			next = plot[newPos.r][newPos.c]
		}
		plot[newPos.r][newPos.c][bliz.dir] += 1
		plot[bliz.pos.r][bliz.pos.c][bliz.dir] -= 1
		if plot[bliz.pos.r][bliz.pos.c][bliz.dir] == 0 {
			delete(plot[bliz.pos.r][bliz.pos.c], bliz.dir)
		}
	}

	return plot, blizzard{newPos, bliz.dir}
}

func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func moveBlizzards(plot [][]map[string]int, blizzards []blizzard) []blizzard {
	newBlizzards := []blizzard{}
	for _, b := range blizzards {
		var bliz blizzard
		plot, bliz = moveBliz(b, plot)
		newBlizzards = append(newBlizzards, bliz)
	}
	return newBlizzards
}

func precomp(plot [][]map[string]int, blizzards []blizzard) map[int]state {
	cache := map[int]state{}
	lcm := LCM(len(plot), len(plot[0]))
	count := 0

	for count < lcm {
		cacheKey := count % lcm
		cache[cacheKey] = state{copyArray(plot), blizzards}
		blizzards = moveBlizzards(plot, blizzards)
		count++
	}

	return cache
}

func simulate(time int, start coord, target coord, plot [][]map[string]int, blizzards []blizzard, cache map[int]state) (int, []blizzard) {
	visited := map[string]bool{}
	lcm := LCM(len(plot), len(plot[0]))
	count := time
	dirs := []coord{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {0, 0}}
	queue := []coord{start}
	moved := [][]int{}
	for _, row := range plot {
		row_ := []int{}
		for range row {
			row_ = append(row_, 0)
		}
		moved = append(moved, row_)
	}

	for len(queue) > 0 {
		queueLen := len(queue)
		cacheKey := count % lcm
		seen, ok := cache[cacheKey]
		if !ok {
			panic("plot state not cached")
		}
		plot = copyArray(seen.plot)
		blizzards = seen.blizzards

		for queueLen > 0 {
			var curr coord
			curr, queue = queue[0], queue[1:]

			if curr.r == target.r && curr.c == target.c {
				return count, blizzards
			}

			for _, d := range dirs {
				newPos := coord{d.r + curr.r, d.c + curr.c}
				visitedKey := fmt.Sprintf("%d,%s", count%lcm, coordKey(newPos))
				inBounds := newPos.c < len(plot[0]) && newPos.c >= 0 && newPos.r < len(plot) && newPos.r >= 0
				if inBounds && canMove(newPos, plot) && !visited[visitedKey] {
					queue = append(queue, newPos)
					moved[newPos.r][newPos.c] = count
					visited[visitedKey] = true
				}
			}
			queueLen--
		}
		count++
	}

	return count, blizzards
}

func copyArray(a [][]map[string]int) [][]map[string]int {
	aCopy := [][]map[string]int{}
	for _, row := range a {
		rowCopy := []map[string]int{}
		for _, col := range row {
			colCopy := map[string]int{}
			for k, v := range col {
				colCopy[k] = v
			}
			rowCopy = append(rowCopy, colCopy)
		}
		aCopy = append(aCopy, rowCopy)
	}

	return aCopy
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func drawMoved(moved [][]int) {
	for _, row := range moved {
		for _, col := range row {
			fmt.Print(col % 10)
		}
		fmt.Println()
	}
}

func draw(plot [][]map[string]int) {
	str := ""
	for _, row := range plot {
		for _, cell := range row {
			if len(cell) == 0 {
				str += "."
			} else if len(cell) == 1 {
				for k := range cell {
					str += k
				}
			} else {
				total := 0
				for _, v := range cell {
					total += v
				}
				str += fmt.Sprint(total)
			}
		}
		str += "\n"
	}
	fmt.Println(str)
}

func parseInput(f string) ([][]map[string]int, []blizzard) {
	lines := strings.Split(f, "\n")
	blizzards := []blizzard{}
	plot := [][]map[string]int{}
	for r, line := range lines {
		row := []map[string]int{}
		for c, char := range line {
			col := map[string]int{}
			if char != '.' {
				col[string(char)] += 1
			}
			if char != '#' && char != '.' {
				blizzards = append(blizzards, blizzard{coord{r, c}, string(char)})
			}
			row = append(row, col)
		}
		plot = append(plot, row)
	}

	return plot, blizzards
}

func coordKey(c coord) string {
	return fmt.Sprintf("%d,%d", c.r, c.c)
}

func parseCoord(s string) coord {
	var c coord
	_, err := fmt.Sscanf(s, "%d,%d", &c.r, &c.c)
	if err != nil {
		panic(err)
	}
	return c
}

func blizKey(b blizzard) string {
	return fmt.Sprintf("%d,%d,%s", b.pos.r, b.pos.c, b.dir)
}

func parseBliz(s string) blizzard {
	var b blizzard
	fmt.Sscanf(s, "%d,%d,%s", &b.pos.r, &b.pos.c, &b.dir)
	return b
}
