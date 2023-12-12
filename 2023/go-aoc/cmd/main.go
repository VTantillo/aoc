package main

import (
	"aoc/puzzles"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day11.txt")

	result := puzzles.Day11Pt2(input, 1000000)

	fmt.Println("result was: ", result)
}
