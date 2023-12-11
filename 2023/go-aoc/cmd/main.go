package main

import (
	"aoc/puzzles"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day10.txt")

	result := puzzles.Day10Pt2(input)

	fmt.Println("result was: ", result)
}
