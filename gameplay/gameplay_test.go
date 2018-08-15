package gameplay

import (
	"testing"
)

func TestGameplayPlay(t *testing.T) {
	gp := &Gameplay{
		State: StateBetting,
		Die: Die{
			Keeping:    []int{},
			Roll:       []int{},
			Qualifiers: []int{},
		},
		Score:     0,
		Qualified: false,
		Turns:     6,
	}

	for {
		err := gp.Play()
		if err != nil {
			if err.Error() != "finished" {
				t.Fatal(err)
			}
			break
		}
	}

	if want, got := 0, gp.Turns; want != got {
		t.Fatalf("want 0 turns left, got %d", got)
	}

}
