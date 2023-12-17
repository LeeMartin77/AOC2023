package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
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

func GetNeighbours(mp [][]int, cord Coordinate) []PriorityCord {
	neighbours := []PriorityCord{}
	limitX := len(mp) - 1
	limitY := len(mp[0]) - 1
	if cord.X > 0 {
		neighbours = append(neighbours, PriorityCord{
			Priority: mp[cord.X-1][cord.Y],
			Position: Coordinate{
				X: cord.X - 1,
				Y: cord.Y,
			}})
	}
	if cord.Y > 0 {
		neighbours = append(neighbours, PriorityCord{
			Priority: mp[cord.X][cord.Y-1],
			Position: Coordinate{
				X: cord.X,
				Y: cord.Y - 1,
			}})
	}
	if cord.X < limitX {
		neighbours = append(neighbours, PriorityCord{
			Priority: mp[cord.X+1][cord.Y],
			Position: Coordinate{
				X: cord.X + 1,
				Y: cord.Y,
			}})
	}
	if cord.Y < limitY {
		neighbours = append(neighbours, PriorityCord{
			Priority: mp[cord.X][cord.Y+1],
			Position: Coordinate{
				X: cord.X,
				Y: cord.Y + 1,
			}})
	}
	return neighbours
}

func ExceedsMaxStraight(next Coordinate, current Coordinate, fromMap map[string]Coordinate, maxStraight int) bool {
	straightCount := 0
	xOffset := 0
	yOffset := 0
	for xOffset == 0 || yOffset == 0 {
		xOffset = xOffset + (next.X - current.X)
		yOffset = yOffset + (next.Y - current.Y)
		next = current
		newCur, ok := fromMap[current.Key()]
		if !ok {
			break
		}
		current = newCur
		straightCount = straightCount + 1
		if maxStraight+1 > straightCount {
			break
		}
	}
	return maxStraight+1 > straightCount
}

func PathFromTo(mp [][]int, from Coordinate, to Coordinate, maxStraight int) int {
	// Work out paths that are less than threshold based on rules
	// we don't have heat from starting cord
	frontier := []PriorityCord{}

	frontier = append(frontier, PriorityCord{0, from})

	cameFrom := map[string]Coordinate{}
	costSoFar := map[string]int{}
	//cameFrom[from.Key()] = nil
	costSoFar[from.Key()] = 0

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
			break
		}
		for _, neighbour := range GetNeighbours(mp, current.Position) {
			newCost := costSoFar[current.Position.Key()] + neighbour.Priority
			// need a condition for neighbour not being fourth straight line
			soFar, ok := costSoFar[neighbour.Position.Key()]
			if !ok || newCost < soFar || !ExceedsMaxStraight(neighbour.Position, current.Position, cameFrom, maxStraight) {
				costSoFar[neighbour.Position.Key()] = newCost
				frontier = append(frontier, PriorityCord{
					Priority: newCost,
					Position: neighbour.Position,
				})
				cameFrom[neighbour.Position.Key()] = current.Position
			}
		}
	}

	return current.Priority
}

func main() {
	// buf, _ := os.ReadFile("data.txt")
	// stringput := string(buf)
	// fmt.Printf("Result: %v", myTestFunction(stringput))
}
