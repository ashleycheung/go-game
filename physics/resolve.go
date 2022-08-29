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
		// If both bodies are static
		// don't resolve
		if c.B1.Static && c.B2.Static {
			continue
		}

		// Pass to their respective resolver
		body1ShapeType := c.B1.Shape.GetType()
		body2ShapeType := c.B2.Shape.GetType()
		if body1ShapeType == CircleType && body2ShapeType == CircleType {
			CircleCircleResolution(c.B1, c.B2)
		} else if body1ShapeType == RectangleType && body2ShapeType == RectangleType {
			RectangleRectangleResolution(c.B1, c.B2)
		} else if body1ShapeType == CircleType && body2ShapeType == RectangleType {
			CircleRectangleResolution(c.B1, c.B2)
		} else if body1ShapeType == RectangleType && body2ShapeType == CircleType {
			CircleRectangleResolution(c.B2, c.B1)
		} else {
			panic(fmt.Sprintf("Resolution between %s and %s not supported", body1ShapeType, body2ShapeType))
		}
	}
}

// Resolves the circle circle collision
// and assumes that only at most one body is static
func CircleCircleResolution(b1, b2 *Body) {
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

	// If any of them are static
	// don't move them
	if b1.Static {
		// Only move b2
		b2.Position = b2.Position.Add(b1ToB2Dir.Scale(overlapAmount))
	} else if b2.Static {
		b1.Position = b1.Position.Subtract(b1ToB2Dir.Scale(overlapAmount))
	} else {
		// Move b2 away by half the overlap amount
		b2.Position = b2.Position.Add(b1ToB2Dir.Scale(overlapAmount / 2))
		// Move b1 away by half overlapped amount in opposite direction
		b1.Position = b1.Position.Subtract(b1ToB2Dir.Scale(overlapAmount / 2))
	}
}

// Assumes that the circle and the rectangle collides
func CircleRectangleResolution(circleBody, rectBody *Body) {
	circle := circleBody.Shape.(Circle)
	rect := rectBody.Shape.(Rectangle)

	xOverlap := circle.Radius + rect.Size.X/2 - math.Abs(circleBody.Position.X-rectBody.Position.X)
	yOverlap := circle.Radius + rect.Size.Y/2 - math.Abs(circleBody.Position.Y-rectBody.Position.Y)

	// Move along the axis with the least overlap
	if xOverlap < yOverlap {
		if circleBody.Position.X < rectBody.Position.X {
			// Circle is left of rectangle
			if circleBody.Static {
				rectBody.Position = rectBody.Position.Add(Vector{X: xOverlap})
			} else if rectBody.Static {
				circleBody.Position = circleBody.Position.Subtract(Vector{X: xOverlap})
			} else {
				circleBody.Position = circleBody.Position.Subtract(Vector{X: xOverlap / 2})
				rectBody.Position = rectBody.Position.Add(Vector{X: xOverlap / 2})
			}
		} else {
			// Circle is right of rectangle
			if circleBody.Static {
				rectBody.Position = rectBody.Position.Subtract(Vector{X: xOverlap})
			} else if rectBody.Static {
				circleBody.Position = circleBody.Position.Add(Vector{X: xOverlap})
			} else {
				circleBody.Position = circleBody.Position.Add(Vector{X: xOverlap / 2})
				rectBody.Position = rectBody.Position.Subtract(Vector{X: xOverlap / 2})
			}
		}
	} else {
		// Circle on top of rectangle
		if circleBody.Position.Y < rectBody.Position.Y {
			if circleBody.Static {
				rectBody.Position = rectBody.Position.Add(Vector{Y: yOverlap})
			} else if rectBody.Static {
				circleBody.Position = circleBody.Position.Subtract(Vector{Y: yOverlap})
			} else {
				circleBody.Position = circleBody.Position.Subtract(Vector{Y: yOverlap / 2})
				rectBody.Position = rectBody.Position.Add(Vector{Y: yOverlap / 2})
			}
		} else {
			// Circle under rectangle
			if circleBody.Static {
				rectBody.Position = rectBody.Position.Subtract(Vector{Y: yOverlap})
			} else if rectBody.Static {
				circleBody.Position = circleBody.Position.Add(Vector{Y: yOverlap})
			} else {
				circleBody.Position = circleBody.Position.Add(Vector{Y: yOverlap / 2})
				rectBody.Position = rectBody.Position.Subtract(Vector{Y: yOverlap / 2})
			}
		}
	}
}

// Resolves collision between two rectangles
func RectangleRectangleResolution(b1, b2 *Body) {
	// If both static, there is no resolution
	if b1.Static && b2.Static {
		return
	}

	b1Rect := b1.Shape.(Rectangle)
	b2Rect := b2.Shape.(Rectangle)

	xOverlap := (b1Rect.Size.X+b2Rect.Size.X)/2 - math.Abs(b1.Position.X-b2.Position.X)
	yOverlap := (b1Rect.Size.Y+b2Rect.Size.Y)/2 - math.Abs(b1.Position.Y-b2.Position.Y)

	// Move along the axis with the least overlap
	if xOverlap < yOverlap {
		if b1.Position.X < b2.Position.X {
			// b1 is on the left side of b2
			if b1.Static {
				b2.Position = b2.Position.Add(Vector{X: xOverlap})
			} else if b2.Static {
				b1.Position = b1.Position.Subtract(Vector{X: xOverlap})
			} else {
				b1.Position = b1.Position.Subtract(Vector{X: xOverlap / 2})
				b2.Position = b2.Position.Add(Vector{X: xOverlap / 2})
			}
		} else {
			// b2 is on the left side of b1
			if b1.Static {
				b2.Position = b2.Position.Subtract(Vector{X: xOverlap})
			} else if b2.Static {
				b1.Position = b1.Position.Add(Vector{X: xOverlap})
			} else {
				b1.Position = b1.Position.Add(Vector{X: xOverlap / 2})
				b2.Position = b2.Position.Subtract(Vector{X: xOverlap / 2})
			}
		}
	} else {
		if b1.Position.Y < b2.Position.Y {
			// b1 is on top of b2
			if b1.Static {
				b2.Position = b2.Position.Add(Vector{Y: yOverlap})
			} else if b2.Static {
				b1.Position = b1.Position.Subtract(Vector{Y: yOverlap})
			} else {
				b1.Position = b1.Position.Subtract(Vector{Y: yOverlap / 2})
				b2.Position = b2.Position.Add(Vector{Y: yOverlap / 2})
			}
		} else {
			// b1 is below b2
			if b1.Static {
				b2.Position = b2.Position.Subtract(Vector{Y: yOverlap})
			} else if b2.Static {
				b1.Position = b1.Position.Add(Vector{Y: yOverlap})
			} else {
				b1.Position = b1.Position.Add(Vector{Y: yOverlap / 2})
				b2.Position = b2.Position.Subtract(Vector{Y: yOverlap / 2})
			}
		}
	}
}
