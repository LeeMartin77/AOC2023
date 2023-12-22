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
	i := 0
	for _, brk := range res {
		if brk.OriginIndex != i {
			t.Errorf("Expected %v got %v", i, brk.OriginIndex)
		}
		i = i + 1
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

func TestOrderingOnZ(t *testing.T) {
	prsd := ParseBricks(example)
	jumbled := []Brick{
		prsd[2],
		prsd[1],
		prsd[0],
		prsd[5],
		prsd[4],
		prsd[3],
		prsd[6],
	}
	ordered := OrderBricksOnZ(jumbled)
	i := 0
	for _, brk := range ordered {
		if brk.OriginIndex != i {
			t.Errorf("Expected %v got %v", i, brk.OriginIndex)
		}
		i = i + 1
	}
}

func TestExampleOne(t *testing.T) {
	prsd := ParseBricks(example)
	OrderBricksOnZ(prsd)
	//dropped := DropOrderedBricks(ordered)
}
