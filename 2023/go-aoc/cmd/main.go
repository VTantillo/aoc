package main

import (
	"aoc/puzzles/day20"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../inputs/day20.txt")
	result := day20.Part2(input)

	fmt.Println("result was: ", result)
}
