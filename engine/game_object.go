// Game object in the engine
package engine

import "github.com/ashleycheung/go-game/event"

// An interface for objects in the game.
// All gameobjects are pointers
type GameObject interface {
	// Returns the event manager of the game object
	GetEventManager() *event.EventManager

	// Sets the event manager
	SetEventManager(m *event.EventManager)

	// This is called every step by the game.
	// Delta is the time passed in milliseconds
	Step(delta float64)

	// Adds a child to the tree
	// and give it an id if it doesn't exist
	AddChild(o GameObject) error

	// Removes a child
	RemoveChild(o GameObject)

	// Returns whether the game object has
	// been assigned a unique id yet. Essentially
	// returns g.Id != 0
	HasId() bool

	// Gets the id of the object
	GetId() int

	// Sets the id of the object
	SetId(id int)

	// Gets the game world of the object
	GetWorld() *GameWorld

	// Sets the game world of the object
	SetWorld(w *GameWorld)

	// Returns a slice of the children of this object.
	// All children gameobjects are pointers
	GetChildren() []GameObject

	// Sets the children
	SetChildren(children []GameObject)
}

// Creates a new game object
func NewGameObject() GameObject {
	obj := &BaseGameObject{}
	InitGameObject(obj)
	return obj
}

// Initialises the private properties of
// the game object
func InitGameObject(o GameObject) {
	o.SetChildren([]GameObject{})
	o.SetEventManager(event.NewEventManager())
}

type GameEvent string

const (
	// Called when a body first enters the scene
	OnSceneEnter GameEvent = "onSceneEnter"

	// Called when the body leaves the scene
	OnSceneExit GameEvent = "onSceneExit"
)
