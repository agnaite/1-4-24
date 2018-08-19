package main

import (
	"github.com/audy/1-4-24/gameplay"
	"time"
)

func main() {

	gameplay := &gameplay.Gameplay{
		State: gameplay.StateBetting,
		Die: gameplay.Die{
			Keeping:    []int{},
			Roll:       []int{},
			Qualifiers: []int{},
		},
		Score:     0,
		Qualified: false,
		Turns:     6,
	}

	for {
		err := gameplay.Play()
		if err != nil && err.Error() == "finished" {
			break
		}

		time.Sleep(time.Second)
	}
}
