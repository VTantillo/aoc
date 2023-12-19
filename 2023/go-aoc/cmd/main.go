package main

import (
	"aoc/puzzles"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day10.txt")
	result := puzzles.Day10Pt2(input, puzzles.Right, false)

	// input := strings.Split(puzzles.Pt2Ex1, "\n")
	// result := puzzles.Day10Pt2(input, puzzles.Left, true)

	fmt.Println("result was: ", result)
}
