package engine

import "github.com/ashleycheung/go-game/event"

// Called when the timer finishes
const OnTimerEndEvent = "onTimerEndEvent"

// Creates a new timer component
func NewTimerComponent() *TimerComponent {
	return &TimerComponent{
		Event: event.NewEventManager(),
	}
}

type TimerComponent struct {
	BaseComponent
	// Timer events
	Event *event.EventManager
	// Whether running or not
	isRunning bool
	// Time passed since start
	// in milliseconds
	timePassed float64
	// Duration of the component
	// in milliseconds
	Duration float64
	// Whether to loop the timer when complete
	// or not
	Loop bool
}

// Starts the timer.
// If the timer is already started, it restarts it
func (tC *TimerComponent) Start() {
	tC.isRunning = true
	tC.timePassed = 0
}

// Overrides
func (tC *TimerComponent) Step(delta float64) {
	// Update timer if running
	if tC.isRunning {
		tC.timePassed += delta
		// Check if duration reached
		if tC.timePassed >= tC.Duration {
			tC.timePassed = 0
			// If not looping then stop
			if !tC.Loop {
				tC.isRunning = false
			}
			// Emit event
			tC.Event.EmitEvent(event.Event{
				Name: OnTimerEndEvent,
			})
		}
	}
}

// Returns whether the timer is running or not
func (tC *TimerComponent) IsRunning() bool {
	return tC.isRunning
}
