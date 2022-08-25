package physics

import "fmt"

// The name of the shape type
type ShapeType string

// The different shape types supported
const (
	CircleType ShapeType = "circle"
)

// A shape
type Shape interface {
	GetType() ShapeType
	String() string
}

// A circle shape
type Circle struct {
	// The radius of the shape
	Radius float64 `json:"radius"`
}

// Returns the type of the shape
func (c Circle) GetType() ShapeType {
	return CircleType
}

// Converts circle to a string
func (c Circle) String() string {
	return "Circle: radius: " + fmt.Sprint(c.Radius)
}
