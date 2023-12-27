package day12

import (
	"fmt"
	"strings"
)

// '.' == operational
// '#' == damaged
// '?' == unknown

type ConditionRecord struct {
	Springs       []rune
	DamagedGroups []int
}

func isRecordValid(springs []rune, groups []int) bool {
	isValid := true
	start := 0
	for _, g := range groups {
		inGroup := false
		remaining := g
		for i := start; i < len(springs); i++ {
			if !inGroup && springs[i] == '#' {
				inGroup = true
			}

			if inGroup && springs[i] == '#' {
				remaining--
			}

			if inGroup && springs[i] != '#' {
				inGroup = false
				start = i
				break
			}
		}

		if remaining != 0 {
			return false
		}
	}

	return isValid
}

func (cr ConditionRecord) FindArrangements() int {
	var unknowns []int

	for x, r := range cr.Springs {
		if r == '?' {
			unknowns = append(unknowns, x)
		}
	}

	var combos [][]rune
	start := cr.Springs[:unknowns[0]]
	combos = append(combos, start)

	for _, s := range cr.Springs {
		if s != '?' {
			for _, c := range combos {
				c = append(c, s)
			}
			continue
		}

		for _, c := range combos {
			c = append(c, '.')
		}

		for _, c := range combos {
			newC := make([]rune, len(c))
			copy(newC, c)
			newC = append(newC, '#')
			combos = append(combos, newC)
		}

	}

	for _, c := range combos {
		fmt.Println(c)
	}

	return len(combos)
}

func parseConditionRecord(line string) ConditionRecord {
	splitLine := strings.Split(line, " ")
	springs := []rune(splitLine[0])

	var groups []int
	for _, g := range strings.Split(splitLine[1], ",") {
		var group int
		fmt.Sscanf(string(g), "%d", &group)
		groups = append(groups, group)
	}

	return ConditionRecord{
		Springs:       springs,
		DamagedGroups: groups,
	}
}

func parseDay12(input []string) []ConditionRecord {
	var records []ConditionRecord
	for _, line := range input {
		r := parseConditionRecord(line)
		records = append(records, r)
	}

	return records
}

func Part1(input []string) int {
	records := parseDay12(input)

	fmt.Println(len(records))
	return 0
}
