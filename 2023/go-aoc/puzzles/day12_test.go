package puzzles

import (
	"strings"
	"testing"
)

var springsEx1 string = `#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1`

var springsEx2 string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestDay12Pt1(t *testing.T) {
	lines := strings.Split(springsEx2, "\n")

	result := Day12Pt1(lines)

	if result != 21 {
		t.Fatalf("Result should be 21, got=%v", result)
	}
}
