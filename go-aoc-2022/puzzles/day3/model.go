package day3

type Rucksack struct {
	Compartment1 Compartment
	Compartment2 Compartment
}

type Compartment struct {
	Items []Item
}

type Item struct {
	ItemType rune
	Priority int
}

func MakeRucksack(s string) Rucksack {
	numItems := len(s)

	var c1 Compartment
	var c2 Compartment

	for i, ch := range s {
		if i < (numItems / 2) {
			c1.Items = append(c1.Items, Item{
				ItemType: ch,
				Priority: getItemPriority(ch),
			})
		} else {
			c2.Items = append(c2.Items, Item{
				ItemType: ch,
				Priority: getItemPriority(ch),
			})
		}
	}

	return Rucksack{
		Compartment1: c1,
		Compartment2: c2,
	}
}

func getItemPriority(ch rune) int {
	chInt := int(ch)
	if chInt > 96 {
		return chInt - 96
	} else {
		return chInt - 38
	}
}
