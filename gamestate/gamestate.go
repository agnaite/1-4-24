package gamestate

type Player struct {
	id   int
	dice *Dice
}

type GameState struct {
	Players       []*Player
	BetSize       int
	currentPlayer *Player
}

func New(players []*Player, betSize int) *GameState {
	return &GameState{
		Players:       players,
		BetSize:       betSize,
		currentPlayer: players[0],
	}
}
