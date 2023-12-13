package main

import (
	"fmt"
	"os"
	"strings"
)

type Pattern struct {
	Lines []string
}

const (
	ash  = '.'
	rock = '#'
)

func (ptrn Pattern) GetHorizontalHashes() []int {
	hashes := []int{}
	for _, ln := range ptrn.Lines {
		mk := 1
		msk := 0
		for _, chr := range ln {
			if chr == rock {
				msk |= mk
			}
			mk = mk * 2
		}
		hashes = append(hashes, msk)
	}
	return hashes
}

func (ptrn Pattern) GetVerticalHashes() []int {
	hashes := []int{}
	for range ptrn.Lines[0] {
		hashes = append(hashes, 0)
	}
	for i := range hashes {
		mk := 1
		for _, ln := range ptrn.Lines {
			if ln[i] == rock {
				hashes[i] |= mk
			}
			mk = mk * 2
		}
	}
	return hashes
}

// lolbruteforcetime
func (ptrn Pattern) GetSmudgedPatterns() []Pattern {
	ptrns := []Pattern{}
	for i, ln := range ptrn.Lines {
		for ii := range ln {
			pt := Pattern{}
			for j, li := range ptrn.Lines {
				pt.Lines = append(pt.Lines, li)
				if j == i {
					// sub a char
					if pt.Lines[j][ii] == ash {
						pt.Lines[j] = pt.Lines[j][:ii] + string(rock) + pt.Lines[j][ii+1:]
					} else {
						pt.Lines[j] = pt.Lines[j][:ii] + string(ash) + pt.Lines[j][ii+1:]
					}
				}
			}
			ptrns = append(ptrns, pt)
		}
	}
	return ptrns
}

func (ptrn Pattern) GetSmudgedReflectionPoint() (bool, int, error) {
	patterns := ptrn.GetSmudgedPatterns()
	for _, pat := range patterns {
		vert, idx, err := GetReflectionPoint(pat)
		if err == nil {
			return vert, idx, err
		}
	}
	return true, 0, fmt.Errorf("No reflection")
}

// vertical, index, notfound
func GetReflectionPoint(ptrn Pattern) (bool, int, error) {
	vert := true
	indx, err := FindReflectionPoint(ptrn.GetVerticalHashes())
	if err != nil {
		vert = false
		indx, err = FindReflectionPoint(ptrn.GetHorizontalHashes())
		if err != nil {
			return true, 0, fmt.Errorf("No reflection")
		}
	}
	return vert, indx, nil
}

// will be index *before* reflection
// this is genuinely fucking horrific
func FindReflectionPoint(nums []int) (int, error) {
outer:
	for i := range nums {
		if i == 0 {
			continue
		}
		firstHalf := nums[:i]
		secondHalf := nums[i:]
		firstHalfRev := []int{}
		for ii := len(firstHalf) - 1; ii >= 0; ii-- {
			firstHalfRev = append(firstHalfRev, firstHalf[ii])
		}
		if len(firstHalfRev) > len(secondHalf) {
			for ii, num := range secondHalf {
				if num != firstHalfRev[ii] {
					continue outer
				}
			}

		} else {
			for ii, num := range firstHalfRev {
				if num != secondHalf[ii] {
					continue outer
				}
			}
		}
		return len(firstHalfRev), nil
	}
	return 0, fmt.Errorf("No Reflection")
}

func ParsePatterns(input string) []Pattern {
	patterns := []Pattern{{}}

	i := 0
	for _, ln := range strings.Split(input, "\n") {
		if len(ln) == 0 {
			patterns = append(patterns, Pattern{})
			i = i + 1
			continue
		}
		patterns[i].Lines = append(patterns[i].Lines, ln)
	}
	return patterns
}

func AccumulateReflections(patterns []Pattern) (int, error) {
	cuml := 0
	for i, ptrn := range patterns {
		vert, indx, err := GetReflectionPoint(ptrn)
		if err != nil {
			return 0, fmt.Errorf("Error on pattern: %v\n", i)
		}
		num := indx
		if vert {
			cuml = cuml + num
		} else {
			cuml = cuml + (num * 100)
		}
	}
	return cuml, nil
}

func AccumulateSmudgedReflections(patterns []Pattern) (int, error) {
	cuml := 0
	for i, ptrn := range patterns {
		vert, indx, err := ptrn.GetSmudgedReflectionPoint()
		if err != nil {
			return 0, fmt.Errorf("Error on pattern: %v\n", i)
		}
		num := indx
		if vert {
			cuml = cuml + num
		} else {
			cuml = cuml + (num * 100)
		}
	}
	return cuml, nil
}

func main() {

	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	ptrns := ParsePatterns(input)
	cuml, err := AccumulateReflections(ptrns)
	if err != nil {

		fmt.Printf("Err: %v\n", err)
	}
	fmt.Printf("Result: %v\n", cuml)
	cuml2, err2 := AccumulateSmudgedReflections(ptrns)
	if err2 != nil {

		fmt.Printf("Err: %v\n", err2)
	}
	fmt.Printf("Result2: %v\n", cuml2)
}
