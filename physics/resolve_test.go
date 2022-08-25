package physics

import (
	"testing"
)

func TestResolve(t *testing.T) {
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

	Resolve(collisions)

	// fmt.Println(w.Bodies())
}
