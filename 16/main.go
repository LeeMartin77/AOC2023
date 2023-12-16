package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type LightBeam struct {
	Position Coordinate
	Velocity Coordinate
	History  []LightBeam
}

func MatchingBeams(beamOne LightBeam, beamTwo LightBeam) bool {
	return beamOne.Position.X == beamTwo.Position.X && beamOne.Position.Y == beamTwo.Position.Y &&
		beamOne.Velocity.X == beamTwo.Velocity.X && beamOne.Velocity.Y == beamTwo.Velocity.Y
}

func ParseMap(input string) ([][]rune, Coordinate) {
	mp := [][]rune{}

	for y, ln := range strings.Split(input, "\n") {
		for x, chr := range ln {
			if y == 0 {
				mp = append(mp, []rune{})
			}
			mp[x] = append(mp[x], chr)
		}
	}
	return mp, Coordinate{X: len(mp[0]) - 1, Y: len(mp) - 1}
}

// beams, energised tiles
func ProgressLight(mp [][]rune, beams []LightBeam, limits Coordinate) ([]LightBeam, map[int]map[int]int) {
	hitTiles := map[int]map[int]int{}
	outBeams := []LightBeam{}
	for _, beam := range beams {
		_, ok := hitTiles[beam.Position.X]
		if !ok {
			hitTiles[beam.Position.X] = map[int]int{}
		}
		hitTiles[beam.Position.X][beam.Position.Y] = hitTiles[beam.Position.X][beam.Position.Y] + 1
		beam.History = append(beam.History, LightBeam{Position: beam.Position, Velocity: beam.Velocity})
		if mp[beam.Position.X][beam.Position.Y] == '.' {
			// we just pass on through
			beam.Position.X = beam.Position.X + beam.Velocity.X
			beam.Position.Y = beam.Position.Y + beam.Velocity.Y
			outBeams = append(outBeams, beam)
		} else if mp[beam.Position.X][beam.Position.Y] == '\\' {
			if math.Abs(float64(beam.Velocity.X)) > 0 {
				// initially moving horizontally
				if beam.Velocity.X > 0 {
					// init moving ->
					beam.Velocity.Y = 1
				} else {
					// init  moving <-
					beam.Velocity.Y = -1
				}
				beam.Velocity.X = 0
			} else {
				// initially moving vertically
				if beam.Velocity.Y > 0 {
					// init moving V
					beam.Velocity.X = 1
				} else {
					// init  moving ^
					beam.Velocity.X = -1
				}
				beam.Velocity.Y = 0
			}
			beam.Position.X = beam.Position.X + beam.Velocity.X
			beam.Position.Y = beam.Position.Y + beam.Velocity.Y
			outBeams = append(outBeams, beam)
		} else if mp[beam.Position.X][beam.Position.Y] == '/' {
			if math.Abs(float64(beam.Velocity.X)) > 0 {
				// initially moving horizontally
				if beam.Velocity.X > 0 {
					// init moving ->
					beam.Velocity.Y = -1
				} else {
					// init  moving <-
					beam.Velocity.Y = 1
				}
				beam.Velocity.X = 0
			} else {
				// initially moving vertically
				if beam.Velocity.Y > 0 {
					// init moving V
					beam.Velocity.X = -1
				} else {
					// init  moving ^
					beam.Velocity.X = 1
				}
				beam.Velocity.Y = 0
			}
			beam.Position.X = beam.Position.X + beam.Velocity.X
			beam.Position.Y = beam.Position.Y + beam.Velocity.Y
			outBeams = append(outBeams, beam)
		} else if mp[beam.Position.X][beam.Position.Y] == '-' {
			if math.Abs(float64(beam.Velocity.X)) > 0 {
				// treat it like a dot
				beam.Position.X = beam.Position.X + beam.Velocity.X
				beam.Position.Y = beam.Position.Y + beam.Velocity.Y
				outBeams = append(outBeams, beam)
			} else {
				// split
				outBeams = append(outBeams, LightBeam{
					Position: Coordinate{beam.Position.X - 1, beam.Position.Y},
					Velocity: Coordinate{-1, 0},
					History:  beam.History,
				})
				outBeams = append(outBeams, LightBeam{
					Position: Coordinate{beam.Position.X + 1, beam.Position.Y},
					Velocity: Coordinate{1, 0},
					History:  beam.History,
				})
			}
		} else if mp[beam.Position.X][beam.Position.Y] == '|' {
			if math.Abs(float64(beam.Velocity.Y)) > 0 {
				// treat it like a dot
				beam.Position.X = beam.Position.X + beam.Velocity.X
				beam.Position.Y = beam.Position.Y + beam.Velocity.Y
				outBeams = append(outBeams, beam)
			} else {
				// split
				outBeams = append(outBeams, LightBeam{
					Position: Coordinate{beam.Position.X, beam.Position.Y - 1},
					Velocity: Coordinate{0, -1},
					History:  beam.History,
				})
				outBeams = append(outBeams, LightBeam{
					Position: Coordinate{beam.Position.X, beam.Position.Y + 1},
					Velocity: Coordinate{0, 1},
					History:  beam.History,
				})
			}
		}
	}

	// prune beams
	prunedOutBeams := []LightBeam{}
outer:
	for _, beam := range outBeams {
		if beam.Position.X > -1 && beam.Position.Y > -1 &&
			beam.Position.X <= limits.X && beam.Position.Y <= limits.Y {
			for _, bm := range beam.History {
				if MatchingBeams(bm, beam) {
					// we're in a loop
					continue outer
				}
			}
			prunedOutBeams = append(prunedOutBeams, beam)
		}
	}
	return prunedOutBeams, hitTiles
}

func IterateMapThrough(mp [][]rune, lmt Coordinate, beams []LightBeam) ([]LightBeam, map[int]map[int]int) {
	totalHistory := map[int]map[int]int{}
	for len(beams) > 0 {
		nbms, history := ProgressLight(mp, beams, lmt)
		for k, kis := range history {
			for kk, cnt := range kis {
				_, ok := totalHistory[k]
				if !ok {
					totalHistory[k] = map[int]int{}
				}
				totalHistory[k][kk] = totalHistory[k][kk] + cnt
			}
		}
		beams = nbms
	}
	return beams, totalHistory
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	res, lmt := ParseMap(input)
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

	fmt.Printf("Result: %v\n", cuml)
}
