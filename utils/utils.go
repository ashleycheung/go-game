package utils

// Create a new sorted queue
func NewSortedQueue[T comparable](
	equalFunc func(elem1 T, elem2 T) int,
) *SortedQueue[T] {
	sq := &SortedQueue[T]{
		Elems:     []T{},
		equalFunc: equalFunc,
	}
	return sq
}

// A sorted queue based on the given less than func
type SortedQueue[T comparable] struct {
	Elems []T
	// Returns 0 if equal
	// -1 if less than
	// 1 if greater than
	equalFunc func(elem1 T, elem2 T) int
}

// Adds to the sorted queue
func (sq *SortedQueue[T]) Push(elem T) {
	// Binary search
	index := BinarySearch(elem, sq.Elems, sq.equalFunc)
	if index == len(sq.Elems) {
		sq.Elems = append(sq.Elems, elem)
	} else {
		// Put element in that index
		sq.Elems = append(sq.Elems[:index+1], sq.Elems[index:]...)
		sq.Elems[index] = elem
	}
}

// Remove an element from a given index
func (sq *SortedQueue[T]) Remove(index int) {
	// Remove the element
	sq.Elems = append(sq.Elems[:index], sq.Elems[index+1:]...)
}

// Remove from sorted queue
func (sq *SortedQueue[T]) Pop(elem T) {
	// No elements do nothing
	if len(sq.Elems) == 0 {
		return
	}

	index := BinarySearch(elem, sq.Elems, sq.equalFunc)

	// Doesnt exist
	if index == sq.Len() {
		return
	}

	// Confirm that the given index
	// is a found elem
	if sq.equalFunc(sq.Elems[index], elem) == 0 {
		// Remove the element
		sq.Remove(index)
	}
}

func (sq *SortedQueue[T]) Len() int {
	return len(sq.Elems)
}

// Searches for index to place a given value
// by using binary search. If the value exists,
// it will return the index where it exists. If value
// doesnt exist, it will return the index where the value
// should go
func BinarySearch[T comparable](
	value T,
	elems []T,
	// Returns 0 if equal
	// -1 if less than
	// 1 if greater than
	equal func(i, j T) int,
) int {
	return bSearch(
		value,
		elems,
		0,
		len(elems)-1,
		equal,
	)
}

// Sub func used by binary search
func bSearch[T comparable](
	value T,
	elems []T,
	start int,
	end int,
	// Returns 0 if equal
	// -1 if less than
	// 1 if greater than
	equal func(i, j T) int,
) int {
	if len(elems) == 0 {
		return 0
	}
	// If given value is less than the start of
	// the array then it should go to the start
	if equal(value, elems[start]) == -1 {
		return start
	} else if equal(elems[end], value) == -1 {
		// Value is more than last value so place 1 more
		// than end
		return end + 1
	}
	// Get middle value floor
	middle := (start + end) / 2

	if equal(elems[middle], value) == 0 {
		return middle
	}

	// Value is less than middle so search
	// start and end
	if equal(value, elems[middle]) == -1 {
		return bSearch(value, elems, start, middle, equal)
	}
	// More than middle
	return bSearch(value, elems, middle+1, end, equal)
}
