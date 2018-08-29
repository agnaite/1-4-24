package main

import (
	"fmt"

	"github.com/agnaite/1-4-24/game"
	"github.com/agnaite/1-4-24/gameplay"
)

func main() {
	players := []string{"Agne", "Sam", "Austin"}

	playerMap := make(map[*game.Player]*gameplay.Gameplay)

	for _, player := range players {
		playerMap[&game.Player{Name: player}] = gameplay.New()
	}

	game := game.Game{
		GameplayPlayer: playerMap,
	}

	if err := game.Start(); err != nil {
		fmt.Println(err)
	}
}
