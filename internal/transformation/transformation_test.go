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
	fullQuarter := transformation.RotationX(math.Pi / 2)
	pHalfQuarter := matrix.MultiplyWithTuple(halfQuarter, p)
	pFullQuarter := matrix.MultiplyWithTuple(fullQuarter, p)
	assert.True(t, trace.Equals2(trace.NewPoint(0, math.Sqrt2/2, math.Sqrt2/2), pHalfQuarter))
	assert.True(t, trace.Equals2(trace.NewPoint(0, 0, 1), pFullQuarter))
}
