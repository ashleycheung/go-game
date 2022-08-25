package event

import "testing"

func TestEvent(t *testing.T) {
	m := NewEventManager()
	l := func(e Event) {}
	m.AddListener("update", l)
	m.EmitEvent(Event{})
}
