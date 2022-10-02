package event

import (
	"testing"
)

func TestEventHistory(t *testing.T) {
	eH := NewEventHistory[string]()
	eM := NewEventManager[string]()
	eH.Track(eM)
	eM.EmitEvent(Event[string]{
		Name: "one",
	})
	eM.EmitEvent(Event[string]{
		Name: "two",
	})
	history := eH.GetHistory()
	if history[0].Name != "two" {
		t.Error("wrong order")
	}
	if history[1].Name != "one" {
		t.Error("wrong order")
	}
}

func TestEventHistoryBuffer(t *testing.T) {
	eH := NewEventHistory[string]()
	eH.BufferSize = 4
	eM := NewEventManager[string]()
	eH.Track(eM)
	eM.EmitEvent(Event[string]{
		Name: "one",
	})
	eM.EmitEvent(Event[string]{
		Name: "two",
	})
	eM.EmitEvent(Event[string]{
		Name: "three",
	})
	eM.EmitEvent(Event[string]{
		Name: "four",
	})
	eM.EmitEvent(Event[string]{
		Name: "five",
	})
	history := eH.GetHistory()
	if history[0].Name != "five" {
		t.Error("wrong order")
	}
	if history[3].Name != "two" {
		t.Error("wrong order")
	}
	if len(history) != 4 {
		t.Error("buffer size not followed")
	}
}

func TestEventHistoryTrackEvents(t *testing.T) {
	eH := NewEventHistory[string]()
	eM := NewEventManager[string]()
	eH.Track(eM)
	eH.TrackEvent("two")
	eM.EmitEvent(Event[string]{
		Name: "one",
	})
	eM.EmitEvent(Event[string]{
		Name: "two",
	})
	eM.EmitEvent(Event[string]{
		Name: "three",
	})
	eM.EmitEvent(Event[string]{
		Name: "four",
	})
	eM.EmitEvent(Event[string]{
		Name: "five",
	})
	history := eH.GetHistory()
	if len(history) != 1 && history[0].Name != "two" {
		t.Error("only event two should be tracked")
	}
}
