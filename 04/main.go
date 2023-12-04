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
	total := 0
	for _, crd := range strings.Split(input, "\n") {
		cr, _ := ParseScratchcard(crd)
		cards = append(cards, cr)
		total = total + cr.ScoreScratchcard()
	}
	fmt.Printf("Part One Total: %d\n", total)
}
