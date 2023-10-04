package puzzles

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func readInput() []string {
	file, err := os.Open("./input/2022_day_1_input.txt")
	utils.Check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	file.Close()

	return fileLines
}

func Day1Puzzle() int {

	fileLines := readInput()

	var currentElfCalories int
	var totalCalories int

	var elfCalories []int

	for _, line := range fileLines {

		if len(line) == 0 {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0
			continue
		}

		mealCalories, err := strconv.Atoi(line)
		utils.Check(err)

		currentElfCalories += mealCalories
	}

	slices.SortFunc(elfCalories, func(a, b int) int {
		return b - a
	})

	top3Elves := elfCalories[0:3]

	fmt.Println(top3Elves)

	for _, elf := range top3Elves {
		totalCalories += elf
	}

	return totalCalories
}

func part1() int {
	fileLines := readInput()

	var highestCalories int
	var currentElfCalories int

	for _, line := range fileLines {

		if len(line) == 0 {
			if currentElfCalories > highestCalories {
				highestCalories = currentElfCalories
			}

			currentElfCalories = 0
			continue
		}

		mealCalories, err := strconv.Atoi(line)
		utils.Check(err)

		currentElfCalories += mealCalories
	}

	return highestCalories
}
