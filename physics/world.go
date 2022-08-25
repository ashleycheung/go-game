package physics

import (
	"fmt"
	"time"
)

func NewWorld() *World {
	return &World{
		bodies:                  map[int]*Body{},
		MaxResolutionIterations: 5,
	}
}

type World struct {
	// The current id of the object
	// last object created
	idIncrement int

	// Maps the id to the body
	// in the world
	bodies map[int]*Body

	// The maximum amount of resolution
	// iterations before giving up
	MaxResolutionIterations int
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
		b.Step(delta)
		w.bodies[b.Id] = b
	}

	// Detect collision and continue
	// to resolve until no more collisions occur
	collisions := FindCollisions(w)

	// Update the velocities from the collisions
	ApplyMomentum(collisions)

	// Resolve the collisions
	resolutionIter := 0
	for len(collisions) != 0 && resolutionIter < w.MaxResolutionIterations {
		Resolve(collisions)
		collisions = FindCollisions(w)
		resolutionIter++
	}
}

// Runs the world at the given
// fps
func (w *World) Run(fps int) {
	currTime := time.Now()
	for {
		// Calculate delta
		newTime := time.Now()
		dif := newTime.Sub(currTime)
		delta := float64(dif.Microseconds()) / 1000.0
		currTime = newTime

		// Step world
		w.Step(delta)

		// Sleep and wait for next tick
		time.Sleep(
			time.Duration(1000.0/float64(fps))*time.Millisecond -
				time.Duration(delta)*time.Millisecond,
		)
	}
}
