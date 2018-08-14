package gameplay

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	StateRolling  State = "rolling"
	StateKeeping        = "keeping"
	StateBetting        = "betting"
	StateFinished       = "finished"

	DieSides int = 6
)

type State string

type Die struct {
	Keeping []int
	Roll    [6]int
}

type Gameplay struct {
	State     State
	Die       Die
	Score     int
	Qualified bool
	Turns     int
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
		if gp.Turns < 3 {
			gp.Keep([]int{0})
		} else {
			gp.Keep([]int{0, 1, 2})
		}
		if gp.isTurn() {
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
	for i := 0; i < gp.Turns; i++ {
		gp.Die.Roll[i] = rand.Intn(DieSides) + 1

		fmt.Printf("roll %d: %d\n", i+1, gp.Die.Roll[i])
	}
	return nil
}

func (gp *Gameplay) Keep(die []int) error {
	for _, k := range die {
		gp.Die.Keeping = append(gp.Die.Keeping, gp.Die.Roll[k])
		gp.Score = gp.Score + gp.Die.Roll[k]
		gp.Turns = gp.Turns - 1
	}

	return nil
}

func (gp *Gameplay) UpdateScore() error {
	return nil
}

func (gp *Gameplay) isTurn() bool {
	return gp.Turns > 0
}
