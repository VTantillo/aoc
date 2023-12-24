package day12

import (
	"fmt"
	"strings"
)

// '.' == operational
// '#' == damaged
// '?' == unknown

type ConditionRecord struct {
	springs       []rune
	damagedGroups []int
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
				start = i
			}

			if inGroup && springs[i] == '#' {
				remaining--
			}

			if remaining == 0 {
				break
			}

			if inGroup && springs[i] != '#' {
				inGroup = false
			}
		}

	}

	return isValid
}

func findArrangements(record ConditionRecord) int {
	total := len(record.springs)
	return total
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
		springs:       springs,
		damagedGroups: groups,
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
