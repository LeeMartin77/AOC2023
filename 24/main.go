package main

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y, Z float64
}

type Hailstone struct {
	Position, Velocity Coordinate
}

func unsafeGetInt(input string) float64 {
	input = strings.Trim(input, " ")
	num, _ := strconv.Atoi(input)
	return float64(num)
}

func ParseHailstone(input string) Hailstone {
	pts := strings.Split(input, "@")
	pos := []float64{}
	for _, str := range strings.Split(pts[0], ",") {
		pos = append(pos, unsafeGetInt(str))
	}
	vel := []float64{}
	for _, str := range strings.Split(pts[1], ",") {
		vel = append(vel, unsafeGetInt(str))
	}
	return Hailstone{
		Coordinate{pos[0], pos[1], pos[2]},
		Coordinate{vel[0], vel[1], vel[2]},
	}
}

func ParseAllHailstones(input string) []Hailstone {
	hstns := []Hailstone{}
	for _, str := range strings.Split(input, "\n") {
		if str != "" {
			hstns = append(hstns, ParseHailstone(str))
		}
	}
	return hstns
}

func CollisionPointXY(one Hailstone, two Hailstone) *Coordinate {
	if one.Velocity.X == two.Velocity.X && one.Velocity.Y == two.Velocity.Y {
		// parallel
		return nil
	}
	denominator := two.Velocity.X*one.Velocity.Y - one.Velocity.X*two.Velocity.Y
	t := ((two.Position.X-one.Position.X)*one.Velocity.Y -
		(two.Position.Y-one.Position.Y)*one.Velocity.X) / denominator

	return &Coordinate{
		X: one.Position.X + t*one.Velocity.X,
		Y: one.Position.Y + t*one.Velocity.Y,
	}
}

func GetXYColisionsInTestArea(stones []Hailstone, xylow float64, xyhigh float64) []Coordinate {
	res := []Coordinate{}
	return res
}

func main() {
	// buf, _ := os.ReadFile("data.txt")
	// stringput := string(buf)
	// fmt.Printf("Result: %v", myTestFunction(stringput))
}
