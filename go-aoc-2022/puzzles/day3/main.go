package day3

import "fmt"

func Step1() {
	inputLines := ReadInput()
	rucksacks := []Rucksack{}

	for _, l := range inputLines {
		rucksacks = append(rucksacks, MakeRucksack(l))
	}

	duplicates := []Item{}
	for _, r := range rucksacks {
		duplicates = append(duplicates, findDuplicateItem(r))
	}

	fmt.Println("Duplicate items sum:", sumPriorities(duplicates))
}

func Step2() {
	// NOTE: This one would be fun to come back and use go routines for
	inputLines := ReadInput()
	rucksacks := []Rucksack{}

	for _, l := range inputLines {
		rucksacks = append(rucksacks, MakeRucksack(l))
	}

	groups := makeGroups(rucksacks)

	badges := []Item{}
	for _, g := range groups {
		badges = append(badges, findBadge(g))
	}

	fmt.Println("Badges sum", sumPriorities(badges))
}

func findBadge(group []Rucksack) Item {
	rucksackItems := []Item{}
	rucksackItems = append(rucksackItems, group[0].Compartment1.Items...)
	rucksackItems = append(rucksackItems, group[0].Compartment2.Items...)

	for _, i := range rucksackItems {
		if checkRucksack(group[1], i) && checkRucksack(group[2], i) {
			return i
		}
	}

	return Item{
		ItemType: '0',
		Priority: 0,
	}
}

func checkRucksack(r Rucksack, item Item) bool {
	for _, i := range r.Compartment1.Items {
		if item.ItemType == i.ItemType {
			return true
		}
	}

	for _, i := range r.Compartment2.Items {
		if item.ItemType == i.ItemType {
			return true
		}
	}

	return false
}

func makeGroups(rucksacks []Rucksack) [][]Rucksack {
	groups := [][]Rucksack{}

	group := []Rucksack{}
	for i, r := range rucksacks {
		if i != 0 && i%3 == 0 {
			groups = append(groups, group)
			group = []Rucksack{r}
		} else {
			group = append(group, r)
		}
	}

	// forgot the last group
	groups = append(groups, group)

	fmt.Println("Num groups", len(groups))
	return groups
}

func findDuplicateItem(r Rucksack) Item {
	for _, i := range r.Compartment1.Items {
		for _, j := range r.Compartment2.Items {
			if i.ItemType == j.ItemType {
				return i
			}
		}
	}

	return Item{
		ItemType: '0',
		Priority: 0,
	}
}

func sumPriorities(items []Item) int {
	sum := 0
	for _, i := range items {
		sum += i.Priority
	}

	return sum
}
