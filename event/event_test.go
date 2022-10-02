package event

import (
	"testing"
)

func TestEvent(t *testing.T) {
	m := NewEventManager[string]()
	l := func(e Event[string]) error { return nil }
	m.AddListener("update", l)
	m.EmitEvent(Event[string]{})
}

func TestMiddleware(t *testing.T) {
	m := NewEventManager[string]()

	called := 0

	delOne := m.Middleware(func(event Event[string]) Event[string] {
		called = 1
		return event
	})

	m.Middleware(func(event Event[string]) Event[string] {
		called = 2
		return event
	})

	delOne()
	m.EmitEvent(Event[string]{})

	if called != 2 {
		t.Error("wrong function deleted")
	}
}

func TestOneTimeListener(t *testing.T) {
	m := NewEventManager[string]()
	m.AddOneTimeListener("event", func(e Event[string]) error {
		return nil
	})
	if len(m.listeners) != 1 {
		t.Error("listener not added")
	}
	m.EmitEvent(Event[string]{
		Name: "event",
	})
	if len(m.listeners["event"]) != 0 {
		t.Error("listener not removed")
	}
	rm := m.AddOneTimeListener("event", func(e Event[string]) error {
		return nil
	})
	rm()
	if len(m.listeners["event"]) != 0 {
		t.Error("listener not removed")
	}
	if len(m.oneTimeIds) != 0 {
		t.Error("one time id not removed")
	}
}
