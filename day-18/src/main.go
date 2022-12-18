package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var cubes = []Cube{}
var inputSet = map[string]Cube{}

func main() {
	f, _ := os.ReadFile("./input.txt")
	setup(string(f))
	resOne := partOne()
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

type Cube struct {
	x int
	y int
	z int
}

func setup(input string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var cube Cube
		fmt.Sscanf(line, "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		cubes = append(cubes, cube)
		inputSet[line] = cube
	}

}

func partOne() int {
	totalSides := 0

	for _, cube := range cubes {
		totalSides += numSides(cube, inputSet)
	}

	return totalSides
}

func numSides(cube Cube, set map[string]Cube) int {
	sides := 6
	neighbors := []Cube{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	}

	for _, neighbor := range neighbors {
		key := fmt.Sprintf("%d,%d,%d", cube.x+neighbor.x, cube.y+neighbor.y, cube.z+neighbor.z)
		_, ok := set[key]
		if ok {
			sides -= 1
		}
	}

	return sides
}

func numSidesBounding(cube Cube, set map[string]Cube) int {
	sides := 6
	neighbors := []Cube{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	}

	for _, neighbor := range neighbors {
		key := fmt.Sprintf("%d,%d,%d", cube.x+neighbor.x, cube.y+neighbor.y, cube.z+neighbor.z)
		_, ok := set[key]
		if !ok {
			sides -= 1
		}
	}

	return sides
}

func getNeighbors(cube Cube, set map[string]Cube) []Cube {
	adjacent := []Cube{}
	neighbors := []Cube{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	}

	for _, neighbor := range neighbors {
		neighborCube := Cube{cube.x + neighbor.x, cube.y + neighbor.y, cube.z + neighbor.z}
		key := cubeToKey(neighborCube)
		_, ok := set[key]
		if ok {
			adjacent = append(adjacent, neighborCube)
		}
	}

	return adjacent
}

func cubeToKey(cube Cube) string {
	return fmt.Sprintf("%d,%d,%d", cube.x, cube.y, cube.z)
}

func partTwo(input string) int {
	volume := make(map[string]Cube)
	minBoundingCube := getMinBoundingCube(inputSet)
	maxBoundingCube := getMaxBoundingCube(inputSet)

	dX := maxBoundingCube.x - minBoundingCube.x + 3
	dY := maxBoundingCube.y - minBoundingCube.y + 3
	dZ := maxBoundingCube.z - minBoundingCube.z + 3

	outerSurfaceArea := 2 * (dX*dY + dX*dZ + dY*dZ)

	for x := minBoundingCube.x - 1; x <= maxBoundingCube.x+1; x++ {
		for y := minBoundingCube.y - 1; y <= maxBoundingCube.y+1; y++ {
			for z := minBoundingCube.z - 1; z <= maxBoundingCube.z+1; z++ {
				cube := Cube{x, y, z}
				volume[cubeToKey(cube)] = cube
			}
		}
	}

	space := setDiff(volume, inputSet)

	curr := Cube{-1, -1, -1}
	currKey := cubeToKey(curr)
	queue := []Cube{curr}
	visited := map[string]bool{currKey: true}
	spaceSurfaceArea := 0

	count := 0

	for len(queue) > 0 {
		count++
		curr, queue = queue[0], queue[1:]
		neighbors := getNeighbors(curr, space)
		spaceSurfaceArea += numSides(curr, space)

		for _, neighbor := range neighbors {
			neighborKey := cubeToKey(neighbor)
			if !visited[neighborKey] {
				queue = append(queue, neighbor)
				visited[neighborKey] = true
			}

		}
	}

	return spaceSurfaceArea - outerSurfaceArea
}

// Subtract a from b, get what is in a that isn't in b
func setDiff(a map[string]Cube, b map[string]Cube) map[string]Cube {
	res := map[string]Cube{}
	for k, v := range a {
		if _, ok := b[k]; !ok {
			res[k] = v
		}
	}

	return res
}

func getMaxBoundingCube(set map[string]Cube) Cube {
	var maxCube Cube
	for _, cube := range set {
		maxCube.x = max(maxCube.x, cube.x)
		maxCube.y = max(maxCube.y, cube.y)
		maxCube.z = max(maxCube.z, cube.z)
	}

	return maxCube
}

func getMinBoundingCube(set map[string]Cube) Cube {
	var minCube Cube
	for _, cube := range set {
		minCube.x = min(minCube.x, cube.x)
		minCube.y = min(minCube.y, cube.y)
		minCube.z = min(minCube.z, cube.z)
	}

	return minCube
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
