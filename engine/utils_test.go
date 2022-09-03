package engine

import "testing"

func TestGameObjectIterator(t *testing.T) {
	root := NewGameObject()

	child1 := NewGameObject()
	child2 := NewGameObject()
	child3 := NewGameObject()

	root.AddChild(child1)
	root.AddChild(child2)
	root.AddChild(child3)

	grandchild1 := NewGameObject()
	grandchild2 := NewGameObject()
	grandchild3 := NewGameObject()

	child1.AddChild(grandchild1)
	child1.AddChild(grandchild2)
	child2.AddChild(grandchild3)

	iter := newBFSIterator(root)

	if iter.Next() != root {
		t.Error("wrong node")
	}
	if iter.Next() != child1 {
		t.Error("wrong node")
	}
	if iter.Next() != child2 {
		t.Error("wrong node")
	}
	if iter.Next() != child3 {
		t.Error("wrong node")
	}
	if iter.Next() != grandchild1 {
		t.Error("wrong node")
	}
	if iter.Next() != grandchild2 {
		t.Error("wrong node")
	}
	if iter.Next() != grandchild3 {
		t.Error("wrong node")
	}
	if iter.HasNext() {
		t.Errorf("should be empty")
	}
}
