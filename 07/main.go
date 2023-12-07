package main

import (
	"strconv"
	"strings"
)

type Hand struct {
	Cards []rune
	Bid   int
}

func ParseHand(input string) Hand {
	parts := strings.Split(input, " ")

	hnd := Hand{}
	for _, chr := range parts[0] {
		hnd.Cards = append(hnd.Cards, chr)
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
