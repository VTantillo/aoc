package main

import (
	"aoc/puzzles/day16"
	"aoc/utils"
	"fmt"
)

var exLayout = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func main() {
	input := utils.ReadInput("../inputs/day16.txt")
	// input := utils.ReadString(exLayout)

	result := day16.Day16(input)

	fmt.Println("result was: ", result)
}
