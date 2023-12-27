package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	file.Close()

	return fileLines
}

func ReadString(input string) []string {
	inputLines := make([]string, 0)
	inputLines = append(inputLines, strings.Split(input, "\n")...)
	return inputLines
}

func PrintRuneMap(pipeMap [][]rune) {
	// Top axis
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(pipeMap[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c ", index[i])
			// fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c%c", '━', '━')
		// fmt.Printf("%c", '━' )
	}
	fmt.Print("\n")

	for y, line := range pipeMap {
		fmt.Printf("%3d%c", y, '┃')
		for _, val := range line {
			fmt.Printf("%c ", val)
			// fmt.Printf("%c", val)
		}
		fmt.Printf("%c\n", '┃')
		// fmt.Print("\n")
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c%c", '━', '━')
		// fmt.Printf("%c", '━' )
	}
	fmt.Print("\n")
}
