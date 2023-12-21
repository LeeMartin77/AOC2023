package main

import (
	"fmt"
	"os"
	"strings"
)

// map, endpoints
func ParseMap(input string) ([][]rune, [][]bool) {
	res := [][]rune{}
	pnts := [][]bool{}
	for y, line := range strings.Split(input, "\n") {
		for x, rn := range line {
			if y == 0 {
				res = append(res, []rune{})
				pnts = append(pnts, []bool{})
			}
			if rn == 'S' {
				res[x] = append(res[x], '.')
				pnts[x] = append(pnts[x], true)
			} else {
				res[x] = append(res[x], rn)
				pnts[x] = append(pnts[x], false)
			}
		}
	}
	return res, pnts
}

func PerformStep(mp [][]rune, endpoints [][]bool) [][]bool {
	newPoints := [][]bool{}
	for _, col := range mp {
		yCol := []bool{}
		for range col {
			yCol = append(yCol, false)
		}
		newPoints = append(newPoints, yCol)
	}
	for x, col := range mp {
		for y := range col {
			if endpoints[x][y] && x-1 > -1 && mp[x-1][y] == '.' {
				newPoints[x-1][y] = true
			}
			if endpoints[x][y] && x+1 < len(endpoints) && mp[x+1][y] == '.' {
				newPoints[x+1][y] = true
			}
			if endpoints[x][y] && y-1 > -1 && mp[x][y-1] == '.' {
				newPoints[x][y-1] = true
			}
			if endpoints[x][y] && y+1 < len(endpoints[0]) && mp[x][y+1] == '.' {
				newPoints[x][y+1] = true
			}
		}
	}
	return newPoints
}

func DrawMap(mp [][]rune, endpoints [][]bool) string {
	res := []string{}
	for range mp {
		res = append(res, "")
	}
	x := 0
	for _, col := range mp {
		y := 0
		for range col {
			if endpoints[x][y] {
				res[y] = res[y] + "O"
			} else {
				res[y] = res[y] + string(mp[x][y])
			}
			y = y + 1
		}
		x = x + 1
	}
	r := ""
	for _, ln := range res {
		r = r + ln + "\n"
	}
	return r
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	mp, stpnt := ParseMap(stringput)

	for i := 0; i < 64; i = i + 1 {
		stpnt = PerformStep(mp, stpnt)
	}

	cuml := 0
	for _, col := range stpnt {
		for _, val := range col {
			if val {
				cuml = cuml + 1
			}
		}
	}

	fmt.Println(DrawMap(mp, stpnt))

	fmt.Printf("Part one: %d\n", cuml)
}
