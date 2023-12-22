package main

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y, Z int
}

type Brick struct {
	Blocks []Coordinate
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
	for _, ln := range strings.Split(input, "\n") {
		bricks = append(bricks, ParseBrick(ln))
	}
	return bricks
}

func main() {
	//buf, _ := os.ReadFile("data.txt")
	//input := string(buf)

}
