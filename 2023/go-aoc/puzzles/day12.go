package puzzles

import (
	"fmt"
	"strings"
)

// '.' == operational
// '#' == damaged
// '?' == unknown

type conditionRecord struct {
	springs       []rune
	damagedGroups []int
}

func findArrangements(record conditionRecord) int {
	total := len(record.springs)
	return total
}

func parseConditionRecord(line string) conditionRecord {
	splitLine := strings.Split(line, " ")
	springs := []rune(splitLine[0])

	var groups []int
	for _, g := range strings.Split(splitLine[1], ",") {
		var group int
		fmt.Sscanf(string(g), "%d", &group)
		groups = append(groups, group)
	}

	return conditionRecord{
		springs:       springs,
		damagedGroups: groups,
	}
}

func parseDay12(input []string) []conditionRecord {
	var records []conditionRecord
	for _, line := range input {
		r := parseConditionRecord(line)
		records = append(records, r)
	}

	return records
}

func Day12Pt1(input []string) int {
	records := parseDay12(input)

	fmt.Println(len(records))
	return 0
}
