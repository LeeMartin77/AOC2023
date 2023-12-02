package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BlockSet struct {
	R int
	G int
	B int
}

type Game struct {
	Id     int
	Rounds []BlockSet
}

func parseGameFromString(data string) Game {
	prts := strings.Split(data, ":")
	gametag := prts[0]
	rounds := strings.Split(prts[1], ";")
	gameId, _ := strconv.Atoi(strings.Split(gametag, " ")[1])
	blockSets := []BlockSet{}
	for _, rnd := range rounds {
		blocks := strings.Split(rnd, ",")
		blkst := BlockSet{
			R: 0,
			G: 0,
			B: 0,
		}
		for _, blk := range blocks {
			bits := strings.Split(blk[1:], " ")
			number, _ := strconv.Atoi(bits[0])
			if bits[1] == "red" {
				blkst.R = number
			}
			if bits[1] == "blue" {
				blkst.B = number
			}
			if bits[1] == "green" {
				blkst.G = number
			}
		}
		blockSets = append(blockSets, blkst)
	}
	return Game{
		Id:     gameId,
		Rounds: blockSets,
	}
}

func (game Game) isGamePossible(prediction BlockSet) bool {
	for _, rnd := range game.Rounds {
		if rnd.B > prediction.B || rnd.G > prediction.G || rnd.R > prediction.R {
			return false
		}
	}
	return true
}

func getTotalOfPossibleGames(data string, prediction BlockSet) int {
	lines := strings.Split(data, "\n")
	games := []Game{}
	for _, line := range lines {
		games = append(games, parseGameFromString(line))
	}
	cuml := 0
	for _, game := range games {
		if game.isGamePossible(prediction) {
			cuml = cuml + game.Id
		}
	}
	return cuml
}

func main() {
	buf, _ := os.ReadFile("data.txt")
	fmt.Printf("Result: %v", getTotalOfPossibleGames(string(buf), BlockSet{R: 12, G: 13, B: 14}))
}
