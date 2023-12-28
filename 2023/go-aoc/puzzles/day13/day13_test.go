package day13

import (
	"aoc/utils"
	"testing"
)

var exInput = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestParseInput(t *testing.T) {
	// input := utils.ReadInput("../../../inputs/day13.txt")
	input := utils.ReadString(exInput)
	patterns := parseInput(input)

	if len(patterns) != 2 {
		t.Errorf("length of patterns should have been 2, got=%v", len(patterns))
	}
}

func TestFindVerticalReflection(t *testing.T) {
	input := utils.ReadString(exInput)
	patterns := parseInput(input)

	ref, ok := patterns[0].findVerticalReflection()

	if !ok {
		t.Errorf("should have found a reflection, got=%v", ok)
	}

	if ref.Direction != ReflectionDirVertical {
		t.Errorf("reflection direction should have been vertical, got=%v", ref.Direction)
	}

	if ref.Axis != 5 {
		t.Errorf("axis should have been 5, got=%v", ref.Axis)
	}
}

func TestFindHorizontalReflection(t *testing.T) {
	input := utils.ReadString(exInput)
	patterns := parseInput(input)

	ref, ok := patterns[1].findHorizontalReflection()

	if !ok {
		t.Errorf("should have found a reflection, got=%v", ok)
	}

	if ref.Direction != ReflectionDirHorizontal {
		t.Errorf("reflection direction should have been horizontal, got=%v", ref.Direction)
	}

	if ref.Axis != 4 {
		t.Errorf("axis should have been 4, got=%v", ref.Axis)
	}
}

func TestFindReflection(t *testing.T) {
	input := utils.ReadString(exInput)
	patterns := parseInput(input)

	for _, p := range patterns {
		p.FindReflection()
	}

	if patterns[0].Reflection.Direction != ReflectionDirVertical {
		t.Errorf("first pattern should have had a vertical reflection, got=%v", patterns[0].Reflection.Direction)
	}

	if patterns[1].Reflection.Direction != ReflectionDirHorizontal {
		t.Errorf("second pattern should have had a horizontal reflection, got=%v", patterns[1].Reflection.Direction)
	}
}

func TestSummarizePatterns(t *testing.T) {
	input := utils.ReadString(exInput)
	patterns := parseInput(input)

	for _, p := range patterns {
		p.FindReflection()
	}

	result := summarizePatterns(patterns)

	if result != 405 {
		t.Errorf("result should have been 405, got=%v", result)
	}
}

func TestFixSmudges(t *testing.T) {
	input := utils.ReadString(exInput)
	patterns := parseInput(input)

	for _, p := range patterns {
		p.FindSumdgedReflection()
	}

	result := summarizePatterns(patterns)

	if result != 400 {
		t.Errorf("result should have been 400, got=%v", result)
	}
}
