package day5

import "fmt"

func Day5(input []string) string {
	startStacks, moves := parseInput(input)
	crateStacks := buildStacks(startStacks)

	fmt.Println("stacks:", startStacks)
	fmt.Println("moves:", moves)
	fmt.Println(len(crateStacks))

	return ""
}

func buildStacks(stacks []string) []*CrateStack {
	var crateStacks []*CrateStack

	return crateStacks
}

func parseInput(input []string) ([]string, []string) {
	is_moves := false

	var startStacks []string
	var moves []string

	for _, l := range input {
		fmt.Println(l, len(l))
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
