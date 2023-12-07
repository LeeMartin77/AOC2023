package main

import (
	"reflect"
	"testing"
)

func xTestAssert(t *testing.T) {
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
		Cards: []rune{'K', 'K', '6', '7', '7'},
		Bid:   28,
	}

	if !reflect.DeepEqual(hands[2], expectedHand) {
		t.Errorf("Expected %v got %v", expectedHand, hands[2])
	}
}

func TestAssert(t *testing.T) {
	input := `32T3K 765`
	hand := ParseHand(input)

	expectedHand := Hand{
		Cards: []rune{'3', '2', 'T', '3', 'K'},
		Bid:   765,
	}

	if !reflect.DeepEqual(hand, expectedHand) {
		t.Errorf("Expected %v got %v", expectedHand, hand)
	}
}
