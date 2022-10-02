package event

// Stores a history of the events
// called
type EventHistory[T comparable] struct {
	// Deletes the middlware
	middlewareDeleter func()
	// The max number of events to
	// track before the oldest events
	// are chucked
	BufferSize int
	// Stores the events
	// stored in reverse order
	// (latest is at the end)
	history []Event[T]
	// Name of all the events tracked
	// If empty, all events are tracked
	trackedEvents map[T]bool
}

// Creates a new event history
func NewEventHistory[T comparable]() *EventHistory[T] {
	return &EventHistory[T]{
		BufferSize:    50,
		history:       []Event[T]{},
		trackedEvents: map[T]bool{},
	}
}

// Returns the history of events with
// the latest one first
func (eH *EventHistory[T]) GetHistory() []Event[T] {
	outHistory := make([]Event[T], len(eH.history))
	for i := 0; i < len(eH.history); i++ {
		outHistory[len(eH.history)-i-1] = eH.history[i]
	}
	return outHistory
}

// Clears the history of events tracked
func (eH *EventHistory[T]) ClearHistory() {
	eH.history = []Event[T]{}
}

// Pushes an event on to the buffer
func (eH *EventHistory[T]) pushEvent(event Event[T]) {
	eH.history = append(eH.history, event)
	// Remove first element if greater than buffer size
	if len(eH.history) > eH.BufferSize {
		eH.history = eH.history[1:]
	}
}

// Tracks the given event manager
func (eH *EventHistory[T]) Track(eM *EventManager[T]) {
	// Stop tracking if already exists
	if eH.middlewareDeleter != nil {
		eH.StopTracking()
	}
	// Add middleware
	eH.middlewareDeleter = eM.Middleware(
		func(event Event[T]) Event[T] {
			// Add event if all events are tracked
			// or the tracked event name is given
			if len(eH.trackedEvents) == 0 || eH.trackedEvents[event.Name] {
				eH.pushEvent(event)
			}
			return event
		},
	)
}

// Stops tracking the current event manager
func (eH *EventHistory[T]) StopTracking() {
	if eH.middlewareDeleter != nil {
		eH.middlewareDeleter()
		eH.middlewareDeleter = nil
	}
}

// Tracks the given event. By default, all events
// are tracked but if an event is given, it will only
// track that one
func (eH *EventHistory[T]) TrackEvent(name T) {
	eH.trackedEvents[name] = true
}

// Stops tracking the given event
func (eH *EventHistory[T]) StopTrackingEvent(name T) {
	delete(eH.trackedEvents, name)
}
