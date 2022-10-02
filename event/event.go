package event

import "fmt"

func NewEventManager[T comparable]() *EventManager[T] {
	return &EventManager[T]{
		listeners:   map[T]map[int]EventListener[T]{},
		middlewares: map[int]func(event Event[T]) Event[T]{},
		oneTimeIds:  map[int]bool{},
	}
}

type EventManager[T comparable] struct {
	// Maps the event name to map of listener id
	// to the listener
	listeners map[T]map[int]EventListener[T]

	// All middle wares for the event
	middlewares map[int]func(event Event[T]) Event[T]

	idIncrement int

	// Set of all listener ids which are one time
	// listeners
	oneTimeIds map[int]bool
}

type Event[T comparable] struct {
	// Name of the event
	Name T `json:"name"`
	// Contains any additional
	// data to the event
	Data any `json:"data"`
}

type EventListener[T comparable] func(e Event[T]) error

// Adds an event listener
func (e *EventManager[T]) AddListener(
	// Event name
	eventName T,
	// The listener for the event
	listener EventListener[T],
) (
	// When called, removes the listener
	removeListener func(),
) {
	e.idIncrement++
	listenerId := e.idIncrement

	if _, exists := e.listeners[eventName]; !exists {
		e.listeners[eventName] = map[int]EventListener[T]{
			listenerId: listener,
		}
	} else {
		e.listeners[eventName][listenerId] = listener
	}
	return func() {
		delete(e.listeners[eventName], listenerId)
	}
}

// Adds a listener that is only called once
// and then removed
func (e *EventManager[T]) AddOneTimeListener(
	eventName T,
	listener EventListener[T],
) (
	removeListener func(),
) {

	// Add as normal
	rm := e.AddListener(eventName, listener)
	// The id is just the current increment
	listenerId := e.idIncrement
	// Add to one time listener
	e.oneTimeIds[listenerId] = true
	// Wrap remove listener
	removeListener = func() {
		// Clear from one time id
		delete(e.oneTimeIds, listenerId)
		// Remove listener
		rm()
	}
	return
}

// Adds a middleware to the event manager
func (e *EventManager[T]) Middleware(fn func(event Event[T]) Event[T]) func() {
	e.idIncrement++
	id := e.idIncrement
	e.middlewares[id] = fn
	return func() {
		delete(e.middlewares, id)
	}
}

func (e *EventManager[T]) EmitEvent(event Event[T]) error {
	// Parse through all middle wares
	for _, fn := range e.middlewares {
		event = fn(event)
	}

	// Call listeners
	for id, listener := range e.listeners[event.Name] {
		err := listener(event)
		// Clear listener if one time
		if e.oneTimeIds[id] {
			// Remove listener
			delete(e.listeners[event.Name], id)
			delete(e.oneTimeIds, id)
		}
		if err != nil {
			return fmt.Errorf("event %v: %w", event.Name, err)
		}
	}
	return nil
}
