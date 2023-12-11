package puzzles

import (
	"fmt"
	"slices"
)

type direction int

const (
	unspecified direction = iota
	north
	east
	south
	west
)

func (d direction) string() string {
	return [...]string{"Unspecified", "North", "East", "South", "West"}[d]
}

type symbol rune

func (s symbol) validDirections() []direction {
	switch s {
	case 'S':
		return []direction{north, east, south, west}
	case '|':
		return []direction{north, south}
	case '-':
		return []direction{east, west}
	case 'L':
		return []direction{north, east}
	case 'J':
		return []direction{north, west}
	case '7':
		return []direction{south, west}
	case 'F':
		return []direction{east, south}
	case '.':
		return []direction{unspecified}
	default:
		return []direction{}
	}
}

type coords struct {
	x int
	y int
}

type cell struct {
	pipe symbol
	coords
}

func (c cell) print() {
	fmt.Printf("%c - (%d, %d), directions: %v\n", c.pipe, c.x, c.y, c.pipe.validDirections())
}

type cellMap [][]cell

func makeCellMap(pipeMap [][]rune) [][]cell {
	var cellMap [][]cell
	for y, line := range pipeMap {
		cellMap = append(cellMap, make([]cell, 0))
		for x, c := range line {
			cellMap[y] = append(cellMap[y], cell{coords: coords{x: x, y: y}, pipe: symbol(c)})
		}
	}

	return cellMap
}

func (m cellMap) width() int {
	return len(m[0])
}

func (m cellMap) height() int {
	return len(m)
}

type navMap struct {
	cellMap       cellMap
	curr          *cell
	prevDirection direction
}

func (m *navMap) init() {
	for _, y := range m.cellMap {
		for _, x := range y {
			if x.pipe == 'S' {
				m.curr = &x
				break
			}
		}
	}
}

func (m *navMap) walk(d direction) {
	if m.canWalk(d) {
		switch d {
		case north:
			m.curr = m.getNorthCell()
			m.prevDirection = south
		case east:
			m.curr = m.getEastCell()
			m.prevDirection = west
		case south:
			m.curr = m.getSouthCell()
			m.prevDirection = north
		case west:
			m.curr = m.getWestCell()
			m.prevDirection = east
		default:
			if m.canWalk(north) {
				m.walk(north)
			} else if m.canWalk(east) {
				m.walk(east)
			} else if m.canWalk(south) {
				m.walk(south)
			} else if m.canWalk(west) {
				m.walk(west)
			}
		}
	} else {
		fmt.Println("Couldn't walk", d.string())
	}
}

func (m *navMap) followPipe() {
	pipeDirections := m.curr.pipe.validDirections()
	remDirections := slices.DeleteFunc(pipeDirections, func(d direction) bool {
		return m.prevDirection == d
	})

	m.walk(remDirections[0])
}

func (m *navMap) getNorthCell() *cell {
	if m.curr.y-1 >= 0 {
		return &m.cellMap[m.curr.y-1][m.curr.x]
	}
	return nil
}

func (m *navMap) getEastCell() *cell {
	if m.curr.x+1 < m.cellMap.width() {
		return &m.cellMap[m.curr.y][m.curr.x+1]
	}
	return nil
}

func (m *navMap) getSouthCell() *cell {
	if m.curr.y+1 < m.cellMap.height() {
		return &m.cellMap[m.curr.y+1][m.curr.x]
	}
	return nil
}

func (m *navMap) getWestCell() *cell {
	if m.curr.x-1 >= 0 {
		return &m.cellMap[m.curr.y][m.curr.x-1]
	}
	return nil
}

func (m *navMap) canWalk(d direction) bool {
	srcDiretions := m.curr.pipe.validDirections()
	var dstDirections []direction
	switch d {
	case north:
		dst := m.getNorthCell()
		if dst != nil {
			dstDirections = dst.pipe.validDirections()
		}
		return slices.Contains(srcDiretions, north) && slices.Contains(dstDirections, south)
	case east:
		dst := m.getEastCell()
		if dst != nil {
			dstDirections = append(dstDirections, dst.pipe.validDirections()...)
		}
		return slices.Contains(srcDiretions, east) && slices.Contains(dstDirections, west)
	case south:
		dst := m.getSouthCell()
		if dst != nil {
			dstDirections = dst.pipe.validDirections()
		}
		return slices.Contains(srcDiretions, south) && slices.Contains(dstDirections, north)
	case west:
		dst := m.getWestCell()
		if dst != nil {
			dstDirections = dst.pipe.validDirections()
		}
		return slices.Contains(srcDiretions, west) && slices.Contains(dstDirections, east)
	default:
		return true
	}
}

