package main

import (
	"testing"
)

var exampleMap string = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`

var expectedParsed string = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..O####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`

var expectedAfterSixSteps string = `...........
.....###.#.
.###.##.O#.
.O#O#O.O#..
O.O.#.#.O..
.##O.O####.
.##.O#O..#.
.O.O.O.##..
.##.#.####.
.##O.##.##.
...........
`

func TestParse(t *testing.T) {
	mp, stpnt := ParseMap(exampleMap)
	res := DrawMap(mp, stpnt)
	if res != expectedParsed {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expectedParsed, res)
	}
}

func TestPartOne(t *testing.T) {
	mp, stpnt := ParseMap(exampleMap)

	for i := 0; i < 6; i = i + 1 {
		stpnt = PerformStep(mp, stpnt)
	}

	res := DrawMap(mp, stpnt)
	if res != expectedAfterSixSteps {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n", expectedAfterSixSteps, res)
	}

	cuml := 0
	for _, col := range stpnt {
		for _, val := range col {
			if val {
				cuml = cuml + 1
			}
		}
	}
	if cuml != 16 {
		t.Errorf("Expected:\n%d\n\nGot:\n%d\n", 16, cuml)
	}
}
