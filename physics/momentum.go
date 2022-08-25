package physics

// Update the velocities of the bodies based on collisions.
// This assumes perfectly elastic collision. The internal
// values of the bodies of the collisions are changed
func ApplyMomentum(collisions []Collision) {
	for _, c := range collisions {
		bA := c.B1
		bB := c.B2

		// Formula
		// vAf = (mA - mB) * vA / (mA + mB) + (2 * mB) * vB / (mA + mB)
		newAVel := bA.Velocity.
			Scale((bA.Mass - bB.Mass) / (bA.Mass + bB.Mass)).
			Add(bB.Velocity.Scale(2 * bB.Mass / (bA.Mass + bB.Mass)))

		// Formula
		// vBf = (2 * mA) * vA / (mA + mB) + (mB - mA) * vB / (mA + mB)
		newBVel := bA.Velocity.
			Scale(2 * bA.Mass / (bA.Mass + bB.Mass)).
			Add(bB.Velocity.Scale((bB.Mass - bA.Mass) / (bA.Mass + bB.Mass)))

		bA.Velocity = newAVel
		bB.Velocity = newBVel
	}
}
