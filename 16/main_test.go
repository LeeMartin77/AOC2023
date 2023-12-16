package main

import (
	"testing"
)

var exampleMap string = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestParseMap(t *testing.T) {
	res, _ := ParseMap(exampleMap)

	if res[0][0] != '.' {
		t.Errorf("Expected . got %v", string(res[2][2]))
	}
	if res[4][1] != '\\' {
		t.Errorf("Expected \\ got %v", string(res[4][1]))
	}
	if res[4][6] != '/' {
		t.Errorf("Expected / got %v", string(res[4][6]))
	}
	if res[1][8] != '|' {
		t.Errorf("Expected | got %v", string(res[1][8]))
	}
}

func TestHitEmpty(t *testing.T) {
	res, _ := ParseMap(exampleMap)
	startBeams := []LightBeam{{
		Position: Coordinate{0, 3},
		Velocity: Coordinate{1, 0},
	}}
	endBeams, history := ProgressLight(res, startBeams, Coordinate{10, 10})

	expectedBeam := LightBeam{
		Position: Coordinate{1, 3},
		Velocity: Coordinate{1, 0},
	}

	if len(endBeams) != 1 {
		t.Errorf("Expected 1 beam got %v", len(endBeams))
	}

	if !MatchingBeams(endBeams[0], expectedBeam) {
		t.Errorf("Expected %v beam got %v", expectedBeam, endBeams[0])
	}

	if history[0][3] != 1 {
		t.Errorf("Expected 1 got %v", history[0][3])
	}
}

func TestHitMirror(t *testing.T) {
	res, lmt := ParseMap(exampleMap)
	startBeams := []LightBeam{{
		Position: Coordinate{X: 4, Y: 1},
		Velocity: Coordinate{X: 1, Y: 0},
	}}
	endBeams, history := ProgressLight(res, startBeams, lmt)

	expectedBeam := LightBeam{
		Position: Coordinate{4, 2},
		Velocity: Coordinate{0, 1},
	}

	if len(endBeams) != 1 {
		t.Errorf("Expected 1 beam got %v", len(endBeams))
	}

	if !MatchingBeams(endBeams[0], expectedBeam) {
		t.Errorf("Expected %v beam got %v", expectedBeam, endBeams[0])
	}

	if history[4][1] != 1 {
		t.Errorf("Expected 1 got %v", history[4][1])
	}
}

func TestHitMirrorOther(t *testing.T) {
	res, lmt := ParseMap(exampleMap)
	startBeams := []LightBeam{{
		Position: Coordinate{X: 4, Y: 6},
		Velocity: Coordinate{X: -1, Y: 0},
	}}
	endBeams, history := ProgressLight(res, startBeams, lmt)

	expectedBeam := LightBeam{
		Position: Coordinate{4, 7},
		Velocity: Coordinate{0, 1},
	}

	if len(endBeams) != 1 {
		t.Errorf("Expected 1 beam got %v", len(endBeams))
	}

	if !MatchingBeams(endBeams[0], expectedBeam) {
		t.Errorf("Expected %v beam got %v", expectedBeam, endBeams[0])
	}

	if history[4][6] != 1 {
		t.Errorf("Expected 1 got %v", history[4][6])
	}
}

func TestHitSplitter(t *testing.T) {
	res, lmt := ParseMap(exampleMap)
	startBeams := []LightBeam{{
		Position: Coordinate{X: 6, Y: 8},
		Velocity: Coordinate{X: 0, Y: -1},
	}}
	endBeams, history := ProgressLight(res, startBeams, lmt)

	expectedBeams := []LightBeam{
		{
			Position: Coordinate{5, 8},
			Velocity: Coordinate{-1, 0},
		},
		{
			Position: Coordinate{7, 8},
			Velocity: Coordinate{1, 0},
		},
	}

	if len(endBeams) != 2 {
		t.Errorf("Expected 2 beam got %v", len(endBeams))
	}

	for i, ebeam := range expectedBeams {

		if !MatchingBeams(endBeams[i], ebeam) {
			t.Errorf("Expected %v beam got %v", ebeam, endBeams[i])
		}
	}

	if history[6][8] != 1 {
		t.Errorf("Expected 1 got %v", history[6][8])
	}
}

func TestHitSplitterOther(t *testing.T) {
	res, lmt := ParseMap(exampleMap)
	startBeams := []LightBeam{{
		Position: Coordinate{X: 7, Y: 7},
		Velocity: Coordinate{X: -1, Y: 0},
	}}
	endBeams, history := ProgressLight(res, startBeams, lmt)

	expectedBeams := []LightBeam{
		{
			Position: Coordinate{7, 6},
			Velocity: Coordinate{0, -1},
		},
		{
			Position: Coordinate{7, 8},
			Velocity: Coordinate{0, 1},
		},
	}

	if len(endBeams) != 2 {
		t.Errorf("Expected 2 beam got %v", len(endBeams))
	}

	for i, ebeam := range expectedBeams {

		if !MatchingBeams(endBeams[i], ebeam) {
			t.Errorf("Expected %v beam got %v", ebeam, endBeams[i])
		}
	}

	if history[7][7] != 1 {
		t.Errorf("Expected 1 got %v", history[7][7])
	}
}

func TestIntegration(t *testing.T) {

	res, lmt := ParseMap(exampleMap)
	beams := []LightBeam{{
		Velocity: Coordinate{1, 0},
		Position: Coordinate{0, 0},
	}}
	_, totalHistory := IterateMapThrough(res, lmt, beams)
	cuml := 0
	for _, hist := range totalHistory {
		for range hist {
			cuml = cuml + 1
		}
	}

	if cuml != 46 {
		t.Errorf("Expected 46 got %v", cuml)
	}
}
