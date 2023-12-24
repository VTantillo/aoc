package day12

import (
	"slices"
	"strings"
	"testing"
)

var springsEx string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestDay12Pt1(t *testing.T) {
	lines := strings.Split(springsEx, "\n")

	result := Part1(lines)

	if result != 21 {
		t.Fatalf("Result should be 21, got=%v", result)
	}
}

func TestParseConditionRecord(t *testing.T) {
	line1 := "???.### 1,1,3"
	line2 := ".??..??...?##. 1,1,3"
	line3 := "?###???????? 3,2,1"

	record1 := parseConditionRecord(line1)
	record2 := parseConditionRecord(line2)
	record3 := parseConditionRecord(line3)

	springs1 := []rune{'?', '?', '?', '.', '#', '#', '#'}
	groups1 := []int{1, 1, 3}

	if !slices.Equal(record1.springs, springs1) {
		t.Fatalf("Springs should have been %v, got=%v", springs1, record1.springs)
	}
	if !slices.Equal(record1.damagedGroups, groups1) {
		t.Fatalf("Damaged groups should have been %v, got=%v", groups1, record1.damagedGroups)
	}

	springs2 := []rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'}
	groups2 := []int{1, 1, 3}

	if !slices.Equal(record2.springs, springs2) {
		t.Fatalf("Springs should have been %v, got=%v", springs2, record2.springs)
	}
	if !slices.Equal(record2.damagedGroups, groups2) {
		t.Fatalf("Damaged groups should have been %v, got=%v", groups2, record2.damagedGroups)
	}

	springs3 := []rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'}
	groups3 := []int{3, 2, 1}

	if !slices.Equal(record3.springs, springs3) {
		t.Fatalf("Springs should have been %v, got=%v", springs3, record3.springs)
	}
	if !slices.Equal(record3.damagedGroups, groups3) {
		t.Fatalf("Damaged groups should have been %v, got=%v", groups3, record3.damagedGroups)
	}
}

func TestFindArrangements(t *testing.T) {
	// ???.### 1,1,3 - 1 arrangement
	record1 := ConditionRecord{
		springs:       []rune{'?', '?', '?', '.', '#', '#', '#'},
		damagedGroups: []int{1, 1, 3},
	}

	// .??..??...?##. 1,1,3 - 4 arrangements
	record2 := ConditionRecord{
		springs:       []rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'},
		damagedGroups: []int{1, 1, 3},
	}

	// ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
	record3 := ConditionRecord{
		springs:       []rune{'?', '#', '?', '#', '?', '#', '?', '#', '?', '#', '?', '#', '?', '#', '?'},
		damagedGroups: []int{1, 3, 1, 6},
	}

	// ????.#...#... 4,1,1 - 1 arrangement
	record4 := ConditionRecord{
		springs:       []rune{'?', '?', '?', '?', '.', '#', '.', '.', '.', '#', '.', '.', '.'},
		damagedGroups: []int{4, 1, 1},
	}

	// ????.######..#####. 1,6,5 - 4 arrangements
	record5 := ConditionRecord{
		springs:       []rune{'?', '?', '?', '?', '.', '#', '#', '#', '#', '#', '#', '.', '.', '#', '#', '#', '#', '#', '.'},
		damagedGroups: []int{1, 6, 5},
	}

	// ?###???????? 3,2,1 - 10 arrangements
	record6 := ConditionRecord{
		springs:       []rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'},
		damagedGroups: []int{3, 2, 1},
	}

	a1 := findArrangements(record1)
	a2 := findArrangements(record2)
	a3 := findArrangements(record3)
	a4 := findArrangements(record4)
	a5 := findArrangements(record5)
	a6 := findArrangements(record6)

	if a1 != 1 {
		t.Fatalf("Num of arrangements should have been 1, got=%v", a1)
	}

	if a2 != 4 {
		t.Fatalf("Num of arrangements should have been 4, got=%v", a2)
	}

	if a3 != 1 {
		t.Fatalf("Num of arrangements should have been 1, got=%v", a3)
	}

	if a4 != 1 {
		t.Fatalf("Num of arrangements should have been 1, got=%v", a4)
	}

	if a5 != 4 {
		t.Fatalf("Num of arrangements should have been 4, got=%v", a5)
	}

	if a6 != 10 {
		t.Fatalf("Num of arrangements should have been 10, got=%v", a6)
	}
}
