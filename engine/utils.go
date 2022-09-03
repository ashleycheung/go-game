package engine

type gameObjectIterator interface {
	// Returns whether there is
	// a next game object
	HasNext() bool

	// Returns the next game object.
	// Panics if there is no next object
	Next() GameObject
}

type bFSIterator struct {
	queue []GameObject
}

func newBFSIterator(root GameObject) gameObjectIterator {
	return &bFSIterator{
		queue: []GameObject{root},
	}
}

func (i *bFSIterator) HasNext() bool {
	return len(i.queue) != 0
}

func (i *bFSIterator) Next() GameObject {
	// Pop next
	nextObj := i.queue[0]
	i.queue[0] = nil
	i.queue = i.queue[1:]

	// Add all children to the queue
	i.queue = append(i.queue, nextObj.GetChildren()...)
	return nextObj
}
