package main

import (
	"reflect"
	"testing"
)

var exampleInput string = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestParsePatterns(t *testing.T) {
	res := ParsePatterns(exampleInput)
	if len(res) != 2 {
		t.Errorf("Expected 2 got %v", len(res))
	}
	// just spot test a bit
	if res[0].Lines[2] != "##......#" {
		t.Errorf("Expected '##......#' got '%v'", res[0].Lines[2])
	}
	// just spot test a bit
	if res[1].Lines[6] != "#....#..#" {
		t.Errorf("Expected '#....#..#' got '%v'", res[1].Lines[6])
	}
}

func TestGetHorizontalHashes(t *testing.T) {

	res := ParsePatterns(exampleInput)[0]
	hsh := res.GetHorizontalHashes()
	exp := []int{205, 180, 259, 259, 180, 204, 181}
	if !reflect.DeepEqual(hsh, exp) {
		t.Errorf("Expected '%v' got '%v'", exp, hsh)
	}
}

func TestGetVerticalHashes(t *testing.T) {

	res := ParsePatterns(exampleInput)[0]
	hsh := res.GetVerticalHashes()
	exp := []int{77, 12, 115, 33, 82, 82, 33, 115, 12}
	if !reflect.DeepEqual(hsh, exp) {
		t.Errorf("Expected '%v' got '%v'", exp, hsh)
	}
}

func TestIntegration(t *testing.T) {
	ptrns := ParsePatterns(exampleInput)
	cuml, _ := AccumulateReflections(ptrns)
	if cuml != 405 {
		t.Errorf("Expected '405' got '%v'", cuml)
	}
}

func TestFindReflectionPoint(t *testing.T) {
	_, err := FindReflectionPoint([]int{205, 180, 259, 259, 180, 204, 181})
	if err == nil {
		t.Errorf("Expected error")
	}
	i, err := FindReflectionPoint([]int{77, 12, 115, 33, 82, 82, 33, 115, 12})
	if err != nil {
		t.Errorf("Unexpected error")
	}
	if i != 5 {
		t.Errorf("Expected '5' got '%v'", i)
	}
}
