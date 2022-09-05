package engine

import (
	"fmt"

	"github.com/ashleycheung/go-game/physics"
)

// Make a custom player component
// which gives the object additional
// functionalities
type PlayerComponent struct {
	BaseComponent
}

// Each step print delta
func (c *PlayerComponent) Step(delta float64) {
	fmt.Println(delta)
}

// Create our own custom game object
func NewPlayerObject() *GameObject {
	// Creates base game object
	obj := NewGameObject()

	// Give it the player component
	obj.AddComponent("player", &PlayerComponent{})

	// Makes object a part of the "player" group
	obj.AddToGroup("player")

	return obj
}

func Example() {
	// Creates a new game world
	world := NewGameWorld()

	// Creates a generic game object
	obj1 := NewGameObject()
	world.Scene.AddChild(obj1)

	// Creates a generic physics object
	obj2 := NewGameObject()
	obj2.AddComponent("physics", NewPhysicsComponent(physics.Circle{Radius: 5}))
	world.Scene.AddChild(obj2)

	// Make a custom player object
	obj3 := NewPlayerObject()
	world.Scene.AddChild(obj3)

	// Efficiently gets all
	// the players in the game
	players := world.GetGroupObjects("player")
	fmt.Println(players) // Prints [obj3]

	// Runs the game at 60fps
	world.Run(60)
}
