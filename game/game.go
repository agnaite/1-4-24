package game

import (
	"fmt"

	"github.com/agnaite/1-4-24/gameplay"
)

type Player struct {
	Name    string
	Balance int
}

type Game struct {
	GameplayPlayer map[Player]*gameplay.Gameplay
	Pot            int
}

func (g *Game) Start() error {

	for p, gp := range g.GameplayPlayer {
		fmt.Printf("It's %v's turn!\n", p.Name)

		for {
			err := gp.Play()
			if err != nil && err.Error() == "finished" {
				fmt.Printf("%v's turn is done.\n", p.Name)
				break
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}
