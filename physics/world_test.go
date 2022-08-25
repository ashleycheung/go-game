package physics

import "testing"

func TestWorldClone(t *testing.T) {
	w := NewWorld()
	b := NewBody(Circle{Radius: 5})
	b.Mass = 7
	w.AddBody(b)

	wClone := w.Clone()

	b.Mass = 5

	clonedB := wClone.GetBody(b.Id)
	if clonedB.Mass != 7 {
		t.Error("Body was not cloned")
	}
}
