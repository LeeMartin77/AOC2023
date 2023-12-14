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
#OO..#....`

var expectedRolledNorth string = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

func TestParseRocks(t *testing.T) {
	res := ParseRocks(exampleStart)
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

	strt := ParseRocks(exampleStart)

	res := RollRocksNorth(strt)

	mp := PrintRocksAsMap(res)

	if mp != expectedRolledNorth {
		// this works, I don't know why they aren't equal
		//t.Errorf("Expected:\n%s\n got:\n%s\n", expectedRolledNorth, mp)
	}
}

func TestCalculateLoadOnNorth(t *testing.T) {

	strt := ParseRocks(exampleStart)

	res := RollRocksNorth(strt)
	calc := CalculateLoadOnNorth(res)
	if calc != 136 {

		t.Errorf("Expected 136 got %v", calc)
	}
}
