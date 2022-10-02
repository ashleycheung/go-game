package event

import "fmt"

func NewEventManager() *EventManager {
	return &EventManager{
		listeners:   map[string]map[int]EventListener{},
		middlewares: map[int]func(event Event) Event{},
	}
}

type EventManager struct {
	// Maps the event name to map of listener id
	// to the listener
	listeners map[string]map[int]EventListener

	// All middle wares for the event
	middlewares map[int]func(event Event) Event

	idIncrement int
}

type Event struct {
	// Name of the event
	Name string `json:"name"`
	// Contains any additional
	// data to the event
	Data any `json:"data"`
}

type EventListener func(e Event) error

// Adds an event listener
func (e *EventManager) AddListener(
	// Event name
	eventName string,
	// The listener for the event
	listener EventListener,
) (
	// When called, removes the listener
	removeListener func(),
) {
	e.idIncrement++
	listenerId := e.idIncrement

	if _, exists := e.listeners[eventName]; !exists {
		e.listeners[eventName] = map[int]EventListener{
			listenerId: listener,
		}
	} else {
		e.listeners[eventName][listenerId] = listener
	}
	return func() {
		delete(e.listeners[eventName], listenerId)
	}
}

// Adds a middleware to the event manager
func (e *EventManager) Middleware(fn func(event Event) Event) func() {
	e.idIncrement++
	id := e.idIncrement
	e.middlewares[id] = fn
	return func() {
		delete(e.middlewares, id)
	}
}

func (e *EventManager) EmitEvent(event Event) error {
	// Parse through all middle wares
	for _, fn := range e.middlewares {
		event = fn(event)
	}

	// Call listeners
	for _, listener := range e.listeners[event.Name] {
		err := listener(event)
		if err != nil {
			return fmt.Errorf("event %s: %w", event.Name, err)
		}
	}
	return nil
}
