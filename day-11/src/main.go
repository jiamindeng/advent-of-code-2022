package main

import (
	"fmt"
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

type Monkey struct {
	items        []int64
	operator     string
	operand      int64
	divisor      int64
	nextTrue     int64
	nextFalse    int64
	numInspected int
	opSelf       bool
}

var calculate = map[string]func(operandA int64, operandB int64) int64{
	"*": func(operandA int64, operandB int64) int64 { return operandA * operandB },
	"-": func(operandA int64, operandB int64) int64 { return operandA - operandB },
	"+": func(operandA int64, operandB int64) int64 { return operandA + operandB },
}

func partOne(input string) int {
	monkeys := parseMonkeys(input)
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			for _, item := range monkey.items {
				monkey.numInspected++
				monkey.items = monkey.items[1:]
				result := int64(0)
				if monkey.opSelf {
					result = calculate[monkey.operator](item, item)
				} else {
					result = calculate[monkey.operator](item, monkey.operand)
				}
				worryLevel := result / 3
				if worryLevel%monkey.divisor == 0 {
					monkeys[monkey.nextTrue].items = append(monkeys[monkey.nextTrue].items, worryLevel)
				} else {
					monkeys[monkey.nextFalse].items = append(monkeys[monkey.nextFalse].items, worryLevel)
				}
			}
		}
	}

	product := 1
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].numInspected > monkeys[j].numInspected
	})
	for _, monkey := range monkeys[:2] {
		product *= monkey.numInspected
	}

	return product
}

func parseMonkeys(input string) []Monkey {
	monkeys := []Monkey{}
	chunks := strings.Split(input, "\n\n")
	for _, chunk := range chunks {
		monkey := Monkey{}
		for i, line := range strings.Split(chunk, "\n") {
			mod := i % 6
			if mod == 0 {
				continue
			} else if mod == 1 {
				items_ := strings.Split(line, ": ")
				items := strings.Split(items_[1], ", ")
				for _, item := range items {
					item_, err := strconv.ParseInt(item, 10, 64)
					if err != nil {
						panic(err)
					}
					monkey.items = append(monkey.items, item_)
				}
			} else if mod == 2 {
				operand := ""
				fmt.Sscanf(line, "  Operation: new = old %s %s", &monkey.operator, &operand)
				if operand == "old" {
					monkey.opSelf = true
				} else {
					monkey.operand, _ = strconv.ParseInt(operand, 10, 64)
				}
			} else if mod == 3 {
				fmt.Sscanf(line, "  Test: divisible by %d", &monkey.divisor)
			} else if mod == 4 {
				fmt.Sscanf(line, "    If true: throw to monkey %d", &monkey.nextTrue)
			} else if mod == 5 {
				fmt.Sscanf(line, "    If false: throw to monkey %d", &monkey.nextFalse)
			}
		}

		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func partTwo(input string) int {
	monkeys := parseMonkeys(input)
	divisors := []int64{}
	for _, monkey := range monkeys {
		divisors = append(divisors, monkey.divisor)
	}
	lcm := LCM(divisors[0], divisors[1], divisors[2:]...)

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			for _, item := range monkey.items {
				monkey.numInspected++
				monkey.items = monkey.items[1:]
				result := int64(0)
				if monkey.opSelf {
					result = calculate[monkey.operator](item, item)
				} else {
					result = calculate[monkey.operator](item, monkey.operand)
				}

				if result%monkey.divisor == 0 {
					monkeys[monkey.nextTrue].items = append(monkeys[monkey.nextTrue].items, result%lcm)
				} else {
					monkeys[monkey.nextFalse].items = append(monkeys[monkey.nextFalse].items, result%lcm)
				}
			}
		}
	}

	product := 1
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].numInspected > monkeys[j].numInspected
	})
	for _, monkey := range monkeys[:2] {
		product *= monkey.numInspected
	}

	return product
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
