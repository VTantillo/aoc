package day16

import (
	"aoc/utils/term"
	"fmt"
	"slices"
	"strconv"
	"time"
)

type Direction int

const (
	DirUnspecified Direction = iota
	DirUp
	DirRight
	DirDown
	DirLeft
)

func (d Direction) String() string {
	return [...]string{"Unspecified", "Up", "Right", "Down", "Left"}[d]
}

func (d Direction) PrintDirection() rune {
	switch d {
	case DirUp:
		return '^'
	case DirRight:
		return '>'
	case DirDown:
		return 'v'
	case DirLeft:
		return '<'
	}

	return '.'
}

type Coords struct {
	Row int
	Col int
}

func (c Coords) String() string {
	return strconv.Itoa(c.Row) + "," + strconv.Itoa(c.Col)
}

type Step struct {
	Coords
	Direction
}

type Beam struct {
	Coords
	Dir     Direction
	IsStuck bool
}

func (b *Beam) PathKey() string {
	return b.String() + "-" + b.Dir.String()
}

func (b *Beam) MoveBeam(d Direction) {
	b.Dir = d
	switch b.Dir {
	case DirUp:
		b.Row -= 1
	case DirRight:
		b.Col += 1
	case DirDown:
		b.Row += 1
	case DirLeft:
		b.Col -= 1
	}
}

type Tile struct {
	Coords
	Tile        rune
	IsEnergized bool
}

type Grid struct {
	Tiles    [][]*Tile
	Beams    []*Beam
	BeamPath map[string]bool
}

func (g *Grid) SimulateConfig() int {
	for len(g.Beams) > 0 {
		g.UpdateBeams()
		g.PruneBeams()
	}

	return g.CountEnergizedTiles()
}

func (g *Grid) UpdateBeams() {
	newBeams := make([]*Beam, 0)
	for _, b := range g.Beams {
		bTile := g.Tiles[b.Row][b.Col]
		switch bTile.Tile {
		case '.':
			g.MoveBeam(b, b.Dir)
		case '/':
			switch b.Dir {
			case DirUp:
				g.MoveBeam(b, DirRight)
			case DirRight:
				g.MoveBeam(b, DirUp)
			case DirDown:
				g.MoveBeam(b, DirLeft)
			case DirLeft:
				g.MoveBeam(b, DirDown)
			}
		case '\\':
			switch b.Dir {
			case DirUp:
				g.MoveBeam(b, DirLeft)
			case DirRight:
				g.MoveBeam(b, DirDown)
			case DirDown:
				g.MoveBeam(b, DirRight)
			case DirLeft:
				g.MoveBeam(b, DirUp)
			}
		case '-':
			if b.Dir == DirUp || b.Dir == DirDown {
				g.SplitBeam(b, b.Dir, '-')
			} else {
				g.MoveBeam(b, b.Dir)
			}
		case '|':
			if b.Dir == DirLeft || b.Dir == DirRight {
				g.SplitBeam(b, b.Dir, '|')
			} else {
				g.MoveBeam(b, b.Dir)
			}
		}
	}
	g.Beams = append(g.Beams, newBeams...)
}

func (g *Grid) MoveBeam(b *Beam, d Direction) {
	valid := false
	switch d {
	case DirUp:
		valid = b.Row-1 >= 0
	case DirRight:
		valid = b.Col+1 < g.Width()
	case DirDown:
		valid = b.Row+1 < g.Height()
	case DirLeft:
		valid = b.Col-1 >= 0
	}

	if valid {
		b.MoveBeam(d)
		g.Tiles[b.Row][b.Col].IsEnergized = true
	} else {
		b.IsStuck = true
	}

	if g.BeamPath[b.PathKey()] {
		b.IsStuck = true
	} else {
		g.BeamPath[b.PathKey()] = true
	}
}

