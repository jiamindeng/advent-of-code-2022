package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	name     string
	size     int
	children map[string]Node
	parent   *Node
}

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func (node Node) print() {
	if node.parent != nil {
		fmt.Printf("{'name': '%s', 'size': '%d', 'children': '%+v', 'parent': '%s' }\n", node.name, node.size, node.children, node.parent.name)
	} else {
		fmt.Printf("{'name': '%s', 'size': '%d', 'children': '%+v', 'parent': 'nil' }\n", node.name, node.size, node.children)
	}
}

func buildTree(input string) Node {
	root := Node{name: "/", children: map[string]Node{}}
	curr := root
	lines := strings.Split(input, "\n")

	for _, line := range lines[1:] {
		if line == "$ ls" {
			continue
		} else if line[:3] == "dir" {
			name := line[4:]
			temp := curr
			childNode := Node{name: name, children: map[string]Node{}, parent: &temp}
			curr.children[name] = childNode
		} else if line[:4] == "$ cd" {
			var name string
			_, err := fmt.Sscanf(line, "$ cd %s", &name)
			if err != nil {
				fmt.Println(err)
			}

			if name == ".." {
				temp := curr
				curr = *temp.parent
			} else if name == "/" {
				curr = root
			} else {
				curr = curr.children[name]
			}
		} else {
			var size int
			var name string
			_, err := fmt.Sscanf(line, "%d %s", &size, &name)
			if err != nil {
				fmt.Println(err)
			}
			temp := curr
			childNode := Node{name: name, size: size, children: map[string]Node{}, parent: &temp}
			curr.children[name] = childNode
		}
	}

	return root
}

func partOne(input string) int {
	sizes := map[string]int{}

	root := buildTree(input)
	dfs(sizes, root, "", 0, 0)

	sum := 0
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func partTwo(input string) int {
	sizes := map[string]int{}

	root := buildTree(input)
	dfs(sizes, root, "", 0, 0)

	total := 70000000
	spaceNeeded := 30000000
	sizes_ := []int{}
	for _, size := range sizes {
		sizes_ = append(sizes_, size)
	}

	sort.Slice(sizes_, func(i, j int) bool {
		return sizes_[i] > sizes_[j]
	})

	freeSize := total - sizes_[0]

	for i := len(sizes_) - 1; i > 0; i-- {
		if sizes_[i] > (spaceNeeded - freeSize) {
			return sizes_[i]
		}
	}

	return -1
}

func dfs(sizes map[string]int, curr Node, filepath string, size int, depth int) int {
	sum := 0
	for name, child := range curr.children {
		filepath := filepath + "/" + name
		sum += dfs(sizes, child, filepath, sum, depth+1)
	}

	if len(curr.children) != 0 {
		curr.size = sum
		sizes[filepath] = curr.size
	}

	return curr.size
}
