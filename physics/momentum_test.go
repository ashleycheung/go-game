package physics

import (
	"fmt"
	"testing"
)

func TestMomentum(t *testing.T) {
	w := NewWorld()
	b1 := NewBody(Circle{Radius: 2})
	b1.Position = Vector{X: -1, Y: 0}
	b1.Velocity = Vector{X: 5, Y: 0}
	w.AddBody(b1)

	b2 := NewBody(Circle{Radius: 2})
	b2.Position = Vector{X: 1, Y: 0}
	w.AddBody(b2)

	collisions := []Collision{{B1: b1, B2: b2}}
	ApplyMomentum(collisions)

	if w.GetBody(b1.Id).Velocity != NewZeroVector() {
		t.Error("expected zero vector got ",
			fmt.Sprint(w.GetBody(b1.Id).Velocity))
	}

	if w.GetBody(b2.Id).Velocity != (Vector{X: 5}) {
		t.Error("expected { X: 5, Y: 0 } got ",
			fmt.Sprint(w.GetBody(b2.Id).Velocity))
	}

}
