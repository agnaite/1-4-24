package gameplay

import (
	"testing"
)

// make sure we can create a Gameplay
func TestGameplayInitiation(t *testing.T) {
	gp := New()

	if want, got := 0, gp.Score(); want != got {
		t.Fatalf("want 0 initial score, got %d", got)
	}
}

func testGamePlay(t *testing.T) {
	gp := New()

	gp.Die.Keeping = []int{6, 6, 6, 6}

	if want, got := 24, gp.Score(); want != got {
		t.Fatalf("want score to be %d, got %d", want, got)
	}

}

func TestIsQualified(t *testing.T) {
	gp := New()

	gp.Die.Qualifiers = []int{1}

	if want, got := false, gp.isQualified(); want != got {
		t.Fatalf("want qualified to be false, got %v", gp.isQualified())
	}

	gp.Die.Qualifiers = []int{1, 4}

	if want, got := true, gp.isQualified(); want != got {
		t.Fatalf("want qualified to be true, got %v", gp.isQualified())
	}

	gp.Die.Qualifiers = []int{6, 6}

	if want, got := false, gp.isQualified(); want != got {
		t.Fatalf("want qualified to be false, got %v", gp.isQualified())
	}

	gp.Die.Qualifiers = []int{4, 2, 0}

	if want, got := false, gp.isQualified(); want != got {
		t.Fatalf("want qualified to be false, got %v", gp.isQualified())
	}
}

func TestGameplayPlay(t *testing.T) {
	gp := New()

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
