package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Connection struct {
	From     string
	To       string
	Distance int
}

func cordToKey(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func keyToCord(key string) (int, int) {
	pts := strings.Split(key, ":")
	x, _ := strconv.Atoi(pts[0])
	y, _ := strconv.Atoi(pts[1])
	return x, y
}

func isNode(x, y int, chr rune, lines []string) bool {
	countDirections := 0
	if chr == '#' {
		return false
	}
	if lines[y][x-1] != '#' {
		countDirections = countDirections + 1
	}
	if lines[y][x+1] != '#' {
		countDirections = countDirections + 1
	}
	if lines[y-1][x] != '#' {
		countDirections = countDirections + 1
	}
	if lines[y+1][x] != '#' {
		countDirections = countDirections + 1
	}
	return countDirections > 2
}

func FollowPathToNode(path []string, lines []string, nodes map[string][]Connection) (string, int) {
	next := ""
	for true {
		last, cur := path[len(path)-2], path[len(path)-1]
		stsNext := next
		lastX, lastY := keyToCord(last)
		curx, cury := keyToCord(cur)
		_, ok := nodes[stsNext]
		if ok {
			// we have reached another node
			return next, len(path)
		}
		if (lines[cury][curx-1] == '.' || lines[cury][curx-1] == '<') && lastX != curx-1 {
			next = cordToKey(curx-1, cury)
		}
		if (lines[cury][curx+1] == '.' || lines[cury][curx+1] == '>') && lastX != curx+1 {
			next = cordToKey(curx+1, cury)
		}
		if (lines[cury-1][curx] == '.') && lastY != cury-1 {
			next = cordToKey(curx, cury-1)
		}
		if (lines[cury+1][curx] == '.' || lines[cury+1][curx] == 'v') && lastY != cury+1 {
			next = cordToKey(curx, cury+1)
		}
		if next == stsNext {
			// we have dead ended
			return "", 0
		}
		path = append(path, next)
	}
	return "", 0
}

func GetConnections(x, y int, lines []string, nodes map[string][]Connection) []Connection {
	connections := []Connection{}
	paths := [][]string{}
	if lines[y][x-1] != '#' && lines[y][x-1] != '>' {
		//
		paths = append(paths, []string{cordToKey(x, y), cordToKey(x-1, y)})
	}
	if lines[y][x+1] != '#' && lines[y][x+1] != '<' {
		//
		paths = append(paths, []string{cordToKey(x, y), cordToKey(x+1, y)})
	}
	if lines[y-1][x] != '#' && lines[y-1][x] != 'v' {
		//
		paths = append(paths, []string{cordToKey(x, y), cordToKey(x, y-1)})
	}
	if lines[y+1][x] != '#' {
		//
		paths = append(paths, []string{cordToKey(x, y), cordToKey(x, y+1)})
	}
	for _, pth := range paths {
		to, len := FollowPathToNode(pth, lines, nodes)
		if to != "" {
			connections = append(connections, Connection{
				From:     cordToKey(x, y),
				To:       to,
				Distance: len,
			})
		}
	}
	return connections
}

func ParseMap(input string) (map[string][]Connection, string, string) {
	lines := strings.Split(input, "\n")
	nodes := map[string][]Connection{}

	endCord := cordToKey(strings.Index(lines[len(lines)-1], "."), len(lines)-1)
	startCord := ""
	nodes[endCord] = []Connection{}
	for y, line := range lines {
		if y-1 < 0 {
			// it's our entry node line
			x := strings.Index(line, ".")
			// pathfind to next connection
			path := []string{cordToKey(x, y), cordToKey(x, y+1)}
			next := ""
			for true {
				last, cur := path[len(path)-2], path[len(path)-1]
				lastX, lastY := keyToCord(last)
				curx, cury := keyToCord(cur)
				countDirections := 0
				if (lines[cury][curx-1] == '.' || lines[cury][curx-1] == '<') && lastX != curx-1 {
					countDirections = countDirections + 1
					next = cordToKey(curx-1, cury)
				}
				if (lines[cury][curx+1] == '.' || lines[cury][curx+1] == '>') && lastX != curx+1 {
					countDirections = countDirections + 1
					next = cordToKey(curx+1, cury)
				}
				if (lines[cury-1][curx] == '.') && lastY != cury-1 {
					countDirections = countDirections + 1
					next = cordToKey(curx, cury-1)
				}
				if (lines[cury+1][curx] == '.' || lines[cury+1][curx] == 'v') && lastY != cury+1 {
					countDirections = countDirections + 1
					next = cordToKey(curx, cury+1)
				}
				if countDirections > 1 {
					//weve reached first other node
					break
				}
				path = append(path, next)
			}
			nodes[cordToKey(x, y)] = []Connection{
				{
					From:     cordToKey(x, y),
					To:       path[len(path)-1],
					Distance: len(path) + 1,
				},
			}
			startCord = cordToKey(x, y)
			continue
		}
		if y+1 >= len(lines) {
			continue
		}
		// check if it's a node
		// if it's a node, start following to destinations
		// remember it's possible a node ends up only having one forward path
		nds := []string{}
		for x, chr := range line {
			if x == 0 || x == len(line)-1 {
				continue
			}
			if isNode(x, y, chr, lines) {
				nds = append(nds, cordToKey(x, y))
				nodes[cordToKey(x, y)] = []Connection{}
			}
		}
		for _, nd := range nds {
			x, y := keyToCord(nd)
			connections := GetConnections(x, y, lines, nodes)
			nodes[cordToKey(x, y)] = connections
		}
	}
	return nodes, startCord, endCord
}

func GetPossibleSeinicPaths(runningPaths [][]Connection, endNode string, nodes map[string][]Connection) [][]Connection {
	//
	stillExploring := false
	nextPaths := [][]Connection{}
	for _, pth := range runningPaths {
		if pth[len(pth)-1].To == endNode {
			nextPaths = append(nextPaths, pth)
			continue
		}
		next := nodes[pth[len(pth)-1].To]
	poss:
		for _, nxt := range next {
			for _, nd := range pth {
				if nxt.To == nd.From {
					// loops/goes on self
					continue poss
				}
			}
			stillExploring = true
			nextPaths = append(nextPaths, append(pth, nxt))
		}
	}
	if !stillExploring {
		return nextPaths
	}
	return GetPossibleSeinicPaths(nextPaths, endNode, nodes)
}

func TotalPathLens(paths [][]Connection) []int {
	lens := []int{}
	for _, pth := range paths {
		cuml := 0
		for _, con := range pth {
			cuml = cuml + con.Distance
		}
		lens = append(lens, cuml)
	}
	return lens
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	input := string(buf)
	res, strt, end := ParseMap(input)
	pths := GetPossibleSeinicPaths([][]Connection{res[strt]}, end, res)
	lens := TotalPathLens(pths)
	sort.SliceStable(lens, func(i, j int) bool {
		return lens[i] > lens[j]
	})

	fmt.Printf("Part one: %d\n", lens[0]+1)
}
