package physics

import (
	"testing"
)

func TestQuadTree(t *testing.T) {
	tree := NewQuadTree(
		BBox{
			TopLeft: NewZeroVector(),
			BottomRight: Vector{
				X: 3000,
				Y: 3000,
			},
		},
		2,
		2,
	)
	body := NewBody(Circle{Radius: 4})
	body.Position = Vector{X: 4, Y: 4}
	tree.AddBody(body)
	tree.AddBody(NewBody(Circle{Radius: 4}))
}
