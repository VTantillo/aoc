package main

import (
	day5 "aoc/puzzles/day-5"
	"aoc/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../sample/day-5.txt")
	result := day5.Day5(input)

	fmt.Printf("Result was: %v\n", result)
}
