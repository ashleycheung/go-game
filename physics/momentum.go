package physics

// Update the velocities of the bodies based on collisions.
// This assumes perfectly elastic collision. The internal
// values of the bodies of the collisions are changed
func ApplyMomentum(collisions []Collision) {
	for _, c := range collisions {
		b1 := c.B1
		b2 := c.B2

		// If positions are the same
		// we cant calculate momentum
		if b1.Position.DistanceSquaredTo(b2.Position) == 0 {
			continue
		}

		newVel1 := b1.Velocity
		newVel2 := b2.Velocity

		if b1.Static && b2.Static {
			continue
		} else if b1.Static {
			newVel2 = NewZeroVector()
			// // Uses the same formula as below
			// // except finds the limit of mass1 approach infinity
			// newVel2 = b2.Velocity.Subtract(
			// 	b2.Position.Subtract(b1.Position).Scale(
			// 		b2.Velocity.Subtract(b1.Velocity).Dot(b2.Position.Subtract(b1.Position)) /
			// 			b2.Position.DistanceSquaredTo(b1.Position) * 2,
			// 	))

		} else if b2.Static {
			newVel1 = NewZeroVector()

			// // Uses the same formula as below
			// // except finds the limit of mass2 approach infinity
			// newVel1 = b1.Velocity.Subtract(
			// 	b1.Position.Subtract(b2.Position).Scale(
			// 		b1.Velocity.Subtract(b2.Velocity).Dot(b1.Position.Subtract(b2.Position)) /
			// 			b1.Position.DistanceSquaredTo(b2.Position) * 2,
			// 	))

		} else {
			// Implemented using the formula
			// described here:
			// https://en.wikipedia.org/wiki/Elastic_collision#Two-dimensional_collision_with_two_moving_objects
			newVel1 = b1.Velocity.Subtract(
				b1.Position.Subtract(b2.Position).Scale(
					b1.Velocity.Subtract(b2.Velocity).
						Dot(b1.Position.Subtract(b2.Position)) /
						b1.Position.DistanceSquaredTo(b2.Position) *
						2 * b2.Mass / (b1.Mass + b2.Mass),
				))

			newVel2 = b2.Velocity.Subtract(
				b2.Position.Subtract(b1.Position).Scale(
					b2.Velocity.Subtract(b1.Velocity).
						Dot(b2.Position.Subtract(b1.Position)) /
						b2.Position.DistanceSquaredTo(b1.Position) *
						2 * b1.Mass / (b1.Mass + b2.Mass),
				))
		}

		if !b1.Static {
			b1.Velocity = newVel1
		}
		if !b2.Static {
			b2.Velocity = newVel2
		}
	}
}
