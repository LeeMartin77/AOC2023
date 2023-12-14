package main

import (
	"fmt"
	"os"
	"strings"
)

type Rock struct {
	Mobile bool
	X, Y   int
}

var (
	RoundRock  = 'O'
	SquareRock = '#'
)

func ParseRocks(input string) []Rock {
	ret := []Rock{}
	for y, ln := range strings.Split(input, "\n") {
		for x, chr := range ln {
			if chr == RoundRock || chr == SquareRock {
				ret = append(ret, Rock{
					Mobile: chr == RoundRock,
					X:      x, Y: y,
				})
			}
		}
	}
	return ret
}

func RollRocksNorth(inp []Rock) []Rock {
	rocks := inp
	columnIndexes := [][]int{}
	maxX := 0
	for _, rk := range rocks {
		if maxX < rk.X {
			maxX = rk.X
		}
	}
	for x := 0; x <= maxX; x = x + 1 {
		columnIndexes = append(columnIndexes, []int{})
		for i, rk := range rocks {
			if rk.X == x {
				columnIndexes[x] = append(columnIndexes[x], i)
			}
		}
	}

	for _, col := range columnIndexes {
		orderedRocks := []*Rock{}
		for _, i := range col {
			if len(orderedRocks) == 0 {
				orderedRocks = append(orderedRocks, &rocks[i])
			} else {
				for ii := range orderedRocks {
					if orderedRocks[ii].Y < rocks[i].Y {
						if len(orderedRocks)-1 == ii {
							orderedRocks = append(orderedRocks, &rocks[i])
						}
						continue
					} else {
						orderedRocks = append([]*Rock{&rocks[i]}, orderedRocks...)
					}
				}
			}
		}
		endpoint := 0
		for _, rks := range orderedRocks {
			if rks.Mobile == false {
				endpoint = rks.Y + 1
			} else {
				rks.Y = endpoint
				endpoint = endpoint + 1
			}
		}
	}
	return rocks
}

func CalculateLoadOnNorth(rocks []Rock) int {

	maxY := 0
	counts := map[int]int{}
	for _, rk := range rocks {
		if rk.Y > maxY {
			maxY = rk.Y
		}
		if rk.Mobile {
			counts[rk.Y] = counts[rk.Y] + 1
		}
	}
	maxY = maxY + 1
	total := 0
	for y, cnt := range counts {
		total = total + ((maxY - y) * cnt)
	}
	return total
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)

	strt := ParseRocks(input)

	res := RollRocksNorth(strt)
	calc := CalculateLoadOnNorth(res)

	fmt.Printf("Result: %d\n", calc)
}
