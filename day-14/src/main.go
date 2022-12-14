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

func partOne(input string) int {
	rockMap := []map[string]string{}
	rocks := parseInput(input)
	for index, rock := range rocks {
		rockMap = append(rockMap, map[string]string{})
		for i := range rock[:len(rock)-1] {
			coordA := rock[i]
			coordB := rock[i+1]
			markRock(coordA, coordB, index, rockMap)
		}
	}

	gridMap := flattenSlice(rockMap)
	points := mapToSlice(gridMap)
	minPoint, maxPoint := getBounds(points)
	sandSpawn := &Coordinate{0, 500}

	curr := sandSpawn
	count := 0
	for curr.c >= minPoint.c && curr.c <= maxPoint.c && curr.r < maxPoint.r {
		curr = sandSpawn
		settled := false
		for !settled && curr.c >= minPoint.c && curr.c <= maxPoint.c && curr.r < maxPoint.r {
			curr, settled = cycle(*curr, gridMap)
		}
		count++
	}

	return count - 1
}

func partTwo(input string) int {
	rockMap := []map[string]string{}
	rocks := parseInput(input)
	for index, rock := range rocks {
		rockMap = append(rockMap, map[string]string{})
		for i := range rock[:len(rock)-1] {
			coordA := rock[i]
			coordB := rock[i+1]
			markRock(coordA, coordB, index, rockMap)
		}
	}

	gridMap := flattenSlice(rockMap)
	points := mapToSlice(gridMap)
	minPoint, maxPoint := getBounds(points)
	grid := pointsToGrid(points, maxPoint)
	drawGrid(grid, minPoint, maxPoint)

	sandSpawn := &Coordinate{0, 500}

	curr := sandSpawn
	count := 0
	for key := fmt.Sprintf("%d,%d", curr.r, curr.c); gridMap[key] != "o"; {
		curr = sandSpawn
		settled := false
		for !settled {
			curr, settled = cycleTwo(*curr, gridMap, maxPoint.r+2)
		}
		count++

	}

	points = mapToSlice(gridMap)
	minPoint, maxPoint = getBounds(points)
	grid = pointsToGrid(points, maxPoint)
	drawGrid(grid, minPoint, maxPoint)

	return count
}

func cycle(curr Coordinate, gridMap map[string]string) (*Coordinate, bool) {
	key := fmt.Sprintf("%d,%d", curr.r, curr.c)
	gridMap[key] = ""

	key = fmt.Sprintf("%d,%d", curr.r+1, curr.c)
	if gridMap[key] == "" {
		gridMap[key] = "+"
		return &Coordinate{curr.r + 1, curr.c}, false
	}

	key = fmt.Sprintf("%d,%d", curr.r+1, curr.c-1)
	if gridMap[key] == "" {
		gridMap[key] = "+"
		return &Coordinate{curr.r + 1, curr.c - 1}, false
	}

	key = fmt.Sprintf("%d,%d", curr.r+1, curr.c+1)
	if gridMap[key] == "" {
		gridMap[key] = "+"
		return &Coordinate{curr.r + 1, curr.c + 1}, false
	}

	key = fmt.Sprintf("%d,%d", curr.r, curr.c)
	gridMap[key] = "o"
	return &curr, true
}

func parseInput(input string) [][]Coordinate {
	rocks := [][]Coordinate{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		rock := []Coordinate{}
		coords := strings.Split(line, " -> ")
		for _, coord := range coords {
			var c Coordinate
			fmt.Sscanf(coord, "%d,%d", &c.c, &c.r)
			rock = append(rock, c)
		}
		rocks = append(rocks, rock)
	}

	return rocks
}

