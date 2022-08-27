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

// Should collide
func TestCircleCircleCollision1(t *testing.T) {
	c1R := 5.0
	c1Pos := Vector{X: 0, Y: 0}

	c2R := 5.0
	c2Pos := Vector{X: 8, Y: 0}
	if !CircleCircleCollision(c1R, c1Pos, c2R, c2Pos) {
		t.Fatalf("Should collide")
	}
}

// Even if on edge, the circles should collide
func TestCircleCircleCollision2(t *testing.T) {
	c1R := 5.0
	c1Pos := Vector{X: 0, Y: 0}

	c2R := 5.0
	c2Pos := Vector{X: 10, Y: 0}
	if !CircleCircleCollision(c1R, c1Pos, c2R, c2Pos) {
		t.Fatalf("Should collide")
	}
}

// Should not collide
func TestCircleCircleCollision3(t *testing.T) {
	c1R := 5.0
	c1Pos := Vector{X: 0, Y: 0}

	c2R := 5.0
	c2Pos := Vector{X: 12, Y: 0}
	if CircleCircleCollision(c1R, c1Pos, c2R, c2Pos) {
		t.Fatalf("Should not collide")
	}
}

// A circle inside another should collide
func TestCircleCircleCollision4(t *testing.T) {
	c1R := 5.0
	c1Pos := Vector{X: 0, Y: 0}

	c2R := 20.0
	c2Pos := Vector{X: 1, Y: 0}
	if !CircleCircleCollision(c1R, c1Pos, c2R, c2Pos) {
		t.Fatalf("Should collide")
	}
}

// A circle on the same origin should collide
func TestCircleCircleCollision5(t *testing.T) {
	c1R := 5.0
	c1Pos := Vector{X: 0, Y: 0}

	c2R := 8.0
	c2Pos := Vector{X: 0, Y: 0}
	if !CircleCircleCollision(c1R, c1Pos, c2R, c2Pos) {
		t.Fatalf("Should collide")
	}
}

// The rectangle should collides
func TestRectangleRectangleCollision1(t *testing.T) {
	r1Size := Vector{X: 4, Y: 4}
	r1Position := Vector{X: 4, Y: 4}

	r2Size := Vector{X: 4, Y: 4}
	r2Position := Vector{X: 2, Y: 2}

	if !RectangleRectangleCollision(r1Size, r1Position, r2Size, r2Position) {
		t.Fatalf("Should collide")
	}
}

// Rectangles just touching on edge
// should collide
func TestRectangleRectangleCollision2(t *testing.T) {
	r1Size := Vector{X: 4, Y: 4}
	r1Position := Vector{X: 0, Y: 0}

	r2Size := Vector{X: 4, Y: 4}
	r2Position := Vector{X: 4, Y: 0}

	if !RectangleRectangleCollision(r1Size, r1Position, r2Size, r2Position) {
		t.Fatalf("Should collide")
	}
}

// Rectangle should collide on corners
func TestRectangleRectangleCollision3(t *testing.T) {
	r1Size := Vector{X: 4, Y: 4}
	r1Position := Vector{X: 0, Y: 0}

	r2Size := Vector{X: 4, Y: 4}
	r2Position := Vector{X: 4, Y: 4}

	if !RectangleRectangleCollision(r1Size, r1Position, r2Size, r2Position) {
		t.Fatalf("Should collide")
	}
}
