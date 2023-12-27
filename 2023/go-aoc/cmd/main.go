package main

import (
	"aoc/puzzles/day13"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day13.txt")
	result := day13.Part2(input)

	fmt.Println("result was: ", result)
}
