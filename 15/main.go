package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label   string
	Stength int
}

func CalculateHash(input string) int {
	cuml := 0
	for _, ch := range input {
		cuml = cuml + int(ch)
		cuml = cuml * 17
		cuml = cuml % 256
	}
	return cuml
}

// returns boxes,
func PerformOperationOnBoxes(boxes map[int][]Lens, op string) map[int][]Lens {

	if strings.Index(op, "-") != -1 { // - : remove
		lbl := strings.TrimSuffix(op, "-")
		hsh := CalculateHash(lbl)
		indx := -1
		for i, lns := range boxes[hsh] {
			if lns.Label == lbl {
				indx = i
			}
		}
		if indx > -1 {
			boxes[hsh] = append(boxes[hsh][:indx], boxes[hsh][indx+1:]...)
		}
	} else { // = : add to end, or if already in box, replace strength
		prts := strings.Split(op, "=")
		lbl := prts[0]
		hsh := CalculateHash(lbl)
		strg, _ := strconv.Atoi(prts[1])

		indx := -1
		for i, lns := range boxes[hsh] {
			if lns.Label == lbl {
				indx = i
			}
		}
		if indx > -1 {
			boxes[hsh][indx].Stength = strg
		} else {
			boxes[hsh] = append(boxes[hsh], Lens{lbl, strg})
		}
	}
	return boxes
}

func CalculateValueFromBoxes(boxes map[int][]Lens) int {
	cuml := 0
	for bxnum, lenses := range boxes {
		for i, lns := range lenses {
			cuml = cuml + ((1 + bxnum) * (i + 1) * lns.Stength)
		}
	}
	return cuml
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	cuml := 0
	for _, com := range strings.Split(input, ",") {
		cuml = cuml + CalculateHash(com)
	}
	fmt.Printf("Result: %v\n", cuml)

	boxes := map[int][]Lens{}
	for _, com := range strings.Split(input, ",") {
		boxes = PerformOperationOnBoxes(boxes, com)
	}

	fmt.Printf("Result: %v\n", CalculateValueFromBoxes(boxes))
}
