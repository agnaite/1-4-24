package gamestate

type Player struct {
	id int
}

type Turn struct {
	diceKept []int
	player   Player
}

type GameState struct {
	currentPlayer Player
	turns         []Turn
}

func New() *GameState {
	return &GameState{
		currentPlayer: Player{
			id: 1,
		},
		turns: []Turn{},
	}
}
