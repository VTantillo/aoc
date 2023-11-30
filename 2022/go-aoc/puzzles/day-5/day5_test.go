package day5

import (
	"aoc/utils"
	"testing"
)

func TestDay5(t *testing.T) {
	input := utils.ReadInput("../../../sample/day-5.txt")

	result := Day5(input)

	if result != "CMZ" {
		t.Errorf("Result did not match 'CMZ', got=%v", result)
	}
}

func TestCrateStack(t *testing.T) {
	stack := NewCrateStack()

	if stack.IsEmpty() != true {
		t.Errorf("Stack should be empty, got=%v", stack.IsEmpty())
	}

	stack.Push("L")
	stack.Push("U")
	stack.Push("X")

	if stack.Size() != 3 {
		t.Errorf("Did not push, size should be 3, got=%v", stack.Size())
	}
}
