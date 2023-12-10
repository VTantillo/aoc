package puzzles

import (
	"fmt"
)

func Day10Pt1(input []string) int {
	steps := 0

	pipeMap := parseDay10(input)

	// for y, l := range pipeMap {
	// 	for x, p := range l {
	// 		switch p {
	// 		case 'S':
	// 			fmt.Println("Start", x, y)
	// 		case '|':
	// 			fmt.Println("Vertical pipe:", x, y)
	// 		case '-':
	// 			fmt.Println("Horizontal pipe:", x, y)
	// 		case 'L':
	// 			fmt.Println("N -> E pipe:", x, y)
	// 		case 'J':
	// 			fmt.Println("N -> W pipe:", x, y)
	// 		case '7':
	// 			fmt.Println("S -> W pipe:", x, y)
	// 		case 'F':
	// 			fmt.Println("S -> E pipe:", x, y)
	// 		}
	// 	}
	// }

	printPipeMap(pipeMap)

	return steps
}

func parseDay10(input []string) [][]rune {
	var pipeMap [][]rune

	for _, l := range input {
		asRunes := []rune(l)
		pipeMap = append(pipeMap, asRunes)
	}
	return pipeMap
}

func printPipeMap(pipeMap [][]rune) {
	// Top axis
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(pipeMap[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c ", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c%c", '━', '━')
	}
	fmt.Print("\n")

	for y, line := range pipeMap {
		fmt.Printf("%3d%c", y, '┃')
		for _, val := range line {
			fmt.Printf("%c ", val)
		}
		fmt.Print("\n")
	}
}
