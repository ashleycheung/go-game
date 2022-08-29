package engine

import "github.com/ashleycheung/go-game/physics"

type PhysicsObject struct {
	BaseGameObject
	Body *physics.Body
}
