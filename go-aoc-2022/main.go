package main

import (
	"fmt"

	"go-aoc-2022/puzzles"
)

func main() {
	step1 := puzzles.Day1Step1()

	fmt.Printf("Highest total: %v\n", step1)

	step2 := puzzles.Day1Step2()

	fmt.Printf("Highest 3 elves total: %v\n", step2)
}
