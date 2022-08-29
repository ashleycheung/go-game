package event

import "fmt"

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: map[string]map[int]EventListener{},
	}
}

type EventManager struct {
	// Maps the event name to map of listener id
	// to the listener
	listeners   map[string]map[int]EventListener
	idIncrement int
}

type Event struct {
	Name string
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

func (e *EventManager) EmitEvent(event Event) error {
	for _, listener := range e.listeners[event.Name] {
		err := listener(event)
		if err != nil {
			return fmt.Errorf("event %s: %w", event.Name, err)
		}
	}
	return nil
}
