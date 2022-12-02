package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readInput() []int {
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
	meals := readInput()
	mealTotals := calculateElfCalories(meals)
	max := mealTotals[len(mealTotals)-1]

	// 	meal := scanner.Text()

	// 	if meal == "" {
	// 		fmt.Printf("Elf's total was %v\n", elfTotal)
	// 		elfCalories = append(elfCalories, elfTotal)
	// 		elfTotal = 0
	// 		continue
	// 	}

	// 	intVar, err := strconv.Atoi(meal)
	// 	if err != nil {
	// 		log.Fatalf("Unexpected input %v\n", err)
	// 	}

	// 	elfTotal += intVar
	// }

	// max := 0
	// for i, c := range elfCalories {
	// 	fmt.Printf("Elf %v had %v calories\n", i, c)
	// 	if max < c {
	// 		max = c
	// 	}
	// }

	return max
}

func Day1Step2() int {
	meals := readInput()
	mealTotals := calculateElfCalories(meals)

	max := 0

	for _, value := range mealTotals[len(mealTotals)-3:] {
		max += value
	}

	return max
}
