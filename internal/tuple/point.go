package tuple

// Point is a point in 3D space
type Point struct {
	Tuple
}

// NewPoint creates a new point
func NewPoint(x, y, z float64) Point {
	return Point{Tuple: NewTuple(x, y, z, 1.0)}
}
