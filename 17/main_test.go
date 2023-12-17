package main

import "testing"

var exampleMap string = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestParseMap(t *testing.T) {
	res := ParseMap(exampleMap)
	if len(res) != 13 {
		t.Errorf("Expected 13 got %v", len(res))
	}
	if len(res[4]) != 13 {
		t.Errorf("Expected 13 got %v", len(res[4]))
	}

	if res[12][12] != 3 {
		t.Errorf("Expected 3 got %v", res[12][12])
	}

	if res[8][10] != 6 {
		t.Errorf("Expected 6 got %v", res[8][10])
	}
}

func TestGetDefault(t *testing.T) {

	res := ParseMap(exampleMap)
	exp := res[0][0] + res[0][1] + res[1][1] + res[1][2] + res[2][2]
	act := GetDefault(res, Coordinate{0, 0}, Coordinate{2, 2})

	if exp != act {
		t.Errorf("Expected %v got %v", exp, act)
	}
}

func TestPathing(t *testing.T) {
	res := ParseMap(exampleMap)
	minHeatLoss, _ := PathFromTo(res, Coordinate{0, 0}, Coordinate{12, 12}, 3)
	if minHeatLoss != 102 {
		t.Errorf("Expected 102 got %v", minHeatLoss)
	}
}
