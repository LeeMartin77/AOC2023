package main

import (
	"testing"
)

var exampleStart string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

var expectedRolledSouth string = `.....#....
....#....#
...O.##...
...#......
O.O....O#O
O.#..O.#.#
O....#....
OO....OO..
#OO..###..
#OO.O#...O
`

var expectedRolledWest string = `O....#....
OOO.#....#
.....##...
OO.#OO....
OO......#.
O.#O...#.#
O....#OO..
O.........
#....###..
#OO..#....
`

var expectedRolledEast string = `....O#....
.OOO#....#
.....##...
.OO#....OO
......OO#.
.O#...O#.#
....O#..OO
.........O
#....###..
#..OO#....
`

var expectedRolledNorth string = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
`

var oneCycle string = `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....
`
var twoCycle string = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O
`
var threeCycle string = `.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O
`

func TestParseRocks(t *testing.T) {
	res, _, _ := ParseRocks(exampleStart)
	if len(res) != 35 {
		t.Errorf("Expected 35 got %v", len(res))
	}
}

func PrintRocksAsMap(rocks []Rock) string {

	maxX := 0
	maxY := 0
	for _, rk := range rocks {
		if rk.X > maxX {
			maxX = rk.X
		}

		if rk.Y > maxY {
			maxY = rk.Y
		}
	}

	str := ""

	for y := 0; y <= maxY; y = y + 1 {
		for x := 0; x <= maxX; x = x + 1 {
			new := "."
			for _, rk := range rocks {
				if rk.X == x && rk.Y == y {
					if rk.Mobile {
						new = "O"
					} else {
						new = "#"
					}
				}
			}
			str = str + new
		}
		str = str + "\n"
	}
	return str
}

func TestRollRocksNorth(t *testing.T) {

	strt, maxX, maxY := ParseRocks(exampleStart)

	res := RollRocksNorth(strt, maxX, maxY)

	mp := PrintRocksAsMap(res)

	if mp != expectedRolledNorth {
		t.Errorf("Expected:\n%s\n got:\n%s\n", expectedRolledNorth, mp)
	}
}

func TestRollRocksSouth(t *testing.T) {

	strt, maxX, maxY := ParseRocks(exampleStart)

	res := RollRocksSouth(strt, maxX, maxY)

	mp := PrintRocksAsMap(res)

	if mp != expectedRolledSouth {
		t.Errorf("Expected:\n%s\n got:\n%s\n", expectedRolledSouth, mp)
	}
}

func TestRollRocksWest(t *testing.T) {

	strt, maxX, maxY := ParseRocks(exampleStart)

	res := RollRocksWest(strt, maxX, maxY)

	mp := PrintRocksAsMap(res)

	if mp != expectedRolledWest {
		t.Errorf("Expected:\n%s\n got:\n%s\n", expectedRolledWest, mp)
	}
}

func TestRollRocksEast(t *testing.T) {

	strt, maxX, maxY := ParseRocks(exampleStart)

	res := RollRocksEast(strt, maxX, maxY)

	mp := PrintRocksAsMap(res)

	if mp != expectedRolledEast {
		t.Errorf("Expected:\n%s\n got:\n%s\n", expectedRolledEast, mp)
	}
}

func TestCalculateLoadOnNorth(t *testing.T) {

	strt, maxX, maxY := ParseRocks(exampleStart)

	res := RollRocksNorth(strt, maxX, maxY)
	calc := CalculateLoadOnNorth(res)
	if calc != 136 {

		t.Errorf("Expected 136 got %v", calc)
	}
}

func TestSpinCycle(t *testing.T) {

	strt, maxX, maxY := ParseRocks(exampleStart)

	res := SpinCycle(strt, maxX, maxY)

	mp := PrintRocksAsMap(res)

	if mp != oneCycle {
		t.Errorf("Expected:\n%s\n got:\n%s\n", oneCycle, mp)
	}

	res = SpinCycle(res, maxX, maxY)

	mp = PrintRocksAsMap(res)

	if mp != twoCycle {
		t.Errorf("Expected:\n%s\n got:\n%s\n", twoCycle, mp)
	}

	res = SpinCycle(res, maxX, maxY)

	mp = PrintRocksAsMap(res)

	if mp != threeCycle {
		t.Errorf("Expected:\n%s\n got:\n%s\n", threeCycle, mp)
	}
}
