package day5

import (
	"aoc/utils"
	"testing"
)

func TestDay5(t *testing.T) {
	input := utils.ReadInput("../../../sample/day-5.txt")

	result := Day5(input)

	if result != "CMZ" {
		t.Fatalf("Result did not match 'CMZ', got=%v", result)
	}
}

func TestCrateStack(t *testing.T) {
	stack := NewCrateStack()

	if stack.IsEmpty() != true {
		t.Fatalf("Stack should be empty, got=%v", stack.IsEmpty())
	}

	stack.Push("L")
	stack.Push("U")
	stack.Push("X")

	if stack.Size() != 3 {
		t.Fatalf("Did not push, size should be 3, got=%v", stack.Size())
	}

	peeked, err := stack.Peek()
	if err != nil {
		t.Fatalf("Didn't peek, there was an error")
	}

	if stack.Size() != 3 {
		t.Fatalf("Stack size should be 3, got=%v", stack.Size())
	}

	if peeked != "X" {
		t.Fatalf("Value returned from peek should be 'X', got=%v", peeked)
	}

	popped, err := stack.Pop()
	if err != nil {
		t.Fatal("Didn't pop, there was an error")
	}

	if stack.Size() != 2 {
		t.Fatalf("Stack size should be 2, got=%v", stack.Size())
	}

	if popped != "X" {
		t.Fatalf("Value returned from pop should be 'X', got=%v", popped)
	}
}
