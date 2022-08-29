package engine

type Scene struct {
	BaseGameObject
}

func NewScene(world *GameWorld) *Scene {
	obj := &Scene{}
	InitGameObject(obj)
	obj.SetWorld(world)
	return obj
}

// Iterates through the children
// and steps them
func (s *Scene) Step(delta float64) {
	objIter := NewBFSIterator(s)
	for objIter.HasNext() {
		nextObj := objIter.Next()
		if nextObj != s {
			nextObj.Step(delta)
		}
	}
}
