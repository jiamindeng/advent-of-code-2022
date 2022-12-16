package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var MAX_INT = 2147483647
var valves = map[string]Valve{}

// valve name : bitmap representing that it is toggled on - for nonzero flowrate-having valves only
var bitmap = map[string]int{}

// distance matrix from node a to node b - populate with floyd warshall
var distances = map[string]map[string]int{}

type Valve struct {
	name      string
	flowRate  int
	neighbors map[string]bool
}

func main() {
	f, _ := os.ReadFile("./input.txt")
	input := string(f)
	valves = parseInput(input)
	setup(valves)
	resOne := partOne()
	fmt.Println(resOne)
	resTwo := partTwo()
	fmt.Println(resTwo)
}

func partOne() int {
	scores := visit("AA", 30, 0, 0, map[int]int{})
	maxScore := 0
	for _, score := range scores {
		maxScore = max(maxScore, score)
	}
	// current time, bitmap of turned on valves, current location = value of total pressure
	return maxScore
}

func partTwo() int {
	scores := visit("AA", 26, 0, 0, map[int]int{})
	maxScore := 0
	for a, scoreA := range scores {
		for b, scoreB := range scores {
			if a&b == 0 {
				maxScore = max(maxScore, scoreA+scoreB)
			}
		}
	}
	// current time, bitmap of turned on valves, current location = value of total pressure
	return maxScore
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func setup(valves map[string]Valve) {
	var i = 0
	for k, valve := range valves {
		if valve.flowRate != 0 {
			bitmap[k] = 1 << i
			i++
		}
	}

	for a, valveA := range valves {
		for b := range valves {
			_, ok := valveA.neighbors[b]
			// nil check cuz golang is weird
			if distances[a] == nil {
				distances[a] = map[string]int{}
			}
			if ok {
				distances[a][b] = 1
			} else {
				// set to inf if not in neighbors for now
				distances[a][b] = MAX_INT
			}
		}
	}

	// fix distance matrix with floyd warshall
	for c := range distances {
		for a := range distances {
			for b := range distances {
				distances[a][b] = min(distances[a][b], distances[a][c]+distances[c][b])
			}
		}
	}
}

// state := bitmap
// scores := bitmap : maximal pressure relieved with valve config
func visit(v string, timeLeft int, state int, maxScore int, scores map[int]int) map[int]int {
	score, ok := scores[state]
	if ok {
		scores[state] = max(score, maxScore)
	} else {
		scores[state] = max(0, maxScore)
	}

	for u, b := range bitmap {
		newTimeLeft := timeLeft - distances[v][u] - 1
		if b&state > 0 || newTimeLeft < 0 {
			continue
		}
		// state | I[u] produces a bitmap with valve u turned on
		visit(u, newTimeLeft, state|bitmap[u], maxScore+newTimeLeft*valves[u].flowRate, scores)
	}

	return scores
}

func parseInput(input string) map[string]Valve {
	var valves = map[string]Valve{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var valve Valve
		valve.neighbors = map[string]bool{}
		var parts []string
		if strings.Contains(line, "valves") {
			parts = strings.Split(line, "to valves ")
		} else {
			parts = strings.Split(line, "to valve ")
		}

		fmt.Sscanf(parts[0], "Valve %s has flow rate=%d;", &valve.name, &valve.flowRate)
		for _, neighbor := range strings.Split(parts[1], ", ") {
			valve.neighbors[neighbor] = true
		}
		valves[valve.name] = valve

	}
	return valves
}
