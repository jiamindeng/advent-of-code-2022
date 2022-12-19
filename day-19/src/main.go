package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var globalMaxGeodes = 0

func main() {
	f, _ := os.ReadFile("./input.txt")
	blueprints := parseInput(string(f))
	fmt.Println(partOne(blueprints))
	fmt.Println(partTwo(blueprints))
}

func partOne(blueprints []Blueprint) int {
	result := 0
	for _, bp := range blueprints {
		res := bp.id * search(bp, 0, 0, 0, 0, 1, 0, 0, 0, 24)
		result += res
		globalMaxGeodes = 0
	}

	return result
}

func partTwo(blueprints []Blueprint) int {
	if len(blueprints) < 3 {
		return -1
	}
	result := 1
	for i := 0; i < 3; i++ {
		result *= search(blueprints[i], 0, 0, 0, 0, 1, 0, 0, 0, 32)
		globalMaxGeodes = 0
	}

	return result
}

func parseInput(input string) []Blueprint {
	lines := strings.Split(input, "\n")
	blueprints := []Blueprint{}

	for _, line := range lines {
		var id int
		oreRobot := OreRobot{}
		clayRobot := ClayRobot{}
		obsRobot := ObsidianRobot{}
		geodeRobot := GeodeRobot{}

		fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreRobot.ore, &clayRobot.ore, &obsRobot.ore, &obsRobot.clay, &geodeRobot.ore, &geodeRobot.obsidian)

		bp := Blueprint{
			id,
			oreRobot,
			clayRobot,
			obsRobot,
			geodeRobot,
		}

		blueprints = append(blueprints, bp)
	}

	return blueprints
}

func search(bp Blueprint, ore, clay, obsidian, geodes, oreRobots, clayRobots, obsRobots, geodeRobots, time int) int {
	if time == 0 || globalMaxGeodes >= geodes+rangeSum(geodeRobots, geodeRobots+time-1) {
		return 0
	}
	if oreRobots >= bp.geodeRobot.ore && obsRobots >= bp.geodeRobot.obsidian {
		return rangeSum(geodeRobots, geodeRobots+time-1)
	}

	oreLimitHit := oreRobots >= max(bp.geodeRobot.ore, max(bp.clayRobot.ore, bp.obsidianRobot.ore))
	clayLimitHit := clayRobots >= bp.obsidianRobot.clay
	obsLimitHit := obsRobots >= bp.geodeRobot.obsidian
	maxGeodes := 0

	if !oreLimitHit {
		maxGeodes = max(
			maxGeodes,
			geodeRobots+search(
				bp, ore+oreRobots, clay+clayRobots, obsidian+obsRobots, geodes+geodeRobots,
				oreRobots, clayRobots, obsRobots, geodeRobots, time-1))
	}
	if ore >= bp.oreRobot.ore && !oreLimitHit {
		maxGeodes = max(
			maxGeodes,
			geodeRobots+search(
				bp, ore-bp.oreRobot.ore+oreRobots, clay+clayRobots, obsidian+obsRobots, geodes+geodeRobots,
				oreRobots+1, clayRobots, obsRobots, geodeRobots, time-1))
	}
	if ore >= bp.clayRobot.ore && !clayLimitHit {
		maxGeodes = max(
			maxGeodes, geodeRobots+search(
				bp, ore-bp.clayRobot.ore+oreRobots, clay+clayRobots, obsidian+obsRobots, geodes+geodeRobots,
				oreRobots, clayRobots+1, obsRobots, geodeRobots, time-1))
	}
	if ore >= bp.obsidianRobot.ore && clay >= bp.obsidianRobot.clay && !obsLimitHit {
		maxGeodes = max(
			maxGeodes, geodeRobots+search(
				bp, ore-bp.obsidianRobot.ore+oreRobots, clay-bp.obsidianRobot.clay+clayRobots, obsidian+obsRobots, geodes+geodeRobots,
				oreRobots, clayRobots, obsRobots+1, geodeRobots, time-1))
	}
	if ore >= bp.geodeRobot.ore && obsidian >= bp.geodeRobot.obsidian {
		maxGeodes = max(
			maxGeodes, geodeRobots+search(
				bp, ore-bp.geodeRobot.ore+oreRobots, clay+clayRobots, obsidian-bp.geodeRobot.obsidian+obsRobots, geodes+geodeRobots,
				oreRobots, clayRobots, obsRobots, geodeRobots+1, time-1))
	}

	globalMaxGeodes = max(maxGeodes, globalMaxGeodes)
	return maxGeodes
}

// Max number of geodes that can be produced for the remaining time, given that
// a geode robot can be built each remaining minute.
func rangeSum(first, last int) int {
	return last*(last+1)/2 - ((first - 1) * first / 2)
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

type Blueprint struct {
	id            int
	oreRobot      OreRobot
	clayRobot     ClayRobot
	obsidianRobot ObsidianRobot
	geodeRobot    GeodeRobot
}

type OreRobot struct {
	ore int
}

type ClayRobot struct {
	ore int
}

type ObsidianRobot struct {
	ore  int
	clay int
}

type GeodeRobot struct {
	ore      int
	obsidian int
}
