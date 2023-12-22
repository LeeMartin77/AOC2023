package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y, Z int
}

type Brick struct {
	Blocks      []Coordinate
	OriginIndex int
}

func ParseCoordinate(cords string) Coordinate {
	xyz := strings.Split(cords, ",")
	xyzints := []int{}
	for _, vl := range xyz {
		num, _ := strconv.Atoi(vl)
		xyzints = append(xyzints, num)
	}
	return Coordinate{X: xyzints[0], Y: xyzints[1], Z: xyzints[2]}
}

func ParseBrick(input string) Brick {
	both := strings.Split(input, "~")
	first, last := both[0], both[1]
	if first == last {
		// one block
		return Brick{
			Blocks: []Coordinate{ParseCoordinate(first)},
		}
	}
	fcord, lcord := ParseCoordinate(first), ParseCoordinate(last)
	cur := fcord
	brck := Brick{Blocks: []Coordinate{fcord}}
	i := 1
	for cur != lcord {
		if cur.X > lcord.X {
			cur = fcord
			cur.X = cur.X - i
		}
		if cur.X < lcord.X {
			cur = fcord
			cur.X = cur.X + i
		}
		if cur.Y > lcord.Y {
			cur = fcord
			cur.Y = cur.Y - i
		}
		if cur.Y < lcord.Y {
			cur = fcord
			cur.Y = cur.Y + i
		}
		if cur.Z > lcord.Z {
			cur = fcord
			cur.Z = cur.Z - i
		}
		if cur.Z < lcord.Z {
			cur = fcord
			cur.Z = cur.Z + i
		}
		i = i + 1
		brck.Blocks = append(brck.Blocks, cur)
	}
	return brck
}

func ParseBricks(input string) []Brick {
	bricks := []Brick{}
	for i, ln := range strings.Split(input, "\n") {
		brk := ParseBrick(ln)
		brk.OriginIndex = i
		bricks = append(bricks, brk)
	}
	return bricks
}

func OrderBricksOnZ(brcks []Brick) []Brick {
	sort.SliceStable(brcks, func(i, j int) bool {
		lowestZi := 0
		lowestZj := 0
		for _, crd := range brcks[i].Blocks {
			if lowestZi < crd.Z {
				lowestZi = crd.Z
			}
		}
		for _, crd := range brcks[j].Blocks {
			if lowestZj < crd.Z {
				lowestZj = crd.Z
			}
		}
		return lowestZi < lowestZj
	})
	return brcks
}

func DropOrderedBricks(brcks []Brick) []Brick {
	nextBricks := []Brick{}
	bricksMoved := false
blks:
	for i, brk := range brcks {
		for _, t := range brk.Blocks {
			if t.Z == 1 {
				nextBricks = append(nextBricks, brk)
				continue blks
			}
			for ii, other := range brcks {
				if i != ii {
					for _, b := range other.Blocks {
						if b.X == t.X && b.Y == t.Y && b.Z == t.Z-1 {
							nextBricks = append(nextBricks, brk)
							continue blks
						}
					}
				}
			}
		}
		for i, cood := range brk.Blocks {
			cood.Z = cood.Z - 1
			brk.Blocks[i] = cood
		}
		nextBricks = append(nextBricks, brk)
		bricksMoved = true
	}
	if !bricksMoved {
		return nextBricks
	}
	return DropOrderedBricks(nextBricks)
}

func (brk Brick) Supports(otherbrk Brick) bool {
	for _, coord := range brk.Blocks {
		for _, ocrd := range otherbrk.Blocks {
			if coord.X == ocrd.X && coord.Y == ocrd.Y && coord.Z == ocrd.Z+1 {
				return true
			}
		}
	}
	return false
}

func CountDestroyableBricks(brcks []Brick) int {
	// supporting a brick, and brick not supported by any other bricks
	count := 0
allblks:
	for _, brck := range brcks {
		supporting := []Brick{}
		for _, obrk := range brcks {
			if brck.OriginIndex == obrk.OriginIndex {
				continue
			}
			if brck.Supports(obrk) {
				supporting = append(supporting, obrk)
			}
		}
		if len(supporting) == 0 {
			count = count + 1
			continue allblks
		}
		countOtherSupported := 0
	sup:
		for _, sup := range supporting {
			for _, other := range brcks {
				if sup.OriginIndex == other.OriginIndex {
					continue
				}
				if other.OriginIndex == brck.OriginIndex {
					continue
				}
				if other.Supports(sup) {
					countOtherSupported = countOtherSupported + 1
					continue sup
				}
			}
		}
		if countOtherSupported == len(supporting) {
			count = count + 1
			continue allblks
		}
	}
	return count
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)

	prsd := ParseBricks(input)
	ordered := OrderBricksOnZ(prsd)
	dropped := DropOrderedBricks(ordered)
	count := CountDestroyableBricks(dropped)
	fmt.Printf("Part one: %d\n", count)
}
