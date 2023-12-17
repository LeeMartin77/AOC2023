package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func DrawMap(mp [][]int, hist []Coordinate) {
	for y, _ := range mp[0] {
	xchr:
		for x, _ := range mp {
			for _, crd := range hist {
				if crd.X == x && crd.Y == y {
					fmt.Printf("#")
					continue xchr
				}
			}
			fmt.Printf(".")
		}
		fmt.Printf("\n")
	}
}

func ParseMap(input string) [][]int {
	mp := [][]int{}
	for y, ln := range strings.Split(input, "\n") {
		for x, chr := range ln {
			if y == 0 {
				mp = append(mp, []int{})
			}
			num, _ := strconv.Atoi(string(chr))
			mp[x] = append(mp[x], num)
		}
	}
	return mp
}

func (cord Coordinate) Equals(other Coordinate) bool {
	return cord.X == other.X && cord.Y == other.Y
}

type PriorityCord struct {
	Priority int
	Position Coordinate
}

func (crd Coordinate) Key() string {
	return fmt.Sprintf("%d:%d", crd.X, crd.Y)
}

func GetPossibleMoves(mp [][]int, cord Coordinate, fromMap map[string]Coordinate, maxStraight int) []PriorityCord {
	neighbours := []PriorityCord{}
	limitX := len(mp) - 1
	limitY := len(mp[0]) - 1
	if cord.X > 0 {
		new := PriorityCord{
			Priority: mp[cord.X-1][cord.Y],
			Position: Coordinate{
				X: cord.X - 1,
				Y: cord.Y,
			}}
		if !ExceedsMaxStraight(new.Position, cord, fromMap, maxStraight) {
			neighbours = append(neighbours, new)
		}
	}
	if cord.Y > 0 {
		new := PriorityCord{
			Priority: mp[cord.X][cord.Y-1],
			Position: Coordinate{
				X: cord.X,
				Y: cord.Y - 1,
			}}
		if !ExceedsMaxStraight(new.Position, cord, fromMap, maxStraight) {
			neighbours = append(neighbours, new)
		}
	}
	if cord.X < limitX {
		new := PriorityCord{
			Priority: mp[cord.X+1][cord.Y],
			Position: Coordinate{
				X: cord.X + 1,
				Y: cord.Y,
			}}
		if !ExceedsMaxStraight(new.Position, cord, fromMap, maxStraight) {
			neighbours = append(neighbours, new)
		}
	}
	if cord.Y < limitY {
		new := PriorityCord{
			Priority: mp[cord.X][cord.Y+1],
			Position: Coordinate{
				X: cord.X,
				Y: cord.Y + 1,
			}}
		if !ExceedsMaxStraight(new.Position, cord, fromMap, maxStraight) {
			neighbours = append(neighbours, new)
		}
	}

	sort.SliceStable(neighbours, func(i, j int) bool {
		return neighbours[i].Priority < neighbours[j].Priority
	})
	return neighbours
}

func ExceedsMaxStraight(next Coordinate, current Coordinate, fromMap map[string]Coordinate, maxStraight int) bool {
	previous := []Coordinate{}
	prev, ok := fromMap[current.Key()]
	for ok && len(previous) < maxStraight+1 {
		previous = append(previous, prev)
		prev, ok = fromMap[prev.Key()]
	}

	if len(previous) < maxStraight {
		return false
	}

	fewMovesAgo := previous[2]

	yStraight := current.X == fewMovesAgo.X && (math.Abs(float64(next.Y-fewMovesAgo.Y))) >= float64(maxStraight-1)
	xStraight := current.Y == fewMovesAgo.Y && (math.Abs(float64(next.X-fewMovesAgo.X))) >= float64(maxStraight-1)

	return (yStraight && current.X == next.X) || (xStraight && current.Y == next.Y)
}

func PathFromTo(mp [][]int, from Coordinate, to Coordinate, maxStraight int) (int, []Coordinate) {
	// Work out paths that are less than threshold based on rules
	// we don't have heat from starting cord
	frontier := []PriorityCord{}

	frontier = append(frontier, PriorityCord{0, from})

	cameFrom := map[string]Coordinate{}
	costSoFar := map[string]int{}
	//cameFrom[from.Key()] = nil
	costSoFar[from.Key()] = 0

	var winner PriorityCord
	var current PriorityCord
	for len(frontier) > 0 {
		// get top priority
		// sort
		sort.SliceStable(frontier, func(i, j int) bool {
			return frontier[i].Priority < frontier[j].Priority
		})
		// pop
		current, frontier = frontier[0], frontier[1:]
		if current.Position.Equals(to) {
			winner = current
			continue
		}
		for _, neighbour := range GetPossibleMoves(mp, current.Position, cameFrom, maxStraight) {
			newCost := costSoFar[current.Position.Key()] + neighbour.Priority
			soFar, ok := costSoFar[neighbour.Position.Key()]
			if !ok || newCost < soFar {
				costSoFar[neighbour.Position.Key()] = newCost
				frontier = append(frontier, PriorityCord{
					Priority: newCost,
					Position: neighbour.Position,
				})
				cameFrom[neighbour.Position.Key()] = current.Position
			}
		}
	}

	history := []Coordinate{winner.Position}

	pos, ok := cameFrom[winner.Position.Key()]
	history = append(history, pos)

	for ok {
		pos, ok = cameFrom[pos.Key()]
		if ok {
			history = append(history, pos)
		}
	}

	return winner.Priority, history
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	stringput := string(buf)

	res := ParseMap(stringput)
	bottom := len(res[0]) - 1
	right := len(res) - 1
	minHeatLoss, history := PathFromTo(res, Coordinate{0, 0}, Coordinate{right, bottom}, 3)
	DrawMap(res, history)
	fmt.Printf("Result: %v\n", minHeatLoss)
}
