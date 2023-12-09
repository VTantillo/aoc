package puzzles

import (
	"aoc/utils"
	"slices"
	"testing"
)

func TestDay9Pt1(t *testing.T) {
	input := utils.ReadInput("../../inputs/day9-ex.txt")

	result := Day9Pt1(input)

	if result != 114 {
		t.Fatalf("Result should be 114, got=%v", result)
	}
}

func TestDay9Pt2(t *testing.T) {
	input := utils.ReadInput("../../inputs/day9-ex.txt")

	result := Day9Pt2(input)

	if result != 2 {
		t.Fatalf("Result should be 2, got=%v", result)
	}
}

func TestParseDay9(t *testing.T) {
	input := utils.ReadInput("../../inputs/day9-ex.txt")

	readings := parseDay9(input)

	result := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	for i, r := range readings {
		if slices.Compare(r, result[i]) != 0 {
			t.Fatal("Reading and result were not the same")
		}
	}
}

func TestDiffLine(t *testing.T) {
	example := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	diff1 := diffLine(example[0])
	if slices.Compare(diff1, []int{3, 3, 3, 3, 3}) != 0 {
		t.Fatal("Diff 1 was not correct")
	}

	diff2 := diffLine(example[1])
	if slices.Compare(diff2, []int{2, 3, 4, 5, 6}) != 0 {
		t.Fatal("Diff 2 was not correct")
	}
}

func TestCheckZeros(t *testing.T) {
	example1 := []int{0, 0, 0}
	if !checkAllZeros(example1) {
		t.Fatal("It no werk")
	}

	example2 := []int{0, 1, 2}
	if checkAllZeros(example2) {
		t.Fatal("It no werk")
	}
}

func TestNextVal(t *testing.T) {
	example := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	nextVal1 := findNextVal(example[0])
	if nextVal1 != 18 {
		t.Fatalf("Result should be 18, got=%v", nextVal1)
	}

	nextVal2 := findNextVal(example[1])
	if nextVal2 != 28 {
		t.Fatalf("Result should be 28, got=%v", nextVal2)
	}

	nextVal3 := findNextVal(example[2])
	if nextVal3 != 68 {
		t.Fatalf("Result should be 68, got=%v", nextVal3)
	}
}

func TestPrevVal(t *testing.T) {
	example := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	prevVal1 := findPrevVal(example[0])
	if prevVal1 != -3 {
		t.Fatalf("Result should be -3, got=%v", prevVal1)
	}

	prevVal2 := findPrevVal(example[1])
	if prevVal2 != 0 {
		t.Fatalf("Result should be 0, got=%v", prevVal2)
	}

	prevVal3 := findPrevVal(example[2])
	if prevVal3 != 5 {
		t.Fatalf("Result should be 5, got=%v", prevVal3)
	}
}
