package puzzles

import (
	"aoc/utils"
	"testing"
)

func TestDay1(t *testing.T) {
	input := utils.ReadInput("../../inputs/ex2-day-1.txt")

	result := Day1(input)

	if result != 281 {
		t.Fatalf("Result should be 281, got=%v", result)
	}
}
