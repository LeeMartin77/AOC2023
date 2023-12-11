package main

import (
	"testing"
)

func TestParseGalaxies(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	res := ParseGalaxies(input)
	if len(res) != 9 {
		t.Errorf("Expected 9 got %v", len(res))
	}
	expected := []Galaxy{
		{3, 0},
		{7, 1},
		{0, 2},
		{6, 4},
		{1, 5},
		{9, 6},
		{7, 8},
		{0, 9},
		{4, 9},
	}
	for i, gl := range expected {
		if gl.X != res[i].X || gl.Y != res[i].Y {
			t.Errorf("Expected %v got %v", gl, res[i])
		}
	}
}

func TestUniverseExpansion(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	initial := ParseGalaxies(input)
	expectedInput := `....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`
	expected := ParseGalaxies(expectedInput)

	res := ExpandUniverse(initial)

	for i, gl := range expected {
		if gl.X != res[i].X || gl.Y != res[i].Y {
			t.Errorf("Expected %v got %v", gl, res[i])
		}
	}
}

func TestCreatePairs(t *testing.T) {

	expectedInput := `....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`
	universe := ParseGalaxies(expectedInput)
	pairs := CreatePairs(universe)
	if len(pairs) != 36 {
		t.Errorf("Expected 36 got %v", len(pairs))
	}
}

func TestDistanceBetweenPair(t *testing.T) {
	pair := GalaxyPair{
		galOne: Galaxy{1, 6},
		galTwo: Galaxy{5, 11},
	}

	res := pair.GetDistance()
	if res != 9 {
		t.Errorf("Expected 9 got %v", res)
	}
	pair = GalaxyPair{
		galOne: Galaxy{4, 0},
		galTwo: Galaxy{9, 10},
	}

	res = pair.GetDistance()
	if res != 15 {
		t.Errorf("Expected 15 got %v", res)
	}
	pair = GalaxyPair{
		galOne: Galaxy{0, 2},
		galTwo: Galaxy{10, 10},
	}

	res = pair.GetDistance()
	if res != 18 {
		t.Errorf("Expected 18 got %v", res)
	}
}

func TestIntegration(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	initial := ParseGalaxies(input)

	expanded := ExpandUniverse(initial)

	pairs := CreatePairs(expanded)
	cuml := 0
	for _, pair := range pairs {
		cuml = cuml + pair.GetDistance()
	}

	if cuml != 374 {
		t.Errorf("Expected 374 got %v", cuml)
	}
}
