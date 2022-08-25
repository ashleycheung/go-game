package physics

import (
	"fmt"
)

// Represents a collision
// between two bodies
type Collision struct {
	// The first body in the collision
	B1 *Body
	// The second body in the collision
	B2 *Body
}

// Returns all pairs of body collision within
// the world. When two shapes just touch
// on the edge they are not considered colliding.
// They have to overlap
func FindCollisions(w *World) []Collision {
	bodies := w.Bodies()

	outCollisions := []Collision{}

	// Compare all bodies
	for i := 0; i < len(bodies); i++ {
		for j := i + 1; j < len(bodies); j++ {
			body1 := bodies[i]
			body2 := bodies[j]
			body1ShapeType := body1.Shape.GetType()
			body2ShapeType := body2.Shape.GetType()
			var doesCollide bool

			// Pass to the correct shape collision detector
			if body1ShapeType == CircleType && body2ShapeType == CircleType {
				doesCollide = CircleCircleCollision(body1, body2)
			} else {
				panic(fmt.Sprintf("collisions between %s and %s not supported", body1ShapeType, body2ShapeType))
			}

			// If collision occurs
			// create a collision
			if doesCollide {
				outCollisions = append(outCollisions, Collision{
					B1: body1, B2: body2})
			}
		}
	}
	return outCollisions
}

// Returns whether there is a collision between
// two circle bodies. Assumes that b1 and b2 have circle
// body shapes
func CircleCircleCollision(b1 *Body, b2 *Body) bool {
	b1Circle := b1.Shape.(Circle)
	b2Circle := b2.Shape.(Circle)

	dist := b1.Position.DistanceTo(b2.Position)
	radiiSum := b1Circle.Radius + b2Circle.Radius
	return dist < radiiSum
}
