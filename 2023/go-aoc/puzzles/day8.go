package puzzles

import (
	"fmt"
	"strings"
)

type NetworkNode struct {
	left  string
	right string
}

func Day8(input []string) int {
	var moves string
	network := make(map[string]NetworkNode)

	var nodes []string

	for i, line := range input {
		if i == 0 {
			moves = line
		}

		if strings.Contains(line, "=") {
			src := line[:3]
			if rune(src[2]) == 'A' {
				nodes = append(nodes, src)
			}
			network[src] = NetworkNode{left: line[7:10], right: line[12:15]}
		}
	}

	fmt.Println("Starting nodes", nodes)

	var nodeSteps []int
	for _, n := range nodes {
		curr := n
		steps := 0
	NodeCycle:
		for rune(curr[2]) != 'Z' {
			for _, dir := range moves {
				switch dir {
				case 'L':
					curr = network[curr].left
				case 'R':
					curr = network[curr].right
				}
				steps++
				if rune(curr[2]) == 'Z' {
					nodeSteps = append(nodeSteps, steps)
					fmt.Println(n, "in", steps)
					break NodeCycle
				}
			}
		}
	}

	sum := lcm(nodeSteps[0], nodeSteps[1], nodeSteps[2:]...)

	return sum
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func isFinished(nodes []string) bool {
	for _, n := range nodes {
		if rune(n[2]) != 'Z' {
			return false
		}
	}
	return true
}