func cycleTwo(curr Coordinate, gridMap map[string]string, floor int) (*Coordinate, bool) {
	key := fmt.Sprintf("%d,%d", curr.r, curr.c)
	if gridMap[key] == "o" {
		return &curr, true
	}
	gridMap[key] = ""

	if curr.r+1 == floor {
		gridMap[key] = "o"
		return &curr, true
	}

	key = fmt.Sprintf("%d,%d", curr.r+1, curr.c)
	if gridMap[key] == "" {
		gridMap[key] = "+"
		return &Coordinate{curr.r + 1, curr.c}, false
	}

	key = fmt.Sprintf("%d,%d", curr.r+1, curr.c-1)
	if gridMap[key] == "" {
		gridMap[key] = "+"
		return &Coordinate{curr.r + 1, curr.c - 1}, false
	}

	key = fmt.Sprintf("%d,%d", curr.r+1, curr.c+1)
	if gridMap[key] == "" {
		gridMap[key] = "+"
		return &Coordinate{curr.r + 1, curr.c + 1}, false
	}

	key = fmt.Sprintf("%d,%d", curr.r, curr.c)
	gridMap[key] = "o"
	return &curr, true
}

type Coordinate struct {
	r int
	c int
}

func markRock(a Coordinate, b Coordinate, index int, grid []map[string]string) {
	vertical := b.r - a.r
	horizontal := b.c - a.c

	if vertical != 0 {
		if vertical > 0 {
			for i := 0; i < vertical+1; i++ {
				key := fmt.Sprintf("%d,%d", a.r+i, a.c)
				grid[index][key] = "#"
			}
		}

		if vertical < 0 {
			for i := 0; i > vertical-1; i-- {
				key := fmt.Sprintf("%d,%d", a.r+i, a.c)
				grid[index][key] = "#"
			}
		}
	}

	if horizontal != 0 {
		if horizontal > 0 {
			for i := 0; i < horizontal+1; i++ {
				key := fmt.Sprintf("%d,%d", a.r, a.c+i)
				grid[index][key] = "#"
			}
		}

		if horizontal < 0 {
			for i := 0; i > horizontal-1; i-- {
				key := fmt.Sprintf("%d,%d", a.r, a.c+i)
				grid[index][key] = "#"
			}
		}
	}
}

func getBounds(points []Point) (Coordinate, Coordinate) {
	maxRow := -2147483647
	minRow := 2147483647
	maxCol := -2147483647
	minCol := 2147483647

	for _, point := range points {
		maxRow = max(maxRow, point.coord.r)
		minRow = min(minRow, point.coord.r)
		maxCol = max(maxCol, point.coord.c)
		minCol = min(minCol, point.coord.c)
	}

	return Coordinate{minRow, minCol}, Coordinate{maxRow, maxCol}
}

func pointsToGrid(points []Point, maxPoint Coordinate) [][]string {
	grid := [][]string{}
	for i := 0; i < maxPoint.r+1; i++ {
		row := make([]string, maxPoint.c+1)
		grid = append(grid, row)
	}

	for _, point := range points {
		grid[point.coord.r][point.coord.c] = point.val
	}

	return grid
}

type Point struct {
	coord Coordinate
	val   string
}

func mapToSlice(rockMap map[string]string) []Point {
	points := []Point{}

	for k, v := range rockMap {
		var coord Coordinate
		fmt.Sscanf(k, "%d,%d", &coord.r, &coord.c)
		points = append(points, Point{coord, v})
	}

	return points
}

func drawGrid(grid [][]string, minPoint Coordinate, maxPoint Coordinate) {
	for i := minPoint.r; i < maxPoint.r+1; i++ {
		for j := minPoint.c; j < maxPoint.c+1; j++ {
			item := grid[i][j]
			if item == "" {
				fmt.Print(".")
			} else {
				fmt.Print(item)
			}
		}
		fmt.Println()
	}
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func flattenSlice(rockMap []map[string]string) map[string]string {
	result := map[string]string{}
	for _, rock := range rockMap {
		for k, v := range rock {
			result[k] = v
		}
	}
	return result
}
