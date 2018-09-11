package gamestate

type Player struct {
	id int
}

type GameState struct {
	currentPlayer Player
	turns         []Turn
}

type Turn struct {
	diceKept []int
	player   Player
}

func New() *GameState {
	return &GameState{
		currentPlayer: Player{
			id: 1,
		},
		turns: []Turn{},
	}
}
