package day13

import (
	"fmt"
	"slices"
)

type ReflectionDir int

const (
	ReflectionDirUndefined ReflectionDir = iota
	ReflectionDirVertical
	ReflectionDirHorizontal
)

func (r ReflectionDir) String() string {
	return [...]string{"undefined", "vertical", "horizontal"}[r]
}

type Reflection struct {
	Direction ReflectionDir
	Axis      int
}

type Pattern struct {
	Diagram    [][]rune
	Reflection *Reflection
}

func (p *Pattern) FindReflection() {
	if r, ok := p.findVerticalReflection(); ok {
		p.Reflection = r
	} else if r, ok := p.findHorizontalReflection(); ok {
		p.Reflection = r
	}

	fmt.Printf("%+v\n", p.Reflection)
}

func (p *Pattern) FindSumdgedReflection() {
	if r, ok := p.findVerticalSumdgeReflection(); ok {
		p.Reflection = r
	} else if r, ok := p.findHorizontalSmudgeReflection(); ok {
		p.Reflection = r
	}
	fmt.Printf("%+v\n", p.Reflection)
}

func (p *Pattern) findVerticalSumdgeReflection() (*Reflection, bool) {
	for i := 0; i < len(p.Diagram); i++ {
		for j := 0; j < len(p.Diagram[i])-1; j++ {
			if p.Diagram[i][j] == p.Diagram[i][j+1] {
				diffs := 0
				for row := 0; row < len(p.Diagram); row++ {
					for start, end := j, j+1; start >= 0 && end < len(p.Diagram[i]); start, end = start-1, end+1 {
						if p.Diagram[row][start] != p.Diagram[row][end] {
							diffs++
						}
					}
				}
				if diffs == 1 {
					ref := Reflection{
						Direction: ReflectionDirVertical,
						Axis:      j + 1,
					}
					return &ref, true
				}
			}
		}
	}

	return nil, false
}

func (p *Pattern) findHorizontalSmudgeReflection() (*Reflection, bool) {
	for i := 0; i < len(p.Diagram)-1; i++ {
		for j := 0; j < len(p.Diagram[i]); j++ {
			if p.Diagram[i][j] == p.Diagram[i+1][j] {
				diffs := 0
				for start, end := i, i+1; start >= 0 && end < len(p.Diagram); start, end = start-1, end+1 {
					for k := 0; k < len(p.Diagram[start]); k++ {
						if p.Diagram[start][k] != p.Diagram[end][k] {
							diffs++
						}
					}
				}
				if diffs == 1 {
					ref := Reflection{
						Direction: ReflectionDirHorizontal,
						Axis:      i + 1,
					}
					return &ref, true
				}
			}
		}
	}
	return nil, false
}

func (p *Pattern) findVerticalReflection() (*Reflection, bool) {
	for i := 0; i < len(p.Diagram); i++ {
		for j := 0; j < len(p.Diagram[i])-1; j++ {
			if p.Diagram[i][j] == p.Diagram[i][j+1] {
				isFound := true
				for row := 0; row < len(p.Diagram); row++ {
					for start, end := j, j+1; start >= 0 && end < len(p.Diagram[i]); start, end = start-1, end+1 {
						if p.Diagram[row][start] != p.Diagram[row][end] {
							isFound = false
						}
					}
				}
				if isFound {
					ref := Reflection{
						Direction: ReflectionDirVertical,
						Axis:      j + 1,
					}
					return &ref, true
				}
			}
		}
	}
	return nil, false
}

func (p *Pattern) findHorizontalReflection() (*Reflection, bool) {
	for i := 0; i < len(p.Diagram)-1; i++ {
		if slices.Equal(p.Diagram[i], p.Diagram[i+1]) {
			isFound := true
			for start, end := i, i+1; start >= 0 && end < len(p.Diagram); start, end = start-1, end+1 {
				if !slices.Equal(p.Diagram[start], p.Diagram[end]) {
					isFound = false
				}
			}
			if isFound {
				ref := Reflection{
					Direction: ReflectionDirHorizontal,
					Axis:      i + 1,
				}
				return &ref, true
			}
		}
	}
	return nil, false
}

func summarizePatterns(patterns []*Pattern) int {
	verticalLines := 0
	horizontalLines := 0

	for _, p := range patterns {
		if p.Reflection != nil {
			switch p.Reflection.Direction {
			case ReflectionDirHorizontal:
				horizontalLines += p.Reflection.Axis
			case ReflectionDirVertical:
				verticalLines += p.Reflection.Axis
			}
		}
	}

	total := horizontalLines * 100
	total = total + verticalLines

	return total
}

func Part1(input []string) int {
	patterns := parseInput(input)

	for _, p := range patterns {
		p.FindReflection()
	}

	result := summarizePatterns(patterns)

	fmt.Printf("Num of patterns: %d\n", len(patterns))
	return result
}

func Part2(input []string) int {
	patterns := parseInput(input)

	for _, p := range patterns {
		p.FindSumdgedReflection()
	}

	result := summarizePatterns(patterns)

	fmt.Printf("Num of patterns: %d\n", len(patterns))
	return result
}

func parseInput(input []string) []*Pattern {
	patterns := make([]*Pattern, 0)

	diagram := make([][]rune, 0)
	for _, l := range input {
		line := []rune(l)

		if len(line) == 0 {
			patternDiagram := make([][]rune, 0)
			patternDiagram = append(patternDiagram, diagram...)
			patterns = append(patterns, &Pattern{
				Diagram: patternDiagram,
			})
			diagram = nil
		} else {
			diagram = append(diagram, line)
		}

	}

	patternDiagram := make([][]rune, 0)
	patternDiagram = append(patternDiagram, diagram...)
	patterns = append(patterns, &Pattern{
		Diagram: patternDiagram,
	})

	return patterns
}
