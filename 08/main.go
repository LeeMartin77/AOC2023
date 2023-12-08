package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name  string
	Left  string
	Right string
}

func ParseNodeFromLine(input string) Node {
	return Node{
		Name:  input[:3],
		Left:  input[7:10],
		Right: input[12:15],
	}
}

func ParseCommandsAndNodes(input string) ([]rune, map[string]Node) {
	com := []rune{}
	nds := map[string]Node{}
	for i, ln := range strings.Split(input, "\n") {
		if i == 0 {
			com = []rune(ln)
			continue
		}
		if i == 1 {
			continue
		}
		nd := ParseNodeFromLine(ln)
		nds[nd.Name] = nd
	}
	return com, nds
}

func CountLengthThroughNodes(com []rune, nds map[string]Node) int {
	count := 0
	position := "AAA"
	comPos := 0
	for position != "ZZZ" {
		if comPos >= len(com) {
			comPos = 0
		}
		if com[comPos] == 'L' {
			position = nds[position].Left
		}
		if com[comPos] == 'R' {
			position = nds[position].Right
		}
		count = count + 1
		comPos = comPos + 1
	}
	return count
}

func AllEndInRune(arr []string, rn rune) bool {
	for _, val := range arr {
		if rune(val[2]) != rn {
			return false
		}
	}

	return true
}

func CountGhostLengthThroughNodes(com []rune, nds map[string]Node) int {
	count := 0
	positions := []string{}
	for k := range nds {
		if k[2] == 'A' {
			positions = append(positions, k)
		}
	}
	comPos := 0
	for !AllEndInRune(positions, 'Z') {
		if comPos >= len(com) {
			comPos = 0
		}
		for i, pos := range positions {
			if com[comPos] == 'L' {
				positions[i] = nds[pos].Left
			}
			if com[comPos] == 'R' {
				positions[i] = nds[pos].Right
			}
		}
		count = count + 1
		comPos = comPos + 1
	}
	return count
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	com, nds := ParseCommandsAndNodes(input)

	res := CountLengthThroughNodes(com, nds)
	fmt.Printf("Pt1: %d\n", res)
	ghst := CountGhostLengthThroughNodes(com, nds)
	fmt.Printf("Pt2: %d\n", ghst)

}
