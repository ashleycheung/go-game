package engine

import (
	"testing"
)

func TestBaseGameObjectGroups(t *testing.T) {
	w := NewGameWorld()
	obj1 := NewGameObject()
	w.Scene.AddChild(obj1)
	obj1.AddToGroup("players")
	if len(w.GetGroupObjects("players")) != 1 {
		t.Error("did not add object")
	}

	obj2 := NewGameObject()
	obj2.AddToGroup("players")
	w.Scene.AddChild(obj2)
	if len(w.GetGroupObjects("players")) != 2 {
		t.Error("did not add object")
	}

	w.Scene.RemoveChild(obj2)
	if len(w.GetGroupObjects("players")) != 1 {
		t.Error("did not remove object")
	}

	obj1.RemoveFromGroup("players")
	if len(w.GetGroupObjects("players")) != 0 {
		t.Error("did not remove object")
	}
}
