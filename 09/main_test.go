package main

import (
	"testing"
)

var example string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestDiffs(t *testing.T) {
	input := []int{0, 3, 6, 9, 12, 15}
	res := GetDiffs(input)
	for _, val := range res {
		if val != 3 {
			t.Errorf("Expected 3's got %v", res)
		}
	}
}

func TestGetNextNumber(t *testing.T) {
	input := []int{0, 3, 6, 9, 12, 15}
	res := GetNextNumber(input)
	if res != 18 {
		t.Errorf("Expected 18 got %v", res)
	}
}

func TestIntegrationPt1(t *testing.T) {
	res := ParseStringToInput(example)
	cuml := 0
	for _, ln := range res {
		cuml = cuml + GetNextNumber(ln)
	}
	if cuml != 114 {
		t.Errorf("Expected 114 got %v", cuml)
	}
}

func TestIntegrationPt2(t *testing.T) {
	res := ParseStringToInput(example)
	cuml := 0
	for _, ln := range res {
		cuml = cuml + GetPreviousNumber(ln)
	}
	if cuml != 2 {
		t.Errorf("Expected 2 got %v", cuml)
	}
}

func TestAllSame(t *testing.T) {
	if !AllSame([]int{3, 3, 3, 3, 3}) {
		t.Errorf("Expected true got false")
	}
	if AllSame([]int{3, 3, 2, 3, 3}) {
		t.Errorf("Expected false got true")
	}
}
