package physics

import (
	"fmt"
	"testing"

	"github.com/ashleycheung/go-game/event"
)

func TestBodyClone(t *testing.T) {
	b := NewBody(Circle{Radius: 5})
	b.Mass = 7
	bClone := b.Clone()
	b.Mass = 5
	if bClone.Mass != 7 {
		t.Error("Mass shouldnt change. Body isn't cloned")
	}
}

func TestCollideEvent(t *testing.T) {
	body1 := NewBody(Circle{Radius: 5})
	body2 := NewBody(Circle{Radius: 10})

	world := NewWorld()
	world.AddBody(body1)
	world.AddBody(body2)

	didCollide := true

	// Add listener for collision
	body1.GetEvent().AddListener(string(BodyCollideEvent), func(e event.Event) error {
		targetBody := e.Data.(BodyCollideEventData).TargetBody
		// Prints true
		didCollide = targetBody == body2
		return nil
	})

	world.Step(1000)

	if !didCollide {
		t.Error("collision event did not occur")
	}
}

func ExampleNewBody() {
	body1 := NewBody(Circle{Radius: 5})
	body2 := NewBody(Circle{Radius: 10})

	world := NewWorld()
	world.AddBody(body1)
	world.AddBody(body2)

	// Add listener for collision
	body1.GetEvent().AddListener(string(BodyCollideEvent), func(e event.Event) error {
		targetBody := e.Data.(BodyCollideEventData).TargetBody
		// Prints true
		fmt.Println(targetBody == body2)
		return nil
	})

	world.Step(1000)
}
