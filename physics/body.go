package physics

import "fmt"

// Creates a new body with a unique id.
// The body is not currently in the world yet.
func NewBody(shape Shape) *Body {
	newBody := Body{
		Id:    0,
		Shape: shape,
		Mass:  1,
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
	// Updates position
	b.Position = b.Position.Add(b.Velocity.Scale(delta / 1000))
}
