package physics

import (
	"fmt"
	"image"
	"math"
)

// Represents a 2d vector
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v Vector) String() string {
	return fmt.Sprintf("{ X: %f, Y: %f}", v.X, v.Y)
}

// Return zero vector
func NewZeroVector() Vector {
	return Vector{
		X: 0,
		Y: 0,
	}
}

// Create a new vector with given angle and magnitude
func NewVector(angle, magnitude float64) Vector {
	return Vector{
		X: math.Cos(angle) * magnitude,
		Y: math.Sin(angle) * magnitude,
	}
}

// Multiplies another vector element wise
func (v Vector) ElementMultiply(v2 Vector) Vector {
	return Vector{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
	}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) Subtract(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

// Converts to an image point
func (v Vector) ToImagePoint() image.Point {
	return image.Point{
		X: int(math.Round(v.X)),
		Y: int(math.Round(v.Y)),
	}
}

// Scales a point
func (v Vector) Scale(s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Returns the midpoint of two vectors
func MidPoint(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: (v1.X + v2.X) / 2,
		Y: (v1.Y + v2.Y) / 2,
	}
}

// Magnitude of the vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Returns the sqred of the magnitude
func (v Vector) MagnitudeSqred() float64 {
	return math.Pow(v.X, 2) + math.Pow(v.Y, 2)
}

// Normalize the vector
func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	return Vector{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

// Returns the square of the distance
func (v Vector) DistanceSquaredTo(v2 Vector) float64 {
	return math.Pow(v.X-v2.X, 2) + math.Pow(v.Y-v2.Y, 2)
}

// Returns distance to target vector
func (v Vector) DistanceTo(v2 Vector) float64 {
	return math.Sqrt(v.DistanceSquaredTo(v2))
}

// Returns the dot product
func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Returns whether this is the zero vector
func (v Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// Clamps a vector within the given bbox
func (v Vector) Clamp(bbox BBox) Vector {
	return Vector{
		X: math.Max(bbox.TopLeft.X, math.Min(bbox.BottomRight.X, v.X)),
		Y: math.Max(bbox.TopLeft.Y, math.Min(bbox.BottomRight.Y, v.Y)),
	}
}
