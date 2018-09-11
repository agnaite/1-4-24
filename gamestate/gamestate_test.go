package gamestate

import (
	"testing"
)

func TestGameStateInitiation(t *testing.T) {
	g := New()

	if want, got := 0, g.currentPlayer; want != got {
		t.Fatalf("want 0 currentPlayer, got %d", got)
	}
}
