package main

import "testing"

// x, y, z
var example string = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

func TestParseBricks(t *testing.T) {
	res := ParseBricks(example)
	if len(res) != 7 {
		t.Errorf("Expected 7 got %v", len(res))
	}
}

func TestParseBrick(t *testing.T) {
	res := ParseBrick("0,2,3~2,2,3")
	if len(res.Blocks) != 3 {
		t.Errorf("Expected 3 got %v", len(res.Blocks))
	}
	expectedBlocks := []Coordinate{{0, 2, 3}, {1, 2, 3}, {2, 2, 3}}
	for i, exp := range expectedBlocks {
		if exp != res.Blocks[i] {
			t.Errorf("Expected %v got %v", exp, res.Blocks[i])
		}
	}
}
