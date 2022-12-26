package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

type Point struct {
	x int
	y int
}

func manhattanDist(a Point, b Point) int {
	return int(math.Abs(float64(b.x-a.x)) + math.Abs(float64(b.y-a.y)))
}

type Pair struct {
	sensor   Point
	beacon   Point
	distance int
}

func min(curr int, val int) int {
	return int(math.Min(float64(curr), float64(val)))
}

func max(curr int, val int) int {
	return int(math.Max(float64(curr), float64(val)))
}

func partOne(input string) int {
	MAX := 2147483647
	minX, minY, maxX, maxY := MAX, MAX, -MAX, -MAX
	maxDist := -MAX

	all := []Point{}
	pairs := []Pair{}
	beacons := map[string]bool{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var sensor Point
		var beacon Point
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		distance := manhattanDist(sensor, beacon)
		pair := Pair{sensor, beacon, distance}
		pairs = append(pairs, pair)
		minX = min(minX, min(beacon.x, sensor.x))
		maxY = max(maxY, max(beacon.y, sensor.y))
		minY = min(minY, min(beacon.y, sensor.y))
		maxX = max(minX, max(beacon.x, sensor.x))
		maxDist = max(maxDist, distance)
		beaconKey := fmt.Sprintf("%d,%d", beacon.x, beacon.y)
		beacons[beaconKey] = true
		all = append(all, beacon, sensor)
	}

	y := 2000000

	count := 0

	for x := minX - 2*maxDist; x <= maxX+2*maxDist; x++ {
		for _, pair := range pairs {
			current := Point{x, y}
			distance := manhattanDist(current, pair.sensor)
			beaconKey := fmt.Sprintf("%d,%d", current.x, current.y)
			if !beacons[beaconKey] && distance <= pair.distance {
				count++
				break
			}
		}
	}

	return count
}

func partTwo(input string) int {
	intervals := map[int][]Interval{}
	all := []Point{}
	pairs := []Pair{}
	beacons := map[string]bool{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var sensor Point
		var beacon Point
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		distance := manhattanDist(sensor, beacon)
		pair := Pair{sensor, beacon, distance}
		pairs = append(pairs, pair)
		beaconKey := fmt.Sprintf("%d,%d", beacon.x, beacon.y)
		beacons[beaconKey] = true
		all = append(all, beacon, sensor)
		for dx := 0; dx <= distance; dx++ {
			dy := distance - dx
			intervals[sensor.y-dy] = append(intervals[sensor.y-dy], Interval{sensor.x - dx, sensor.x + dx})
			intervals[sensor.y+dy] = append(intervals[sensor.y+dy], Interval{sensor.x - dx, sensor.x + dx})
		}

	}

	for y, row := range intervals {
		sort.Slice(row, func(i, j int) bool {
			return row[i].left < row[j].left
		})
		merged := []Interval{row[0]}
		for _, interval := range row[1:] {
			if interval.left-merged[len(merged)-1].right > 1 {
				merged = append(merged, interval)
			} else {
				merged[len(merged)-1].right = max(merged[len(merged)-1].right, interval.right)
			}
		}
		intervals[y] = merged
	}

	for y := 0; y <= 4000000; y++ {
		row := intervals[y]
		if len(row) == 2 {
			return y + 4000000*(row[0].right+1)
		}
	}

	return -1
}

type Interval struct {
	left  int
	right int
}
