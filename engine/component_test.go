package engine

import (
	"testing"
)

func TestIsInScene1(t *testing.T) {
	type TestComponent struct {
		BaseComponent
	}

	tComp := &TestComponent{}

	if tComp.IsInScene() {
		t.Error("not in scene yet")
	}

	obj := NewGameObject()
	obj.AddComponent("test", tComp)
	if tComp.IsInScene() {
		t.Error("not in scene yet")
	}

	w := NewGameWorld()
	w.Scene.AddChild(obj)
	if !tComp.IsInScene() {
		t.Error("should be in scene now")
	}
}

func TestIsInScene2(t *testing.T) {
	type TestComponent struct {
		BaseComponent
	}

	tComp := &TestComponent{}

	if tComp.IsInScene() {
		t.Error("not in scene yet")
	}

	obj := NewGameObject()
	w := NewGameWorld()
	w.Scene.AddChild(obj)
	obj.AddComponent("test", tComp)
	if !tComp.IsInScene() {
		t.Error("should be in scene now")
	}
}

func TestCallWhenInScene1(t *testing.T) {
	type TestComponent struct {
		BaseComponent
	}

	hasCalled := false
	tComp := &TestComponent{}
	tComp.CallWhenInScene(func() {
		hasCalled = true
	})
	if hasCalled {
		t.Error("should not be called yet")
	}

	obj := NewGameObject()
	obj.AddComponent("test", tComp)
	if hasCalled {
		t.Error("should not be called yet")
	}

	w := NewGameWorld()
	w.Scene.AddChild(obj)
	if !hasCalled {
		t.Error("should be called now")
	}
}

func TestCallWhenInScene2(t *testing.T) {
	type TestComponent struct {
		BaseComponent
	}

	hasCalled := false
	tComp := &TestComponent{}
	tComp.CallWhenInScene(func() {
		hasCalled = true
	})
	if hasCalled {
		t.Error("should not be called yet")
	}

	obj := NewGameObject()
	w := NewGameWorld()
	w.Scene.AddChild(obj)
	obj.AddComponent("test", tComp)
	if !hasCalled {
		t.Error("should be called now")
	}
}

type TestComponent struct {
	BaseComponent
	IsAttached         bool
	IsCurrentlyInScene bool
}

func (tComp *TestComponent) OnSceneEnter() {
	if tComp.IsCurrentlyInScene {
		panic("on scene enter called multiple times")
	}
	tComp.IsCurrentlyInScene = true
}

func (tComp *TestComponent) OnSceneExit() {
	if !tComp.IsCurrentlyInScene {
		panic("on scene exit called when not already in scene")
	}
	tComp.IsCurrentlyInScene = false
}

func (tComp *TestComponent) OnGameObjectAttach() {
	if tComp.IsAttached {
		panic("on attached called multiple times")
	}
	tComp.IsAttached = true
}

func (tComp *TestComponent) OnGameObjectDetach() {
	if !tComp.IsAttached {
		panic("on detached called multiple times")
	}
	tComp.IsAttached = false
}

func TestOnAttachCall1(t *testing.T) {
	tComp := &TestComponent{}
	if tComp.IsAttached {
		t.Error("should not be attached")
	}
	obj := NewGameObject()
	obj.AddComponent("test", tComp)
	if !tComp.IsAttached {
		t.Error("should be attached")
	}
	if tComp.IsCurrentlyInScene {
		t.Error("on scene enter should not be called")
	}
	w := NewGameWorld()
	w.Scene.AddChild(obj)
	if !tComp.IsCurrentlyInScene {
		t.Error("on scene enter should be called")
	}
	w.Scene.RemoveChild(obj)
	if tComp.IsCurrentlyInScene {
		t.Error("on scene exit has not been called")
	}
}

func TestOnAttachCall2(t *testing.T) {
	tComp := &TestComponent{}
	if tComp.IsAttached {
		t.Error("should not be attached")
	}
	obj := NewGameObject()
	w := NewGameWorld()
	w.Scene.AddChild(obj)
	obj.AddComponent("test", tComp)
	if !tComp.IsCurrentlyInScene {
		t.Error("on scene enter should be called")
	}
	obj.RemoveComponent("test")
	if tComp.IsCurrentlyInScene {
		t.Error("on scene exit has not been called")
	}
}
