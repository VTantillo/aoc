package day12

import (
	"slices"
	"strings"
	"testing"
)

var springsEx = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

var kr1 = ConditionRecord{
	Springs:       []rune{'#', '.', '#', '.', '#', '#', '#'},
	DamagedGroups: []int{1, 1, 3},
}

var kr2 = ConditionRecord{
	Springs:       []rune{'.', '#', '.', '.', '.', '#', '.', '.', '.', '.', '#', '#', '#', '.'},
	DamagedGroups: []int{1, 1, 3},
}

// ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
var kr3 = ConditionRecord{
	Springs:       []rune{'.', '#', '.', '#', '#', '#', '.', '#', '.', '#', '#', '#', '#', '#', '#'},
	DamagedGroups: []int{1, 3, 1, 6},
}

// ????.#...#... 4,1,1 - 1 arrangement
var kr4 = ConditionRecord{
	Springs:       []rune{'#', '#', '#', '#', '.', '#', '.', '.', '.', '#', '.', '.', '.'},
	DamagedGroups: []int{4, 1, 1},
}

// ????.######..#####. 1,6,5 - 4 arrangements
var kr5 = ConditionRecord{
	Springs:       []rune{'#', '.', '.', '.', '.', '#', '#', '#', '#', '#', '#', '.', '.', '#', '#', '#', '#', '#', '.'},
	DamagedGroups: []int{1, 6, 5},
}

// ?###???????? 3,2,1 - 10 arrangements
var kr6 = ConditionRecord{
	Springs:       []rune{'.', '#', '#', '#', '.', '#', '#', '.', '.', '.', '.', '#'},
	DamagedGroups: []int{3, 2, 1},
}

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

	if !slices.Equal(record1.Springs, springs1) {
		t.Fatalf("Springs should have been %v, got=%v", springs1, record1.Springs)
	}
	if !slices.Equal(record1.DamagedGroups, groups1) {
		t.Fatalf("damaged groups should have been %v, got=%v", groups1, record1.DamagedGroups)
	}

	springs2 := []rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'}
	groups2 := []int{1, 1, 3}

	if !slices.Equal(record2.Springs, springs2) {
		t.Fatalf("Springs should have been %v, got=%v", springs2, record2.Springs)
	}
	if !slices.Equal(record2.DamagedGroups, groups2) {
		t.Fatalf("Damaged groups should have been %v, got=%v", groups2, record2.DamagedGroups)
	}

	springs3 := []rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'}
	groups3 := []int{3, 2, 1}

	if !slices.Equal(record3.Springs, springs3) {
		t.Fatalf("Springs should have been %v, got=%v", springs3, record3.Springs)
	}
	if !slices.Equal(record3.DamagedGroups, groups3) {
		t.Fatalf("Damaged groups should have been %v, got=%v", groups3, record3.DamagedGroups)
	}
}

func TestIsRecordValid(t *testing.T) {
	kr1Valid := isRecordValid(kr1.Springs, kr1.DamagedGroups)
	kr2Valid := isRecordValid(kr2.Springs, kr2.DamagedGroups)
	kr3Valid := isRecordValid(kr3.Springs, kr3.DamagedGroups)
	kr4Valid := isRecordValid(kr4.Springs, kr4.DamagedGroups)
	kr5Valid := isRecordValid(kr5.Springs, kr5.DamagedGroups)
	kr6Valid := isRecordValid(kr6.Springs, kr6.DamagedGroups)

	if !kr1Valid {
		t.Fatalf("Known record should have been valid, got=%v", kr1Valid)
	}
	if !kr2Valid {
		t.Fatalf("Known record should have been valid, got=%v", kr2Valid)
	}
	if !kr3Valid {
		t.Fatalf("Known record should have been valid, got=%v", kr3Valid)
	}
	if !kr4Valid {
		t.Fatalf("Known record should have been valid, got=%v", kr4Valid)
	}
	if !kr5Valid {
		t.Fatalf("Known record should have been valid, got=%v", kr5Valid)
	}
	if !kr6Valid {
		t.Fatalf("Known record should have been valid, got=%v", kr6Valid)
	}
}

func TestFindArrangements(t *testing.T) {
	// ???.### 1,1,3 - 1 arrangement
	record1 := ConditionRecord{
		Springs:       []rune{'?', '?', '?', '.', '#', '#', '#'},
		DamagedGroups: []int{1, 1, 3},
	}

	// .??..??...?##. 1,1,3 - 4 arrangements
	// record2 := ConditionRecord{
	// 	Springs:       []rune{'.', '?', '?', '.', '.', '?', '?', '.', '.', '.', '?', '#', '#', '.'},
	// 	DamagedGroups: []int{1, 1, 3},
	// }
	//
	// // ?#?#?#?#?#?#?#? 1,3,1,6 - 1 arrangement
	// record3 := ConditionRecord{
	// 	Springs:       []rune{'?', '#', '?', '#', '?', '#', '?', '#', '?', '#', '?', '#', '?', '#', '?'},
	// 	DamagedGroups: []int{1, 3, 1, 6},
	// }
	//
	// // ????.#...#... 4,1,1 - 1 arrangement
	// record4 := ConditionRecord{
	// 	Springs:       []rune{'?', '?', '?', '?', '.', '#', '.', '.', '.', '#', '.', '.', '.'},
	// 	DamagedGroups: []int{4, 1, 1},
	// }
	//
	// // ????.######..#####. 1,6,5 - 4 arrangements
	// record5 := ConditionRecord{
	// 	Springs:       []rune{'?', '?', '?', '?', '.', '#', '#', '#', '#', '#', '#', '.', '.', '#', '#', '#', '#', '#', '.'},
	// 	DamagedGroups: []int{1, 6, 5},
	// }
	//
	// // ?###???????? 3,2,1 - 10 arrangements
	// record6 := ConditionRecord{
	// 	Springs:       []rune{'?', '#', '#', '#', '?', '?', '?', '?', '?', '?', '?', '?'},
	// 	DamagedGroups: []int{3, 2, 1},
	// }

	a1 := record1.FindArrangements()
	// a2 := record2.FindArrangements()
	// a3 := record3.FindArrangements()
	// a4 := record4.FindArrangements()
	// a5 := record5.FindArrangements()
	// a6 := record6.FindArrangements()

	if a1 != 1 {
		t.Fatalf("Num of arrangements should have been 1, got=%v", a1)
	}

	// if a2 != 4 {
	// 	t.Fatalf("Num of arrangements should have been 4, got=%v", a2)
	// }
	//
	// if a3 != 1 {
	// 	t.Fatalf("Num of arrangements should have been 1, got=%v", a3)
	// }
	//
	// if a4 != 1 {
	// 	t.Fatalf("Num of arrangements should have been 1, got=%v", a4)
	// }
	//
	// if a5 != 4 {
	// 	t.Fatalf("Num of arrangements should have been 4, got=%v", a5)
	// }
	//
	// if a6 != 10 {
	// 	t.Fatalf("Num of arrangements should have been 10, got=%v", a6)
	// }
}
