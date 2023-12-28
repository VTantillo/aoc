package main

import (
	"aoc/puzzles/day15"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day15.txt")
	result := day15.Day15(input)

	fmt.Println("result was: ", result)
}
