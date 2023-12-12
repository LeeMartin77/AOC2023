package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	operational = '.'
	damaged     = '#'
	unknown     = '?'
)

type Record struct {
	Entries  []rune
	Checksum []int
}

func ParseLine(input string) Record {
	parts := strings.Split(input, " ")
	record := parts[0]
	checksum := parts[1]
	chk := []int{}
	for _, ch := range strings.Split(checksum, ",") {
		num, _ := strconv.Atoi(ch)
		chk = append(chk, num)
	}
	return Record{
		Entries:  []rune(record),
		Checksum: chk,
	}
}

func ChecksumBitmask(checksum []int) int {
	basemask := 0
	mskVal := 1
	for _, chk := range checksum {
		for i := 0; i < chk; i = i + 1 {
			basemask |= mskVal
			mskVal = mskVal * 2
		}
		mskVal = mskVal * 2
	}
	return basemask
}

func GetEntryBitmask(entries []rune) int {
	basemask := 0
	mskVal := 1
	lastWasOperational := true
	for _, chk := range entries {
		if lastWasOperational && chk == operational {
			continue
		}
		if chk == damaged {
			basemask |= mskVal
			mskVal = mskVal * 2
			lastWasOperational = false
			continue
		}
		mskVal = mskVal * 2
		lastWasOperational = true
	}
	return basemask
}

func FulfilsChecksum(entries []rune, checksum []int) bool {
	basemask := GetEntryBitmask(entries)
	chksm := ChecksumBitmask(checksum)
	return basemask == chksm
}

func PossibleConfigurations(input Record) int {
	// screw it, brute force time
	poss := []string{}
	for ii, chr := range input.Entries {
		if chr != unknown {
			if ii == 0 {
				poss = append(poss, string(chr))
			} else {
				for i := range poss {
					poss[i] = poss[i] + string(chr)
				}
			}
		} else {
			if ii == 0 {
				poss = append(poss, string(operational))
				poss = append(poss, string(damaged))
			} else {
				newPosses := []string{}
				for i := range poss {
					new := poss[i] + string(operational)
					newPosses = append(newPosses, new)
					poss[i] = poss[i] + string(damaged)
				}
				poss = append(poss, newPosses...)
			}
		}
	}

	count := 0

	chksm := ChecksumBitmask(input.Checksum)
	for _, p := range poss {
		if GetEntryBitmask([]rune(p)) == chksm {
			count = count + 1
		}
	}

	return count
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	cuml := 0
	for _, str := range strings.Split(stringput, "\n") {
		cuml = cuml + PossibleConfigurations(ParseLine(str))
	}
	fmt.Println(cuml)
}
