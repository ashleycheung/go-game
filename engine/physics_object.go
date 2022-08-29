package engine

import (
	"github.com/ashleycheung/go-game/event"
	"github.com/ashleycheung/go-game/physics"
)

type PhysicsObject struct {
	BaseGameObject
	Body *physics.Body
}

func NewPhysicsObject(shape physics.Shape) *PhysicsObject {
	obj := &PhysicsObject{
		Body: physics.NewBody(shape),
	}

	// Initialise game object
	InitGameObject(obj)

	// On enter add physics body
	obj.GetEventManager().
		AddListener(string(OnSceneEnter), func(e event.Event) error {
			obj.World.Physics.AddBody(obj.Body)
			return nil
		})

	// On exit remove physics body
	obj.GetEventManager().
		AddListener(string(OnSceneExit), func(e event.Event) error {
			obj.World.Physics.RemoveBody(obj.Body)
			return nil
		})

	return obj
}
