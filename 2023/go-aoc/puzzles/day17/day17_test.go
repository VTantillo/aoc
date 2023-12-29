package day17

import (
	"aoc/utils"
	"slices"
	"testing"
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

func TestParseMap(t *testing.T) {
	firstRow := []int{2, 4, 1, 3, 4, 3, 2, 3, 1, 1, 3, 2, 3}

	blockMap := parseMap(testInput)
	mapFirstRow := blockMap[0]

	if !slices.Equal(firstRow, mapFirstRow) {
		t.Fatalf("first rows don't match expected %v, got=%v", firstRow, mapFirstRow)
	}
}
