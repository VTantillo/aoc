package main

import (
	"aoc/puzzles"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day-1.txt")

	result := puzzles.Day1(input)

	fmt.Println("result was: ", result)
}
