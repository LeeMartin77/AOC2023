package main

import (
	"reflect"
	"strings"
	"testing"
)

// pt 1 score:

func TestParseScratchcard(t *testing.T) {
	res, err := ParseScratchcard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	expected := Scratchcard{
		CardNumber:     1,
		WinningNumbers: []int{41, 48, 83, 86, 17},
		CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v | Got %v", expected, res)
	}
}

func TestScoreScratchcard(t *testing.T) {
	res, err := ParseScratchcard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	if err != nil {
		t.Errorf("%v", err)
	}
	if res.ScoreScratchcard() != 8 {
		t.Errorf("Expected 8 | Got %v", res.ScoreScratchcard())
	}
}

func TestIntegrationWithBothPt1(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	cards := []Scratchcard{}
	total := 0
	for _, crd := range strings.Split(input, "\n") {
		cr, _ := ParseScratchcard(crd)
		cards = append(cards, cr)
		total = total + cr.ScoreScratchcard()
	}
	if 13 != total {
		t.Errorf("Expected 13 | Got %v", total)
	}
}

func TestIntegrationWithBothPt2(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	cards := [][]Scratchcard{}
	for _, crd := range strings.Split(input, "\n") {
		cr, _ := ParseScratchcard(crd)
		cards = append(cards, []Scratchcard{cr})
	}

	_, total := CollectCountOfCards(cards, 0, 0)

	// for i, cal := range culm {
	// 	fmt.Printf("index: %d | count %d\n", i, len(cal))
	// }

	if 30 != total {
		t.Errorf("Expected 30 | Got %v", total)
	}
}
