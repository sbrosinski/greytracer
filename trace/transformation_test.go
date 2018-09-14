package trace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyingByTanslation(t *testing.T) {
	trans := Translation(5, -3, 2)
	p := NewPoint(-3, 4, 5)
	result := trans.MultiplyWithTuple(p)
	expected := NewPoint(2, 1, 7)
	assert.Equal(t, expected, result)
}

func TestRotationPointAroundX(t *testing.T) {
	pHalfQuarter := RotationX(math.Pi / 4).MultiplyWithTuple(NewPoint(0, 1, 0))
	expectedHalfQuart := NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)
	assert.True(t, expectedHalfQuart.Equals(pHalfQuarter))

	fullQuarter := RotationX(math.Pi / 2)
	pFullQuarter := fullQuarter.MultiplyWithTuple(NewPoint(0, 1, 0))
	expectedFullQuart := NewPoint(0, 0, 1)
	assert.True(t, expectedFullQuart.Equals(pFullQuarter))
}

func TestInverseXRotationRotatesOppositeDirection(t *testing.T) {
	v := NewPoint(0, 1, 0)
	halfQuarter := RotationX(math.Pi / 4)
	inv := halfQuarter.Inverse()
	pHalfQuarterInvers := inv.MultiplyWithTuple(v)
	expected := NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)
	assert.True(t, expected.Equals(pHalfQuarterInvers))
}

func TestRotationAroundY(t *testing.T) {
	p := NewPoint(0, 0, 1)
	halfQuarter := RotationY(math.Pi / 4)
	pHalfQuarter := halfQuarter.MultiplyWithTuple(p)
	expectedHalfQuart := NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)
	assert.True(t, expectedHalfQuart.Equals(pHalfQuarter))

	fullQuarter := RotationY(math.Pi / 2)
	pFullQuarter := fullQuarter.MultiplyWithTuple(p)
	expectedFullQuart := NewPoint(1, 0, 0)
	assert.True(t, expectedFullQuart.Equals(pFullQuarter))
}

func TestRotationAroundZ(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := RotationZ(math.Pi / 4)
	pHalfQuarter := halfQuarter.MultiplyWithTuple(p)
	expectedHalfQuart := NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0)
	assert.True(t, expectedHalfQuart.Equals(pHalfQuarter))

	fullQuarter := RotationZ(math.Pi / 2)
	pFullQuarter := fullQuarter.MultiplyWithTuple(p)
	expectedFullQuart := NewPoint(-1, 0, 0)
	assert.True(t, expectedFullQuart.Equals(pFullQuarter))
}

func TestViewTransformDefault(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, -1)
	up := NewVector(0, 1, 0)
	vt := ViewTransform(from, to, up)
	assert.True(t, Identidy4x4.Equal(vt))
}

func TestViewTransformLookingPositiveZ(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, 1)
	up := NewVector(0, 1, 0)
	vt := ViewTransform(from, to, up)
	assert.True(t, Scaling(-1, 1, -1).Equal(vt))
}

func TestViewTransformRandomView(t *testing.T) {
	from := NewPoint(1, 3, 2)
	to := NewPoint(4, -2, 8)
	up := NewVector(1, 1, 0)
	vt := ViewTransform(from, to, up)
	expected := Parse(`
		| -0.50709 | 0.50709 | 0.67612 | -2.36643 |
		| 0.76772 | 0.60609 | 0.12122 | -2.82843 |
		| -0.35857 | 0.59761 | -0.71714 | 0.00000 |
		| 0.00000 | 0.00000 | 0.00000 | 1.00000 |
	`)
	assert.True(t, expected.Equal(vt))
}
