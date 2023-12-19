package puzzles

import (
	"fmt"
	"slices"
	"time"
)

var Pt2Ex1 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

var Pt2Ex2 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

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

type side int

const (
	left side = iota + 1
	right
)

func (s side) string() string {
	return [...]string{"Left", "Right"}[s]
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

func (s symbol) sideDirections(side side, to direction) direction {
	switch s {
	case '|':
		switch to {
		case north:
			switch side {
			case left:
				return west
			case right:
				return east
			}
		case south:
			switch side {
			case left:
				return east
			case right:
				return west
			}
		}
	case '-':
		switch to {
		case east:
			switch side {
			case left:
				return north
			case right:
				return south
			}
		case west:
			switch side {
			case left:
				return south
			case right:
				return north
			}
		}
	case 'L':
		return unspecified
	case 'J':
		return unspecified
	case '7':
		return unspecified
	case 'F':
		return unspecified
	default:
		return unspecified
	}

	return unspecified
}

type coords struct {
	x int
	y int
}

type cell struct {
	pipe symbol
	coords
}

type cellMap [][]*cell

func makeCellMap(pipeMap [][]rune) [][]*cell {
	var cellMap [][]*cell
	for y, line := range pipeMap {
		cellMap = append(cellMap, make([]*cell, 0))
		for x, c := range line {
			cellMap[y] = append(cellMap[y], &cell{coords: coords{x: x, y: y}, pipe: symbol(c)})
		}
	}

	return cellMap
}

func makeLoopMap(pipeMap [][]rune, loop []coords) [][]cell {
	var cellMap [][]cell
	for y, line := range pipeMap {
		cellMap = append(cellMap, make([]cell, 0))
		for x, c := range line {
			if slices.Contains(loop, coords{x: x, y: y}) {
				cellMap[y] = append(cellMap[y], cell{coords: coords{x: x, y: y}, pipe: symbol(c)})
			} else {
				cellMap[y] = append(cellMap[y], cell{coords: coords{x: x, y: y}, pipe: symbol('.')})
			}
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
				m.curr = x
				break
			}
		}
	}
}

