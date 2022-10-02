package event

import (
	"testing"
)

func TestEvent(t *testing.T) {
	m := NewEventManager()
	l := func(e Event) error { return nil }
	m.AddListener("update", l)
	m.EmitEvent(Event{})
}

func TestMiddleware(t *testing.T) {
	m := NewEventManager()

	called := 0

	delOne := m.Middleware(func(event Event) Event {
		called = 1
		return event
	})

	m.Middleware(func(event Event) Event {
		called = 2
		return event
	})

	delOne()
	m.EmitEvent(Event{})

	if called != 2 {
		t.Error("wrong function deleted")
	}
}
