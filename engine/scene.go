package engine

func NewScene(world *GameWorld) *GameObject {
	obj := NewGameObject()
	obj.World = world
	obj.AddComponent("scene", &SceneComponent{
		obj: obj,
	})
	return obj
}

type SceneComponent struct {
	BaseComponent
	obj *GameObject
}

// Iterates through the children
// and steps them
func (s *SceneComponent) Step(delta float64) {
	objIter := newBFSIterator(s.obj)
	for objIter.HasNext() {
		nextObj := objIter.Next()
		if nextObj != s.obj {
			nextObj.Step(delta)
		}
	}
}
