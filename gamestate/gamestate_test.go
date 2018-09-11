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
	g := New()

	if want, got := 1, g.currentPlayer.id; want != got {
		t.Fatalf("want 0 currentPlayer, got %d", got)
	}

	// g.Players should return a slice of Player structs:
	//
	// g.Players = [ Player{id: 123} ]

	// g.Turns should return a slice of Turn structs:
	// g.Turns = [
	//   { player: Player{}, diceKept: []Dice }
	// ]
}

func TestProcessAction(t *testing.T) {
	g := New()

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
