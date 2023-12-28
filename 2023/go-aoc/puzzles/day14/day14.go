package day14

import "fmt"

type Direction int

const (
	DirectionNorth Direction = iota
	DirectionEast
	DirectionSouth
	DirectionWest
)

type Rock rune

const (
	RockGround Rock = '.'
	RockRound  Rock = 'O'
	RockCube   Rock = '#'
)

func (ro Rock) AsRunePtr() *rune {
	r := rune(ro)
	return &r
}

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Platform struct {
	Rocks [][]*rune
}

func (p *Platform) CalcLoad() int {
	load := 0
	for i, row := range p.Rocks {
		numRocks := 0
		for _, rock := range row {
			r := *rock
			if r == rune(RockRound) {
				numRocks++
			}
		}
		distance := len(p.Rocks) - i
		load += numRocks * distance
	}
	return load
}

func (p *Platform) TiltNorth() {
	for i, row := range p.Rocks {
		for j, rock := range row {
			r := *rock
			if r == rune(RockRound) {
				for k := i - 1; k >= 0; k-- {
					dstRock := *p.Rocks[k][j]
					if dstRock == rune(RockGround) {
						p.Rocks[k+1][j] = RockGround.AsRunePtr()
						p.Rocks[k][j] = RockRound.AsRunePtr()
					} else {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) TiltWest() {
	for i, row := range p.Rocks {
		for j, rock := range row {
			r := *rock
			if r == rune(RockRound) {
				for k := j - 1; k >= 0; k-- {
					dstRock := *p.Rocks[i][k]
					if dstRock == rune(RockGround) {
						p.Rocks[i][k+1] = RockGround.AsRunePtr()
						p.Rocks[i][k] = RockRound.AsRunePtr()
					} else {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) TiltSouth() {
	for i := len(p.Rocks) - 1; i >= 0; i-- {
		for j := 0; j < len(p.Rocks[i]); j++ {
			r := *p.Rocks[i][j]
			if r == rune(RockRound) {
				for k := i + 1; k < len(p.Rocks); k++ {
					dstRock := *p.Rocks[k][j]
					if dstRock == rune(RockGround) {
						p.Rocks[k-1][j] = RockGround.AsRunePtr()
						p.Rocks[k][j] = RockRound.AsRunePtr()
					} else {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) TiltEast() {
	for i := 0; i < len(p.Rocks); i++ {
		for j := len(p.Rocks[i]) - 1; j >= 0; j-- {
			r := *p.Rocks[i][j]
			if r == rune(RockRound) {
				for k := j + 1; k < len(p.Rocks[i]); k++ {
					dstRock := *p.Rocks[i][k]
					if dstRock == rune(RockGround) {
						p.Rocks[i][k-1] = RockGround.AsRunePtr()
						p.Rocks[i][k] = RockRound.AsRunePtr()
					} else {
						break
					}
				}
			}
		}
	}
}

func (p *Platform) Spin() {
	p.TiltNorth()
	p.TiltWest()
	p.TiltSouth()
	p.TiltEast()
}

func (p *Platform) PrintPlatform() {
	for i := 0; i < 3; i++ {
		fmt.Print("    ")
		for j := 0; j < len(p.Rocks[0]); j++ {
			index := fmt.Sprintf("%3d", j)
			fmt.Printf("%c", index[i])
		}
		fmt.Print("\n")
	}

	fmt.Print("   ┏")
	for j := 0; j < len(p.Rocks[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┓\n")

	for row, rocks := range p.Rocks {
		fmt.Printf("%3d%c", row, '┃')
		for _, rock := range rocks {
			fmt.Printf("%c", *rock)
		}
		fmt.Printf("%c\n", '┃')
	}

	fmt.Print("   ┗")
	for j := 0; j < len(p.Rocks[0]); j++ {
		fmt.Printf("%c", '━')
	}
	fmt.Print("┛\n")
}

func parseInput(input []string) Platform {
	var r [][]*rune

	for i, l := range input {
		r = append(r, make([]*rune, 0))
		for _, c := range l {
			rock := c
			r[i] = append(r[i], &rock)
		}
	}

	return Platform{
		Rocks: r,
	}
}

func Day14(input []string, isPart2 bool) int {
	p := parseInput(input)
	p.PrintPlatform()

	if !isPart2 {
		p.TiltNorth()
	} else {
		for i := 1; i <= 1000000000; i++ {
			p.Spin()
			load := p.CalcLoad()
			if 1000000000%i == 0 {
				fmt.Println(i, load)
			}
		}
	}

	p.PrintPlatform()
	return p.CalcLoad()
}
