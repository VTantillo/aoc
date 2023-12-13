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

func parseDay12(input []string) []conditionRecord {
	var records []conditionRecord
	for _, line := range input {
		splitLine := strings.Split(line, " ")
		fmt.Print(splitLine)
	}

	return records
}

func Day12Pt1(input []string) int {
	records := parseDay12(input)

	fmt.Println(len(records))
	return 0
}
