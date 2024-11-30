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

func (d direction) reverse() direction {
	switch d {
	case dirUp:
		return dirDown
	case dirRight:
		return dirLeft
	case dirDown:
		return dirUp
	case dirLeft:
		return dirRight
	}
	return dirNone
}

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

type coords struct {
	row int
	col int
}

func (c coords) key() string {
	return fmt.Sprintf("%d-%d", c.row, c.col)
}

func distance(src, dst coords) int {
	rowDist := math.Pow(float64(dst.row-src.row), 2)
	colDist := math.Pow(float64(dst.col-src.col), 2)
	dist := math.Sqrt(colDist + rowDist)
	return int(dist)
}

type node struct {
	coords
	weight        int
	distance      int
	straightCount int
	step          int
	visited       bool
	neighbors     map[direction]*node
	prevDir       direction
	cantGo        []direction
	prev          *node
}

type pathQueue []*node

func (p pathQueue) Len() int { return len(p) }

func (p pathQueue) IsEmpty() bool { return len(p) == 0 }

func (p *pathQueue) Push(n *node) {
	*p = append(*p, n)
	slices.SortStableFunc(*p, func(a, b *node) int {
		return a.distance - b.distance
	})
}

func (p *pathQueue) Pop() *node {
	curr := *p
	item := curr[0]
	curr[0] = nil
	*p = curr[1:]
	return item
}

type blockMap struct {
	blocks   [][]*node
	curr     *node
	queue    pathQueue
	showPath bool
	delay    time.Duration
}

func (b *blockMap) findPathAStar(src, dst coords) {
	b.queue.Push(b.blocks[src.row][src.col])

	steps := 0
	for !b.queue.IsEmpty() {
		b.curr = b.queue.Pop()

		// if b.curr.key() == dst.key() {
		// 	b.curr.visited = true
		// 	b.printMap()
		// 	fmt.Println("Reached destination")
		// 	break
		// }

		if b.curr.visited || b.curr.straightCount > 3 {
			continue
		}

		steps++
		for d, n := range b.curr.neighbors {
			if !n.visited {
				dist := b.curr.distance + n.weight
				count := 0
				if b.curr.prevDir == d.reverse() || b.curr.prevDir == dirNone {
					count = b.curr.straightCount + 1
				}

				if dist <= n.distance && count < 3 {
					n.distance = dist
					n.prev = b.curr
					n.prevDir = d.reverse()
					n.straightCount = count
					n.step = steps

					// noGo := []direction{n.prevDir}
					// if b.curr.prev != nil && b.curr.prevDir == d.reverse() &&
					// 	(b.curr.prevDir == b.curr.prev.prevDir || b.curr.prev.prevDir == dirNone) {
					// 	noGo = append(noGo, d)
					// }
					// n.cantGo = noGo

					b.queue.Push(n)
				}

			}
		}
		b.curr.visited = true
		b.printMap()
	}

	b.curr = b.blocks[dst.row][dst.col]
	b.printMap()
}

func (b *blockMap) printMap() {
	if !b.showPath {
		return
	}

	fmt.Print(term.ClearScreen)

	pathKeys := make([]string, 0)

	pathPtr := b.curr
	for pathPtr != nil {
		pathKeys = append(pathKeys, pathPtr.key())
		pathPtr = pathPtr.prev
	}

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
			} else if slices.Contains(pathKeys, val.key()) {
				fmt.Printf("%s%s%c%s", term.BlueBackground, term.Black, val.prevDir.reverse().printDirection(), term.Reset)
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

	fmt.Print("Current node:\n")
	fmt.Printf("   key: %s\n", b.curr.key())
	fmt.Printf("   distance: %d\n", b.curr.distance)
	fmt.Printf("   straight count: %d\n", b.curr.straightCount)
	if b.curr.prev != nil {
		fmt.Printf("   prev: %s\n", b.curr.prev.key())
	}
	fmt.Printf("   prev direction: %s\n", b.curr.prevDir.String())
	fmt.Printf("   can't go: %s\n", b.curr.cantGo)
	fmt.Print("\n")
	fmt.Printf("in queue: %d\n", b.queue.Len())
	// for _, q := range b.queue {
	// 	fmt.Print("[   ")
	// 	fmt.Printf("key: %s, weight: %d dist: %d, straight: %d  ", q.key(), q.weight, q.distance, q.straightCount)
	// 	fmt.Print("   ]\n")
	// }

	time.Sleep(b.delay * time.Millisecond)
}

func Day17(input []string, showPath bool) int {
	blockWeights := parseMap(input)
	blocks := make([][]*node, 0)
	dst := coords{
		row: len(blockWeights) - 1,
		col: len(blockWeights[0]) - 1,
	}
	for r, row := range blockWeights {
		blockRow := make([]*node, 0)
		for c, col := range row {
			initDistance := 90000000
			if r == 0 && c == 0 {
				initDistance = 0
			}
			src := coords{
				row: r,
				col: c,
			}
			n := node{
				coords:    src,
				weight:    col,
				distance:  initDistance,
				visited:   false,
				neighbors: make(map[direction]*node),
				cantGo:    make([]direction, 0),
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

	b := blockMap{
		blocks:   blocks,
		curr:     blocks[0][0],
		showPath: false,
		queue:    make(pathQueue, 0),
	}

	b.showPath = showPath
	b.delay = 100

	b.findPathAStar(coords{0, 0}, dst)

	return b.curr.distance
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
