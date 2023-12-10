package main

import (
	"fmt"
	"os"
	"strings"
)

type Tile struct {
	X    int
	Y    int
	Rune rune
}

type InputOutput struct {
	X, Y int
}

func (tl Tile) ConnectsTo() []InputOutput {
	// we return +/-1 of the directions of travel
	// so, for example, []InputOutput{{-1, 0}, {1, 0}}
	// is left-right on the X axis from {0,0}
	res := []InputOutput{}
	if tl.Rune == '|' {
		res = append(res, []InputOutput{{tl.X, tl.Y - 1}, {tl.X, tl.Y + 1}}...)
	}
	if tl.Rune == '-' {
		res = append(res, []InputOutput{{tl.X - 1, tl.Y}, {tl.X + 1, tl.Y}}...)
	}
	if tl.Rune == 'L' {
		res = append(res, []InputOutput{{tl.X, tl.Y - 1}, {tl.X + 1, tl.Y}}...)
	}
	if tl.Rune == 'J' {
		res = append(res, []InputOutput{{tl.X, tl.Y - 1}, {tl.X - 1, tl.Y}}...)
	}
	if tl.Rune == '7' {
		res = append(res, []InputOutput{{tl.X, tl.Y + 1}, {tl.X - 1, tl.Y}}...)
	}
	if tl.Rune == 'F' {
		res = append(res, []InputOutput{{tl.X, tl.Y + 1}, {tl.X + 1, tl.Y}}...)
	}
	return res
}

func ParseTiles(input string) (InputOutput, [][]Tile) {
	tiles := [][]Tile{}
	startPos := InputOutput{}
	for y, ln := range strings.Split(input, "\n") {

		for x, chr := range ln {
			tile := Tile{Rune: chr, X: x, Y: y}
			if y == 0 {
				tiles = append(tiles, []Tile{tile})
			} else {
				tiles[x] = append(tiles[x], tile)
			}
			if chr == 'S' {
				startPos = InputOutput{x, y}
			}
		}
	}
	return startPos, tiles
}

func includes(inputs []InputOutput, mtch InputOutput) bool {
	for _, inp := range inputs {
		if inp.X == mtch.X && inp.Y == mtch.Y {
			return true
		}
	}
	return false
}

func GetShuttlesAndRoutes(start InputOutput, tiles [][]Tile) ([]InputOutput, [][]InputOutput) {
	shuttles := []InputOutput{}
	routes := [][]InputOutput{}
	// see which tiles around start can connect to start
	for _, x := range []int{-1, 1, 0} {
		for _, y := range []int{-1, 1, 0} {
			if x == 0 && y == 0 {
				continue
			}
			//ignore diagonals
			if x != 0 && y != 0 {
				continue
			}
			neighbourConnection := tiles[x+start.X][y+start.Y].ConnectsTo()
			if includes(neighbourConnection, start) {
				shuttles = append(shuttles, InputOutput{x + start.X, y + start.Y})
				routes = append(routes, []InputOutput{{start.X, start.Y}})
			}
		}
	}
	return shuttles, routes
}

func RecurrAlongRoute(tiles [][]Tile, start InputOutput, shuttle InputOutput, route []InputOutput) (bool, []InputOutput) {

	if shuttle.X == start.X && shuttle.Y == start.Y {
		return true, route
	}
	lastLocation := route[len(route)-1]
	currentConnections := tiles[shuttle.X][shuttle.Y].ConnectsTo()
	nextConnection := lastLocation

	for _, con := range currentConnections {
		if con.X != nextConnection.X || con.Y != nextConnection.Y {
			nextConnection = con
			break
		}
	}
	if nextConnection.X == lastLocation.X && nextConnection.Y == lastLocation.Y {
		//dead end
		return false, route
	}
	route = append(route, InputOutput{shuttle.X, shuttle.Y})

	return RecurrAlongRoute(tiles, start, InputOutput{nextConnection.X, nextConnection.Y}, route)
}

func GetLoop(start InputOutput, tiles [][]Tile) ([]InputOutput, error) {
	shuttles, routes := GetShuttlesAndRoutes(start, tiles)
	isLoop := []bool{}
	for i, shuttle := range shuttles {
		isLoop = append(isLoop, false)

		isL, route := RecurrAlongRoute(tiles, start, shuttle, routes[i])
		isLoop[i] = isL
		routes[i] = route
	}
	for i, success := range isLoop {
		if success {
			return routes[i], nil
		}
	}
	return []InputOutput{}, fmt.Errorf("Couldn't find a route")
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	startPos, tiles := ParseTiles(input)
	loop, err := GetLoop(startPos, tiles)
	if err != nil {
		panic("OH NO")
	}
	fmt.Printf("Result: %v\n", len(loop)/2)
}
