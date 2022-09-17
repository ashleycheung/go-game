package utils

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	l := []float64{1, 2, 3, 4, 6, 7, 8}
	value := 5.0
	i := BinarySearch(value, l, func(i, j float64) int {
		if i < j {
			return -1
		} else if i > j {
			return 1
		}
		return 0
	})
	if i != 4 {
		t.Error("wrong index")
	}
}

func TestBinarySearchExists(t *testing.T) {
	l := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	value := 6.0
	i := BinarySearch(value, l, func(i, j float64) int {
		if i < j {
			return -1
		} else if i > j {
			return 1
		}
		return 0
	})
	if i != 5 {
		t.Error("wrong index")
	}
}

func TestSortedQueue(t *testing.T) {
	sq := NewSortedQueue(func(i, j float64) int {
		if i < j {
			return -1
		} else if i > j {
			return 1
		}
		return 0
	})
	sq.Push(100)
	sq.Push(200)
	sq.Push(50)
	sq.Push(1100)
	sq.Pop(100)
	sq.Pop(324234)
}
