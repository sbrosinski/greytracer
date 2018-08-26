package trace

import (
	"math"
)

// TODO
// func (tpl *Tuple) Multiply(t Tuple) Tuple
// func (tpl *Tuple) Divide(t Tuple) Tuple

// Tuple describes a point in 3 dimensional space
type Tuple struct {
	X, Y, Z float64
	W       float64
}

// Add adds a tuple to this tuple
func Add(a Tuple, b Tuple) Tuple {
	return Tuple{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
		a.W + b.W}
}

// Add adds a tuple to this tuple
func (tpl *Tuple) Add(t Tuple) Tuple {
	return Tuple{
		tpl.X + t.X,
		tpl.Y + t.Y,
		tpl.Z + t.Z,
		tpl.W + t.W}
}

// Subtract substracts a tuple from this tuple
func (tpl *Tuple) Subtract(t Tuple) Tuple {
	return Tuple{
		tpl.X - t.X,
		tpl.Y - t.Y,
		tpl.Z - t.Z,
		tpl.W - t.W}
}

// Negate negates this tuple, subtracting it from the zero tuple
func (tpl *Tuple) Negate() Tuple {
	return Tuple{
		0 - tpl.X,
		0 - tpl.Y,
		0 - tpl.Z,
		0 - tpl.W}
}

// Magnitude calculates the magnitude of the vector described by t
func (tpl *Tuple) Magnitude() float64 {
	return math.Sqrt(tpl.X*tpl.X + tpl.Y*tpl.Y + tpl.Z*tpl.Z + tpl.W*tpl.W)
}

// Normalize normalizes a vector
func (tpl *Tuple) Normalize() Tuple {
	var mag = tpl.Magnitude()
	return Tuple{
		tpl.X / mag,
		tpl.Y / mag,
		tpl.Z / mag,
		tpl.W / mag}
}

// Equals checks if this tuple is equal to t
func (tpl *Tuple) Equals(t Tuple) bool {
	return tpl.X == t.X && tpl.Y == t.Y && tpl.Z == t.Z && tpl.W == t.W
}

// Equals checks if this tuple is equal to t
func Equals2(a Tuple, b Tuple) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}

// NewPoint creates a new tuple which is a point
func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// NewVector creates a new tuple which is a vector
func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
