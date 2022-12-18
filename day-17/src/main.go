package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOneTwo(string(f))
	fmt.Println(resOne)
}

func countLeadingZeros(n int) int {
	if n == 0 {
		return 7
	}
	max := 128
	count := 0
	for n < max {
		n = n << 1
		count++
	}
	return count - 1
}

func countTrailingZeros(n int) int {
	if n == 0 {
		return 7
	}

	count := 0
	for (n & 1) == 0 {
		n = n >> 1
		count += 1
	}
	return count
}

func parseInput(input string) []string {
	chars := []string{}
	for _, char := range input {
		chars = append(chars, string(char))
	}
	return chars
}

func partOneTwo(input string) int {
	rowCount := 0
	shift := parseInput(input)
	rocks := [][]int{
		{1111}, {10, 111, 10}, {111, 1, 1}, {1, 1, 1, 1}, {11, 11},
	}
	for i, rock := range rocks {
		for j, row := range rock {
			rocks[i][j] = binaryToDecimal(row)
		}
	}

	space := []int{127, 0, 0, 0, 0, 0, 0, 0}
	start := 4
	shiftcount := 0
	cache := map[string][]int{}

	for i := 0; i < 500000; i++ {
		if i > 100000 {
			key := fmt.Sprintf("%d,%d", i%(len(rocks)), shiftcount)
			data, ok := cache[key]
			if ok {
				y2 := float64(rowCount + countRows(space) - 1)
				y1 := float64(data[1])
				x2 := float64(i)
				x1 := float64(data[0])
				m := (y2 - y1) / (x2 - x1)
				b := y1 - m*x1
				x := float64(1000000000000)
				fmt.Printf("Cache hit, estimated height: %.12f\n", m*x+b)
			} else {
				key := fmt.Sprintf("%d,%d", i%(len(rocks)), shiftcount)
				cache[key] = []int{i, rowCount + countRows(space) - 1}
			}
		}
		rock := rocks[i%(len(rocks))]
		rock = spawnrock(rock)
		shifted := true
		space_ := make([]int, len(space))
		copy(space_, space)
		for shifted && start >= 1 {
			rock = shiftrock(rock, 1, shift[shiftcount])
			shiftcount = (shiftcount + 1) % (len(shift))
			temp, shifted := shiftdown(space_, rock, start)
			if shifted {
				copy(space, temp)
				temp, shifted = shiftdown(space_, rock, start-1)
				if shifted {
					copy(space, temp)
				} else {
					break
				}
			} else {
				rock = shiftrock(rock, 1, unshift[shift[(shiftcount-1)%(len(shift))]])
				temp, shifted = shiftdown(space_, rock, start)
				if shifted {
					copy(space, temp)
					temp, shifted = shiftdown(space_, rock, start-1)
					if shifted {
						copy(space, temp)
					} else {
						break
					}
				} else {
					break
				}
			}

			start--
		}

		if len(space) > 100 && len(space)%100 == 0 {
			rowCount += len(space) - 100

			temp := make([]int, len(space[len(space)-100:]))
			copy(temp, space[len(space)-100:])
			space = temp
		}

		paddedRowCount := countRows(space) + 7
		spaceLen := len(space)
		for i := spaceLen; i < paddedRowCount; i++ {
			space = append(space, 0)
		}

		start = startingheight(space)
	}

	return rowCount + countRows(space) - 1
}

func countRows(space []int) int {
	count := len(space)
	for i := len(space) - 1; i >= 0; i-- {
		if space[i] > 0 {
			break
		}
		count--

	}
	return count
}

func shiftdown(space_ []int, rock []int, start int) ([]int, bool) {
	space := make([]int, len(space_))
	copy(space, space_)
	for i := 0; i < len(rock); i++ {
		if space[start+i]&rock[i] != 0 {
			return space_, false
		}
		space[start+i] = space[start+i] | rock[i]
	}

	return space, true
}

var unshift = map[string]string{"<": ">", ">": "<"}

func startingheight(space []int) int {
	for i := len(space) - 1; i >= 0; i-- {
		if space[i] > 0 {
			return i + 4
		}
	}
	return -1
}

func spawnrock(rock_ []int) []int {
	rock := make([]int, len(rock_))
	copy(rock, rock_)
	zeroes := 8
	for zeroes != 2 {
		for r := 0; r < len(rock); r++ {
			rock[r] = rock[r] << 1
			zeroes = int(math.Min(float64(zeroes), float64(countLeadingZeros(rock[r]))))
		}
	}

	return rock
}

func shiftrock(rock_ []int, count int, direction string) []int {
	rock := make([]int, len(rock_))
	copy(rock, rock_)

	zeroes := 8
	for r := 0; r < len(rock); r++ {
		if direction == "<" {
			zeroes = int(math.Min(float64(zeroes), float64(countLeadingZeros(rock[r]))))
		} else if direction == ">" {
			zeroes = int(math.Min(float64(zeroes), float64(countTrailingZeros(rock[r]))))
		}
	}

	for zeroes != 0 && count > 0 {
		for r := 0; r < len(rock); r++ {
			if direction == "<" {
				rock[r] = rock[r] << 1
				zeroes = int(math.Min(float64(zeroes), float64(countLeadingZeros(rock[r]))))
			} else if direction == ">" {
				rock[r] = rock[r] >> 1
				zeroes = int(math.Min(float64(zeroes), float64(countTrailingZeros(rock[r]))))
			}
		}
		count--
	}

	return rock
}

func bitarrayprint(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		bitprint(arr[i])
	}
	fmt.Println()
}

func bitprint(n int) {
	fmt.Printf("%07s\n", strconv.FormatInt(int64(n), 2))
}

func bitcount(n int) int {
	count := 0
	for n > 0 {
		count = count + 1
		n = n & (n - 1)
	}
	return count
}

func binaryToDecimal(n int) int {
	i, err := strconv.ParseInt(fmt.Sprint(n), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
