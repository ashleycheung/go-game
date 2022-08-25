package physics

import (
	"fmt"
	"math"
	"math/rand"
)

// Resolves all the collisions in the given world.
// This is not guaranteed to resolve every collision
// so this must be run iteratively with detect collision
func Resolve(collisions []Collision) {
	for _, c := range collisions {
		body1ShapeType := c.B1.Shape.GetType()
		body2ShapeType := c.B2.Shape.GetType()
		if body1ShapeType == CircleType && body2ShapeType == CircleType {
			CircleCircleResolution(c.B1, c.B2)
		} else {
			panic(fmt.Sprintf("Resolution between %s and %s not supported", body1ShapeType, body2ShapeType))
		}
	}
}

// Resolves the circle circle collision
func CircleCircleResolution(b1 *Body, b2 *Body) {
	b1Circle := b1.Shape.(Circle)
	b2Circle := b2.Shape.(Circle)

	overlapAmount := b1Circle.Radius + b2Circle.Radius - b1.Position.DistanceTo(b2.Position)
	if overlapAmount <= 0 {
		return
	}

	b1ToB2Dir := b2.Position.Subtract(b1.Position)
	// If b1 directly on b2, resolve in random direction
	if b1ToB2Dir.Magnitude() == 0 {
		angle := rand.Float64() * math.Pi * 2
		b1ToB2Dir = Vector{
			X: math.Cos(angle),
			Y: math.Sin(angle),
		}
	}
	// Normalize direction
	b1ToB2Dir = b1ToB2Dir.Normalize()

	// Move b2 away by half the overlap amount
	b2.Position = b2.Position.Add(b1ToB2Dir.Scale(overlapAmount / 2))
	// Move b1 away by half overlapped amount in opposite direction
	b1.Position = b1.Position.Subtract(b1ToB2Dir.Scale(overlapAmount / 2))
}
