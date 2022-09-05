package engine

import (
	"github.com/ashleycheung/go-game/event"
	"github.com/ashleycheung/go-game/physics"
)

type PhysicsComponent struct {
	BaseComponent
	Body *physics.Body
}

// On object attach add physics body
func (pC *PhysicsComponent) OnGameObjectAttach() {
	obj := pC.GetGameObject()
	// If already in scene add body
	if obj.World != nil {
		obj.World.Physics.AddBody(pC.Body)
	}

	// On enter add physics body
	obj.Event.
		AddListener(string(OnSceneEnterEvent), func(e event.Event) error {
			obj.World.Physics.AddBody(pC.Body)
			return nil
		})

	// On exit remove physics body
	obj.Event.
		AddListener(string(OnSceneExitEvent), func(e event.Event) error {
			obj.World.Physics.RemoveBody(pC.Body)
			return nil
		})
}

func NewPhysicsComponent(shape physics.Shape) *PhysicsComponent {
	component := &PhysicsComponent{
		Body: physics.NewBody(shape),
	}
	return component
}
