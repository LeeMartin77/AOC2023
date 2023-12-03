package main

import (
	"fmt"
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

func myTestFunction() int {
	return 1
}

func ParseGridToEntities(data string, spacerSymbol rune) ([]ValueLocation, []SymbolLocation) {
	lines := strings.Split(data, "\n")

	values := []ValueLocation{}
	symbols := []SymbolLocation{}

	numbuff := ""
	numbuffloc := []Location{}

	for y, line := range lines {
		if y == 0 && len(numbuff) > 0 {
			// "purge" the buffer
			val, _ := strconv.Atoi(numbuff)
			values = append(values, ValueLocation{
				Value:     val,
				Locations: numbuffloc,
			})
			numbuff = ""
			numbuffloc = []Location{}
		}
		for x, char := range line {
			if char == spacerSymbol && len(numbuff) > 0 {
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

func main() {
	//buf, _ := os.ReadFile("data.txt")
	//stringput := string(buf)

	fmt.Printf("Result: %v", myTestFunction())
}
