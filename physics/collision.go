package physics

import (
	"fmt"
	"math"
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
// on the edge they are considered colliding.
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
			} else if body1ShapeType == RectangleType && body2ShapeType == RectangleType {
				doesCollide = RectangleRectangleCollision(body1, body2)
			} else if body1ShapeType == CircleType && body2ShapeType == RectangleType {
				doesCollide = CircleRectangleCollision(body1, body2)
			} else if body1ShapeType == RectangleType && body2ShapeType == CircleType {
				doesCollide = CircleRectangleCollision(body2, body1)
			} else {
				panic(fmt.Sprintf("collisions between %s and %s not supported", body1ShapeType, body2ShapeType))
			}

			// If collision occurs
			// create a collision
			if doesCollide {
				// Add bodies to their respective ids
				body1.CollisionBodyIds[body2.Id] = true
				body2.CollisionBodyIds[body1.Id] = true

				// Add to out collisions
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
func CircleCircleCollision(b1, b2 *Body) bool {
	b1Circle := b1.Shape.(Circle)
	b2Circle := b2.Shape.(Circle)

	dist := b1.Position.DistanceTo(b2.Position)
	radiiSum := b1Circle.Radius + b2Circle.Radius
	return dist <= radiiSum
}

// Based off the first solution here
// https://stackoverflow.com/questions/401847/circle-rectangle-collision-detection-intersection
func CircleRectangleCollision(circleBody, rectangleBody *Body) bool {
	circle := circleBody.Shape.(Circle)
	rect := rectangleBody.Shape.(Rectangle)
	circleDist := Vector{
		X: math.Abs(circleBody.Position.X - rectangleBody.Position.X),
		Y: math.Abs(circleBody.Position.Y - rectangleBody.Position.Y),
	}

	if circleDist.X > (rect.Size.X/2 + circle.Radius) {
		return false
	}
	if circleDist.Y > (rect.Size.Y/2 + circle.Radius) {
		return false
	}

	if circleDist.X <= (rect.Size.X / 2) {
		return true
	}
	if circleDist.Y <= (rect.Size.Y / 2) {
		return true
	}

	cornerDistSqred := circleDist.DistanceSquaredTo(rect.Size.Scale(0.5))

	return cornerDistSqred <= math.Pow(circle.Radius, 2)
}

// Logic taken from here
// https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection
func RectangleRectangleCollision(b1, b2 *Body) bool {
	b1Rect := b1.Shape.(Rectangle)
	b1TopLeft := b1.Position.Subtract(b1Rect.Size.Scale(0.5))
	b1BottomRight := b1.Position.Add(b1Rect.Size.Scale(0.5))

	b2Rect := b2.Shape.(Rectangle)
	b2TopLeft := b2.Position.Subtract(b2Rect.Size.Scale(0.5))
	b2BottomRight := b2.Position.Add(b2Rect.Size.Scale(0.5))

	return b1TopLeft.X < b2BottomRight.X &&
		b1BottomRight.X > b2TopLeft.X &&
		b1TopLeft.Y < b2BottomRight.Y &&
		b1BottomRight.Y > b2TopLeft.Y
}
