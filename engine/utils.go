package engine

type GameObjectIterator interface {
	// Returns whether there is
	// a next game object
	HasNext() bool

	// Returns the next game object.
	// Panics if there is no next object
	Next() GameObject
}

type BFSIterator struct {
	queue []GameObject
}

func NewBFSIterator(root GameObject) GameObjectIterator {
	return &BFSIterator{
		queue: []GameObject{root},
	}
}

func (i *BFSIterator) HasNext() bool {
	return len(i.queue) != 0
}

func (i *BFSIterator) Next() GameObject {
	// Pop next
	nextObj := i.queue[0]
	i.queue[0] = nil
	i.queue = i.queue[1:]

	// Add all children to the queue
	i.queue = append(i.queue, nextObj.GetChildren()...)
	return nextObj
}
