package engine

import (
	"fmt"
	"time"

	"github.com/ashleycheung/go-game/event"
	"github.com/ashleycheung/go-game/physics"
)

// Represents a game world
type GameWorld struct {
	// The increment that the
	// game world is up to
	idIncrement int

	// Game world specific events
	Event *event.EventManager[WorldEvent]

	// The scene is the root
	// of the game tree
	Scene *GameObject

	// Physics world
	Physics *physics.World

	// Whether the world is currently running
	running bool

	// Maps the group name to a set of
	// game objects
	groupsMap map[string]map[*GameObject]bool

	// Stores a slice of functions to call
	// at the next tick
	funcQueue []func()

	// The time the world started
	worldStartTime time.Time
}

// Gets the world time which is basically the milliseconds
// since the start
func (w *GameWorld) GetWorldTime() float64 {
	return float64(time.Since(w.worldStartTime).Milliseconds())
}

// Queues a function to be called at the start
// of the next tick. This is needed to prevent
// concurrent write errors
func (w *GameWorld) QueueFunction(f func()) {
	w.funcQueue = append(w.funcQueue, f)
}

// Process the functions in the queue
func (w *GameWorld) processFunctions() {
	currQueue := w.funcQueue
	w.funcQueue = []func(){}
	for _, f := range currQueue {
		f()
	}
}

// Return a slice of all the group objects that
// belong in the given group. This operation is extremely fast
// as it is cached. O(g) where g is the number of objects in the group
func (w *GameWorld) GetGroupObjects(groupName string) []*GameObject {
	groupObjects := []*GameObject{}
	groupSet, exists := w.groupsMap[groupName]
	if exists {
		for o := range groupSet {
			groupObjects = append(groupObjects, o)
		}
	}
	return groupObjects
}

// Internal use
// called by the object itself to add to world
func (w *GameWorld) addObjectToGroup(obj *GameObject, groupName string) {
	groupSet, exists := w.groupsMap[groupName]
	if exists {
		groupSet[obj] = true
	} else {
		w.groupsMap[groupName] = map[*GameObject]bool{
			obj: true,
		}
	}
}

// Internal use
// called by object to remove from world
func (w *GameWorld) removeObjectFromGroup(obj *GameObject, groupName string) {
	groupSet, exists := w.groupsMap[groupName]
	if exists {
		delete(groupSet, obj)
	}
}

// Game world step
func (w *GameWorld) Step(delta float64) {
	// Start step
	w.Event.EmitEvent(event.Event[WorldEvent]{
		Name: BeforeGameStepEvent,
	})
	// Process functions
	w.processFunctions()
	// Increment scene
	w.Scene.Step(delta)
	// Update physics
	w.Physics.Step(delta)
	// Emit step finish event
	w.Event.EmitEvent(event.Event[WorldEvent]{
		Name: AfterGameStepEvent,
	})
}

// Runs the world at the given fps.
// If fps is -1, it runs at highest
// possible fps
func (w *GameWorld) Run(fps int) {
	currTime := time.Now()
	w.worldStartTime = time.Now()
	if w.running {
		fmt.Println("can't run world as it is already running")
		return
	}

	w.running = true
	for w.running {
		// Calculate delta
		newTime := time.Now()
		dif := newTime.Sub(currTime)
		delta := float64(dif.Microseconds()) / 1000.0
		currTime = newTime

		startProcessTime := time.Now()

		// Step world
		w.Step(delta)

		// The time it took to process the world
		processingTime := time.Since(startProcessTime)

		// Sleep and wait for next tick
		if fps != -1 {
			time.Sleep(
				(time.Duration(1000.0/float64(fps)))*time.Millisecond - processingTime,
			)
		}
	}
}

// Stops the game world
func (w *GameWorld) Stop() {
	w.running = false
}

// Creates a new game world
func NewGameWorld() *GameWorld {
	w := &GameWorld{
		Event:     event.NewEventManager[WorldEvent](),
		funcQueue: []func(){},
	}
	w.groupsMap = map[string]map[*GameObject]bool{}
	w.Scene = NewScene(w)
	w.Physics = physics.NewWorld()
	return w
}
