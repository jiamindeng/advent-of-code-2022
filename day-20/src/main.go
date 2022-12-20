package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(parseInput(string(f)))
	resTwo := partTwo(parseInput(string(f)))
	fmt.Println(resOne)
	fmt.Println(resTwo)
}

func partOne(lookup []*Node) int {
	_, zero := shuffle(lookup)
	total := 0

	for _, n := range []int{1000, 2000, 3000} {
		total += findNth(n, zero, len(lookup)).value
	}

	return total
}

func partTwo(lookup []*Node) int {
	var zero *Node
	key := 811589153
	for _, node := range lookup {
		node.value = node.value * key
	}

	for i := 0; i < 10; i++ {
		_, zero = shuffle(lookup)
	}

	total := 0
	for _, n := range []int{1000, 2000, 3000} {
		total += findNth(n, zero, len(lookup)).value
	}

	return total
}

func shuffle(lookup []*Node) ([]*Node, *Node) {
	var zero *Node

	for i := range lookup {
		toMove := lookup[i]
		dest := lookup[i]

		if toMove.value == 0 {
			zero = toMove
			continue
		}

		if toMove.value > 0 {
			for j := 0; j < toMove.value%(len(lookup)-1); j++ {
				dest = dest.next
				if dest == toMove {
					dest = dest.next
				}
			}
		} else {
			for j := 0; j <= -toMove.value%(len(lookup)-1); j++ {
				dest = dest.prev
				if dest == toMove {
					dest = dest.prev
				}
			}
		}

		if toMove != dest {
			dest.insert(toMove)
		}
	}

	return lookup, zero
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (list LinkedList) String() string {
	str := "["
	curr := list.head
	for i := 0; i < list.length && curr != nil; i++ {
		str += fmt.Sprintf("%d,", curr.value)
		curr = curr.next
	}
	str += "]"
	return str
}

func (node Node) String() string {
	if node.prev == nil || node.next == nil {
		return fmt.Sprintf("{prev: <nil>, value: %d, next: <nil>}", node.value)
	}
	return fmt.Sprintf("{prev: %d, value: %d, next: %d}", node.prev.value, node.value, node.next.value)
}

func (node *Node) remove() *Node {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.next = nil
	node.prev = nil
	return node
}

func (node *Node) insert(new *Node) {
	new.remove()
	next := node.next
	node.next = new
	next.prev = new
	new.prev = node
	new.next = next
}

func findNth(n int, zero *Node, length int) *Node {
	curr := zero

	for i := 0; i < n%length; i++ {
		curr = curr.next
	}

	return curr
}

func parseInput(input string) []*Node {
	lookup := []*Node{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}

		node := &Node{value: int(num)}
		lookup = append(lookup, node)
	}

	for i := 1; i < len(lookup)-1; i++ {
		curr := lookup[i]
		prev := lookup[i-1]
		next := lookup[i+1]
		curr.next = next
		curr.prev = prev
		lookup[i] = curr
	}

	last := lookup[len(lookup)-1]
	next := lookup[0]
	prev := lookup[len(lookup)-2]
	last.next = next
	last.prev = prev
	next.prev = last
	prev.next = last

	start := next
	prev = last
	next = lookup[1]
	start.next = next
	start.prev = prev
	next.prev = start
	prev.next = start

	return lookup
}
