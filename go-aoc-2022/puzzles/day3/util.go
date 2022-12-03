package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput() []string {
	f, err := os.Open("./res/day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rucksacks := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, line)
	}

	fmt.Println("Number of rucksacks:", len(rucksacks))

	return rucksacks
}
