package main

import "testing"

var example string = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

var expectedTrenches string = `#######
#.....#
###...#
..#...#
..#...#
###.###
#...#..
##..###
.#....#
.######
`

var expectedFilled string = `#######
#######
#######
..#####
..#####
#######
#####..
#######
.######
.######
`

func DrawMap(trncs map[int]map[int]string, minX int, maxX int, minY int, maxY int) string {
	ret := ""
	for y := minY; y <= maxY; y = y + 1 {
		for x := minX; x <= maxX; x = x + 1 {
			if trncs[x][y] == "" {
				ret = ret + "."
			} else {
				ret = ret + "#"
			}
		}
		ret = ret + "\n"
	}
	return ret
}

func TestDigTrenches(t *testing.T) {
	res, minX, maxX, minY, maxY := DigTrenches(example)

	mapOfRes := DrawMap(res, minX, maxX, minY, maxY)
	if mapOfRes != expectedTrenches {
		t.Errorf("Expected:\n%s\n\ngot:\n%s", expectedTrenches, mapOfRes)
	}
}

func TestFillTrenches(t *testing.T) {
	res, minX, maxX, minY, maxY := DigTrenches(example)

	filled := FillTrenches(res, "WOLOLOLOL")

	mapOfRes := DrawMap(filled, minX, maxX, minY, maxY)
	if mapOfRes != expectedFilled {
		t.Errorf("Expected:\n%s\n\ngot:\n%s", expectedTrenches, mapOfRes)
	}
}

func TestPartOne(t *testing.T) {
	res, _, _, _, _ := DigTrenches(example)

	filled := FillTrenches(res, "WOLOLOLOL")

	countFilled := filled.CountDug()
	if countFilled != 62 {
		t.Errorf("Expected:\n%d\n\ngot:\n%d", 62, countFilled)
	}
}
