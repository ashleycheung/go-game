package engine

import (
	"fmt"
	"time"

	"github.com/ashleycheung/go-game/physics"
)

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

	// Whether the world is currently running
	running bool

	// Maps the group name to a set of
	// game objects
	groupsMap map[string]map[GameObject]bool
}

// Return a slice of all the group objects that
// belong in the given group. This operation is extremely fast
// as it is cached. O(g) where g is the number of objects in the group
func (w *GameWorld) GetGroupObjects(groupName string) []GameObject {
	groupObjects := []GameObject{}
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
func (w *GameWorld) addObjectToGroup(obj GameObject, groupName string) {
	groupSet, exists := w.groupsMap[groupName]
	if exists {
		groupSet[obj] = true
	} else {
		w.groupsMap[groupName] = map[GameObject]bool{
			obj: true,
		}
	}
}

// Internal use
// called by object to remove from world
func (w *GameWorld) removeObjectFromGroup(obj GameObject, groupName string) {
	groupSet, exists := w.groupsMap[groupName]
	if exists {
		delete(groupSet, obj)
	}
}

// Game world step
func (w *GameWorld) Step(delta float64) {
	// Increment scene
	w.Scene.Step(delta)
	// Update physics
	w.Physics.Step(delta)
}

// Runs the world at the given fps.
// If fps is -1, it runs at highest
// possible fps
func (w *GameWorld) Run(fps int) {
	currTime := time.Now()
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

		// Step world
		w.Step(delta)

		// Sleep and wait for next tick
		if fps != -1 {
			time.Sleep(
				time.Duration(1000.0/float64(fps))*time.Millisecond -
					time.Duration(delta)*time.Millisecond,
			)
		}
	}
	fmt.Println("World stopped")
}

// Creates a new game world
func NewGameWorld() *GameWorld {
	w := &GameWorld{}
	w.groupsMap = map[string]map[GameObject]bool{}
	w.Scene = NewScene(w)
	w.Physics = physics.NewWorld()
	return w
}
