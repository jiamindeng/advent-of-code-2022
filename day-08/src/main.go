package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

type Max struct {
	north int
	east  int
	south int
	west  int
}

func partOne(input string) int {
	rows := strings.Split(input, "\n")
	trees := [][]int64{}
	visible := 0
	for _, row := range rows {
		rowInt := []int64{}
		for _, height := range row {
			heightInt, _ := strconv.ParseInt(string(height), 10, 64)
			rowInt = append(rowInt, heightInt)
		}
		trees = append(trees, rowInt)
	}

	for r := 1; r < len(trees) && r > 0 && r < len(trees)-1; r++ {
		for c := 1; c < len(trees[0]) && c > 0 && c < len(trees[0])-1; c++ {

			curr := trees[r][c]

			_, eastMax := MinMax(trees[r][c+1:])

			_, westMax := MinMax(trees[r][:c])

			treesNorth := []int64{}
			for i := 0; i < len(trees) && i < r; i++ {
				treesNorth = append(treesNorth, trees[i][c])
			}
			_, northMax := MinMax(treesNorth)

			treesSouth := []int64{}
			for i := r + 1; i < len(trees); i++ {
				treesSouth = append(treesSouth, trees[i][c])
			}
			_, southMax := MinMax(treesSouth)

			if westMax < curr || eastMax < curr || northMax < curr || southMax < curr {
				visible++
			}

		}
	}
	return visible + len(trees[0])*4 - 4
}

func partTwo(input string) float64 {
	max := float64(-1)
	rows := strings.Split(input, "\n")
	trees := [][]int64{}
	for _, row := range rows {
		rowInt := []int64{}
		for _, height := range row {
			heightInt, _ := strconv.ParseInt(string(height), 10, 64)
			rowInt = append(rowInt, heightInt)
		}
		trees = append(trees, rowInt)
	}

	for r := 1; r < len(trees) && r > 0 && r < len(trees)-1; r++ {
		for c := 1; c < len(trees[0]) && c > 0 && c < len(trees[0])-1; c++ {

			curr := trees[r][c]

			numEast := NumVisibleTrees(curr, trees[r][c+1:])
			fmt.Println("e", numEast)
			numWest := NumVisibleTrees(curr, reverse(trees[r][:c]))
			fmt.Println("w", numWest)
			treesNorth := []int64{}
			for i := 0; i < len(trees) && i < r; i++ {
				treesNorth = append(treesNorth, trees[i][c])
			}
			numNorth := NumVisibleTrees(curr, reverse(treesNorth))
			fmt.Println("n", numNorth)
			treesSouth := []int64{}
			for i := r + 1; i < len(trees); i++ {
				treesSouth = append(treesSouth, trees[i][c])
			}
			numSouth := NumVisibleTrees(curr, treesSouth)
			fmt.Println("s", numSouth)
			max = math.Max(max, float64(numNorth*numEast*numSouth*numWest))

		}
	}
	return max
}

func NumVisibleTrees(current int64, array []int64) int {
	fmt.Println(current, array)
	for i, num := range array {
		if num >= current {
			return i + 1
		}
	}
	return len(array)
}

func reverse(array []int64) []int64 {
	array_ := make([]int64, len(array))
	copy(array_, array)
	for i, j := 0, len(array_)-1; i < j; i, j = i+1, j-1 {
		array_[i], array_[j] = array_[j], array_[i]
	}
	return array_
}

func MinMax(array []int64) (int64, int64) {
	if len(array) == 0 {
		return int64(math.Inf(1)), int64(math.Inf(-1))
	}
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
