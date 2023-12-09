package main

import (
	"aoc/puzzles"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day9.txt")
	// input := utils.ReadInput("../inputs/day9-ex.txt")

	result := puzzles.Day9Pt1(input)

	fmt.Println("result was: ", result)
}
