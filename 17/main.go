package main

import (
	"fmt"
	"math"
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

// goes straight through, gets a "minimum"
func GetDefault(mp [][]int, from Coordinate, to Coordinate) int {
	cuml := 0
	for from.X != to.X || from.Y != to.Y {
		cuml = cuml + mp[from.X][from.Y]
		if math.Abs(float64(from.X-to.X)) > math.Abs(float64(from.Y-to.Y)) {
			// move on horizontal
			if from.X > to.X {
				from.X = from.X - 1
			} else {
				from.X = from.X + 1
			}
		} else {
			// move on vertical
			if from.Y > to.Y {
				from.Y = from.Y - 1
			} else {
				from.Y = from.Y + 1
			}
		}
	}
	cuml = cuml + mp[from.X][from.Y]
	return cuml
}

type Path struct {
	History  []Coordinate
	HeatLoss int
}

func (cord Coordinate) Equals(other Coordinate) bool {
	return cord.X == other.X && cord.Y == other.Y
}

func RecurrThroughPaths(mp [][]int, paths []Path, target Coordinate, threshold int, maxStraight int) Path {
	continuingPaths := []Path{}
	stillWalking := false
	limitX := len(mp) - 1
	limitY := len(mp[0]) - 1
	for _, path := range paths {
		if path.HeatLoss > threshold {
			// too cold
			continue
		}
		curPos := path.History[len(path.History)-1]
		if curPos.Equals(target) {
			// still in the running and at end
			if path.HeatLoss < threshold {
				threshold = path.HeatLoss
			}
			continuingPaths = append(continuingPaths, path)
			continue
		}
		stillWalking = true
		if len(path.History) < 2 {
			// if it's the first move, we can move in any possible direction
			if curPos.X > 0 {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X-1][curPos.Y],
					History:  append(path.History, Coordinate{curPos.X - 1, curPos.Y}),
				})
			}
			if curPos.Y > 0 {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X][curPos.Y-1],
					History:  append(path.History, Coordinate{curPos.X, curPos.Y - 1}),
				})
			}
			if curPos.X < limitX {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X+1][curPos.Y],
					History:  append(path.History, Coordinate{curPos.X + 1, curPos.Y}),
				})
			}
			if curPos.Y < limitY {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X][curPos.Y+1],
					History:  append(path.History, Coordinate{curPos.X, curPos.Y + 1}),
				})
			}
		} else if len(path.History) < maxStraight+2 {
			// if it's the first few moves, we can move in any possible direction but backward
			lastPos := path.History[len(path.History)-2]
			if curPos.X > 0 && lastPos.X != curPos.X-1 {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X-1][curPos.Y],
					History:  append(path.History, Coordinate{curPos.X - 1, curPos.Y}),
				})
			}
			if curPos.Y > 0 && lastPos.Y != curPos.Y-1 {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X][curPos.Y-1],
					History:  append(path.History, Coordinate{curPos.X, curPos.Y - 1}),
				})
			}
			if curPos.X < limitX && lastPos.X != curPos.X+1 {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X+1][curPos.Y],
					History:  append(path.History, Coordinate{curPos.X + 1, curPos.Y}),
				})
			}
			if curPos.Y < limitY && lastPos.Y != curPos.Y+1 {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X][curPos.Y+1],
					History:  append(path.History, Coordinate{curPos.X, curPos.Y + 1}),
				})
			}
		} else {
			// we generate paths for all possible paths from this path
			lastPos := path.History[len(path.History)-2]
			manyMovesBack := path.History[len(path.History)-(maxStraight+1)]
			if curPos.X > 0 && lastPos.X != curPos.X-1 && !(manyMovesBack.X == curPos.X-maxStraight && manyMovesBack.Y == curPos.Y) {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X-1][curPos.Y],
					History:  append(path.History, Coordinate{curPos.X - 1, curPos.Y}),
				})
			}
			if curPos.Y > 0 && lastPos.Y != curPos.Y-1 && !(manyMovesBack.Y == curPos.Y-maxStraight && manyMovesBack.X == curPos.X) {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X][curPos.Y-1],
					History:  append(path.History, Coordinate{curPos.X, curPos.Y - 1}),
				})
			}
			if curPos.X < limitX && lastPos.X != curPos.X+1 && !(manyMovesBack.X == curPos.X+maxStraight && manyMovesBack.Y == curPos.Y) {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X+1][curPos.Y],
					History:  append(path.History, Coordinate{curPos.X + 1, curPos.Y}),
				})
			}
			if curPos.Y < limitY && lastPos.Y != curPos.Y+1 && !(manyMovesBack.Y == curPos.Y+maxStraight && manyMovesBack.X == curPos.X) {
				continuingPaths = append(continuingPaths, Path{
					HeatLoss: path.HeatLoss + mp[curPos.X][curPos.Y+1],
					History:  append(path.History, Coordinate{curPos.X, curPos.Y + 1}),
				})
			}
		}
	}
	fmt.Println(len(continuingPaths))
	if stillWalking {
		return RecurrThroughPaths(mp, continuingPaths, target, threshold, maxStraight)
	}
	// the paths will all have the same heatloss
	return continuingPaths[0]
}

func PathFromTo(mp [][]int, from Coordinate, to Coordinate, maxStraight int) (int, []Coordinate) {
	// Calc a "threshold" heat loss, of an unoptimised diagonal
	threshold := GetDefault(mp, from, to)
	// Work out paths that are less than threshold based on rules
	// we don't have heat from starting cord
	final := RecurrThroughPaths(mp, []Path{{HeatLoss: 0, History: []Coordinate{{from.X, from.Y}}}}, to, threshold, maxStraight)
	// get least heat loss path (maybe shortest if tie)
	return final.HeatLoss, final.History
}

func main() {
	// buf, _ := os.ReadFile("data.txt")
	// stringput := string(buf)
	// fmt.Printf("Result: %v", myTestFunction(stringput))
}
