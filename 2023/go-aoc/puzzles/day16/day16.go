package day16

import (
	"aoc/utils/term"
	"fmt"
	"slices"
	"time"
)

// .|...\....
// |.-.\.....
// .....|-...
// ........|.
// ..........
// .........\
// ..../.\\..
// .-.-/..|..
// .|....-|.\
// ..//.|....`
//
// Beam enters top left corner, from the left heading to the right
// - Beam hits empty space, it continues in same direction
// - Beam hits a mirror, beam is reflected 90 degrees
// - Beam hits pointy end of splitter, acts as if it was empty space
// - Beam hits flat side of splitter the beam splits into two beams
//
// Tile is energized if that tile has at least one beam passing through it

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

type Beam struct {
	Coords
	Dir Direction
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
	Tile  rune
	Beams []*Beam
}

func (t Tile) IsEnergized() bool {
	return len(t.Beams) > 0
}

type Grid struct {
	Tiles [][]*Tile
	Beams []*Beam
}

func (g *Grid) SimulateConfig() int {
	prevEnergized := 0
	turnCount := 0
	waitTurns := 1

	for start := time.Now(); time.Since(start) < time.Second*30; {
		g.UpdateBeams()

		if prevEnergized == g.CountEnergizedTiles() {
			if turnCount < waitTurns {
				turnCount++
			} else {
				break
			}
		} else {
			prevEnergized = g.CountEnergizedTiles()
			turnCount = 0
		}
	}

	// fmt.Println("final:")
	// g.PrintGrid()

	return g.CountEnergizedTiles()
}

func (g *Grid) UpdateBeams() {
	newBeams := make([]*Beam, 0)
	for _, b := range g.Beams {
		bTile := g.Tiles[b.Row][b.Col]
		switch bTile.Tile {
		case '.':
			if g.CheckDir(b.Coords, b.Dir) {
				b.MoveBeam(b.Dir)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			}
		case '/':
			if b.Dir == DirUp && g.CheckDir(b.Coords, DirRight) {
				b.MoveBeam(DirRight)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			} else if b.Dir == DirRight && g.CheckDir(b.Coords, DirUp) {
				b.MoveBeam(DirUp)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			} else if b.Dir == DirDown && g.CheckDir(b.Coords, DirLeft) {
				b.MoveBeam(DirLeft)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			} else if b.Dir == DirLeft && g.CheckDir(b.Coords, DirDown) {
				b.MoveBeam(DirDown)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			}

		case '\\':
			if b.Dir == DirUp && g.CheckDir(b.Coords, DirLeft) {
				b.MoveBeam(DirLeft)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			} else if b.Dir == DirRight && g.CheckDir(b.Coords, DirDown) {
				b.MoveBeam(DirDown)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			} else if b.Dir == DirDown && g.CheckDir(b.Coords, DirRight) {
				b.MoveBeam(DirRight)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			} else if b.Dir == DirLeft && g.CheckDir(b.Coords, DirUp) {
				b.MoveBeam(DirUp)
				g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
			}
		case '-':
			if b.Dir == DirUp || b.Dir == DirDown {
				if g.CheckDir(b.Coords, DirLeft) {
					newBeam := Beam{
						Coords: b.Coords,
					}
					newBeam.MoveBeam(DirLeft)
					newBeams = append(newBeams, &newBeam)
					g.Tiles[newBeam.Row][newBeam.Col].Beams = append(g.Tiles[newBeam.Row][newBeam.Col].Beams, b)
				}

				if g.CheckDir(b.Coords, DirRight) {
					b.MoveBeam(DirRight)
					g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
				}
			} else {
				if g.CheckDir(b.Coords, b.Dir) {
					b.MoveBeam(b.Dir)
					g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
				}
			}
		case '|':
			if b.Dir == DirLeft || b.Dir == DirRight {
				if g.CheckDir(b.Coords, DirUp) {
					newBeam := Beam{
						Coords: b.Coords,
					}
					newBeam.MoveBeam(DirUp)
					newBeams = append(newBeams, &newBeam)
					g.Tiles[newBeam.Row][newBeam.Col].Beams = append(g.Tiles[newBeam.Row][newBeam.Col].Beams, b)
				}

				if g.CheckDir(b.Coords, DirDown) {
					b.MoveBeam(DirDown)
					g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
				}
			} else {
				if g.CheckDir(b.Coords, b.Dir) {
					b.MoveBeam(b.Dir)
					g.Tiles[b.Row][b.Col].Beams = append(g.Tiles[b.Row][b.Col].Beams, b)
				}
			}

		}
	}
	g.Beams = append(g.Beams, newBeams...)
}

func (g *Grid) CheckDir(c Coords, d Direction) bool {
	switch d {
	case DirUp:
		return c.Row-1 >= 0
	case DirRight:
		return c.Col+1 < g.Width()
	case DirDown:
		return c.Row+1 < g.Height()
	case DirLeft:
		return c.Col-1 >= 0
	}

	return false
}

func (g *Grid) CountEnergizedTiles() int {
	count := 0
	for _, row := range g.Tiles {
		for _, t := range row {
			if t.IsEnergized() {
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
			if val.IsEnergized() {
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
	fmt.Printf("%d [%d, %d] %s %d %d\n", cr.Id, cr.Row, cr.Col, cr.Direction.String(), cr.Duration.Milliseconds(), cr.NumTiles)
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
		Tiles: parseInput(input),
		Beams: beams,
	}

	g.Tiles[config.Row][config.Col].Beams = beams

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
