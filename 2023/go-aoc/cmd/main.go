package main

import (
	"aoc/puzzles/day14"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day14.txt")
	result := day14.Day14(input, true)

	fmt.Println("result was: ", result)
}
