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
	// array of  clean up funcs called when clean up
	cleanUpFuncs []func()
}

// On object attach add physics body
func (pC *PhysicsComponent) OnGameObjectAttach() {
	obj := pC.GetGameObject()
	// If already in scene add body
	if obj.World != nil {
		obj.World.Physics.AddBody(pC.Body)
	}

	// On enter add physics body
	rmEnter := obj.Event.
		AddListener(string(OnSceneEnterEvent), func(e event.Event) error {
			obj.World.Physics.AddBody(pC.Body)
			return nil
		})

	// On exit remove physics body
	rmExit := obj.Event.
		AddListener(string(OnSceneExitEvent), func(e event.Event) error {
			pC.cleanUp()
			return nil
		})

	// Wrap the clean up func
	// Remove body on clean up
	pC.cleanUpFuncs = append(pC.cleanUpFuncs, func() {
		obj.World.Physics.RemoveBody(pC.Body)
	})
	pC.cleanUpFuncs = append(pC.cleanUpFuncs, rmEnter)
	pC.cleanUpFuncs = append(pC.cleanUpFuncs, rmExit)
}

func (pC *PhysicsComponent) OnGameObjectDetach() {
	pC.cleanUp()
}

func (pC *PhysicsComponent) cleanUp() {
	// Call all the clean up funcs
	for _, f := range pC.cleanUpFuncs {
		f()
	}
}

// Creates new physics component
func NewPhysicsComponent(shape physics.Shape) *PhysicsComponent {
	component := &PhysicsComponent{
		Body:         physics.NewBody(shape),
		Event:        event.NewEventManager(),
		cleanUpFuncs: []func(){},
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
