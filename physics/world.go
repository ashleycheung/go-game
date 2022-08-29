package physics

import (
	"fmt"
	"time"

	"github.com/ashleycheung/go-game/event"
)

func NewWorld() *World {
	return &World{
		bodies: map[int]*Body{},
		Event:  event.NewEventManager(),
		Config: DefaultWorldConfig(),
	}
}

type World struct {
	// The current id of the object
	// last object created
	idIncrement int

	// The config of the world
	Config WorldConfig

	// Manages the events in the world
	Event *event.EventManager

	// Maps the id to the body
	// in the world
	bodies map[int]*Body

	// Whether the world is running or not
	running bool
}

// Adds a body into the world
func (w *World) AddBody(b *Body) {
	_, exists := w.bodies[b.Id]
	if exists {
		panic(fmt.Sprintf("body with id %d already exists in the world", b.Id))
	}
	// Increment id
	w.idIncrement++
	// Set body to new id
	b.Id = w.idIncrement
	b.world = w
	// Add to map
	w.bodies[b.Id] = b
}

// Gets a body of the given id.
// Body will be nil if it doesnt exist
func (w *World) GetBody(id int) (body *Body) {
	body = w.bodies[id]
	return
}

// Removes a body
func (w *World) RemoveBody(b *Body) bool {
	_, exists := w.bodies[b.Id]
	if !exists {
		return false
	}
	b.world = nil
	delete(w.bodies, b.Id)
	return true
}

// Returns all bodies in the world
func (w *World) Bodies() []*Body {
	outBodies := []*Body{}
	for _, b := range w.bodies {
		outBodies = append(outBodies, b)
	}
	return outBodies
}

// Runs a step in the world
// where delta is the time passed in milliseconds
func (w *World) Step(delta float64) {
	// Update bodies
	for _, b := range w.bodies {
		// Clear collision ids
		b.CollisionBodyIds = map[int]bool{}
		b.Step(delta)
		w.bodies[b.Id] = b
	}

	// Detect collision and continue
	// to resolve until no more collisions occur
	collisions := FindCollisions(w)

	// Resolve the collisions
	Resolve(collisions)

	// Update the velocities from the collisions
	ApplyMomentum(collisions)

	// Call step finish event
	err := w.Event.EmitEvent(event.Event{
		Name: string(StepEndEvent),
	})
	if err != nil {
		panic(err)
	}
}

// Makes a deep clone of this game world
func (w *World) Clone() *World {
	clonedWorld := NewWorld()
	clonedWorld.idIncrement = w.idIncrement
	// A cloned world will be paused by default
	clonedWorld.running = false
	for _, body := range w.bodies {
		clonedWorld.bodies[body.Id] = body.Clone()
	}
	return clonedWorld
}

// Runs the world at the given
// fps. If fps is -1, it runs at max possible
// fps
func (w *World) Run(fps int) {
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

// Stops the world.
// If the world isn't running, does nothing
func (w *World) Stop() {
	w.running = false
}

type WorldEvent string

const (
	// Called when a step has finished
	StepEndEvent WorldEvent = "stepend"
)
