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
	s := Sphere{}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := NewIntersections(i1, i2)
	i, _ := xs.hit()
	assert.Equal(t, i1, i)
}

func TestHitSomeNegative(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := NewIntersections(i1, i2)
	i, _ := xs.hit()
	assert.Equal(t, i2, i)
}

func TestHitAllNegative(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := NewIntersections(i1, i2)
	_, hasHit := xs.hit()
	assert.False(t, hasHit)
}

func TestHitAlwaysLowestNonNegative(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{5, s}
	i2 := Intersection{7, s}
	i3 := Intersection{-3, s}
	i4 := Intersection{2, s}
	xs := NewIntersections(i1, i2, i3, i4)
	i, _ := xs.hit()
	assert.Equal(t, i4, i)
}
