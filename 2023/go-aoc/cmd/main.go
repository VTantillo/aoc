package main

import (
	"aoc/puzzles"
	"fmt"
	"strings"
)

func main() {
	// input := utils.ReadInput("../inputs/day10.txt")
	input := strings.Split(puzzles.Pt2Ex1, "\n")

	result := puzzles.Day10Pt2(input)

	fmt.Println("result was: ", result)
}
