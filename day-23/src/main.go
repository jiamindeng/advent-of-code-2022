package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(parseInput(string(f)))
	fmt.Println(resOne)
	resTwo := partTwo(parseInput(string(f)))
	fmt.Println(resTwo)
}

func partOne(elves map[string]bool) int {
	order := []string{"n", "s", "w", "e"}
	for i := 0; i < 10; i++ {
		newPositions := map[string][]string{}
		for k := range elves {
			var key string
			pos := parseCoordKey(k)
			newPos := decide(pos, elves, order)
			if newPos != nil {
				key = coordKey(*newPos)
				newPositions[key] = append(newPositions[key], coordKey(pos))
			}
		}

		for k, v := range newPositions {
			if len(v) == 1 {
				delete(elves, v[0])
				elves[k] = true
			}
		}

		tmp := order[0]
		order = order[1:]
		order = append(order, tmp)
	}

	draw(elves)

	minPos, maxPos := minMaxCoord(elves)
	return (maxPos.r-minPos.r+1)*(maxPos.c-minPos.c+1) - len(elves)
}

func partTwo(elves map[string]bool) int {
	order := []string{"n", "s", "w", "e"}
	i := 0
	for true {
		added := false
		newPositions := map[string][]string{}
		for k := range elves {
			var key string
			pos := parseCoordKey(k)
			newPos := decide(pos, elves, order)
			if newPos != nil {
				key = coordKey(*newPos)
				newPositions[key] = append(newPositions[key], coordKey(pos))
			}
		}

		for k, v := range newPositions {
			if len(v) == 1 {
				added = true
				delete(elves, v[0])
				elves[k] = true
			}
		}

		if !added {
			fmt.Print()
		}

		tmp := order[0]
		order = order[1:]
		order = append(order, tmp)
		i++

		draw(elves)
	}

	return -1
}

type Coord struct {
	r int
	c int
}

func decide(pos Coord, elves map[string]bool, order []string) *Coord {
	dirs := getDirs()
	north := search(pos, elves, dirs)["n"]()
	south := search(pos, elves, dirs)["s"]()
	west := search(pos, elves, dirs)["w"]()
	east := search(pos, elves, dirs)["e"]()

	if !north && !south && !west && !east {
		return nil
	}

	check := map[string]func() *Coord{
		"n": func() *Coord {
			if !north {
				return &Coord{pos.r - 1, pos.c}
			}
			return nil
		},
		"s": func() *Coord {
			if !south {
				return &Coord{pos.r + 1, pos.c}
			}
			return nil
		},
		"w": func() *Coord {
			if !west {
				return &Coord{pos.r, pos.c - 1}
			}
			return nil
		},
		"e": func() *Coord {
			if !east {
				return &Coord{pos.r, pos.c + 1}
			}
			return nil
		},
	}

	for _, dir := range order {
		newPos := check[dir]()
		if newPos != nil {
			return newPos
		}
	}

	return nil
}

func search(pos Coord, elves map[string]bool, dirs []Coord) map[string]func() bool {
	hasElf := map[string]func() bool{
		"n": func() bool {
			for _, d := range dirs {
				if d.r == -1 {
					coord := Coord{d.r + pos.r, d.c + pos.c}
					if elves[coordKey(coord)] {
						return true
					}
				}
			}
			return false
		},
		"e": func() bool {
			for _, d := range dirs {
				if d.c == 1 {
					coord := Coord{d.r + pos.r, d.c + pos.c}
					if elves[coordKey(coord)] {
						return true
					}
				}
			}
			return false
		},
		"s": func() bool {
			for _, d := range dirs {
				if d.r == 1 {
					coord := Coord{d.r + pos.r, d.c + pos.c}
					if elves[coordKey(coord)] {
						return true
					}
				}
			}
			return false
		},
		"w": func() bool {
			for _, d := range dirs {
				if d.c == -1 {
					coord := Coord{d.r + pos.r, d.c + pos.c}
					if elves[coordKey(coord)] {
						return true
					}
				}
			}
			return false
		},
	}

	return hasElf
}

func draw(elves map[string]bool) {
	str := ""
	minPos, maxPos := minMaxCoord(elves)

	for i := minPos.r; i <= maxPos.r; i++ {
		for j := minPos.c; j <= maxPos.c; j++ {
			if elves[coordKey(Coord{i, j})] {
				str += "🟩"
			} else {
				str += "⬛"
			}
		}
		str += "\n"
	}

	fmt.Printf("\033[%d;%dH", 2, 0)
	fmt.Print(str)
	time.Sleep(1 * time.Millisecond)
}

func minMaxCoord(elves map[string]bool) (Coord, Coord) {
	minPos := Coord{math.MaxInt, math.MaxInt}
	maxPos := Coord{math.MinInt, math.MinInt}
	for k := range elves {
		parsed := parseCoordKey(k)
		minPos.c = min(minPos.c, parsed.c)
		minPos.r = min(minPos.r, parsed.r)
		maxPos.c = max(maxPos.c, parsed.c)
		maxPos.r = max(maxPos.r, parsed.r)
	}
	return minPos, maxPos
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func coordKey(c Coord) string {
	return fmt.Sprintf("%d,%d", c.r, c.c)
}

func parseCoordKey(key string) Coord {
	var coord Coord
	_, err := fmt.Sscanf(key, "%d,%d", &coord.r, &coord.c)
	if err != nil {
		panic(err)
	}
	return coord
}

func parseInput(f string) map[string]bool {
	elves := map[string]bool{}
	rows := strings.Split(f, "\n")
	for r, col := range rows {
		for c, char := range col {
			if string(char) == "#" {
				elves[coordKey(Coord{r, c})] = true
			}
		}
	}
	return elves
}

func getDirs() []Coord {
	dirs := []Coord{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				dirs = append(dirs, Coord{i, j})
			}
		}
	}
	return dirs
}
