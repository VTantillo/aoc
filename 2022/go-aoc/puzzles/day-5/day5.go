package day5

import "fmt"

func Day5(input []string) string {
	startStacks, _ := parseInput(input)
	buildStacks(startStacks)

	return ""
}

func buildStacks(stacks []string) []*CrateStack {
	var crateStacks []*CrateStack

	var colRow int
	for i, s := range stacks {
		if len(s) > 1 && string(s[1]) == "1" {
			colRow = i
		}
	}

	numStacks := 0
	for _, c := range stacks[colRow] {
		if string(c) != " " {
			numStacks++
		}
	}

	for i := 0; i < numStacks; i++ {
		crateStacks = append(crateStacks, NewCrateStack())
	}

	fmt.Println("num of stacks", numStacks)
	for _, line := range stacks {
		idx := 1
		for i := 0; i < numStacks; i++ {
			crateStacks[i].Push(string(line[idx]))
			idx = idx + 4
		}
	}

	return crateStacks
}

func parseInput(input []string) ([]string, []string) {
	is_moves := false
	var startStacks []string
	var moves []string

	for _, l := range input {
		// fmt.Println(l, len(l))
		if !is_moves {
			startStacks = append(startStacks, l)
		} else {
			moves = append(moves, l)
		}

		if len(l) == 0 {
			is_moves = true
		}
	}

	return startStacks, moves
}
