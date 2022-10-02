package physics

// All the physics events
type PhysicsBodyEvent string

const (
	// Called when a body collides with another body
	// and returns that body.
	BodyCollideEvent PhysicsBodyEvent = "bodycollide"
)

type PhysicsWorldEvent string

const (
	// Called when a step has finished
	StepEndEvent                  PhysicsWorldEvent = "stepend"
	BeforeCollisionDetectionEvent PhysicsWorldEvent = "beforeCollisionDetection"
)
