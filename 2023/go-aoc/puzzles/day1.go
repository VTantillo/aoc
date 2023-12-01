package puzzles

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day1(input []string) int {
	var lineDigits []string
	for _, l := range input {
		var digits strings.Builder
		for _, c := range l {
			if unicode.IsDigit(c) {
				digits.WriteRune(c)
			}
		}
		lineDigits = append(lineDigits, digits.String())
		digits.Reset()
	}

	total := 0
	for i, d := range lineDigits {
		var numBuilder strings.Builder
		numBuilder.WriteByte(d[0])
		numBuilder.WriteByte(d[len(d)-1])

		num, err := strconv.Atoi(numBuilder.String())
		if err != nil {
			panic("It wasn't just numbers")
		}

		total += num
		fmt.Printf("%v - num: %v, total: %v \n", i, num, total)
	}
	return total
}
