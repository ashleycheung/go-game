package physics

func DefaultWorldConfig() WorldConfig {
	return WorldConfig{
		AirResistance: 10,
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
}
