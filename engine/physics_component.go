package engine

import (
	"github.com/ashleycheung/go-game/event"
	"github.com/ashleycheung/go-game/physics"
)

const (
	OnPhysicsComponentCollide = "onPhysicsComponentCollide"
)

// The data type during collision
type OnPhysicsComponentCollideData struct {
	Target *PhysicsComponent
}

// Manages physics
type PhysicsComponent struct {
	BaseComponent
	Body  *physics.Body
	Event *event.EventManager
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

// Creates new physics component
func NewPhysicsComponent(shape physics.Shape) *PhysicsComponent {
	component := &PhysicsComponent{
		Body:  physics.NewBody(shape),
		Event: event.NewEventManager(),
	}
	// Stores the component in the body
	// as metadata
	component.Body.Metadata = component

	// Call collision
	component.Body.GetEvent().AddListener(
		string(physics.BodyCollideEvent),
		func(e event.Event) error {
			component.Event.EmitEvent(event.Event{
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
