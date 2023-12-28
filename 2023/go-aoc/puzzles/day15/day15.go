package day15

import (
	"fmt"
	"strings"
)

type (
	Lens struct {
		Label         string
		FocalLength   int
		FocusingPower int
	}
	Boxes map[int][]*Lens
	Step  struct {
		Label       string
		Operation   rune
		FocalLength int
	}
)

func (b Boxes) AddLens(l Lens) {
	boxNum := hash(l.Label)

	for i, lens := range b[boxNum] {
		if lens != nil && lens.Label == l.Label {
			b[boxNum][i] = &l
			return
		}
	}

	b[boxNum] = append(b[boxNum], &l)
}

func (b Boxes) RemoveLens(l Lens) {
	boxNum := hash(l.Label)
	var lenses []*Lens
	lenses = append(lenses, b[boxNum]...)

	for i, lens := range lenses {
		if lens != nil && lens.Label == l.Label {
			copy(lenses[i:], lenses[i+1:])
			lenses[len(lenses)-1] = nil
			lenses = lenses[:len(lenses)-1]
		}
	}
	b[boxNum] = lenses
}

func (b Boxes) CalcFocusingPower() int {
	for k, v := range b {
		lensSum := 0
		for i, l := range v {
			if l != nil {
				lensSum = (k + 1) * (i + 1) * l.FocalLength
				fmt.Printf("Lens: %d, %d, %d = %d\n", k+1, i+1, l.FocalLength, lensSum)
				l.FocusingPower = lensSum
			}
		}
	}

	b.PrintBoxes()

	sum := 0
	for _, v := range b {
		for _, l := range v {
			if l != nil {
				sum += l.FocusingPower
			}
		}
	}

	return sum
}

func (b Boxes) PrintBoxes() {
	for k, v := range b {
		fmt.Printf("Box: %d ", k)
		for _, l := range v {
			if l != nil {
				fmt.Printf("[%s %d %d]", l.Label, l.FocalLength, l.FocusingPower)
			}
		}
		fmt.Print("\n")
	}
}

func hash(input string) int {
	val := 0

	for _, c := range input {
		ascii := int(c)
		val += ascii
		val *= 17
		val %= 256
	}

	return val
}

func Day15(input []string) int {
	b := make(Boxes)
	steps := parseInput(input)
	for _, s := range steps {
		fmt.Printf("%s %c %d\n", s.Label, s.Operation, s.FocalLength)
		lens := Lens{
			Label:       s.Label,
			FocalLength: s.FocalLength,
		}

		switch s.Operation {
		case '=':
			b.AddLens(lens)
		case '-':
			b.RemoveLens(lens)
		}
	}
	return b.CalcFocusingPower()
}

func parseInput(input []string) []Step {
	var steps []Step
	for _, line := range input {
		for _, s := range strings.Split(line, ",") {
			op := []rune(s)

			var label string
			var focalLength int
			var operation rune

			for i, c := range op {
				if c == '-' || c == '=' {
					operation = c
					label = string(op[:i])
					if c == '=' {
						fmt.Sscanf(string(op[i+1:]), "%d", &focalLength)
					}
					break
				}
			}

			steps = append(steps, Step{
				Label:       label,
				Operation:   operation,
				FocalLength: focalLength,
			})
		}
	}
	return steps
}
