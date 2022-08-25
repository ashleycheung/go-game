package event

import "testing"

func TestEvent(t *testing.T) {
	m := NewEventManager()
	l := func(e Event) error { return nil }
	m.AddListener("update", l)
	m.EmitEvent(Event{})
}
