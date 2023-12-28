package day14

import (
	"aoc/utils"
	"testing"
)

var exInput = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

var exResult = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

func TestPart1(t *testing.T) {
	input := utils.ReadString(exInput)
	result := Day14(input, false)

	if result != 136 {
		t.Errorf("the total load should have been 136, got=%v", result)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadString(exInput)
	result := Day14(input, true)

	if result != 64 {
		t.Errorf("the total load should have been 64, got=%v", result)
	}
}

func TestParseInput(t *testing.T) {
	input := utils.ReadString(exInput)
	p := parseInput(input)

	if len(p.Rocks) != 10 {
		t.Errorf("number of rows of rocks should have been 10, got=%v", len(p.Rocks))
	}
	if len(p.Rocks[1]) != 10 {
		t.Errorf("number of cols of rocks should have been 10, got=%v", len(p.Rocks))
	}
}

func TestCalcLoad(t *testing.T) {
	input := utils.ReadString(exResult)
	p := parseInput(input)

	result := p.CalcLoad()

	if result != 136 {
		t.Errorf("the total load should have been 136, got=%v", result)
	}
}

func TestTiltNorth(t *testing.T) {
	initial := utils.ReadString(exInput)
	result := utils.ReadString(exResult)

	initialPlatform := parseInput(initial)
	resultPlatform := parseInput(result)

	initialPlatform.TiltNorth()

	for row, rocks := range initialPlatform.Rocks {
		for col, rock := range rocks {
			myRock := *rock
			resultRock := *resultPlatform.Rocks[row][col]
			if myRock != resultRock {
				t.Errorf("didn't tilt correctly, expected=%s, got=%s",
					string(resultRock), string(myRock))
			}
		}
	}
}
