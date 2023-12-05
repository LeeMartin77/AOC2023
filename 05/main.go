package main

import (
	"fmt"
	"os"
)

type Almanac struct {
}

func (almnc *Almanac) GetLowestLocationNumberThatCanTakeAnySeed() int {
	return 0
}

func ParseAlmanac(input string) (*Almanac, error) {
	return nil, fmt.Errorf("Not implemented")
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	almanac, _ := ParseAlmanac(stringput)

	res := almanac.GetLowestLocationNumberThatCanTakeAnySeed()
	fmt.Printf("Pt 1 Result: %v", res)
}
