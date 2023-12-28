package day16

import (
	"aoc/utils/term"
	"fmt"
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

	time.Sleep(200 * time.Millisecond)
	fmt.Print(term.ClearScreen)
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

func newGrid(tiles [][]*Tile) Grid {
	beams := make([]*Beam, 0)
	beams = append(beams, &Beam{
		Dir: DirRight,
	})

	g := Grid{
		Tiles: tiles,
		Beams: beams,
	}

	g.Tiles[0][0].Beams = beams

	return g
}

func Day16(input []string) int {
	tiles := parseInput(input)
	g := newGrid(tiles)

	fmt.Println("initial:")
	g.PrintGrid()

	prevEnergized := 0
	turnCount := 0
	waitTurns := 1

	for {
		g.UpdateBeams()
		g.PrintGrid()

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

	return g.CountEnergizedTiles()
}
