package main

import (
	"fmt"
	"sort"
)

type GameLog struct {
	Name  string
	Score int
}

type Player struct {
	Name  string
	Score int
	Level string
}

func main() {
	gameLog := []GameLog{
		{Name: "John", Score: 10},
		{Name: "Alex", Score: 7},
		{Name: "John", Score: 5},
		{Name: "Julia", Score: 15},
		{Name: "Julia", Score: 2},
		{Name: "Saimon", Score: 3},
	}

	players := []Player{}

	m := make(map[string]int)
	for _, log := range gameLog {
		m[log.Name] += log.Score
	}

	for key, val := range m {
		status := ""

		if val >= 11 {
			status = "senior"
		} else if val >= 7 && val <= 10 {
			status = "middle"
		} else if val >= 0 && val <= 6 {
			status = "junior"
		}

		player := Player{
			Name:  key,
			Score: val,
			Level: status,
		}

		players = append(players, player)
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Score > players[j].Score
	})

	for _, val := range players {
		fmt.Printf("%s - %s [%d]\n",val.Name,val.Level,val.Score)
	}
	
}


