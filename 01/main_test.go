package main

import "testing"

func TestAssert(t *testing.T) {
	testInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	res := getConfigurationValue(testInput)
	if res != 142 {
		t.Errorf("Expected 142 got %v", res)
	}
}

func TestJustStraightMessMeUp(t *testing.T) {
	testInput := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	res := getConfigurationValue(testInput)
	if res != 281 {
		t.Errorf("Expected 281 got %v", res)
	}
}
