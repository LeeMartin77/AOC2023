package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scratchcard struct {
	CardNumber     int
	WinningNumbers []int
	CardNumbers    []int
}

func (scrd Scratchcard) ScoreScratchcard() int {
	cuml := 0
	for _, wm := range scrd.WinningNumbers {
		for _, cn := range scrd.CardNumbers {
			if cn == wm {
				if cuml == 0 {
					cuml = 1
				} else {
					cuml = cuml * 2
				}
				continue
			}
		}
	}
	return cuml
}

func (scrd Scratchcard) CardIndexesWon() []int {
	cuml := []int{}
	for _, wm := range scrd.WinningNumbers {
		for _, cn := range scrd.CardNumbers {
			if cn == wm {
				cuml = append(cuml, len(cuml)+scrd.CardNumber)
				continue
			}
		}
	}
	return cuml
}

// this is hilariously slow and brute force
func CollectCountOfCards(cards [][]Scratchcard, index int, culm int) ([][]Scratchcard, int) {
	if index >= len(cards) {
		return cards, culm
	}
	icards := cards[index]

	for _, crd := range icards {
		culm = culm + 1
		indexes := crd.CardIndexesWon()
		for _, val := range indexes {
			if val < len(cards) {
				cards[val] = append(cards[val], cards[val][0])
			}
		}
	}

	return CollectCountOfCards(cards, index+1, culm)
}

func ParseScratchcard(input string) (Scratchcard, error) {
	scard := Scratchcard{}
	splt := strings.Split(input, ":")
	nameNum := strings.Split(splt[0], " ")[len(strings.Split(splt[0], " "))-1]
	numbers := strings.Split(splt[1], "|")
	winning := strings.Split(numbers[0], " ")
	cardNumbers := strings.Split(numbers[1], " ")

	nn, err := strconv.Atoi(nameNum)
	if err != nil {
		return scard, err
	}
	scard.CardNumber = nn

	for _, wn := range winning {
		if wn != "" {
			num, err := strconv.Atoi(wn)
			if err != nil {
				return scard, err
			}
			scard.WinningNumbers = append(scard.WinningNumbers, num)
		}
	}

	for _, wn := range cardNumbers {
		if wn != "" {
			num, err := strconv.Atoi(wn)
			if err != nil {
				return scard, err
			}
			scard.CardNumbers = append(scard.CardNumbers, num)
		}
	}
	return scard, nil
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	cards := []Scratchcard{}
	multiplyingcards := [][]Scratchcard{}
	total := 0
	for _, crd := range strings.Split(input, "\n") {
		cr, _ := ParseScratchcard(crd)
		cards = append(cards, cr)
		multiplyingcards = append(multiplyingcards, []Scratchcard{cr})
		total = total + cr.ScoreScratchcard()
	}
	fmt.Printf("Part One Total: %d\n", total)
	_, multitotal := CollectCountOfCards(multiplyingcards, 0, 0)
	fmt.Printf("Part Two Total: %d\n", multitotal)
}
