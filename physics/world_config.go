package physics

func DefaultWorldConfig() WorldConfig {
	return WorldConfig{
		AirResistance: 50,
		Gravity:       NewZeroVector(),
	}
}

// Configures the world
type WorldConfig struct {
	// This is the magnitude
	// of the deceleration that will
	// be applied to every object in the world
	// every second. For example, an air resistance
	// of 5 will decrease the velocity by 5 units each
	// second given that the drag resistance of the body is 1.
	// The formula for velocity decrease is
	// airResistance x body.DragCoefficient
	AirResistance float64

	// The gravity vector to apply
	// to the velocity of every nom static body
	// each second. The gravity vector is
	// added to the velocity of the body.
	// So a positive gravity will make
	// the body go down
	Gravity Vector
}
