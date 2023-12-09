package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AllSame(input []int) bool {
	common := input[0]
	for _, i := range input {
		if i != common {
			return false
		}
	}
	return true
}

func GetDiffs(input []int) []int {
	diffs := []int{}
	top := len(input) - 1
	for i, val := range input {
		if i == top {
			break
		}
		diffs = append(diffs, input[i+1]-val)
	}
	return diffs
}

func GetNextNumber(input []int) int {
	diffs := GetDiffs(input)
	if AllSame(diffs) {
		return input[len(input)-1] + diffs[0]
	} else {
		return input[len(input)-1] + GetNextNumber(diffs)
	}
}

func GetPreviousNumber(input []int) int {
	diffs := GetDiffs(input)
	if AllSame(diffs) {
		return input[0] - diffs[0]
	} else {
		return input[0] - GetPreviousNumber(diffs)
	}
}

func ParseStringToInput(str string) [][]int {
	input := [][]int{}
	for _, ln := range strings.Split(str, "\n") {
		line := []int{}
		for _, strnum := range strings.Split(ln, " ") {
			num, _ := strconv.Atoi(strnum)
			line = append(line, num)
		}
		input = append(input, line)
	}
	return input
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)
	input := ParseStringToInput(stringput)
	cuml := 0
	for _, ln := range input {
		cuml = cuml + GetNextNumber(ln)
	}
	revcuml := 0
	for _, ln := range input {
		revcuml = revcuml + GetPreviousNumber(ln)
	}
	fmt.Printf("Pt1: %v\n", cuml)
	fmt.Printf("Pt2: %v\n", revcuml)
}
