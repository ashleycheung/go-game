package engine

import "github.com/ashleycheung/go-game/physics"

// Represents a game world
type GameWorld struct {
	// The increment that the
	// game world is up to
	idIncrement int

	// The scene is the root
	// of the game tree
	Scene *Scene

	// Physics world
	Physics *physics.World
}

// Game world step
func (w *GameWorld) Step(delta float64) {
	// Increment scene
	w.Scene.Step(delta)
	// Update physics
	w.Physics.Step(delta)
}

// Creates a new game world
func NewGameWorld() *GameWorld {
	w := &GameWorld{}
	w.Scene = NewScene(w)
	w.Physics = physics.NewWorld()
	return w
}
