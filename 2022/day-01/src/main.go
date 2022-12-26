package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func partOne(input string) float64 {
	groups := strings.Split(input, "\n\n")
	max := float64(-1)
	for _, group := range groups {
		nums := strings.Split(group, "\n")
		sum := float64(0)
		for _, num := range nums {
			parsedNum, _ := strconv.ParseFloat(num, 64)
			sum += parsedNum
		}
		max = math.Max(max, sum)
	}
	return max
}

func partTwo(input string) float64 {
	var sums []float64
	var res float64
	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		nums := strings.Split(group, "\n")
		sum := float64(0)
		for _, num := range nums {
			parsedNum, _ := strconv.ParseFloat(num, 64)
			sum += parsedNum
		}
		sums = append(sums, sum)
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	for i := 0; i < 3; i++ {
		res += sums[i]
	}
	return res
}
