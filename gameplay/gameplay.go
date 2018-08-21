package gameplay

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
		gp.Keep()

		if gp.isTurn() {
			gp.State = StateRolling
		} else {
			gp.State = StateFinished
			// TODO: clean up func
		}
	case StateFinished:
		gp.ShowStats()
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
		rand.Seed(time.Now().UnixNano())
		gp.Die.Roll = append(gp.Die.Roll, rand.Intn(DieSides)+1)

		fmt.Printf("roll %d: %d\n", i+1, gp.Die.Roll[i])
	}
	return nil
}

func (gp *Gameplay) Keep() error {

	dice, err := gp.getDiceForKeeping()
	if err != nil {
		return err
	}

	if dice == nil {
		return errors.New("need to choose at least one die.")
	}

	for _, k := range dice {
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

func (gp *Gameplay) ShowStats() {
	if gp.Qualified {
		fmt.Println("\nYou have qualified!")
	} else {
		fmt.Println("\nYou did not qualify.")
	}
	fmt.Printf("Your final score is: %d\n", gp.Score)
	fmt.Printf("Your dice are: %v %v\n", gp.Die.Keeping, gp.Die.Qualifiers)
}

func (gp *Gameplay) numQualify(num int) error {
	if !numInSlice(num, gp.Die.Qualifiers) {
		gp.Die.Qualifiers = append(gp.Die.Qualifiers, num)
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

func (gp *Gameplay) getDiceForKeeping() ([]int, error) {
	dice := []int{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Choose dice to keep (enter to exit): ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if text == "" {
			return dice, nil
		}

		i, err := strconv.Atoi(text)
		if err != nil {
			return []int{}, err
		}

    if i == 0 {
      fmt.Printf("roll 0 is not an option: \n")
      continue
    }

    if i > len(gp.Die.Roll) {
      if len(gp.Die.Roll) == 1 {
        fmt.Printf("You only have 1 die to choose from: \n")
      } else {
        fmt.Printf("You only have %d dice to choose from: \n", len(gp.Die.Roll))
      }
    } else {
      dice = append(dice, i-1)
    }
	}
	return []int{}, nil
}
