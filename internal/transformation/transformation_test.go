package transformation_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sbrosinski/greytracer/internal/matrix"

	"github.com/sbrosinski/greytracer/internal/trace"
	"github.com/sbrosinski/greytracer/internal/transformation"
)

func TestMultiplyingByTanslation(t *testing.T) {
	trans := transformation.Translation(5, -3, 2)
	p := trace.NewPoint(-3, 4, 5)
	result := matrix.MultiplyWithTuple(trans, p)
	expected := trace.NewPoint(2, 1, 7)
	assert.Equal(t, expected, result)
}

func TestRotationPointAroundX(t *testing.T) {
	p := trace.NewPoint(0, 1, 0)
	halfQuarter := transformation.RotationX(math.Pi / 4)
	pHalfQuarter := matrix.MultiplyWithTuple(halfQuarter, p)
	expectedHalfQuart := trace.NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)
	assert.True(t, expectedHalfQuart.Equals(pHalfQuarter))

	fullQuarter := transformation.RotationX(math.Pi / 2)
	pFullQuarter := matrix.MultiplyWithTuple(fullQuarter, p)
	expectedFullQuart := trace.NewPoint(0, 0, 1)
	assert.True(t, expectedFullQuart.Equals(pFullQuarter))
}

func TestInverseXRotationRotatesOppositeDirection(t *testing.T) {
	v := trace.NewPoint(0, 1, 0)
	halfQuarter := transformation.RotationX(math.Pi / 4)
	inv := matrix.Inverse(halfQuarter)
	pHalfQuarterInvers := matrix.MultiplyWithTuple(inv, v)
	expected := trace.NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)
	assert.True(t, expected.Equals(pHalfQuarterInvers))
}

func TestRotationAroundY(t *testing.T) {
	p := trace.NewPoint(0, 0, 1)
	halfQuarter := transformation.RotationY(math.Pi / 4)
	pHalfQuarter := matrix.MultiplyWithTuple(halfQuarter, p)
	expectedHalfQuart := trace.NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)
	assert.True(t, expectedHalfQuart.Equals(pHalfQuarter))

	fullQuarter := transformation.RotationY(math.Pi / 2)
	pFullQuarter := matrix.MultiplyWithTuple(fullQuarter, p)
	expectedFullQuart := trace.NewPoint(1, 0, 0)
	assert.True(t, expectedFullQuart.Equals(pFullQuarter))
}

func TestRotationAroundZ(t *testing.T) {
	p := trace.NewPoint(0, 1, 0)
	halfQuarter := transformation.RotationZ(math.Pi / 4)
	pHalfQuarter := matrix.MultiplyWithTuple(halfQuarter, p)
	expectedHalfQuart := trace.NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0)
	assert.True(t, expectedHalfQuart.Equals(pHalfQuarter))

	fullQuarter := transformation.RotationZ(math.Pi / 2)
	pFullQuarter := matrix.MultiplyWithTuple(fullQuarter, p)
	expectedFullQuart := trace.NewPoint(-1, 0, 0)
	assert.True(t, expectedFullQuart.Equals(pFullQuarter))
}
