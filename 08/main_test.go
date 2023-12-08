package main

import (
	"reflect"
	"testing"
)

func TestPt1(t *testing.T) {
	input := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	com, nds := ParseCommandsAndNodes(input)

	res := CountLengthThroughNodes(com, nds)
	if res != 2 {
		t.Errorf("Expected 2 got %v", res)
	}
}

func TestPt1Alt(t *testing.T) {
	input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	com, nds := ParseCommandsAndNodes(input)

	res := CountLengthThroughNodes(com, nds)
	if res != 6 {
		t.Errorf("Expected 6 got %v", res)
	}
}

func TestPt2(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	com, nds := ParseCommandsAndNodes(input)

	res := CountGhostLengthThroughNodes(com, nds)
	if res != 6 {
		t.Errorf("Expected 6 got %v", res)
	}
}

func TestParseNodesAndCommands(t *testing.T) {
	input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	com, nds := ParseCommandsAndNodes(input)

	expectedCom := []rune{'L', 'L', 'R'}
	expectedNodes := []Node{
		{
			Name:  "AAA",
			Left:  "BBB",
			Right: "BBB",
		},
		{
			Name:  "BBB",
			Left:  "AAA",
			Right: "ZZZ",
		}, {
			Name:  "ZZZ",
			Left:  "ZZZ",
			Right: "ZZZ",
		},
	}

	for i, rn := range expectedCom {
		if rn != com[i] {
			t.Errorf("Expected %v got %v", rn, com[i])
		}
	}

	for _, expected := range expectedNodes {
		if !reflect.DeepEqual(nds[expected.Name], expected) {
			t.Errorf("Expected %v got %v", expected, nds[expected.Name])
		}
	}
}

func TestParseNodeFromLine(t *testing.T) {
	res := ParseNodeFromLine("AAA = (BBB, CCC)")
	expected := Node{
		Name:  "AAA",
		Left:  "BBB",
		Right: "CCC",
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v got %v", expected, res)
	}
}
