package puzzles

import (
	"fmt"
	"strings"
)

func Day9Pt1(input []string) int {
	readings := parseDay9(input)

	var nextVals []int
	for _, r := range readings {
		fmt.Println("start line:", r)
		nextVals = append(nextVals, findNextVal(r))
		fmt.Println("---")
	}

	sum := 0
	for _, v := range nextVals {
		sum += v
	}

	return sum
}

func Day9Pt2(input []string) int {
	readings := parseDay9(input)

	var nextVals []int
	for _, r := range readings {
		fmt.Println("start line:", r)
		nextVals = append(nextVals, findPrevVal(r))
		fmt.Println("---")
	}

	sum := 0
	for _, v := range nextVals {
		sum += v
	}

	return sum
}

func parseDay9(input []string) [][]int {
	var allReadings [][]int
	for _, l := range input {
		var lineReadings []int
		for _, r := range strings.Split(l, " ") {
			var reading int
			fmt.Sscanf(r, "%d", &reading)
			lineReadings = append(lineReadings, reading)
		}
		allReadings = append(allReadings, lineReadings)
	}

	return allReadings
}

func diffLine(readings []int) []int {
	var diff []int
	for i := 1; i < len(readings); i++ {
		diff = append(diff, readings[i]-readings[i-1])
	}
	return diff
}

func findNextVal(readings []int) int {
	diff := diffLine(readings)
	fmt.Println("diff line: ", diff)

	if !checkAllZeros(diff) {
		nextVal := findNextVal(diff)
		return readings[len(readings)-1] + nextVal
	} else {
		fmt.Println("Next diff was 0, end is:", readings[len(readings)-1])
		return readings[len(readings)-1]
	}
}

func findPrevVal(readings []int) int {
	diff := diffLine(readings)
	fmt.Println("diff line: ", diff)

	if !checkAllZeros(diff) {
		prevVal := findPrevVal(diff)
		return readings[0] - prevVal
	} else {
		fmt.Println("Next diff was 0, end is:", readings[0])
		return readings[0]
	}
}

func checkAllZeros(line []int) bool {
	for _, l := range line {
		if l != 0 {
			return false
		}
	}
	return true
}
