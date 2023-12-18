package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func myTestFunction(input string) int {
	return 1
}

type TrenchMap map[int]map[int]string

// trenches, minX, maxX, minY, maxY
func DigTrenches(input string) (TrenchMap, int, int, int, int) {
	positionX := 0
	positionY := 0

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	trenches := map[int]map[int]string{}

	trenches[0] = map[int]string{0: "(#000000)"}

	for _, command := range strings.Split(input, "\n") {
		direction := strings.Split(command, " ")[0]
		numberOfMoves, _ := strconv.Atoi(strings.Split(command, " ")[1])
		color := strings.Split(command, " ")[2]
		// R
		for numberOfMoves > 0 {
			numberOfMoves = numberOfMoves - 1
			if direction == "R" {
				positionX = positionX + 1
				_, ok := trenches[positionX]
				if !ok {
					trenches[positionX] = map[int]string{}
				}
				trenches[positionX][positionY] = color
			}
			// L
			if direction == "L" {
				positionX = positionX - 1
				_, ok := trenches[positionX]
				if !ok {
					trenches[positionX] = map[int]string{}
				}
				trenches[positionX][positionY] = color
			}
			// D
			if direction == "D" {
				positionY = positionY + 1
				trenches[positionX][positionY] = color
			}
			// U
			if direction == "U" {
				positionY = positionY - 1
				trenches[positionX][positionY] = color
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

func (trnchs TrenchMap) FloodFill(posX int, posY int, fill string) {
	val := trnchs[posX][posY]
	if val != "" {
		return
	}
	trnchs[posX][posY] = fill
	trnchs.FloodFill(posX+1, posY, fill)
	trnchs.FloodFill(posX-1, posY, fill)
	trnchs.FloodFill(posX, posY+1, fill)
	trnchs.FloodFill(posX, posY-1, fill)
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

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	res, _, _, _, _ := DigTrenches(input)

	filled := FillTrenches(res, "WOLOLOLOL")

	countFilled := filled.CountDug()
	fmt.Printf("Part one: %d\n", countFilled)
}
