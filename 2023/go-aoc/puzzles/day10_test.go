package puzzles

import (
	"aoc/utils"
	"slices"
	"testing"
)

var ex1a = [][]rune{
	{'.', '.', '.', '.', '.'},
	{'.', 'S', '-', '7', '.'},
	{'.', '|', '.', '|', '.'},
	{'.', 'L', '-', 'J', '.'},
	{'.', '.', '.', '.', '.'},
}

var ex1b = [][]rune{
	{'-', 'L', '|', 'F', '7'},
	{'7', 'S', '-', '7', '|'},
	{'.', '|', '.', '|', '|'},
	{'.', 'L', '-', 'J', '|'},
	{'L', '|', '-', 'J', 'F'},
}

var ex2a = [][]rune{
	{'.', '.', 'F', '7', '.'},
	{'.', 'F', 'J', '|', '.'},
	{'S', 'J', '.', 'L', '7'},
	{'|', 'F', '-', '-', 'J'},
	{'L', 'J', '.', '.', '.'},
}

var ex2b = [][]rune{
	{'7', '-', 'F', '7', '-'},
	{'.', 'F', 'J', '|', '7'},
	{'S', 'J', 'L', 'L', '7'},
	{'|', 'F', '-', '-', 'J'},
	{'L', 'J', '.', 'L', 'J'},
}

func TestDay10Pt1(t *testing.T) {
	input := utils.ReadInput("../../inputs/day10-ex2.txt")

	result := Day10Pt1(input)

	if result != 8 {
		t.Fatalf("Total loop steps should be 8, got=%v", result)
	}
}

func TestFindLoop(t *testing.T) {
	myMap := navMap{cellMap: makeCellMap(ex2a)}
	myMap.init()

	result := myMap.findLoop(east)

	if len(result) != 16 {
		t.Fatalf("Total loop steps should be 16, got=%v", len(result))
	}
}

func TestParseDay10(t *testing.T) {
	input := utils.ReadInput("../../inputs/day10-ex1.txt")

	pipeMap := parseDay10(input)

	for i, r := range pipeMap {
		if slices.Compare(r, ex1a[i]) != 0 {
			t.Fatal("Map and result were not the same")
		}
	}
}
