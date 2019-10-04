package manager

import (
	"math"
	"math/rand"
	"time"
)

/*
PlayerList : Slice for players
*/
type PlayerList []Player

/*
Player : struct to stock player data
	ID    int     `jons:"id"`
*/
type Player struct {
	ID    int     `json:"id"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Color Color   `json:"color"`
	time  time.Time
}

/*
Color : color type
*/
type Color [3]int

var (
	id      int
	players PlayerList
)

/*
CreatePlayer : create a new player
*/
func CreatePlayer() Player {
	player := Player{
		id,
		float64(rand.Intn(100)),
		float64(rand.Intn(100)),
		Color{
			rand.Intn(255),
			rand.Intn(255),
			rand.Intn(255),
		},
		time.Now(),
	}
	players = append(players, player)
	id++
	return player
}

/*
Refresh : refresh player in list and get all other players
*/
func Refresh(id int, x, y float64) (PlayerList, int) {
	var others PlayerList
	for i := range players {
		if players[i].ID == id {
			players[i].X = x
			players[i].Y = y
			players[i].time = time.Now()
		} else {
			others = append(others, players[i])
		}
	}

	go func() {
		var refreshedPlayer PlayerList
		for i := range players {
			if float64(time.Now().Sub(players[i].time))*math.Pow(10, -9) < 7 {
				refreshedPlayer = append(refreshedPlayer, players[i])
			}
		}
		players = refreshedPlayer
	}()

	return others, len(others)
}
