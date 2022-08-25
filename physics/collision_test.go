package physics

import (
	"testing"
)

func TestCollision(t *testing.T) {
	w := NewWorld()

	b1 := NewBody(Circle{2})
	b1.Position = Vector{4, 4}
	w.AddBody(b1)

	b2 := NewBody(Circle{2})
	b2.Position = Vector{2, 1}
	w.AddBody(b2)

	b3 := NewBody(Circle{2})
	b3.Position = Vector{10, 2}
	w.AddBody(b3)

	b4 := NewBody(Circle{2})
	b4.Position = Vector{1, 4}
	w.AddBody(b4)

	b5 := NewBody(Circle{2})
	b5.Position = Vector{10, -2}
	w.AddBody(b5)

	collisions := FindCollisions(w)

	hasCollision := func(b1 *Body, b2 *Body) bool {
		for _, c := range collisions {
			if c.B1 == b1 && c.B2 == b2 {
				return true
			}
			if c.B2 == b1 && c.B1 == b2 {
				return true
			}
		}
		return false
	}

	if !hasCollision(b4, b1) {
		t.Fatalf("Expecting collision between %s and %s", b4, b1)
	}
	if !hasCollision(b4, b2) {
		t.Fatalf("Expecting collision between %s and %s", b4, b2)
	}
	if !hasCollision(b2, b1) {
		t.Fatalf("Expecting collision between %s and %s", b2, b1)
	}
}