func (g *Grid) SplitBeam(b *Beam, d Direction, s rune) {
	switch s {
	case '-':
		rightValid := b.Col+1 < g.Width()
		leftValid := b.Col-1 >= 0
		if rightValid && leftValid {
			newBeam := Beam{
				Coords:  b.Coords,
				Dir:     DirRight,
				IsStuck: false,
			}
			newBeam.MoveBeam(DirRight)
			b.MoveBeam(DirLeft)

			if g.BeamPath[newBeam.PathKey()] {
				b.IsStuck = true
			} else {
				g.BeamPath[newBeam.PathKey()] = true
				g.Beams = append(g.Beams, &newBeam)
				g.Tiles[newBeam.Row][newBeam.Col].IsEnergized = true
			}

			if g.BeamPath[b.PathKey()] {
				b.IsStuck = true
			} else {
				g.BeamPath[b.PathKey()] = true
				g.Tiles[b.Row][b.Col].IsEnergized = true
			}

		} else if rightValid || leftValid {
			if rightValid {
				b.MoveBeam(DirRight)
			} else {
				b.MoveBeam(DirLeft)
			}

			if g.BeamPath[b.PathKey()] {
				b.IsStuck = true
			} else {
				g.BeamPath[b.PathKey()] = true
				g.Tiles[b.Row][b.Col].IsEnergized = true
			}
		} else {
			b.IsStuck = true
		}
	case '|':
		upValid := b.Row-1 >= 0
		downValid := b.Row+1 < g.Height()
		if upValid && downValid {
			newBeam := Beam{
				Coords:  b.Coords,
				Dir:     DirUp,
				IsStuck: false,
			}
			newBeam.MoveBeam(DirUp)
			b.MoveBeam(DirDown)

			if g.BeamPath[newBeam.PathKey()] {
				b.IsStuck = true
			} else {
				g.BeamPath[newBeam.PathKey()] = true
				g.Beams = append(g.Beams, &newBeam)
				g.Tiles[newBeam.Row][newBeam.Col].IsEnergized = true
			}

			if g.BeamPath[b.PathKey()] {
				b.IsStuck = true
			} else {
				g.BeamPath[b.PathKey()] = true
				g.Tiles[b.Row][b.Col].IsEnergized = true
			}

		} else if upValid || downValid {
			if upValid {
				b.MoveBeam(DirUp)
			} else {
				b.MoveBeam(DirDown)
			}

			if g.BeamPath[b.PathKey()] {
				b.IsStuck = true
			} else {
				g.BeamPath[b.PathKey()] = true
				g.Tiles[b.Row][b.Col].IsEnergized = true
			}
		} else {
			b.IsStuck = true
		}
	}
}

func (g *Grid) PruneBeams() {
	g.Beams = slices.DeleteFunc(g.Beams, func(b *Beam) bool {
		return b.IsStuck
	})
}

func (g *Grid) CountEnergizedTiles() int {
	count := 0
	for _, row := range g.Tiles {
		for _, t := range row {
			if t.IsEnergized {
				count++
			}
		}
	}

	return count
}

func (g *Grid) Width() int {
	return len(g.Tiles[0])
}

func (g *Grid) Height() int {
	return len(g.Tiles)
}

