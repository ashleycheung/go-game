package engine

// Stores all the events in the game engine
type GameObjectEvent string

const (
	// Called when a body first enters the scene
	OnSceneEnterEvent GameObjectEvent = "onSceneEnter"

	// Called when the body leaves the scene
	OnSceneExitEvent GameObjectEvent = "onSceneExit"

	// Called when a step begins for the given game object
	OnGameObjectStepEvent GameObjectEvent = "onGameObjectStepEvent"
)

type WorldEvent string

// Events that
// are called by the game world
const (
	// Runs before all game steps by the world
	BeforeGameStepEvent WorldEvent = "beforeGameStep"
	// Called after the game step finishes
	AfterGameStepEvent WorldEvent = "afterGameStepEvent"
)

type PhysicsComponentEvent string

const (
	OnPhysicsComponentCollide PhysicsComponentEvent = "onPhysicsComponentCollide"
)

// The data type during collision
type OnPhysicsComponentCollideData struct {
	Target *PhysicsComponent
}

type TimerComponentEvent string

// Called when the timer finishes
const (
	OnTimerEndEvent   TimerComponentEvent = "onTimerEndEvent"
	OnTimerStartEvent TimerComponentEvent = "onTimerStartEvent"
)
