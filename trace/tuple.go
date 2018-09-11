package trace

import (
	"math"
)

// TODO
// func (t *Tuple) Divide(t Tuple) Tuple

// Tuple describes a point in 3 dimensional space
type Tuple struct {
	X, Y, Z float64
	W       float64
}

// NewPoint creates a new tuple which is a point
func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// NewVector creates a new tuple which is a vector
func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

// Add adds a tuple to this tuple
func (t *Tuple) Add(a Tuple) Tuple {
	return Tuple{
		t.X + a.X,
		t.Y + a.Y,
		t.Z + a.Z,
		t.W + a.W}
}

// Subtract substracts a tuple from this tuple
func (t *Tuple) Subtract(a Tuple) Tuple {
	return Tuple{
		t.X - a.X,
		t.Y - a.Y,
		t.Z - a.Z,
		t.W - a.W}
}

// Multiply multiplies a tuple from this tuple
func (t *Tuple) Multiply(a float64) Tuple {
	return Tuple{
		t.X * a,
		t.Y * a,
		t.Z * a,
		t.W * a}
}

// Negate negates this tuple, subtracting it from the zero tuple
func (t *Tuple) Negate() Tuple {
	return Tuple{
		0 - t.X,
		0 - t.Y,
		0 - t.Z,
		0 - t.W}
}

// Magnitude calculates the magnitude of the vector described by t
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Normalize normalizes a vector
func (t *Tuple) Normalize() Tuple {
	var mag = t.Magnitude()
	return Tuple{
		t.X / mag,
		t.Y / mag,
		t.Z / mag,
		t.W / mag}
}

// Dot calculates the dot project with another tuple
func (t *Tuple) Dot(a Tuple) float64 {
	return a.X*t.X + a.Y*t.Y + a.Z*t.Z + a.W*t.W
}

// Equals checks if this tuple is mostly equal to t
func (t *Tuple) Equals(a Tuple) bool {
	return floatEquals(a.X, t.X) &&
		floatEquals(a.Y, t.Y) &&
		floatEquals(a.Z, t.Z) &&
		floatEquals(a.W, t.W)
}

func (t *Tuple) Reflect(normal Tuple) Tuple {
	doubled := normal.Multiply(2.0)
	dot := t.Dot(normal)
	multiplyByDot := doubled.Multiply(dot)
	return t.Subtract(multiplyByDot)
}
