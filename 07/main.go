package main

import (
	"strconv"
	"strings"
)

// 6 Five of a kind, where all five cards have the same label: AAAAA
// 5 Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 4 Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 3 Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 2 Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 1 One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 0 High card, where all cards' labels are distinct: 23456

type Hand struct {
	Cards    []rune
	Bid      int
	Strength int
}

func ParseHand(input string) Hand {
	parts := strings.Split(input, " ")

	hnd := Hand{}
	counts := map[rune]int{}
	for _, chr := range parts[0] {
		hnd.Cards = append(hnd.Cards, chr)
		counts[chr] = counts[chr] + 1
	}
	countValues := []int{}
	max := 0
	for _, val := range counts {
		countValues = append(countValues, val)
		if val > max {
			max = val
		}
	}
	if max == 5 {
		hnd.Strength = 6
	} else if max == 4 {
		hnd.Strength = 5
	} else if len(countValues) == 5 {
		hnd.Strength = 0
	} else if len(countValues) == 4 {
		hnd.Strength = 1
	} else if len(countValues) == 3 && max == 2 {
		hnd.Strength = 2
	} else if len(countValues) == 3 && max == 3 {
		hnd.Strength = 3
	} else {
		hnd.Strength = 4
	}

	bid, _ := strconv.Atoi(parts[1])
	hnd.Bid = bid
	return hnd
}

func ParseHandList(input string) []Hand {
	ret := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		ret = append(ret, ParseHand(line))
	}
	return ret
}

func main() {
	// buf, _ := os.ReadFile("data.txt")
	// stringput := string(buf)
	// fmt.Printf("Result: %v", myTestFunction(stringput))
}
