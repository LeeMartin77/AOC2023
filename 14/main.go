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

func ParseRocks(input string) ([]Rock, int, int) {
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
	maxX := 0
	maxY := 0
	for _, rk := range ret {
		if rk.X > maxX {
			maxX = rk.X
		}

		if rk.Y > maxY {
			maxY = rk.Y
		}
	}
	return ret, maxX, maxY
}

func RollRocksNorth(inp []Rock, maxX int, maxY int) []Rock {
	rocks := inp
	columnIndexes := [][]int{}
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

func RollRocksSouth(inp []Rock, maxX int, maxY int) []Rock {
	rocks := inp
	columnIndexes := [][]int{}
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
					if len(orderedRocks)-1 == ii {
						orderedRocks = append([]*Rock{&rocks[i]}, orderedRocks...)
						break
					}
					if orderedRocks[ii].Y > rocks[i].Y {
						orderedRocks = append(orderedRocks, &rocks[i])
						break
					} else {

						continue
					}
				}
			}
		}
		endpoint := maxY
		for _, rks := range orderedRocks {
			if rks.Mobile == false {
				endpoint = rks.Y - 1
			} else {
				rks.Y = endpoint
				endpoint = endpoint - 1
			}
		}
	}
	return rocks
}

func RollRocksWest(inp []Rock, maxX int, maxY int) []Rock {
	rocks := inp
	columnIndexes := [][]int{}
	for y := 0; y <= maxY; y = y + 1 {
		columnIndexes = append(columnIndexes, []int{})
		for i, rk := range rocks {
			if rk.Y == y {
				columnIndexes[y] = append(columnIndexes[y], i)
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
					if orderedRocks[ii].X < rocks[i].X {
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
				endpoint = rks.X + 1
			} else {
				rks.X = endpoint
				endpoint = endpoint + 1
			}
		}
	}
	return rocks
}

func RollRocksEast(inp []Rock, maxX int, maxY int) []Rock {
	rocks := inp
	columnIndexes := [][]int{}
	for y := 0; y <= maxY; y = y + 1 {
		columnIndexes = append(columnIndexes, []int{})
		for i, rk := range rocks {
			if rk.Y == y {
				columnIndexes[y] = append(columnIndexes[y], i)
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
					if len(orderedRocks)-1 == ii {
						orderedRocks = append([]*Rock{&rocks[i]}, orderedRocks...)
						break
					}
					if orderedRocks[ii].X > rocks[i].X {
						orderedRocks = append(orderedRocks, &rocks[i])
						break
					} else {

						continue
					}
				}
			}
		}
		endpoint := maxX
		for _, rks := range orderedRocks {
			if rks.Mobile == false {
				endpoint = rks.X - 1
			} else {
				rks.X = endpoint
				endpoint = endpoint - 1
			}
		}
	}
	return rocks
}

// screw it

func SpinCycle(inp []Rock, maxX int, maxY int) []Rock {
	inp2 := RollRocksNorth(inp, maxX, maxY)
	inp3 := RollRocksWest(inp2, maxX, maxY)
	inp4 := RollRocksSouth(inp3, maxX, maxY)
	inp5 := RollRocksEast(inp4, maxX, maxY)
	return inp5
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

	strt, maxX, maxY := ParseRocks(input)

	res := RollRocksNorth(strt, maxX, maxY)
	calc := CalculateLoadOnNorth(res)

	fmt.Printf("Result: %d\n", calc)
}
