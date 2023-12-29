package main

import (
	"aoc/puzzles/day17"
	"aoc/utils"
)

var exMap = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

var testInput = utils.ReadString(exMap)

func main() {
	// input := utils.ReadInput("../inputs/day17.txt")
	// day17.Day17(input)

	// start := time.Now()
	// result := day17.Day17(input)
	// end := time.Now()

	day17.Day17(testInput)
}
