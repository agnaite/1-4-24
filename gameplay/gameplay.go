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
	Keeping    []int
	Roll       []int
	Qualifiers []int
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
		gp.Keep([]int{0})

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
	gp.Die.Roll = []int{}

	for i := 0; i < gp.Turns; i++ {
		gp.Die.Roll = append(gp.Die.Roll, rand.Intn(DieSides)+1)

		fmt.Printf("roll %d: %d\n", i+1, gp.Die.Roll[i])
	}
	return nil
}

func (gp *Gameplay) Keep(die []int) error {
	for _, k := range die {
		if !gp.Qualified && (gp.Die.Roll[k] == 1 || gp.Die.Roll[k] == 4) {
			if err := gp.numQualify(gp.Die.Roll[k]); err != nil {
				return err
			}
		} else {
			gp.Die.Keeping = append(gp.Die.Keeping, gp.Die.Roll[k])
			gp.Score = gp.Score + gp.Die.Roll[k]
		}
		gp.Turns = gp.Turns - 1
	}

	return nil
}

func (gp *Gameplay) numQualify(num int) error {
	if !numInSlice(num, gp.Die.Qualifiers) {
		gp.Die.Qualifiers = append(gp.Die.Qualifiers, num)

		fmt.Printf("qualifiers: %#v\n", gp.Die.Qualifiers)
	}
	if len(gp.Die.Qualifiers) == 2 {
		gp.Qualified = true
	}
	return nil
}

func (gp *Gameplay) isTurn() bool {
	return gp.Turns > 0
}

func numInSlice(n int, slice []int) bool {
	for _, num := range slice {
		if num == n {
			return true
		}
	}
	return false
}
