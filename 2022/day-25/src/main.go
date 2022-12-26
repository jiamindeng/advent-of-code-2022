package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("./input.txt")
	dict, _ := os.ReadFile("./dict.txt")
	input, toSnafu, fromSnafu := parseInput(string(in), string(dict))
	resOne := partOne(input, toSnafu, fromSnafu)
	fmt.Println(resOne)
}

func partOne(input []string, toSnafu map[int]string, fromSnafu map[string]int) string {
	sum := 0
	for _, snafu := range input {
		decimal := snafuToDecimal(snafu, fromSnafu)
		sum += decimal
	}

	return decimalToSnafu(sum, toSnafu)
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}

	return res
}

func ceil(n int, base int) int {
	for mod(n, base) != 0 {
		n++
	}
	return n
}

func decimalToSnafu(decimal int, toSnafu map[int]string) string {
	snafu := ""
	quotient := decimal
	base := 5
	for quotient != 0 {
		newQuotient := quotient / base
		remainder := mod(quotient, base)
		if remainder >= 3 {
			remainder -= base
			newQuotient = ceil(quotient, base) / base
		}
		quotient = newQuotient
		snafu = toSnafu[remainder] + snafu
	}
	return snafu
}

func snafuToDecimal(snafu string, fromSnafu map[string]int) int {
	places := len(snafu) - 1
	base := 5
	decimal := float64(0)
	for i := 0; i < len(snafu); i++ {
		decimal += math.Pow(float64(base), float64(places)) * float64(fromSnafu[string(snafu[i])])
		places -= 1
	}

	return int(decimal)
}

type word struct {
	decimal string
	snafu   string
}

func parseInput(in string, dict string) ([]string, map[int]string, map[string]int) {
	dictPairs := strings.Split(dict, "\n")
	lines := strings.Split(in, "\n")
	toSnafu := map[int]string{}
	fromSnafu := map[string]int{}

	for _, pair := range dictPairs {
		words := strings.Split(pair, ",")
		quinary, err := strconv.ParseInt(words[0], 10, 64)
		if err != nil {
			panic(err)
		}

		toSnafu[int(quinary)] = words[1]
		fromSnafu[words[1]] = int(quinary)
	}

	return lines, toSnafu, fromSnafu
}
