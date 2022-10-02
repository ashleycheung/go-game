package event

// Stores a history of the events
// called
type EventHistory struct {
	// Deletes the middlware
	middlewareDeleter func()
	// The max number of events to
	// track before the oldest events
	// are chucked
	BufferSize int
	// Stores the events
	// stored in reverse order
	// (latest is at the end)
	history []Event
	// Name of all the events tracked
	// If empty, all events are tracked
	trackedEvents map[string]bool
}

// Creates a new event history
func NewEventHistory() *EventHistory {
	return &EventHistory{
		BufferSize:    50,
		history:       []Event{},
		trackedEvents: map[string]bool{},
	}
}

// Returns the history of events with
// the latest one first
func (eH *EventHistory) GetHistory() []Event {
	outHistory := make([]Event, len(eH.history))
	for i := 0; i < len(eH.history); i++ {
		outHistory[len(eH.history)-i-1] = eH.history[i]
	}
	return outHistory
}

// Clears the history of events tracked
func (eH *EventHistory) ClearHistory() {
	eH.history = []Event{}
}

// Pushes an event on to the buffer
func (eH *EventHistory) pushEvent(event Event) {
	eH.history = append(eH.history, event)
	// Remove first element if greater than buffer size
	if len(eH.history) > eH.BufferSize {
		eH.history = eH.history[1:]
	}
}

// Tracks the given event manager
func (eH *EventHistory) Track(eM *EventManager) {
	// Stop tracking if already exists
	if eH.middlewareDeleter != nil {
		eH.StopTracking()
	}
	// Add middleware
	eH.middlewareDeleter = eM.Middleware(
		func(event Event) Event {
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
func (eH *EventHistory) StopTracking() {
	if eH.middlewareDeleter != nil {
		eH.middlewareDeleter()
		eH.middlewareDeleter = nil
	}
}

// Tracks the given event. By default, all events
// are tracked but if an event is given, it will only
// track that one
func (eH *EventHistory) TrackEvent(name string) {
	eH.trackedEvents[name] = true
}

// Stops tracking the given event
func (eH *EventHistory) StopTrackingEvent(name string) {
	delete(eH.trackedEvents, name)
}
