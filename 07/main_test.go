package main

import (
	"reflect"
	"testing"
)

func TestParseList(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	hands := ParseHandList(input)

	if len(hands) != 5 {
		t.Errorf("Expected 5 got %v", len(hands))
	}

	expectedHand := Hand{
		Cards:         []rune{'K', 'K', '6', '7', '7'},
		Bid:           28,
		Strength:      2,
		JokerStrength: 2,
	}

	if !reflect.DeepEqual(hands[2], expectedHand) {
		t.Errorf("Expected %v got %v", expectedHand, hands[2])
	}
}

func TestPartOneResult(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	hands := ParseHandList(input)

	res := GetPartOneResult(hands)
	if res != 6440 {
		t.Errorf("Expected 6440 got %d", res)
	}
}

func TestPartTwoResult(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	hands := ParseHandList(input)

	res := GetPartTwoResult(hands)
	if res != 5905 {
		t.Errorf("Expected 5905 got %d", res)
	}
}

func TestParsingStrengths(t *testing.T) {
	input := `23456 0
23455 0
22344 0
33345 0
33344 0
QQQQA 0
QQQQQ 0`
	hands := ParseHandList(input)

	if len(hands) != 7 {
		t.Errorf("Expected 7 got %v", len(hands))
	}

	for i, hnd := range hands {
		if hnd.Strength != i {
			t.Errorf("Expected %d got %d", i, hnd.Strength)
		}
	}
}

func TestDoesItBeat_DiffStrength(t *testing.T) {
	handOne := Hand{
		Cards:    []rune{'K', 'K', '6', '7', '7'},
		Bid:      28,
		Strength: 2,
	}

	handTwo := Hand{
		Cards:    []rune{'K', 'K', '7', '7', '7'},
		Bid:      28,
		Strength: 4,
	}

	if handOne.DoesItBeat(handTwo) {
		t.Errorf("Expected false got true")
	}
	if !handTwo.DoesItBeat(handOne) {
		t.Errorf("Expected true got false")
	}
}

func TestDoesItBeat_SameStrength(t *testing.T) {
	handOne := Hand{
		Cards:    []rune{'K', 'K', '6', '7', '7'},
		Bid:      28,
		Strength: 2,
	}

	handTwo := Hand{
		Cards:    []rune{'A', 'A', '6', '7', '7'},
		Bid:      28,
		Strength: 2,
	}

	if handOne.DoesItBeat(handTwo) {
		t.Errorf("Expected false got true")
	}
	if !handTwo.DoesItBeat(handOne) {
		t.Errorf("Expected true got false")
	}
}

func TestDoesItBeat_SameStrength_NotFirst(t *testing.T) {
	handOne := Hand{
		Cards:    []rune{'K', 'K', '6', '7', '7'},
		Bid:      28,
		Strength: 2,
	}

	handTwo := Hand{
		Cards:    []rune{'K', 'K', '7', '6', '7'},
		Bid:      28,
		Strength: 2,
	}

	if handOne.DoesItBeat(handTwo) {
		t.Errorf("Expected false got true")
	}
	if !handTwo.DoesItBeat(handOne) {
		t.Errorf("Expected true got false")
	}
}

func TestParse(t *testing.T) {
	input := `32T3K 765`
	hand := ParseHand(input)

	expectedHand := Hand{
		Cards:         []rune{'3', '2', 'T', '3', 'K'},
		Bid:           765,
		Strength:      1,
		JokerStrength: 1,
	}

	if !reflect.DeepEqual(hand, expectedHand) {
		t.Errorf("Expected %v got %v", expectedHand, hand)
	}
}
