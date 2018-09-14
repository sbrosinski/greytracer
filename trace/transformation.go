package trace

import "math"

// Translation returns a translation matrix for x, y, z
func Translation(x, y, z float64) Matrix {
	return New4X4(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
}

// Scaling returns a scaling matrix for x, y, z
func Scaling(x, y, z float64) Matrix {
	return New4X4(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
}

// RotationX returns a rotation matrix around the x axis by rad degrees
func RotationX(rad float64) Matrix {
	return New4X4(
		1, 0, 0, 0,
		0, math.Cos(rad), -math.Sin(rad), 0,
		0, math.Sin(rad), math.Cos(rad), 0,
		0, 0, 0, 1,
	)
}

// RotationY returns a rotation matrix around the y axis by rad degrees
func RotationY(rad float64) Matrix {
	return New4X4(
		math.Cos(rad), 0, math.Sin(rad), 0,
		0, 1, 0, 0,
		-math.Sin(rad), 0, math.Cos(rad), 0,
		0, 0, 0, 1,
	)
}

// RotationZ returns a rotation matrix around the z axis by rad degrees
func RotationZ(rad float64) Matrix {
	return New4X4(
		math.Cos(rad), -math.Sin(rad), 0, 0,
		math.Sin(rad), math.Cos(rad), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func ViewTransform(from, to, up Tuple) Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)
	orientaton := New4X4(
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-forward.X, -forward.Y, -forward.Z, 0,
		0, 0, 0, 1)
	return orientaton.Multiply(Translation(-from.X, -from.Y, -from.Z))
}
