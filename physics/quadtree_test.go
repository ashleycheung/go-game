package physics

import (
	"fmt"
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

// This test shouldnt cause an infinite loop
func TestQuadTreeInfiniteLoop(t *testing.T) {

	body1 := NewBody(Circle{
		Radius: 30,
	})
	body1.Position = Vector{X: 15.342564, Y: 25.779948}

	body2 := NewBody(Rectangle{
		Size: Vector{
			X: 2500,
			Y: 2500,
		},
	})
	body2.Position = Vector{X: 15.342564, Y: 25.779948}

	body3 := NewBody(Circle{
		Radius: 30,
	})
	body3.Position = Vector{X: -15.342564, Y: -25.779948}

	body4 := NewBody(Rectangle{
		Size: Vector{
			X: 2500,
			Y: 2500,
		},
	})
	body4.Position = Vector{X: -15.342564, Y: -25.779948}

	q := NewQuadTreeFromBodies([]*Body{
		body1, body2, body3, body4,
	}, 4, 20)
	fmt.Println(q)
}
