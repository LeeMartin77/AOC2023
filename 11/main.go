package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type Galaxy struct {
	X, Y int
}

type GalaxyPair struct {
	indexOne, indexTwo int
	galOne, galTwo     Galaxy
}

func (pr GalaxyPair) GetDistance() int {
	xDiff := pr.galOne.X - pr.galTwo.X
	yDiff := pr.galOne.Y - pr.galTwo.Y

	return int(math.Abs(float64(xDiff)) + math.Abs(float64(yDiff)))
}

func ParseGalaxies(input string) []Galaxy {
	galaxies := []Galaxy{}
	for y, ln := range strings.Split(input, "\n") {
		for x, chr := range ln {
			if chr == '#' {
				galaxies = append(galaxies, Galaxy{x, y})
			}
		}
	}
	return galaxies
}

func ExpandUniverse(universe []Galaxy) []Galaxy {
	expanded := universe
	columns := []int{}
	rows := []int{}
	minCol := 9999999
	maxCol := 0
	minRow := 9999999
	maxRow := 0
	for _, gal := range universe {
		if !slices.Contains(columns, gal.X) {
			columns = append(columns, gal.X)
			if maxCol < gal.X {
				maxCol = gal.X
			}
			if minCol > gal.X {
				minCol = gal.X
			}
		}
		if !slices.Contains(rows, gal.Y) {
			rows = append(rows, gal.Y)
			if maxRow < gal.Y {
				maxRow = gal.Y
			}
			if minRow > gal.Y {
				minRow = gal.Y
			}
		}
	}
	sort.SliceStable(columns, func(i, j int) bool {
		return columns[i] > columns[j]
	})
	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i] > rows[j]
	})
	for i := maxCol; i > 0; i = i - 1 {
		if !slices.Contains(columns, i) {
			for ii, gal := range universe {
				if gal.X > i {
					expanded[ii].X = expanded[ii].X + 1
				}
			}
		}
	}
	for i := maxRow; i > -1; i = i - 1 {
		if !slices.Contains(rows, i) {
			for ii, gal := range universe {
				if gal.Y > i {
					expanded[ii].Y = expanded[ii].Y + 1
				}
			}
		}
	}
	return expanded
}

func CreatePairs(universe []Galaxy) []GalaxyPair {
	pairs := []GalaxyPair{}
	for i, gal := range universe {
		otherGalaxies := universe[i+1:]
		for ii, oGal := range otherGalaxies {
			pairs = append(pairs, GalaxyPair{
				indexOne: i,
				indexTwo: ii + i + 1,
				galOne:   gal,
				galTwo:   oGal,
			})
		}
	}
	return pairs
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	initial := ParseGalaxies(input)

	expanded := ExpandUniverse(initial)

	pairs := CreatePairs(expanded)
	cuml := 0
	for _, pair := range pairs {
		cuml = cuml + pair.GetDistance()
	}

	fmt.Printf("Result pt1: %d\n", cuml)
}
