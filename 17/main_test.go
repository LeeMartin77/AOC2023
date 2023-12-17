package main

import (
	"testing"
)

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

func TestPathing(t *testing.T) {
	res := ParseMap(exampleMap)
	minHeatLoss, history := PathFromTo(res, Coordinate{0, 0}, Coordinate{12, 12}, 3)
	if minHeatLoss != 102 {
		t.Errorf("Expected 102 got %v", minHeatLoss)
		DrawMap(res, history)
	}
}
