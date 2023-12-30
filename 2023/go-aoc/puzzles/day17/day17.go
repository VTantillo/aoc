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

func (c coords) distance(dst coords) int {
	rowDist := math.Exp((float64(dst.row - c.col)))
	colDist := math.Exp((float64(dst.row - c.col)))
	dist := math.Sqrt(colDist + rowDist)
	return int(dist)
}

type node struct {
	coords
	weight        int
	distance      int
	cost          int
	visited       bool
	neighbors     map[direction]*node
	prevDir       direction
	prev          *node
	straightCount int
	index         int
}

func (n *node) calcCost(dst coords) {
	normalDist := n.coords.distance(dst)
	dist := normalDist + n.distance
	// dist := n.distance
	n.cost = dist + n.index

	// if n.straightCount >= 3 {
	// 	n.cost = n.cost * 100000000
	// }
}

type pathQueue []*node

func (p pathQueue) Len() int { return len(p) }

func (p pathQueue) IsEmpty() bool { return len(p) == 0 }

func (p *pathQueue) Push(n *node) {
	*p = append(*p, n)
	n.index = p.Len()

	slices.SortFunc(*p, func(a, b *node) int {
		return a.cost - b.cost
	})
	slices.DeleteFunc(*p, func(a *node) bool {
		return a.visited
	})
}

func (p *pathQueue) Fix() {
	if p.Len() > 1 {
		slices.SortFunc(*p, func(a, b *node) int {
			return a.cost - b.cost
		})
	}
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
	showPath bool
	delay    time.Duration
	queue    pathQueue
}

func (b *blockMap) findPathAStar(src, dst coords) {
	nodes := make(map[string]*node)
	for _, row := range b.blocks {
		for _, n := range row {
			nodes[n.key()] = n
		}
	}

	srcNode := nodes[src.key()]
	b.queue.Push(srcNode)

	for !b.queue.IsEmpty() {
		b.curr = b.queue.Pop()

		if b.curr.key() == dst.key() {
			b.curr.visited = true
			b.printMap()
			fmt.Println("Reached destination")
			break
		}

		if b.curr.visited {
			continue
		}

		for d, n := range b.curr.neighbors {
			if !n.visited {
				dist := b.curr.distance + n.weight

				if dist < n.distance {
					n.distance = dist
					n.prev = b.curr
					n.prevDir = d.reverse()
					if b.curr.prevDir == n.prevDir || b.curr.key() == src.key() {
						n.straightCount = b.curr.straightCount + 1
					}
				}
				if n.straightCount < 3 {
					n.calcCost(dst)
					b.queue.Push(n)
				}
			}
		}
		b.curr.visited = true
		b.printMap()
	}

	b.curr = nodes[dst.key()]
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
				fmt.Printf("%s%s%d%s", term.BlueBackground, term.Black, val.weight, term.Reset)
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
	fmt.Printf("\tkey: %s\n", b.curr.key())
	fmt.Printf("\tdistance: %d\n", b.curr.distance)
	fmt.Printf("\tcost: %d\n", b.curr.distance)
	fmt.Printf("\tstraight count: %d\n", b.curr.straightCount)
	fmt.Printf("\tprev direction: %s\n", b.curr.prevDir.String())
	fmt.Print("\n")
	fmt.Printf("In Queue: %d\n", b.queue.Len())
	for _, q := range b.queue {
		fmt.Print("[\t")
		fmt.Printf("key: %s, cost: %d, dist: %d, straight: %d, ", q.key(), q.cost, q.distance, q.straightCount)
		fmt.Print("\t]\n")
	}

	time.Sleep(b.delay * time.Millisecond)
}

func Day17(input []string, showPath bool) int {
	blockWeights := parseMap(input)
	blocks := make([][]*node, 0)
	for r, row := range blockWeights {
		blockRow := make([]*node, 0)
		for c, col := range row {
			initDistance := 90000000
			// straightCount := 0
			if r == 0 && c == 0 {
				initDistance = 0
				// straightCount = 1
			}
			n := node{
				coords: coords{
					row: r,
					col: c,
				},
				weight:        col,
				distance:      initDistance,
				visited:       false,
				straightCount: 0,
				neighbors:     make(map[direction]*node),
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
	b.delay = 200

	dst := coords{
		row: len(blockWeights) - 1,
		col: len(blockWeights[0]) - 1,
	}

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
