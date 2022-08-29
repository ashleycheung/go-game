package physics

import (
	"fmt"
	"math"
)

// Creates a new body with a unique id.
// The body is not currently in the world yet.
func NewBody(shape Shape) *Body {
	newBody := Body{
		Id:               0,
		Shape:            shape,
		Mass:             1,
		DragCoefficient:  1,
		CollisionBodyIds: map[int]bool{},
	}
	return &newBody
}

// Represents a physics body
type Body struct {
	// The unique id that the physics body has in the world.
	// The id will be 0 until it is place in the world, after
	// which it will have a positive id
	Id int `json:"id"`

	// The shape of the physics body
	Shape Shape `json:"shape"`

	// The mass of the body
	Mass float64 `json:"mass"`

	// The position of the body
	Position Vector `json:"position"`

	// Velocity in units per second
	Velocity Vector `json:"velocity"`

	// Acceleration in unit per second sqred
	Acceleration Vector `json:"acceleration"`

	// The amount that the world's air resistance
	// affects this body. A drag coefficient of 1
	// will lead to their air resistance of the world
	// having a normal effect. A drag coefficent of 0
	// means the air resistance has no effect
	DragCoefficient float64 `json:"dragCoefficient"`

	// If set to true, this body can't
	// be knocked back
	Static bool `json:"static"`

	// The ids of all the bodies
	// that this body is currently
	// colliding with
	CollisionBodyIds map[int]bool `json:"collisionBodyIds"`

	// A Reference to the world
	// the body is in
	world *World
}

// Converts the body to a string
func (b Body) String() string {
	return fmt.Sprintf("{\n Id: %d \n Shape: %s \n Position: %s \n Velocity: %s \n Acceleration: %s \n}",
		b.Id, b.Shape, b.Position, b.Velocity, b.Acceleration)
}

// Steps the body forward delta
// where delta is the time in milliseconds
func (b *Body) Step(delta float64) {
	// Updates velocity
	b.Velocity = b.Velocity.Add(b.Acceleration.Scale(delta / 1000))

	// Apply gravity
	// if not static
	if b.world != nil && !b.world.Config.Gravity.IsZero() && !b.Static {
		b.Velocity = b.Velocity.Add(b.world.Config.Gravity.Scale(delta / 1000))
	}

	// Apply air resistance
	if b.world != nil &&
		b.world.Config.AirResistance != 0 &&
		!b.Velocity.IsZero() {

		// The amount to resist in this time frame
		airResistAmount := b.DragCoefficient *
			b.world.Config.AirResistance *
			delta / 1000

		// If the air resistance is more than
		// the current velocity magnitude
		// the velocity should be set to zero
		if b.Velocity.MagnitudeSqred() < math.Pow(airResistAmount, 2) {
			b.Velocity = NewZeroVector()
		} else {
			airResistVel := b.Velocity.Normalize().Scale(airResistAmount)
			b.Velocity = b.Velocity.Subtract(airResistVel)
		}
	}

	// Updates position
	b.Position = b.Position.Add(b.Velocity.Scale(delta / 1000))
}

// Makes a deep clone of the given body
// with the exact same id
func (b *Body) Clone() *Body {
	clonedBody := *b
	return &clonedBody
}
