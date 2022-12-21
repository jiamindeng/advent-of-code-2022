package main

import (
	"fmt"
	"os"
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

func partOne(input string) int {
	lookup := parseInput(input)
	return dfs(lookup["root"], lookup)
}

func partTwo(input string) int {
	lookup := parseInput(input)
	markHumn(lookup["root"], lookup)
	left := lookup["root"].left
	right := lookup["root"].right
	var res int
	if lookup[lookup["root"].left].humn {
		res = calculate(lookup[left], lookup, dfs(lookup[right], lookup))
	} else {
		res = calculate(lookup[right], lookup, dfs(lookup[left], lookup))
	}
	return res
}

type Node struct {
	name  string
	val   int
	left  string
	right string
	op    string
	humn  bool
}

func (node *Node) String() string {
	if node.left == "" && node.right == "" && node.op == "" {
		return fmt.Sprintf("{name: %s, value: %d, humn: %t}", node.name, node.val, node.humn)
	}
	return fmt.Sprintf("{name: %s, value: %d, left: %s, op: %s, right: %s, humn: %t}", node.name, node.val, node.left, node.op, node.right, node.humn)
}

func parseInput(input string) map[string]*Node {
	lines := strings.Split(input, "\n")
	lookup := map[string]*Node{}

	for _, line := range lines {
		node := &Node{}
		var children []string
		line_ := strings.Split(line, ": ")
		op := parseOp(line_[1])
		name := line_[0]
		if op != "" {
			children = strings.Split(line_[1], fmt.Sprintf(" %s ", op))
			node.left = children[0]
			node.right = children[1]
			node.op = op
		} else {
			v, err := strconv.ParseInt(line_[1], 10, 64)
			if err != nil {
				panic(err)
			}
			node.val = int(v)
		}
		node.name = name
		lookup[name] = node
	}

	return lookup
}

func parseOp(line string) string {
	operations := []string{"+", "-", "/", "*"}
	for _, op := range operations {
		if strings.Contains(line, op) {
			return op
		}
	}
	return ""
}

func dfs(node *Node, lookup map[string]*Node) int {
	if node.op == "" {
		return node.val
	}

	res := operate[node.op](dfs(lookup[node.left], lookup), dfs(lookup[node.right], lookup))
	node.val = res
	return res
}

func markHumn(node *Node, lookup map[string]*Node) bool {
	if node.name == "humn" {
		node.humn = true
		return true
	}

	if node.op == "" {
		return false
	}

	node.humn = markHumn(lookup[node.left], lookup) || markHumn(lookup[node.right], lookup)
	return node.humn
}

var operate = map[string]func(a int, b int) int{
	"+": func(a int, b int) int {
		return a + b
	},
	"-": func(a int, b int) int {
		return a - b
	},
	"*": func(a int, b int) int {
		return a * b
	},
	"/": func(a int, b int) int {
		return a / b
	},
}

var solve = map[string]map[string]func(res int, a int) int{
	"right": {
		"+": func(res int, a int) int {
			return res - a
		},
		"-": func(res int, a int) int {
			return a - res
		},
		"*": func(res int, a int) int {
			return res / a
		},
		"/": func(res int, a int) int {
			return a / res
		},
	},

	"left": {
		"+": func(res int, a int) int {
			return res - a
		},
		"-": func(res int, a int) int {
			return res + a
		},
		"*": func(res int, a int) int {
			return res / a
		},
		"/": func(res int, a int) int {
			return res * a
		},
	},
}

func calculate(node *Node, lookup map[string]*Node, target int) int {
	if node.left == "" || node.right == "" {
		return target
	}

	var newTarget int
	left := lookup[node.left]
	right := lookup[node.right]
	if right.humn {
		newTarget = solve["right"][node.op](target, dfs(left, lookup))
		return calculate(right, lookup, newTarget)
	}

	if left.humn {
		newTarget = solve["left"][node.op](target, dfs(right, lookup))
		return calculate(left, lookup, newTarget)
	}

	return 0
}
