package main

import (
	"sort"
	"testing"
)

var exampleMap string = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

func TestQuickCheckMapParsing(t *testing.T) {
	res, strt, end := ParseMap(exampleMap)
	if len(res) != 9 {
		t.Errorf("Expected 9 got %v", len(res))
	}
	if strt != "1:0" {
		t.Errorf("Expected 1:0 got %v", strt)
	}
	if end != "21:22" {
		t.Errorf("Expected 21:22 got %v", end)
	}
}

func TestPossiblePaths(t *testing.T) {
	res, strt, end := ParseMap(exampleMap)
	pths := GetPossibleSeinicPaths([][]Connection{res[strt]}, end, res)
	lens := TotalPathLens(pths)
	sort.SliceStable(lens, func(i, j int) bool {
		return lens[i] > lens[j]
	})

	for i, exp := range []int{94, 90, 86} {
		if lens[i]+1 != exp {
			t.Errorf("Expected %v got %v", exp, lens[i]+1)
		}
	}
}
