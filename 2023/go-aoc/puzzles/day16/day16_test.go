package day16

import (
	"aoc/utils"
	"testing"
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

func TestPart1(t *testing.T) {
	input := utils.ReadString(exLayout)
	result := Day16(input)

	expected := 51

	if result != expected {
		t.Errorf("number of energized tiles should have been %d, got=%d", expected, result)
	}
}
