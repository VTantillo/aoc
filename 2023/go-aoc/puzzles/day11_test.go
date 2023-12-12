package puzzles

import (
	"aoc/utils"
	"testing"
)

var ex1 = [][]rune{
	{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
	{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
	{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
}

func TestDay11Pt1(t *testing.T) {
	input := utils.ReadInput("../../inputs/day11-ex.txt")

	result := Day11Pt1(input)

	if result != 374 {
		t.Fatalf("Sum should have been 374, got=%v", result)
	}
}

func TestDay11Pt2(t *testing.T) {
	universe := universe{
		starMap: ex1,
	}

	universe.findGalaxies()

	result1 := universe.findDistanceSum(2)
	result2 := universe.findDistanceSum(10)
	result3 := universe.findDistanceSum(100)

	if result1 != 374 {
		t.Fatalf("Sum should have been 374, got=%v", result1)
	}
	if result2 != 1030 {
		t.Fatalf("Sum should have been 1030, got=%v", result2)
	}
	if result3 != 8410 {
		t.Fatalf("Sum should have been 8410, got=%v", result3)
	}
}

func TestEmptySpaceCols(t *testing.T) {
	universe := universe{
		starMap: ex1,
	}
	galaxy1 := galaxy{
		coords: coords{
			x: 3,
			y: 0,
		},
	}
	galaxy7 := galaxy{
		coords: coords{
			x: 7,
			y: 8,
		},
	}

	utils.PrintRuneMap(universe.starMap)
	cols1 := universe.getEmptySpaceCols(galaxy1.coords, galaxy7.coords)

	if cols1 != 1 {
		t.Fatalf("Cols should have been 1, got=%v", cols1)
	}

	// cols2 := universe.getEmptySpaceCols(galaxy1, galaxy3)
	//
	// if cols2 != 3 {
	// 	t.Fatalf("Cols should have been 1, got=%v", cols2)
	// }
}

func TestEmptySpaceRows(t *testing.T) {
	universe := universe{
		starMap: ex1,
	}
	galaxy1 := coords{
		x: 1,
		y: 5,
	}
	galaxy2 := coords{
		x: 4,
		y: 9,
	}
	galaxy3 := coords{
		x: 3,
		y: 0,
	}

	rows1 := universe.getEmptySpaceRows(galaxy1, galaxy2)
	if rows1 != 1 {
		t.Fatalf("rows should have been 1, got=%v", rows1)
	}

	rows2 := universe.getEmptySpaceRows(galaxy3, galaxy2)
	if rows2 != 2 {
		t.Fatalf("rows should have been 2, got=%v", rows2)
	}
}

func TestGetDistance(t *testing.T) {
	universe := universe{
		starMap: ex1,
	}
	galaxy1 := galaxy{
		coords: coords{
			x: 3,
			y: 0,
		},
	}
	// galaxy2 := galaxy{
	// 	coords: coords{
	// 		x: 7,
	// 		y: 1,
	// 	},
	// }
	galaxy3 := galaxy{
		coords: coords{
			x: 0,
			y: 2,
		},
	}
	// galaxy4 := galaxy{
	// 	coords: coords{
	// 		x: 6,
	// 		y: 4,
	// 	},
	// }
	galaxy5 := galaxy{
		coords: coords{
			x: 1,
			y: 5,
		},
	}
	galaxy6 := galaxy{
		coords: coords{
			x: 9,
			y: 6,
		},
	}
	galaxy7 := galaxy{
		coords: coords{
			x: 7,
			y: 8,
		},
	}
	galaxy8 := galaxy{
		coords: coords{
			x: 0,
			y: 9,
		},
	}
	galaxy9 := galaxy{
		coords: coords{
			x: 4,
			y: 9,
		},
	}

	d1 := universe.findDistance(galaxy5.coords, galaxy9.coords, 2)
	if d1 != 9 {
		t.Fatalf("Distance should be 9, got=%v", d1)
	}

	d2 := universe.findDistance(galaxy1.coords, galaxy7.coords, 2)
	if d2 != 15 {
		t.Fatalf("Distance should be 15, got=%v", d2)
	}

	d3 := universe.findDistance(galaxy3.coords, galaxy6.coords, 2)
	if d3 != 17 {
		t.Fatalf("Distance should be 17, got=%v", d3)
	}

	d4 := universe.findDistance(galaxy8.coords, galaxy9.coords, 2)
	if d4 != 5 {
		t.Fatalf("Distance should be 5, got=%v", d4)
	}
}
