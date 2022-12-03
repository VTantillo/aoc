package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readDay1Input() []int {
	f, err := os.Open("./res/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elfMeals := []int{}
	totalElves := 0

	for scanner.Scan() {
		meal := scanner.Text()

		if meal == "" {
			elfMeals = append(elfMeals, -1)
			totalElves++
			continue
		}

		intVar, err := strconv.Atoi(meal)
		if err != nil {
			log.Fatalf("Unexpected input %v\n", err)
		}

		elfMeals = append(elfMeals, intVar)
	}

	fmt.Printf("Number of meals read in: %v\n", len(elfMeals))
	fmt.Printf("Number of elves: %v\n", totalElves)

	return elfMeals
}

func calculateElfCalories(meals []int) []int {
	mealTotals := []int{}

	total := 0
	for _, value := range meals {
		if value < 0 {
			mealTotals = append(mealTotals, total)
			total = 0
			continue
		}

		total += value
	}

	sort.Ints(mealTotals)

	return mealTotals
}

func Day1Step1() int {
	meals := readDay1Input()
	mealTotals := calculateElfCalories(meals)
	max := mealTotals[len(mealTotals)-1]

	return max
}

func Day1Step2() int {
	meals := readDay1Input()
	mealTotals := calculateElfCalories(meals)

	max := 0

	for _, value := range mealTotals[len(mealTotals)-3:] {
		max += value
	}

	return max
}
