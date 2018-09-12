package gamestate

import (
	"testing"
)

func TestGameStateInitiation(t *testing.T) {

	// TODO: initializer should take some parameters:
	//
	// - list of current players
	// - bet size
	//
	players := []*Player{&Player{id: 1}, &Player{id: 2}}
	g := New(players, 1)

	if want, got := 2, len(g.Players); want != got {
		t.Fatalf("want %d players, got %d players", want, got)
	}

	if want, got := players[0], g.currentPlayer; want != got {
		t.Fatalf("currentPlayer should be first player %v. Got %v", want, got)
	}
}

func TestProcessAction(t *testing.T) {
	//g := New()

	// this will initialize diceRolled with 6 random integers:
	// g.processAction({ Player, "action": "roll" })

	// sets diceKept to [1, 4]
	// g.processAction({ Player, "action": "keep", "selection": [1, 4] })

	// this will initialize diceRolled with 6 random integers:
	// g.processAction({ Player, "action": "roll" })
	// (initializes remaining dice in diceRolled to new random integers. 4 in total)

	// g.processAction({ Player, "action": "keep", "selection": [6] })
	// g.processAction({ Player, "action": "roll" })

	// g.processAction({ Player, "action": "keep", "selection": [5] })
	// g.processAction({ Player, "action": "roll" })

	// g.processAction({ Player, "action": "keep", "selection": [4] })
	// g.processAction({ Player, "action": "roll" })

	// g.processAction({ Player, "action": "keep", "selection": [3] })
	// g.processAction({ Player, "action": "roll" })

	// - if the player tries to keep dice that weren't in diceRolled, there
	// should be an error

	// at this point the turn has ended and g.currentPlayer should be set to the
	// next player

	// - if player 1 sends another processAction, there should be an error

}
