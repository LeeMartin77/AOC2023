package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
			break
		}
	}
	if sdm == nil {
		return source
	}
	offset := sdm.Size - (sdm.GetSourceEnd() - source) - 1
	return sdm.DestinationStart + offset
}

type Almanac struct {
	Seeds         []int
	MappingTitles []string
	Mappings      []SourceDestinationRange
}

func (almnc *Almanac) GetLowestLocationNumberThatCanTakeAnySeed() int {
	res := 0
	for _, seed := range almnc.Seeds {
		sres := seed
		for _, mping := range almnc.Mappings {
			sres = mping.GetDesinationForSource(sres)
		}
		if sres < res || res == 0 {
			res = sres
		}
	}
	return res
}

func ParseAlmanac(input string) (*Almanac, error) {
	almnc := Almanac{}
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			// parse seeds
			seedNumbers := strings.Split(strings.Replace(line, "seeds: ", "", 1), " ")
			for _, num := range seedNumbers {
				n, err := strconv.Atoi(num)
				if err != nil {
					return nil, err
				}
				almnc.Seeds = append(almnc.Seeds, n)
			}
			continue
		}
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			almnc.MappingTitles = append(almnc.MappingTitles, strings.Replace(line, " map:", "", 1))
			almnc.Mappings = append(almnc.Mappings, SourceDestinationRange{})
			continue
		}

		nums := []int{}
		for _, num := range strings.Split(line, " ") {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			nums = append(nums, n)
		}
		if len(nums) < 3 {
			return nil, fmt.Errorf("Something went wrong, we seem to be missing numbers")
		}
		almnc.Mappings[len(almnc.Mappings)-1].Ranges = append(almnc.Mappings[len(almnc.Mappings)-1].Ranges, SourceDestinationMapping{
			DestinationStart: nums[0],
			SourceStart:      nums[1],
			Size:             nums[2],
		})
	}
	return &almnc, nil
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	almanac, _ := ParseAlmanac(stringput)

	res := almanac.GetLowestLocationNumberThatCanTakeAnySeed()
	fmt.Printf("Pt 1 Result: %v", res)
}
