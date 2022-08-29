package engine

import (
	"testing"

	"github.com/ashleycheung/go-game/physics"
)

// Adds itself to physics world
// when entering game world
// and removes itself when exiting
// game world
func TestPhysicsObject(t *testing.T) {
	o := NewPhysicsObject(physics.Circle{Radius: 5})
	w := NewGameWorld()
	w.Scene.AddChild(o)
	if len(w.Physics.Bodies()) != 1 {
		t.Error("body not added")
	}
	w.Scene.RemoveChild(o)
	if len(w.Physics.Bodies()) != 0 {
		t.Error("body not removed")
	}
}
