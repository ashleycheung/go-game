package physics

import "fmt"

// The name of the shape type
type ShapeType string

// The different shape types supported
const (
	CircleType    ShapeType = "circle"
	RectangleType ShapeType = "rectangle"
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

// The rectangle position is
// the center
type Rectangle struct {
	// Represents the width
	// and height of the rectangle
	// respectively
	Size Vector `json:"size"`
}

func (r Rectangle) GetType() ShapeType {
	return RectangleType
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: width %f, height: %f", r.Size.X, r.Size.Y)
}
