package engine

import (
	"github.com/ashleycheung/go-game/event"
	"github.com/ashleycheung/go-game/physics"
)

// Manages physics
type PhysicsComponent struct {
	BaseComponent
	Body  *physics.Body
	Event *event.EventManager[PhysicsComponentEvent]
}

func (pC *PhysicsComponent) OnSceneEnter() {
	obj := pC.GetGameObject()
	obj.World.Physics.AddBody(pC.Body)
}

func (pC *PhysicsComponent) OnSceneExit() {
	obj := pC.GetGameObject()
	obj.World.Physics.RemoveBody(pC.Body)
}

// Creates new physics component
func NewPhysicsComponent(shape physics.Shape) *PhysicsComponent {
	component := &PhysicsComponent{
		Body:  physics.NewBody(shape),
		Event: event.NewEventManager[PhysicsComponentEvent](),
	}
	// Stores the component in the body
	// as metadata
	component.Body.Metadata = component

	// Call collision
	component.Body.GetEvent().AddListener(
		physics.BodyCollideEvent,
		func(e event.Event[physics.PhysicsBodyEvent]) error {
			component.Event.EmitEvent(event.Event[PhysicsComponentEvent]{
				Name: OnPhysicsComponentCollide,
				// Add the target body's physic component
				// which is stored in the meta data
				Data: OnPhysicsComponentCollideData{
					Target: e.Data.(physics.BodyCollideEventData).TargetBody.Metadata.(*PhysicsComponent),
				},
			})
			return nil
		},
	)
	return component
}