func (m *navMap) printMap() {
	// Top axis
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(m.cellMap[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("   ┏")
	for j := 0; j < len(m.cellMap[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┓\n")

	for y, line := range m.cellMap {
		fmt.Printf("%3d%c", y, '┃')
		for _, val := range line {
			if val == m.curr {
				fmt.Printf("%c", '╳')
			} else {
				fmt.Printf("%c", swapSymbol(rune(val.pipe)))
			}
		}
		fmt.Printf("%c\n", '┃')
	}

	fmt.Print("   ┗")
	for j := 0; j < len(m.cellMap[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┛\n")
}

func (m *navMap) walk(d direction, showMap bool) {
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
				m.walk(north, showMap)
			} else if m.canWalk(east) {
				m.walk(east, showMap)
			} else if m.canWalk(south) {
				m.walk(south, showMap)
			} else if m.canWalk(west) {
				m.walk(west, showMap)
			}
		}
	} else {
		fmt.Println("Couldn't walk", d.string())
	}

	if showMap {
		fmt.Print("\033c")
		m.printMap()
		time.Sleep(250 * time.Millisecond)

	}
}

func (m *navMap) followPipe(showMap bool) {
	pipeDirections := m.curr.pipe.validDirections()
	remDirections := slices.DeleteFunc(pipeDirections, func(d direction) bool {
		return m.prevDirection == d
	})

	m.walk(remDirections[0], showMap)
}

func (m *navMap) getNorthCell() *cell {
	if m.curr.y-1 >= 0 {
		return m.cellMap[m.curr.y-1][m.curr.x]
	}
	return nil
}

func (m *navMap) getEastCell() *cell {
	if m.curr.x+1 < m.cellMap.width() {
		return m.cellMap[m.curr.y][m.curr.x+1]
	}
	return nil
}

func (m *navMap) getSouthCell() *cell {
	if m.curr.y+1 < m.cellMap.height() {
		return m.cellMap[m.curr.y+1][m.curr.x]
	}
	return nil
}

func (m *navMap) getWestCell() *cell {
	if m.curr.x-1 >= 0 {
		return m.cellMap[m.curr.y][m.curr.x-1]
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

	m.walk(d, false)
	visited = append(visited, m.curr.coords)

	for {
		m.followPipe(false)
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

func (m *navMap) clearLoop(loop []coords) {
	for _, line := range m.cellMap {
		for _, cell := range line {
			if !slices.Contains(loop, cell.coords) {
				cell.pipe = symbol('.')
			}
		}
	}
}

func (m *navMap) findInner(loop []coords, side side, d direction) {
	fmt.Println("Finding inner loop")
	start := m.curr.coords

	m.walk(d, true)

	for start.x != m.curr.x || start.y != m.curr.y {
		pipeDirections := m.curr.pipe.validDirections()
		remDirections := slices.DeleteFunc(pipeDirections, func(d direction) bool {
			return m.prevDirection == d
		})
		checkDir := m.curr.pipe.sideDirections(side, remDirections[0])

		fmt.Println(checkDir.string())

		switch checkDir {
		case north:
			cell := m.getNorthCell()
			if cell != nil && cell.pipe == '.' {
				cell.pipe = 'I'
			}
		case south:
			cell := m.getSouthCell()
			if cell != nil && cell.pipe == '.' {
				cell.pipe = 'I'
			}
		case east:
			cell := m.getEastCell()
			if cell != nil && cell.pipe == '.' {
				cell.pipe = 'I'
			}
		case west:
			cell := m.getWestCell()
			if cell != nil && cell.pipe == '.' {
				cell.pipe = 'I'
			}
		}

		m.followPipe(true)
	}
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

	myMap := navMap{cellMap: makeCellMap(pipeMap)}
	myMap.init()

	var startingDirections []direction
	for _, d := range myMap.curr.pipe.validDirections() {
		if myMap.canWalk(d) {
			startingDirections = append(startingDirections, d)
		}
	}

	loop := myMap.findLoop(startingDirections[0])
	myMap.clearLoop(loop)

	myMap.findInner(loop, left, startingDirections[0])

	return 0
}

func parseDay10(input []string) [][]rune {
	var pipeMap [][]rune

	for _, l := range input {
		asRunes := []rune(l)
		pipeMap = append(pipeMap, asRunes)
	}
	return pipeMap
}

// loopMap := navMap{cellMap: makeLoopMap(pipeMap, loop)}
// loopMap.printMap()
// printPipeLoop(loop, pipeMap)

func swapSymbol(s rune) rune {
	switch s {
	case 'S':
		return '█'
	case '|':
		return '│'
	case '-':
		return '─'
	case 'L':
		return '└'
	case 'J':
		return '┘'
	case '7':
		return '┐'
	case 'F':
		return '┌'
	case '.':
		return '.'
	default:
		return s
	}
}

func printPipeMap(pipeMap [][]rune) {
	// Top axis
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(pipeMap[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("\n")

	for y, line := range pipeMap {
		fmt.Printf("%3d%c", y, '┃')
		for _, val := range line {
			fmt.Printf("%c", swapSymbol(val))
		}
		fmt.Printf("%c\n", '┃')
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("\n")
}

func printPipeLoop(loop []coords, pipeMap [][]rune) {
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(pipeMap[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("\n")

	for y, line := range pipeMap {
		fmt.Printf("%3d%c", y, '┃')
		for x, val := range line {
			if slices.Contains(loop, coords{x: x, y: y}) {
				fmt.Printf("%c", swapSymbol(val))
			} else {
				fmt.Printf("%c", '╳')
			}
		}
		fmt.Printf("%c\n", '┃')
	}

	fmt.Print("    ")
	for j := 0; j < len(pipeMap[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("\n")
}
