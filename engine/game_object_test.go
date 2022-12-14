package engine

import "testing"

// Tests that an object not in a world
// can be added as children
func TestGameObjectAddChild1(t *testing.T) {
	obj1 := NewGameObject()
	obj2 := NewGameObject()
	err := obj1.AddChild(obj2)
	if err != nil {
		t.Error(err)
	}
	if len(obj1.GetChildren()) != 1 {
		t.Error("object not added")
	}
	if obj1.HasId() || obj2.HasId() {
		t.Error("Objects should not have id")
	}
}

func TestGameObjectAddChild2(t *testing.T) {
	obj1 := NewGameObject()
	obj2 := NewGameObject()
	obj1.AddChild(obj2)

	w := NewGameWorld()
	err := w.Scene.AddChild(obj1)
	if err != nil {
		t.Error(err)
	}

	if !obj1.HasId() {
		t.Error("object should have id")
	}
	if !obj2.HasId() {
		t.Error("object should have id")
	}
}

func TestGameObjectRemoveChild1(t *testing.T) {
	obj1 := NewGameObject()
	obj2 := NewGameObject()
	obj1.AddChild(obj2)

	obj1.RemoveChild(obj2)

	if len(obj1.GetChildren()) != 0 {
		t.Error("object not removed")
	}
}

func TestGameObjectGroups(t *testing.T) {
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
