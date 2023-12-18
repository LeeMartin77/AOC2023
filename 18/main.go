package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func myTestFunction(input string) int {
	return 1
}

type TrenchMap map[int64]map[int64]bool

// trenches, minX, maxX, minY, maxY
func DigTrenches(input string) (TrenchMap, int64, int64, int64, int64) {
	positionX := int64(0)
	positionY := int64(0)

	minX := int64(0)
	maxX := int64(0)
	minY := int64(0)
	maxY := int64(0)

	trenches := TrenchMap{}

	trenches[0] = map[int64]bool{0: true}

	for _, command := range strings.Split(input, "\n") {
		direction := strings.Split(command, " ")[0]
		numberOfMoves, _ := strconv.Atoi(strings.Split(command, " ")[1])
		//color := strings.Split(command, " ")[2]
		// R
		for numberOfMoves > 0 {
			numberOfMoves = numberOfMoves - 1
			if direction == "R" {
				positionX = positionX + 1
				_, ok := trenches[positionX]
				if !ok {
					trenches[positionX] = map[int64]bool{}
				}
				trenches[positionX][positionY] = true
			}
			// L
			if direction == "L" {
				positionX = positionX - 1
				_, ok := trenches[positionX]
				if !ok {
					trenches[positionX] = map[int64]bool{}
				}
				trenches[positionX][positionY] = true
			}
			// D
			if direction == "D" {
				positionY = positionY + 1
				trenches[positionX][positionY] = true
			}
			// U
			if direction == "U" {
				positionY = positionY - 1
				trenches[positionX][positionY] = true
			}

			if minX > positionX {
				minX = positionX
			}

			if maxX < positionX {
				maxX = positionX
			}

			if minY > positionY {
				minY = positionY
			}

			if maxY < positionY {
				maxY = positionY
			}
		}
	}

	return trenches, minX, maxX, minY, maxY
}

func DigTrenchesHex(input string) ([]Coord, int64) {
	positionX := int64(0)
	positionY := int64(0)

	trenches := []Coord{{0, 0}}

	trenchesDug := int64(0)

	for _, command := range strings.Split(input, "\n") {
		hexcode := strings.Split(command, " ")[2]
		hexcode = strings.TrimPrefix(hexcode, "(#")
		hexcode = strings.TrimSuffix(hexcode, ")")
		hexcodeDist := hexcode[:5]
		numberOfMoves, _ := strconv.ParseInt(hexcodeDist, 16, 64) //convert hexcode
		direction := hexcode[5:]
		// R
		trenchesDug = trenchesDug + numberOfMoves

		if direction == "0" {
			positionX = positionX + numberOfMoves
		}
		// L
		if direction == "2" {
			positionX = positionX - numberOfMoves
		}
		// D
		if direction == "1" {
			positionY = positionY + numberOfMoves
		}
		// U
		if direction == "3" {
			positionY = positionY - numberOfMoves
		}
		trenches = append(trenches, Coord{positionX, positionY})
	}

	return trenches, trenchesDug
}

type Coord struct{ X, Y int64 }

func (trnchs TrenchMap) FloodFill(posX int64, posY int64, fill string) {
	queue := []Coord{{posX, posY}}
	var current Coord
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		trnchs[current.X][current.Y] = true
		if !trnchs[current.X+1][current.Y] {
			queue = append(queue, Coord{current.X + 1, current.Y})
		}
		if !trnchs[current.X-1][current.Y] {
			queue = append(queue, Coord{current.X - 1, current.Y})
		}
		if !trnchs[current.X][current.Y+1] {
			queue = append(queue, Coord{current.X, current.Y + 1})
		}
		if !trnchs[current.X][current.Y-1] {
			queue = append(queue, Coord{current.X, current.Y - 1})
		}
	}
}

func FillTrenches(trenches TrenchMap, fill string) TrenchMap {
	trenches.FloodFill(1, 1, fill)
	return trenches
}

func (trnchs TrenchMap) CountDug() int {
	cuml := 0
	for _, col := range trnchs {
		for range col {
			cuml = cuml + 1
		}
	}
	return cuml
}

func ShoelaceArea(vertices []Coord) int64 {
	fmt.Println("Shoelacing")
	numberOfVertices := len(vertices)
	sum1 := int64(0)
	sum2 := int64(0)
	for i := 0; i < numberOfVertices-1; i = i + 1 {
		sum1 = sum1 + vertices[i].X*vertices[i+1].Y
		sum2 = sum2 + vertices[i].Y*vertices[i+1].X
	}
	sum1 = sum1 + vertices[numberOfVertices-1].X*vertices[0].Y
	sum2 = sum2 + vertices[0].X*vertices[numberOfVertices-1].Y
	return int64(math.Abs(float64(sum1-sum2)) / 2)
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	// res, _, _, _, _ := DigTrenches(input)

	// filled := FillTrenches(res, "WOLOLOLOL")

	// countFilled := filled.CountDug()
	// fmt.Printf("Part one: %d\n", countFilled)

	resVert, trenchlength := DigTrenchesHex(input)
	bigCountFilled := ShoelaceArea(resVert)
	fmt.Printf("Part two: %d\n", bigCountFilled+(trenchlength/2+1))
}
