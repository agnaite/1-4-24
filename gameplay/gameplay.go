package gameplay

import (
	"errors"
	"sync"
)

const (
	StateRolling  State = "rolling"
	StateKeeping        = "keeping"
	StateBetting        = "betting"
	StateFinished       = "finished"
)

type State string

type Die struct {
	Keeping [6]int
	Roll    [6]int
}

type Gameplay struct {
	State     State
	Die       Die
	Score     int
	Qualified bool
	Turns     int

	wg sync.WaitGroup
}

func (gp *Gameplay) Play() error {

	switch gp.State {
	case StateBetting:
		gp.Bet()
		gp.State = StateRolling
	case StateRolling:
		gp.Roll()
		gp.State = StateKeeping
	case StateKeeping:
		gp.Keep()
		gp.UpdateScore()
		if gp.Turns > 0 {
			gp.State = StateRolling
		} else {
			gp.State = StateFinished
			// TODO: clean up func
		}
	case StateFinished:
		return errors.New("finished")
	default:
		return errors.New("unknown state")
	}
	return nil
}

func (gp *Gameplay) Bet() error {
	return nil
}

func (gp *Gameplay) Roll() error {
	gp.Turns = gp.Turns - 1
	return nil
}

func (gp *Gameplay) Keep() error {
	return nil
}

func (gp *Gameplay) UpdateScore() error {
	return nil
}
