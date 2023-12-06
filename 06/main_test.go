package main

import (
	"reflect"
	"testing"
)

func TestAssert(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	res := ParseRaces(input)
	expectedRaces := []Race{
		{
			Duration: 7,
			Record:   9,
		},
		{
			Duration: 15,
			Record:   40,
		},
		{
			Duration: 30,
			Record:   200,
		},
	}
	for i := range res {
		if !reflect.DeepEqual(expectedRaces[i], res[i]) {
			t.Errorf("Expected %v got %v", expectedRaces[i], res[i])
		}
	}
}

func TestMakeMegarace(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	rcs := ParseRaces(input)

	res := GenerateMegarace(rcs)

	expected := Race{
		Duration: 71530,
		Record:   940200,
	}
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("Expected %v got %v", expected, res)
	}
}

func TestGetVictoryBounds(t *testing.T) {
	testRace := Race{
		Duration: 7,
		Record:   9,
	}

	lower, upper := testRace.GetVictoryBounds()

	if lower != 2 {
		t.Errorf("Expected 2 got %v", lower)
	}
	if upper != 5 {
		t.Errorf("Expected 5 got %v", upper)
	}
}

func TestExample(t *testing.T) {

	input := `Time:      7  15   30
Distance:  9  40  200`
	res := ParseRaces(input)

	total := 1

	for _, rc := range res {
		lower, upper := rc.GetVictoryBounds()
		total = total * (upper - lower + 1)
	}

	if total != 288 {
		t.Errorf("Expected 288 got %v", total)
	}
}
