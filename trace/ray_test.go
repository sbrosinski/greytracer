package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointFromDistance(t *testing.T) {
	r := Ray{NewPoint(2, 3, 4), NewVector(1, 0, 0)}
	assert.Equal(t, NewPoint(2, 3, 4), r.Position(0))
	assert.Equal(t, NewPoint(3, 3, 4), r.Position(1))
	assert.Equal(t, NewPoint(1, 3, 4), r.Position(-1))
	assert.Equal(t, NewPoint(4.5, 3, 4), r.Position(2.5))
}

func TestHitAllPositive(t *testing.T) {
	s := sphere{}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := NewIntersections(i1, i2)
	i, _ := xs.Hit()
	assert.Equal(t, i1, i)
}

func TestHitSomeNegative(t *testing.T) {
	s := sphere{}
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := NewIntersections(i1, i2)
	i, _ := xs.Hit()
	assert.Equal(t, i2, i)
}

func TestHitAllNegative(t *testing.T) {
	s := sphere{}
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := NewIntersections(i1, i2)
	_, hasHit := xs.Hit()
	assert.False(t, hasHit)
}

func TestHitAlwaysLowestNonNegative(t *testing.T) {
	s := sphere{}
	i1 := Intersection{5, s}
	i2 := Intersection{7, s}
	i3 := Intersection{-3, s}
	i4 := Intersection{2, s}
	xs := NewIntersections(i1, i2, i3, i4)
	i, _ := xs.Hit()
	assert.Equal(t, i4, i)
}

func TestTranslatingRay(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	m := Translation(3, 4, 5)
	r2 := r.Transform(m)
	assert.Equal(t, NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, NewVector(0, 1, 0), r2.Direction)
}

func TestScalingRay(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	m := Scaling(2, 3, 4)
	r2 := r.Transform(m)
	assert.Equal(t, NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, NewVector(0, 3, 0), r2.Direction)
}
