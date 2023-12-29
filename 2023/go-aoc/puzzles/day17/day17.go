package day17

import (
	"aoc/utils/term"
	"fmt"
	"math"
	"slices"
	"time"
)

type direction int

const (
	dirNone direction = iota
	dirUp
	dirDown
	dirLeft
	dirRight
)

func (d direction) String() string {
	return [...]string{"none", "up", "down", "left", "right"}[d]
}

func (d direction) printDirection() rune {
	switch d {
	case dirUp:
		return '^'
	case dirRight:
		return '>'
	case dirDown:
		return 'v'
	case dirLeft:
		return '<'
	}

	return '.'
}

type node struct {
	row       int
	col       int
	weight    int
	distance  int
	visited   bool
	neighbors map[direction]*node
}

func (n node) key() string {
	return fmt.Sprintf("%d-%d", n.row, n.col)
}

type neighbor struct {
	key      string
	dir      direction
	distance int
}

type blockMap struct {
	blocks [][]*node
	curr   *node
}

func newBlockMap(weights [][]int) blockMap {
	blocks := make([][]*node, 0)
	for r, row := range weights {
		blockRow := make([]*node, 0)
		for c, col := range row {
			initDistance := math.MaxInt16
			if r == 0 && c == 0 {
				initDistance = 0
			}
			n := node{
				row:       r,
				col:       c,
				weight:    col,
				distance:  initDistance,
				visited:   false,
				neighbors: make(map[direction]*node),
			}
			blockRow = append(blockRow, &n)
		}
		blocks = append(blocks, blockRow)
	}

	for r, row := range blocks {
		for c, node := range row {
			if r-1 >= 0 {
				node.neighbors[dirUp] = blocks[r-1][c]
			}
			if r+1 < len(blocks) {
				node.neighbors[dirDown] = blocks[r+1][c]
			}
			if c-1 >= 0 {
				node.neighbors[dirLeft] = blocks[r][c-1]
			}
			if c+1 < len(blocks[0]) {
				node.neighbors[dirRight] = blocks[r][c+1]
			}
		}
	}

	return blockMap{
		blocks: blocks,
		curr:   blocks[0][0],
	}
}

func (b *blockMap) findPathDikstra(dstKey string) {
	unvisited := make(map[string]bool)
	nodes := make([]*node, 0)
	for _, row := range b.blocks {
		for _, n := range row {
			unvisited[n.key()] = true
			nodes = append(nodes, n)
		}
	}

	fmt.Println("num unvisited:", len(unvisited))

	for unvisited[dstKey] {
		fmt.Printf("Curr key: %s, dist: %d\n", b.curr.key(), b.curr.distance)
		b.printMap()
		for _, n := range b.curr.neighbors {
			if unvisited[n.key()] {
				dist := b.curr.distance + n.weight

				if dist < n.distance {
					n.distance = dist
				}

			}
		}

		b.curr.visited = true
		delete(unvisited, b.curr.key())

		slices.SortFunc(nodes, func(a, b *node) int {
			return a.distance - b.distance
		})

		for _, n := range nodes {
			if unvisited[n.key()] {
				b.curr = n
				break
			}
		}
	}
}

func (b *blockMap) printMap() {
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(b.blocks[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("   ┏")
	for j := 0; j < len(b.blocks[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┓\n")

	for y, line := range b.blocks {
		fmt.Printf("%3d%c", y, '┃')
		for _, val := range line {
			if b.curr.key() == val.key() {
				fmt.Printf("%s%s%d%s", term.RedBackground, term.Black, val.weight, term.Reset)
			} else if val.visited {
				fmt.Printf("%s%s%d%s", term.YellowBackground, term.Black, val.weight, term.Reset)
			} else {
				fmt.Printf("%d", val.weight)
			}
		}
		fmt.Printf("%c\n", '┃')
	}

	fmt.Print("   ┗")
	for j := 0; j < len(b.blocks[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┛\n")
	time.Sleep(250 * time.Millisecond)
}

func Day17(input []string) int {
	blockWeights := parseMap(input)
	m := newBlockMap(blockWeights)

	m.findPathDikstra("12-12")
	m.printMap()

	return 0
}

func parseMap(input []string) [][]int {
	blockMap := make([][]int, 0)
	for _, line := range input {
		row := make([]int, 0)
		for _, c := range line {
			b := 0
			fmt.Sscanf(string(c), "%d", &b)
			row = append(row, b)
		}
		blockMap = append(blockMap, row)
	}
	return blockMap
}
