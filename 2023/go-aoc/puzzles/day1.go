package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1(input []string) int {
	total := 0
	for j, l := range input {
		var digits strings.Builder
		var letters strings.Builder

		for i := 0; i < len(l); i++ {
			if l[i] >= '0' && l[i] <= '9' {
				digits.WriteByte(l[i])
				break
			}

			letters.WriteByte(l[i])
			found, c := digitWord(letters.String())

			if found {
				digits.WriteRune(c)
				break
			}
		}

		letters.Reset()
		for i := len(l) - 1; i > -1; i-- {
			if l[i] >= '0' && l[i] <= '9' {
				digits.WriteByte(l[i])
				break
			}

			letters.WriteByte(l[i])
			found, c := digitWord(reverseString(letters.String()))

			if found {
				digits.WriteRune(c)
				break
			}
		}

		fmt.Printf("%v: %s - %s \n", j, l, digits.String())
		num, err := strconv.Atoi(digits.String())
		if err != nil {
			panic("It wasn't just numbers")
		}
		total += num

		digits.Reset()
		letters.Reset()
	}

	return total
}

func digitWord(input string) (bool, rune) {
	switch {
	case strings.Contains(input, "one"):
		return true, '1'
	case strings.Contains(input, "two"):
		return true, '2'
	case strings.Contains(input, "three"):
		return true, '3'
	case strings.Contains(input, "four"):
		return true, '4'
	case strings.Contains(input, "five"):
		return true, '5'
	case strings.Contains(input, "six"):
		return true, '6'
	case strings.Contains(input, "seven"):
		return true, '7'
	case strings.Contains(input, "eight"):
		return true, '8'
	case strings.Contains(input, "nine"):
		return true, '9'
	default:
		return false, '0'
	}
}

func reverseString(s string) string {
	r := []rune(s)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}
