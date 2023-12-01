package puzzles

import (
	"aoc/utils"
	"testing"
)

func TestDay1(t *testing.T) {
	input := utils.ReadInput("../../inputs/ex-day-1.txt")

	result := Day1(input)

	if result != 142 {
		t.Fatalf("Result should be 142, got=%v", result)
	}
}
