package gamestate

type GameState struct {
	currentPlayer int
}

func New() *GameState {
	return &GameState{}
}
