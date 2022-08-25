package physics

import "testing"

func TestBodyClone(t *testing.T) {
	b := NewBody(Circle{Radius: 5})
	b.Mass = 7
	bClone := b.Clone()
	b.Mass = 5
	if bClone.Mass != 7 {
		t.Error("Mass shouldnt change. Body isn't cloned")
	}
}
