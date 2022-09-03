package physics

import (
	"fmt"
	"math"

	"github.com/ashleycheung/go-game/event"
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

	// Build quadtree
	w.QuadTree = NewQuadTreeFromBodies(bodies, w.QuadTree.splitAmount, w.QuadTree.maxDepth)

	outCollisions := []Collision{}

	// Compare all bodies
	for i := 0; i < len(bodies); i++ {
		// Get neighbours
		body1 := bodies[i]
		neighbours := w.QuadTree.GetNeighbours(body1)

		// Check if colliding with neighbours
		for _, body2 := range neighbours {
			// If they already have collided ignore
			if body1.CollisionBodyIds[body2.Id] {
				continue
			}

			body1ShapeType := body1.Shape.GetType()
			body2ShapeType := body2.Shape.GetType()
			var doesCollide bool

			// Pass to the correct shape collision detector
			if body1ShapeType == CircleType && body2ShapeType == CircleType {
				b1Circle := body1.Shape.(Circle)
				b2Circle := body2.Shape.(Circle)
				doesCollide = CircleCircleCollision(
					b1Circle.Radius,
					body1.Position,
					b2Circle.Radius,
					body2.Position)
			} else if body1ShapeType == RectangleType && body2ShapeType == RectangleType {
				b1Rect := body1.Shape.(Rectangle)
				b2Rect := body2.Shape.(Rectangle)
				doesCollide = RectangleRectangleCollision(
					b1Rect.Size,
					body1.Position,
					b2Rect.Size,
					body2.Position)
			} else if body1ShapeType == CircleType && body2ShapeType == RectangleType {
				circle := body1.Shape.(Circle)
				rect := body2.Shape.(Rectangle)
				doesCollide = CircleRectangleCollision(circle.Radius, body1.Position, rect.Size, body2.Position)
			} else if body1ShapeType == RectangleType && body2ShapeType == CircleType {
				rect := body1.Shape.(Rectangle)
				circle := body2.Shape.(Circle)
				doesCollide = CircleRectangleCollision(circle.Radius, body2.Position, rect.Size, body1.Position)
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

	// For each pair of collisions
	// call the event
	for _, c := range outCollisions {
		// Call b1 event
		c.B1.GetEvent().EmitEvent(event.Event{
			Name: string(BodyCollideEvent),
			Data: BodyCollideEventData{
				TargetBody: c.B2,
			},
		})

		// Call b2 event
		c.B2.GetEvent().EmitEvent(event.Event{
			Name: string(BodyCollideEvent),
			Data: BodyCollideEventData{
				TargetBody: c.B1,
			},
		})

	}

	return outCollisions
}

// Returns whether there is a collision between
// two circles
func CircleCircleCollision(
	circle1Radius float64,
	circle1Position Vector,
	circle2Radius float64,
	circle2Position Vector,
) (didCollide bool) {
	dist := circle1Position.DistanceTo(circle2Position)
	radiiSum := circle1Radius + circle2Radius
	return dist <= radiiSum
}

// Detects circle and rectangle collision
// Based off the first solution here
// https://stackoverflow.com/questions/401847/circle-rectangle-collision-detection-intersection
func CircleRectangleCollision(
	circleRadius float64,
	circlePosition Vector,
	rectSize Vector,
	rectPosition Vector,
) (didCollide bool) {
	circleDist := Vector{
		X: math.Abs(circlePosition.X - rectPosition.X),
		Y: math.Abs(circlePosition.Y - rectPosition.Y),
	}

	if circleDist.X > (rectSize.X/2 + circleRadius) {
		return false
	}
	if circleDist.Y > (rectSize.Y/2 + circleRadius) {
		return false
	}

	if circleDist.X <= (rectSize.X / 2) {
		return true
	}
	if circleDist.Y <= (rectSize.Y / 2) {
		return true
	}

	cornerDistSqred := circleDist.DistanceSquaredTo(rectSize.Scale(0.5))

	return cornerDistSqred <= math.Pow(circleRadius, 2)
}

// Detects whether two rectangles collide
// Logic taken from here
// https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection
func RectangleRectangleCollision(
	// The size of rectangle 1
	rect1Size Vector,
	// The position of rectangle 1.
	// The position is the center of
	// the rectangle
	rect1Position Vector,
	// The size of the rectangle 2
	rect2Size Vector,
	// The position of the rectangle 2
	// The position is the center of
	// the rectangle
	rect2Position Vector,
) (didCollide bool) {
	b1TopLeft := rect1Position.Subtract(rect1Size.Scale(0.5))
	b1BottomRight := rect1Position.Add(rect1Size.Scale(0.5))

	b2TopLeft := rect2Position.Subtract(rect2Size.Scale(0.5))
	b2BottomRight := rect2Position.Add(rect2Size.Scale(0.5))

	return b1TopLeft.X <= b2BottomRight.X &&
		b1BottomRight.X >= b2TopLeft.X &&
		b1TopLeft.Y <= b2BottomRight.Y &&
		b1BottomRight.Y >= b2TopLeft.Y
}
