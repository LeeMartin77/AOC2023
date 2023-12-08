package main

import (
	"fmt"
	"os"
	"sort"
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
	Cards         []rune
	Bid           int
	Strength      int
	JokerStrength int
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

	if counts['J'] == 5 {
		hnd.JokerStrength = 0
	} else if counts['J'] > 0 {
		jkcountValues := []int{}
		max := 0
		for k, val := range counts {
			if k == 'J' {
				continue
			}
			jkcountValues = append(jkcountValues, val)
			if val > max {
				max = val
			}
		}
		sort.SliceStable(jkcountValues, func(i, j int) bool {
			return jkcountValues[i] > jkcountValues[j]
		})
		jkcountValues[0] = jkcountValues[0] + counts['J']
		max = jkcountValues[0]
		if max == 5 {
			hnd.JokerStrength = 6
		} else if max == 4 {
			hnd.JokerStrength = 5
		} else if len(countValues) == 5 {
			hnd.JokerStrength = 0
		} else if len(countValues) == 4 {
			hnd.JokerStrength = 1
		} else if len(countValues) == 3 && max == 2 {
			hnd.JokerStrength = 2
		} else if len(countValues) == 3 && max == 3 {
			hnd.JokerStrength = 3
		} else {
			hnd.JokerStrength = 4
		}
	} else {
		hnd.JokerStrength = hnd.Strength
	}

	bid, _ := strconv.Atoi(parts[1])
	hnd.Bid = bid
	return hnd
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order

var CardStrengths map[rune]int = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func (hnd Hand) DoesItBeat(other Hand) bool {
	if hnd.Strength > other.Strength {
		return true
	}
	if hnd.Strength == other.Strength {
		for i := range hnd.Cards {
			if CardStrengths[hnd.Cards[i]] > CardStrengths[other.Cards[i]] {
				return true
			} else if CardStrengths[hnd.Cards[i]] < CardStrengths[other.Cards[i]] {
				return false
			}
		}
	}
	return false
}
func (hnd Hand) DoesItBeatJoker(other Hand) bool {
	if hnd.JokerStrength > other.JokerStrength {
		return true
	}
	if hnd.JokerStrength == other.JokerStrength {
		for i := range hnd.Cards {
			if hnd.Cards[i] == 'J' && other.Cards[i] != 'J' {
				return false
			}
			if hnd.Cards[i] != 'J' && other.Cards[i] == 'J' {
				return true
			}
			if CardStrengths[hnd.Cards[i]] > CardStrengths[other.Cards[i]] {
				return true
			} else if CardStrengths[hnd.Cards[i]] < CardStrengths[other.Cards[i]] {
				return false
			}
		}
	}
	return false
}

func ParseHandList(input string) []Hand {
	ret := []Hand{}
	for _, line := range strings.Split(input, "\n") {
		ret = append(ret, ParseHand(line))
	}
	return ret
}

//Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.

func GetPartOneResult(hands []Hand) int {
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].DoesItBeat(hands[j])
	})
	sum := 0
	currentRank := len(hands)
	for _, hnd := range hands {
		sum = sum + currentRank*hnd.Bid
		currentRank = currentRank - 1
	}
	return sum
}

func GetPartTwoResult(hands []Hand) int {
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].DoesItBeatJoker(hands[j])
	})
	sum := 0
	currentRank := len(hands)
	for _, hnd := range hands {
		sum = sum + currentRank*hnd.Bid
		currentRank = currentRank - 1
	}
	return sum
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)

	hands := ParseHandList(stringput)

	res := GetPartOneResult(hands)
	fmt.Printf("Result: %v\n", res)
	res = GetPartTwoResult(hands)
	fmt.Printf("Result: %v\n", res)
}