func (g *Grid) PrintGrid() {
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(g.Tiles[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("   ┏")
	for j := 0; j < len(g.Tiles[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┓\n")

	for y, line := range g.Tiles {
		fmt.Printf("%3d%c", y, '┃')
		for _, val := range line {
			if val.IsEnergized {
				fmt.Printf("%s%s%c%s", term.YellowBackground, term.Black, val.Tile, term.Reset)
			} else {
				fmt.Printf("%c", val.Tile)
			}
		}
		fmt.Printf("%c\n", '┃')
	}

	fmt.Print("   ┗")
	for j := 0; j < len(g.Tiles[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┛\n")
}

func parseInput(input []string) [][]*Tile {
	tiles := make([][]*Tile, 0)
	for row, line := range input {
		tiles = append(tiles, make([]*Tile, 0))
		for col, r := range line {
			tile := Tile{
				Coords: Coords{Row: row, Col: col},
				Tile:   r,
			}
			tiles[row] = append(tiles[row], &tile)
		}
	}

	return tiles
}

type BeamConfig struct {
	Coords
	Direction Direction
}

type ConfigResult struct {
	Id int
	BeamConfig
	NumTiles int
	Duration time.Duration
}

func (cr ConfigResult) PrintResult() {
	fmt.Printf("%d [%d, %d] %s %s %d\n", cr.Id, cr.Row, cr.Col, cr.Direction.String(), cr.Duration, cr.NumTiles)
}

type SimulationResults struct {
	results []ConfigResult
	Min     int
	Max     int
}

func (sr SimulationResults) PrintSummary() {
	fmt.Printf("Count: %d, Min: %d, Max: %d\n", len(sr.results), sr.Min, sr.Max)
}

func (s *SimulationResults) Add(r ConfigResult) {
	s.results = append(s.results, r)

	slices.SortFunc(s.results, func(a, b ConfigResult) int {
		return a.NumTiles - b.NumTiles
	})
	if r.NumTiles < s.Min {
		s.Min = r.NumTiles
	}

	if r.NumTiles > s.Max {
		s.Max = r.NumTiles
	}
}

func newGrid(input []string, config BeamConfig) Grid {
	beams := make([]*Beam, 0)
	beams = append(beams, &Beam{
		Coords: config.Coords,
		Dir:    config.Direction,
	})

	g := Grid{
		Tiles:    parseInput(input),
		Beams:    beams,
		BeamPath: make(map[string]bool),
	}

	g.Tiles[config.Row][config.Col].IsEnergized = true

	return g
}

func Day16(input []string) int {
	tiles := parseInput(input)
	maxTiles := 0

	results := SimulationResults{}

	// Top Edge
	for i := 0; i < len(tiles[0]); i++ {

		config := BeamConfig{
			Coords{Row: 0, Col: i},
			DirDown,
		}
		g := newGrid(input, config)

		start := time.Now()
		numTiles := g.SimulateConfig()
		end := time.Now()

		result := ConfigResult{
			i + 1,
			config,
			numTiles,
			start.Sub(end),
		}
		results.Add(result)

		if numTiles > maxTiles {
			maxTiles = numTiles
			g.PrintGrid()
			fmt.Println("New max")
		}
		result.PrintResult()
	}

	fmt.Print("Done with top edge, summary:")
	results.PrintSummary()

	// Bottom edge
	for i := 0; i < len(tiles[0]); i++ {
		config := BeamConfig{
			Coords{Row: len(tiles) - 1, Col: i},
			DirUp,
		}
		g := newGrid(input, config)

		start := time.Now()
		numTiles := g.SimulateConfig()
		end := time.Now()

		result := ConfigResult{
			i + 1,
			config,
			numTiles,
			start.Sub(end),
		}
		results.Add(result)

		if numTiles > maxTiles {
			maxTiles = numTiles
			g.PrintGrid()
			fmt.Println("New max")
		}
		result.PrintResult()
	}

	fmt.Println("Done with bottom edge, current max is:", maxTiles)

	// Left edge
	for i := 0; i < len(tiles); i++ {
		config := BeamConfig{
			Coords{Row: i, Col: 0},
			DirRight,
		}
		g := newGrid(input, config)

		start := time.Now()
		numTiles := g.SimulateConfig()
		end := time.Now()

		result := ConfigResult{
			i + 1,
			config,
			numTiles,
			start.Sub(end),
		}
		results.Add(result)

		if numTiles > maxTiles {
			maxTiles = numTiles
			g.PrintGrid()
			fmt.Println("New max")
		}
		result.PrintResult()
	}

	fmt.Println("Done with left edge, current max is:", maxTiles)

	// Right Edge
	for i := 0; i < len(tiles); i++ {
		config := BeamConfig{
			Coords{Row: i, Col: len(tiles[0]) - 1},
			DirLeft,
		}
		g := newGrid(input, config)

		start := time.Now()
		numTiles := g.SimulateConfig()
		end := time.Now()

		result := ConfigResult{
			i + 1,
			config,
			numTiles,
			start.Sub(end),
		}
		results.Add(result)

		if numTiles > maxTiles {
			maxTiles = numTiles
			g.PrintGrid()
			fmt.Println("New max")
		}
		result.PrintResult()
	}

	return maxTiles
}

func RunSingle(input []string, config BeamConfig) ConfigResult {
	g := newGrid(input, config)

	g.PrintGrid()

	start := time.Now()
	numTiles := g.SimulateConfig()
	end := time.Now()

	result := ConfigResult{
		1,
		config,
		numTiles,
		start.Sub(end),
	}

	return result
}
