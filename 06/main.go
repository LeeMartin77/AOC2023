package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Race struct {
	Duration int
	Record   int
}

func (rc Race) GetVictoryBounds() (int, int) {
	// work out lowest
	lower := -1
	i := 0
	for lower < 0 {
		if (rc.Duration-i)*i > rc.Record {
			lower = i
		}
		i = i + 1
	}
	// work out highest
	upper := -1
	i = rc.Duration
	for upper < 0 {
		if (rc.Duration-i)*i > rc.Record {
			upper = i
		}
		i = i - 1
	}
	return lower, upper
}

func ParseRaces(input string) []Race {
	lines := strings.Split(input, "\n")
	durations := []int{}
	records := []int{}

	buffer := ""

	for _, chr := range lines[0] {
		if unicode.IsNumber(chr) {
			buffer = buffer + string(chr)
		} else if len(buffer) > 0 {
			i, _ := strconv.Atoi(buffer)
			durations = append(durations, i)
			buffer = ""
		}
	}
	if len(buffer) > 0 {
		i, _ := strconv.Atoi(buffer)
		durations = append(durations, i)
		buffer = ""
	}
	for _, chr := range lines[1] {
		if unicode.IsNumber(chr) {
			buffer = buffer + string(chr)
		} else if len(buffer) > 0 {
			i, _ := strconv.Atoi(buffer)
			records = append(records, i)
			buffer = ""
		}
	}
	if len(buffer) > 0 {
		i, _ := strconv.Atoi(buffer)
		records = append(records, i)
		buffer = ""
	}
	races := []Race{}
	for i, dur := range durations {
		races = append(races, Race{
			Duration: dur,
			Record:   records[i],
		})
	}
	return races
}

func GenerateMegarace(rcs []Race) Race {
	durBuff := ""
	recBuff := ""

	for _, rc := range rcs {
		durBuff = durBuff + fmt.Sprint(rc.Duration)
		recBuff = recBuff + fmt.Sprint(rc.Record)
	}

	dur, _ := strconv.Atoi(durBuff)
	rec, _ := strconv.Atoi(recBuff)

	return Race{
		Duration: dur,
		Record:   rec,
	}
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	res := ParseRaces(stringput)

	total := 1

	for _, rc := range res {
		lower, upper := rc.GetVictoryBounds()
		total = total * (upper - lower + 1)
	}

	fmt.Printf("Result 1: %v\n", total)

	mega := GenerateMegarace(res)

	lower, upper := mega.GetVictoryBounds()

	fmt.Printf("Result 2: %v\n", upper-lower+1)
}
