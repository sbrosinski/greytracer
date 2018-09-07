package trace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	var p = NewPoint(4.3, 4.2, 3.1)
	if p.X != 4.3 && p.Y != -4.2 {
		t.Fail()
	}
}

func TestAddingTwoTuples(t *testing.T) {
	var t1 = Tuple{3, -2, 5, 1}
	var t2 = Tuple{-2, 3, 1, 0}
	var result = t1.Add(t2)
	if !result.Equals(Tuple{1, 1, 6, 1}) {
		t.Fail()
	}
}

func TestAddingTwoTuples2(t *testing.T) {
	var t1 = Tuple{3, -2, 5, 1}
	var t2 = Tuple{-2, 3, 1, 0}
	var result = t1.Add(t2)
	assert.Equal(t, Tuple{1, 1, 6, 1}, result)
}

func TestSubtractingTwoPoints(t *testing.T) {
	var p1 = NewPoint(3, 2, 1)
	var p2 = NewPoint(5, 6, 7)
	var result = p1.Subtract(p2)
	if !result.Equals(NewVector(-2, -4, -6)) {
		t.Fail()
	}
}

func TestMultiplyingTuple(t *testing.T) {
	var a = Tuple{1, -2, 3, -4}
	aResult := a.Multiply(3.5)
	assert.Equal(t, Tuple{3.5, -7, 10.5, -14}, aResult)

	var b = Tuple{1, -2, 3, -4}
	bResult := b.Multiply(0.5)
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, bResult)
}

func TestNegatingATuple(t *testing.T) {
	var tuple = Tuple{1, -2, 3, -4}
	var negated = tuple.Negate()
	assert.Equal(t, Tuple{-1, 2, -3, 4}, negated)
}

func TestMagnitude(t *testing.T) {
	var v1 = NewVector(1, 0, 0)
	assert.Equal(t, 1.0, v1.Magnitude())

	var v2 = NewVector(-1, -2, -3)
	assert.Equal(t, math.Sqrt(14), v2.Magnitude())
}

func TestNormalize(t *testing.T) {
	var v1 = NewVector(4, 0, 0)
	var v1Norm = v1.Normalize()
	assert.Equal(t, NewVector(1, 0, 0), v1Norm)
}