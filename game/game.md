package game

import "github.com/agnaite/1-4-24/gameplay"

func receiveTurn(game, turn) {
	// updates game

	// make sure that the user who sent the turn is the correct user
	// if the user doesn't qualify, make sure that they bet another $1
}

func sendGameState() {
	// dump all the game data to JSON
}


type Game struct {
	Gameplays []gameplay.Gameplay
	State
}

/*

for gp in Game.Gameplays:
	gp.Play()
	Game.RecordScore()

*/

type Match struct {
	gameplay_player Map[Player]Gameplay 
	pot int // pot of money, default =0
}

type Player struct {
	name string
	balance int
}

/*

# start the game


match = Match()

default_bet = 1

# have each player play the game
for player in players:

	gameplay = Gameplay.new()

	Match[player] = gameplay

	match.pot += default_bet
	player.balance -= default_bet

	// send and receive a bunch of JSON
	gameplay.Play()

# check and see who won

*/
