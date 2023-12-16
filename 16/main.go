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
func ProgressLight(mp [][]rune, beams []LightBeam, limits Coordinate, visited map[int]map[int][]Coordinate) ([]LightBeam, map[int]map[int]int) {
	hitTiles := map[int]map[int]int{}
	outBeams := []LightBeam{}
	for _, beam := range beams {
		_, ok := hitTiles[beam.Position.X]
		if !ok {
			hitTiles[beam.Position.X] = map[int]int{}
		}
		hitTiles[beam.Position.X][beam.Position.Y] = hitTiles[beam.Position.X][beam.Position.Y] + 1
		_, ok = visited[beam.Position.X]
		if !ok {
			visited[beam.Position.X] = map[int][]Coordinate{}
		}
		visited[beam.Position.X][beam.Position.Y] = append(visited[beam.Position.X][beam.Position.Y], beam.Velocity)
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
				})
				outBeams = append(outBeams, LightBeam{
					Position: Coordinate{beam.Position.X + 1, beam.Position.Y},
					Velocity: Coordinate{1, 0},
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
				})
				outBeams = append(outBeams, LightBeam{
					Position: Coordinate{beam.Position.X, beam.Position.Y + 1},
					Velocity: Coordinate{0, 1},
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
			for _, bm := range visited[beam.Position.X][beam.Position.Y] {
				if beam.Velocity.X == bm.X && beam.Velocity.Y == bm.Y {
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
	visitedWithDirection := map[int]map[int][]Coordinate{}
	for x, col := range mp {
		visitedWithDirection[x] = map[int][]Coordinate{}
		for y := range col {
			visitedWithDirection[x][y] = []Coordinate{}
		}
	}
	maxTotalHistory := 0
	for len(beams) > 0 {
		nbms, history := ProgressLight(mp, beams, lmt, visitedWithDirection)
		for k, kis := range history {
			for kk, cnt := range kis {
				_, ok := totalHistory[k]
				if !ok {
					totalHistory[k] = map[int]int{}
				}
				totalHistory[k][kk] = totalHistory[k][kk] + cnt
				if maxTotalHistory < totalHistory[k][kk] {
					maxTotalHistory = totalHistory[k][kk]
				}
			}
		}
		beams = nbms
	}
	return beams, totalHistory
}

// Like a gorilla with a keyboard
func FindLargest(mp [][]rune, lmt Coordinate) int {
	top := 0

	for x := 0; x <= lmt.X; x = x + 1 {
		beams := []LightBeam{{
			Velocity: Coordinate{0, 1},
			Position: Coordinate{x, 0},
		}}
		_, totalHistory := IterateMapThrough(mp, lmt, beams)
		cuml := 0
		for _, hist := range totalHistory {
			for range hist {
				cuml = cuml + 1
			}
		}
		if cuml > top {
			top = cuml
		}

		beams = []LightBeam{{
			Velocity: Coordinate{0, -1},
			Position: Coordinate{x, lmt.Y},
		}}
		_, totalHistory = IterateMapThrough(mp, lmt, beams)
		cuml = 0
		for _, hist := range totalHistory {
			for range hist {
				cuml = cuml + 1
			}
		}
		if cuml > top {
			top = cuml
		}
	}
	for y := 0; y <= lmt.Y; y = y + 1 {
		beams := []LightBeam{{
			Velocity: Coordinate{1, 0},
			Position: Coordinate{0, y},
		}}
		_, totalHistory := IterateMapThrough(mp, lmt, beams)
		cuml := 0
		for _, hist := range totalHistory {
			for range hist {
				cuml = cuml + 1
			}
		}
		if cuml > top {
			top = cuml
		}

		beams = []LightBeam{{
			Velocity: Coordinate{-1, 0},
			Position: Coordinate{lmt.X, 0},
		}}
		_, totalHistory = IterateMapThrough(mp, lmt, beams)
		cuml = 0
		for _, hist := range totalHistory {
			for range hist {
				cuml = cuml + 1
			}
		}
		if cuml > top {
			top = cuml
		}
	}
	return top
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

	top := FindLargest(res, lmt)

	fmt.Printf("Result: %v\n", top)
}
