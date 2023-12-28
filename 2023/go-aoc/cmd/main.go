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

	// result := day16.RunSingle(input, day16.BeamConfig{
	// 	Coords:    day16.Coords{Row: 0, Col: 75},
	// 	Direction: day16.DirDown,
	// })

	fmt.Println("result was: ", result)
}