func (m *navMap) findLoop(d direction) []coords {
	var visited []coords

	visited = append(visited, m.curr.coords)

	m.walk(d)
	visited = append(visited, m.curr.coords)

	for {
		m.followPipe()
		if !checkFoundLoop(visited, m.curr.coords) {
			visited = append(visited, m.curr.coords)
			continue
		}
		break
	}
	return visited
}

func checkFoundLoop(visited []coords, next coords) bool {
	return slices.Contains(visited, next)
}

func expandLoop(pipeMap [][]rune, loop []coords) [][]rune {
	var expandedMap [][]rune

	newCols := len(pipeMap[0]) + (len(pipeMap[0]) - 1)
	for y, row := range pipeMap {
		newRow := make([]rune, 0)
		for x, col := range row {
			if slices.Contains(loop, coords{x: x, y: y}) {
				newRow = append(newRow, col)
			} else {
				newRow = append(newRow, '.')
			}

			if x != newCols-1 {
				newRow = append(newRow, ' ')
			}
		}
		expandedMap = append(expandedMap, newRow)

		if y != len(pipeMap)-1 {
			emptyRow := make([]rune, 0)

			for i := 0; i < newCols; i++ {
				emptyRow = append(emptyRow, ' ')
			}
			expandedMap = append(expandedMap, emptyRow)
		}
	}

	return expandedMap
}

func parseDay10(input []string) [][]rune {
	var pipeMap [][]rune

	for _, l := range input {
		asRunes := []rune(l)
		pipeMap = append(pipeMap, asRunes)
	}
	return pipeMap
}

func Day10Pt1(input []string) int {
	pipeMap := parseDay10(input)
	printPipeMap(pipeMap)

	loops := make(map[direction][]coords)
	myMap := navMap{cellMap: makeCellMap(pipeMap)}
	myMap.init()

	var startingDirections []direction
	for _, d := range myMap.curr.pipe.validDirections() {
		if myMap.canWalk(d) {
			startingDirections = append(startingDirections, d)
		}
	}

	for _, d := range startingDirections {
		loops[d] = myMap.findLoop(d)
	}

	steps := 1
	for i := 1; i < len(loops[startingDirections[0]]); i++ {
		if loops[startingDirections[0]][i] == loops[startingDirections[1]][i] {
			break
		}
		steps++
	}

	return steps
}

func Day10Pt2(input []string) int {
	pipeMap := parseDay10(input)

	// ex3 := [][]rune{
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// 	{'.', 'S', '-', '-', '-', '-', '-', '-', '7', '.'},
	// 	{'.', '|', 'F', '-', '-', '-', '-', '7', '|', '.'},
	// 	{'.', '|', '|', '.', '.', '.', '.', '|', '|', '.'},
	// 	{'.', '|', '|', '.', '.', '.', '.', '|', '|', '.'},
	// 	{'.', '|', 'L', '-', '7', 'F', '-', 'J', '|', '.'},
	// 	{'.', '|', '.', '.', '|', '|', '.', '.', '|', '.'},
	// 	{'.', 'L', '-', '-', 'J', 'L', '-', '-', 'J', '.'},
	// 	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	// }

	printPipeMap(pipeMap)

	myMap := navMap{cellMap: makeCellMap(pipeMap)}
	myMap.init()

	var startingDirections []direction
	for _, d := range myMap.curr.pipe.validDirections() {
		if myMap.canWalk(d) {
			startingDirections = append(startingDirections, d)
		}
	}

	loop := myMap.findLoop(startingDirections[0])

	fmt.Println("--------")
	printPipeLoop(loop, pipeMap)

	// expandedMap := expandLoop(ex3, loop)

	// printPipeMap(expandedMap)
	fmt.Println("--------")

	return 0
}

func printPipeMap(pipeMap [][]rune) {
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

func printPipeLoop(loop []coords, pipeMap [][]rune) {
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
		for x, val := range line {
			if slices.Contains(loop, coords{x: x, y: y}) {
				fmt.Printf("%c ", val)
				// fmt.Printf("%c", val)
			} else {
				fmt.Printf("%c ", ' ')
				// fmt.Printf("%c", '▐')
			}
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
