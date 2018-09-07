package game

import (
	"fmt"
	"github.com/audy/1-4-24/gameplay"
)

type Player struct {
	Name string
	Bet  int
}

type Game struct {
	GameplayPlayer map[*Player]*gameplay.Gameplay
	Pot            int
	BestScore      int
	WinningPlayer  *Player
	Tie            bool
	Round          int
}

func (g *Game) Start() error {
	g.Round = g.Round + 1

	if g.Round > 1 {
		g.BestScore = 0
		g.Tie = false
	}

	fmt.Printf("Game round: %d\n", g.Round)

	for p, _ := range g.GameplayPlayer {
		p.Bet = p.Bet + 1
		g.Pot = g.Pot + 1

		if g.Round > 1 {
			g.GameplayPlayer[p] = gameplay.New()
		}
	}

	fmt.Printf("Pot is $%d\n\n", g.Pot)

	for p, gp := range g.GameplayPlayer {
		fmt.Printf("It's %v's turn!\n\n", p.Name)

		for {
			err := gp.Play()
			if err != nil && err.Error() == "finished" {
				fmt.Printf("%v's turn is done.\n\n", p.Name)

				if !gp.Qualified {
					p.Bet = p.Bet + 1
					g.Pot = g.Pot + 1
					fmt.Printf("Player %v adds $1 to the pot, for a total of $%d\n", p.Name, g.Pot)
				}
				if gp.Qualified {
					if gp.Score() > g.BestScore {
						g.BestScore = gp.Score()
						g.WinningPlayer = p
						g.Tie = false
					} else if gp.Score() == g.BestScore {
						g.Tie = true
					}
				}
				break
			}
			if err != nil {
				return err
			}
		}
	}

	if g.Tie {
		fmt.Println("\nIt's a tie! We keep playing.")
		err := g.Start()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("\nGame is over")
		fmt.Printf("Player %v wins $%d\n\n", g.WinningPlayer.Name, g.Pot)

		for p, _ := range g.GameplayPlayer {
			if p != g.WinningPlayer {
				fmt.Printf("Player %v owes Player %v $%d.\n", p.Name, g.WinningPlayer.Name, p.Bet)
			}
		}
	}
	return nil
}
