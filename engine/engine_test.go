package engine

import (
	"fmt"

	"github.com/ashleycheung/go-game/physics"
)

func Example() {
	// Creates a new game world
	world := NewGameWorld()

	// Creates a generic game object
	obj1 := NewGameObject()
	world.Scene.AddChild(obj1)

	// Creates a generic physics object
	obj2 := NewPhysicsObject(physics.Circle{Radius: 5})
	world.Scene.AddChild(obj2)

	// Make the object part of the "player" group
	obj2.AddToGroup("player")

	// Efficiently gets all
	// the players in the game
	players := world.GetGroupObjects("player")
	fmt.Println(players) // Prints [obj2]

	// Runs the game at 60fps
	world.Run(60)
}
