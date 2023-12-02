package main

import "testing"

func TestParsing(t *testing.T) {

	gm := parseGameFromString("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	if gm.Id != 1 {
		t.Errorf("Expected 1 got %v", gm.Id)
	}

	if gm.Rounds[0].R != 4 || gm.Rounds[0].B != 3 || gm.Rounds[0].G != 0 {
		t.Errorf("Expected round one got %v", gm.Rounds[0])
	}
}

func TestAssert(t *testing.T) {
	data := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	res := getTotalOfPossibleGames(data, BlockSet{
		R: 12,
		G: 13,
		B: 14,
	})
	if res != 8 {
		t.Errorf("Expected 8 got %v", res)
	}
}
