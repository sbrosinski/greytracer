package transformation

import "math"
import "github.com/sbrosinski/greytracer/internal/matrix"

// Translation returns a translation matrix for x, y, z
func Translation(x, y, z float64) matrix.Matrix {
	return matrix.New4X4(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
}

// Scaling returns a scaling matrix for x, y, z
func Scaling(x, y, z float64) matrix.Matrix {
	return matrix.New4X4(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
}

// RotationX returns a rotation matrix around the x axis by rad degrees
func RotationX(rad float64) matrix.Matrix {
	return matrix.New4X4(
		1, 0, 0, 0,
		0, math.Cos(rad), -math.Sin(rad), 0,
		0, math.Sin(rad), math.Cos(rad), 0,
		0, 0, 0, 1,
	)
}

// RotationY returns a rotation matrix around the y axis by rad degrees
func RotationY(rad float64) matrix.Matrix {
	return matrix.New4X4(
		math.Cos(rad), 0, math.Sin(rad), 0,
		0, 1, 0, 0,
		-math.Sin(rad), 0, math.Cos(rad), 0,
		0, 0, 0, 1,
	)
}

// RotationZ returns a rotation matrix around the z axis by rad degrees
func RotationZ(rad float64) matrix.Matrix {
	return matrix.New4X4(
		math.Cos(rad), -math.Sin(rad), 0, 0,
		math.Sin(rad), math.Cos(rad), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}
