package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type ValueLocation struct {
	Value     int
	Locations []Location
}

type SymbolLocation struct {
	Symbol   string
	Location Location
}

type Location struct {
	X int
	Y int
}

func (lc Location) IsSame(other Location) bool {
	return lc.X == other.X && lc.Y == other.Y
}

func (lc Location) IsAdjacent(other Location) bool {
	adjacents := []Location{
		{X: other.X - 1, Y: other.Y - 1},
		{X: other.X + 1, Y: other.Y + 1},
		{X: other.X - 1, Y: other.Y + 1},
		{X: other.X + 1, Y: other.Y - 1},
		{X: other.X - 1, Y: other.Y},
		{X: other.X, Y: other.Y - 1},
		{X: other.X + 1, Y: other.Y},
		{X: other.X, Y: other.Y + 1},
	}
	for _, adj := range adjacents {
		if adj.IsSame(lc) {
			return true
		}
	}
	return false
}

func (val ValueLocation) AdjacentToSymbol(symbols []SymbolLocation) bool {
	for _, vl := range val.Locations {
		for _, sbl := range symbols {
			if vl.IsAdjacent(sbl.Location) {
				return true
			}
		}
	}
	return false
}

func ParseGridToEntities(data string, spacerSymbol rune) ([]ValueLocation, []SymbolLocation) {
	lines := strings.Split(data, "\n")

	values := []ValueLocation{}
	symbols := []SymbolLocation{}

	numbuff := ""
	numbuffloc := []Location{}

	for y, line := range lines {
		for x, char := range line {
			if (x == 0 || !unicode.IsNumber(char)) && len(numbuff) > 0 {
				// "purge" the buffer
				val, _ := strconv.Atoi(numbuff)
				values = append(values, ValueLocation{
					Value:     val,
					Locations: numbuffloc,
				})
				numbuff = ""
				numbuffloc = []Location{}
			}
			if char == spacerSymbol {
				continue
			}
			if unicode.IsNumber(char) {
				numbuff = numbuff + string(char)
				numbuffloc = append(numbuffloc, Location{X: x, Y: y})
				continue
			}
			symbols = append(symbols, SymbolLocation{
				Symbol:   string(char),
				Location: Location{X: x, Y: y},
			})
		}
	}

	return values, symbols
}

func GetSumAdjacentToSymbols(values []ValueLocation, symbols []SymbolLocation) int {
	cml := 0
	for _, value := range values {
		if value.AdjacentToSymbol(symbols) {
			cml = cml + value.Value
		}
	}
	return cml
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)

	val, sym := ParseGridToEntities(stringput, '.')

	fmt.Println("Result: ", GetSumAdjacentToSymbols(val, sym))
}
