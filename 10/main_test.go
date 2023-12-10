package main

import (
	"strings"
	"testing"
)

func TestTileParsing(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....`
	startPos, res := ParseTiles(input)
	for y, ln := range strings.Split(input, "\n") {
		for x, chr := range ln {

			if chr != res[x][y].Rune {
				t.Errorf("Expected %v got %v", chr, res)
			}
			if x != res[x][y].X {
				t.Errorf("Expected %v got %v", x, res)
			}
			if y != res[x][y].Y {
				t.Errorf("Expected %v got %v", y, res)
			}
		}
	}
	if startPos.X != 1 && startPos.Y != 1 {

		t.Errorf("Expected %v got %v", InputOutput{1, 1}, startPos)
	}
}

func TestPartOneRouting(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....`
	startPos, tiles := ParseTiles(input)
	loop, _ := GetLoop(startPos, tiles)
	if len(loop) != 8 {
		t.Errorf("Expected 8 got %v", len(loop))
	}
}

func TestShuttleGenerating(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....`
	startPos, tiles := ParseTiles(input)
	shuttle, routes := GetShuttlesAndRoutes(startPos, tiles)
	if shuttle[0].X != 2 && shuttle[0].Y != 1 {

		t.Errorf("Expected asd got %v", shuttle[0])
	}
	if shuttle[1].X != 1 && shuttle[1].Y != 2 {

		t.Errorf("Expected asd got %v", shuttle[0])
	}
	for i, rt := range routes {
		if !matches(rt, []InputOutput{startPos}) {
			t.Errorf("Expected %d got %v", i, rt)

		}
	}
}

func matches(one []InputOutput, two []InputOutput) bool {
	contains := []bool{}
	for _, o := range one {
		for _, t := range two {
			if o.X == t.X && o.Y == t.Y {
				contains = append(contains, true)
			}
		}
	}
	return len(contains) == len(one) && len(one) == len(two)
}

func TestConnectsTo(t *testing.T) {
	// | is a vertical pipe connecting north and south.
	tile := Tile{
		X:    0,
		Y:    0,
		Rune: '|',
	}
	expected := []InputOutput{{0, -1}, {0, 1}}
	actual := tile.ConnectsTo()
	if !matches(expected, actual) {

		t.Errorf("Expected %v got %v", expected, actual)
	}
	// - is a horizontal pipe connecting east and west.
	tile.Rune = '-'
	expected = []InputOutput{{-1, 0}, {1, 0}}
	actual = tile.ConnectsTo()
	if !matches(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
	// L is a 90-degree bend connecting north and east.
	tile.Rune = 'L'
	expected = []InputOutput{{0, -1}, {1, 0}}
	actual = tile.ConnectsTo()
	if !matches(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
	// J is a 90-degree bend connecting north and west.
	tile.Rune = 'J'
	expected = []InputOutput{{0, -1}, {-1, 0}}
	actual = tile.ConnectsTo()
	if !matches(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
	// 7 is a 90-degree bend connecting south and west.
	tile.Rune = '7'
	expected = []InputOutput{{0, 1}, {-1, 0}}
	actual = tile.ConnectsTo()
	if !matches(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
	// F is a 90-degree bend connecting south and east.
	tile.Rune = 'F'
	expected = []InputOutput{{0, 1}, {1, 0}}
	actual = tile.ConnectsTo()
	if !matches(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
	// . is ground; there is no pipe in this tile.
	tile.Rune = '.'
	expected = []InputOutput{}
	actual = tile.ConnectsTo()
	if !matches(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
}
