package main

import "testing"

func TestParsingOfGrid(t *testing.T) {
	grid := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	values, symbols := ParseGridToEntities(grid, '.')
	if len(values) != 10 {
		t.Errorf("Expected 10 got %v", len(values))
	}
	if len(symbols) != 6 {
		t.Errorf("Expected 6 got %v", len(symbols))
	}
	if symbols[2].Location.X != 3 && symbols[2].Location.Y != 4 {
		t.Errorf("Expected 4,3 got %v", symbols[2].Location)
	}
	if values[6].Value != 592 {
		t.Errorf("Expected 592 got %v", values[6].Value)
	}
	if len(values[6].Locations) != 3 {
		t.Errorf("Expected 3 got %v", len(values[6].Locations))
	}

	if values[6].Locations[0].X != 2 && values[6].Locations[0].Y != 6 {
		t.Errorf("Expected 2,6 got %v", symbols[2].Location)
	}
	if values[6].Locations[1].X != 3 && values[6].Locations[1].Y != 6 {
		t.Errorf("Expected 3,6 got %v", symbols[2].Location)
	}
	if values[6].Locations[2].X != 4 && values[6].Locations[2].Y != 6 {
		t.Errorf("Expected 4,6 got %v", symbols[2].Location)
	}
}

func TestGetTotalAdjacentToSymbols(t *testing.T) {
	grid := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	values, symbols := ParseGridToEntities(grid, '.')

	res := GetSumAdjacentToSymbols(values, symbols)

	if res != 4361 {

		t.Errorf("Expected 4361 got %v", res)
	}
}

func TestGetSumOfGearRatios(t *testing.T) {
	grid := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	values, symbols := ParseGridToEntities(grid, '.')

	res := GetSumOfGearRatios(values, symbols)

	if res != 467835 {

		t.Errorf("Expected 467835 got %v", res)
	}
}

func TestAdjacencyCode(t *testing.T) {
	loc := Location{X: 4, Y: 7}
	if !loc.IsAdjacent(Location{X: 3, Y: 8}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 3, Y: 7}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 3, Y: 6}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 4, Y: 8}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 4, Y: 6}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 5, Y: 8}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 5, Y: 7}) {
		t.Error("Expected adjacent")
	}
	if !loc.IsAdjacent(Location{X: 5, Y: 6}) {
		t.Error("Expected adjacent")
	}

	if loc.IsAdjacent(Location{X: 4, Y: 7}) {
		t.Error("Expected not adjacent")
	}
	if loc.IsAdjacent(Location{X: 12, Y: 10}) {
		t.Error("Expected not adjacent")
	}
}
