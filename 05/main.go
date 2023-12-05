package main

import (
	"fmt"
	"os"
)

type SourceDestinationMapping struct {
	SourceStart      int
	DestinationStart int
	Size             int
}

func (sdm SourceDestinationMapping) GetSourceEnd() int {
	// it's inclusive so
	if sdm.Size < 1 {
		return -1
	}
	return sdm.Size + sdm.SourceStart - 1
}

type SourceDestinationRange struct {
	Ranges []SourceDestinationMapping
}

func (sdr SourceDestinationRange) GetDesinationForSource(source int) int {
	var sdm *SourceDestinationMapping
	for _, mping := range sdr.Ranges {
		if source >= mping.SourceStart && source <= mping.GetSourceEnd() {
			sdm = &mping
		}
	}
	if sdm == nil {
		return source
	}
	offset := sdm.Size - (sdm.GetSourceEnd() - source) - 1
	return sdm.DestinationStart + offset
}

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
