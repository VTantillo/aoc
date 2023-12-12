package puzzles

import (
	"aoc/utils"
	"fmt"
)

type galaxy struct {
	coords
	num int
}

type universe struct {
	starMap  [][]rune
	galaxies *[]galaxy
}

func (u *universe) findGalaxies() {
	var galaxies []galaxy
	numGalaxies := 0

	for row, line := range u.starMap {
		for col, val := range line {
			if val == '#' {
				galaxies = append(galaxies, galaxy{
					num: numGalaxies, coords: coords{x: col, y: row},
				})
				numGalaxies++
			}
		}
	}

	u.galaxies = &galaxies
}

func (u *universe) findDistanceSum(factor int) int {
	var distances []int
	for _, curr := range *u.galaxies {
		for _, g := range *u.galaxies {
			d := u.findDistance(curr.coords, g.coords, factor)
			if curr.num < g.num {
				distances = append(distances, d)
			}
		}
	}

	sum := 0
	for _, d := range distances {
		sum += d
	}
	return sum
}

func (u *universe) findDistance(a coords, b coords, factor int) int {
	var x, y, emptyRows, emptyCols int
	if a.x > b.x {
		x = a.x - b.x
	} else {
		x = b.x - a.x
	}

	if a.y > b.y {
		y = a.y - b.y
	} else {
		y = b.y - a.y
	}

	emptyRows = u.getEmptySpaceRows(a, b)
	emptyCols = u.getEmptySpaceCols(a, b)

	expandedRows := emptyRows * factor
	expandedCols := emptyCols * factor

	x = x - emptyCols
	y = y - emptyRows

	return x + y + expandedRows + expandedCols
}

func (u *universe) getEmptySpaceRows(a coords, b coords) int {
	var yStart, yEnd, rows int

	if a.y > b.y {
		yStart = b.y
		yEnd = a.y
	} else {
		yStart = a.y
		yEnd = b.y
	}

	var shouldExpand bool
	for y := yStart; y < yEnd; y++ {
		shouldExpand = true
		for _, val := range u.starMap[y] {
			if val != '.' {
				shouldExpand = false
			}
		}

		if shouldExpand {
			rows++
		}
	}

	return rows
}

func (u *universe) getEmptySpaceCols(a coords, b coords) int {
	var xStart, xEnd, cols int

	if a.x > b.x {
		xStart = b.x
		xEnd = a.x
	} else {
		xStart = a.x
		xEnd = b.x
	}

	var shouldExpand bool
	for x := xStart; x < xEnd; x++ {
		shouldExpand = true
		for y := 0; y < len(u.starMap); y++ {
			if u.starMap[y][x] != '.' {
				shouldExpand = false
			}
		}

		if shouldExpand {
			cols++
		}
	}

	return cols
}

func Day11Pt1(input []string) int {
	starMap := parseDay11(input)

	universe := universe{
		starMap: starMap,
	}

	universe.findGalaxies()
	utils.PrintRuneMap(universe.starMap)

	sum := universe.findDistanceSum(2)

	return sum
}

func Day11Pt2(input []string, factor int) int {
	fmt.Println(factor)
	starMap := parseDay11(input)

	universe := universe{
		starMap: starMap,
	}

	universe.findGalaxies()
	utils.PrintRuneMap(universe.starMap)

	sum := universe.findDistanceSum(factor)

	return sum
}

func parseDay11(input []string) [][]rune {
	var universe [][]rune
	for _, l := range input {
		lineRunes := []rune(l)
		universe = append(universe, lineRunes)
	}
	return universe
}
