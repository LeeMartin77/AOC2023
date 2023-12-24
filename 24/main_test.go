package main

import "testing"

var exampleData string = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

func TestParsing(t *testing.T) {
	res := ParseHailstone("18, 19, 22 @ -1, -1, -2")
	expPos := Coordinate{18, 19, 22}
	expVel := Coordinate{-1, -1, -2}
	if res.Position != expPos {
		t.Errorf("Expected %v got %v", expPos, res.Position)
	}
	if res.Velocity != expVel {
		t.Errorf("Expected %v got %v", expVel, res.Velocity)
	}
}

func TestCollisionPointXY(t *testing.T) {
	stns := ParseAllHailstones(exampleData)
	res := CollisionPointXY(stns[0], stns[1])
	if res == nil {

		t.Errorf("Expected res got nil")
	}
	exp := Coordinate{X: 14.333, Y: 15.333}
	if res != nil && exp != *res {

		t.Errorf("Expected %v got %v", exp, res)
	}
}

func TestExamplePt1(t *testing.T) {
	stns := ParseAllHailstones(exampleData)
	res := GetXYColisionsInTestArea(stns, 7, 27)
	if len(res) != 2 {

		t.Errorf("Expected %v got %v", 2, len(res))
	}
}
