package main

import (
	"aoc/puzzles"
	"fmt"
	"strings"
)

func main() {
	springsEx2 := `???.### 1,1,3
  .??..??...?##. 1,1,3
  ?#?#?#?#?#?#?#? 1,3,1,6
  ????.#...#... 4,1,1
  ????.######..#####. 1,6,5
  ?###???????? 3,2,1`

	input := strings.Split(springsEx2, "\n")

	// input := utils.ReadInput("../inputs/day11.txt")

	result := puzzles.Day12Pt1(input)

	fmt.Println("result was: ", result)
}
