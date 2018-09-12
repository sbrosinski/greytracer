package trace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayIntersectsSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	intersections := s.Intersect(r)
	assert.Equal(t, 4., intersections.xs[0].T)
	assert.Equal(t, 6., intersections.xs[1].T)
	assert.Equal(t, &s, intersections.xs[0].Object)
	assert.Equal(t, &s, intersections.xs[1].Object)
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	r := Ray{NewPoint(0, 1, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	intersections := s.Intersect(r)
	assert.Equal(t, 5., intersections.xs[0].T)
	assert.Equal(t, 5., intersections.xs[1].T)
}

func TestRayMissesSphere(t *testing.T) {
	r := Ray{NewPoint(0, 2, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	intersections := s.Intersect(r)
	assert.Equal(t, 0, len(intersections.xs))
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	s := NewSphere()
	intersections := s.Intersect(r)
	assert.Equal(t, -1., intersections.xs[0].T)
	assert.Equal(t, 1., intersections.xs[1].T)
}

func TestSphereIsBehindRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, 5), NewVector(0, 0, 1)}
	s := NewSphere()
	intersections := s.Intersect(r)
	assert.Equal(t, -6., intersections.xs[0].T)
	assert.Equal(t, -4., intersections.xs[1].T)
}

func TestIntersectingScaledSphereWithRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := sphere{Transform: Scaling(2, 2, 2)}
	intersections := s.Intersect(r)
	assert.Equal(t, 3., intersections.xs[0].T)
	assert.Equal(t, 7., intersections.xs[1].T)
}

func TestIntersectingTranslatedSphereWithRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := sphere{Transform: Translation(5, 0, 0)}
	intersections := s.Intersect(r)
	assert.True(t, len(intersections.xs) == 0)
}

func TestNormalOnSphereXAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(1, 0, 0))
	assert.Equal(t, NewVector(1, 0, 0), n)
}
func TestNormalOnSphereYAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(0, 1, 0))
	assert.Equal(t, NewVector(0, 1, 0), n)
}
func TestNormalOnSphereZAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(0, 0, 1))
	assert.Equal(t, NewVector(0, 0, 1), n)
}
func TestNormalOnSphereNonAxial(t *testing.T) {
	s := NewSphere()
	sqrt3 := math.Sqrt(3) / 3
	n := s.NormalAt(NewPoint(sqrt3, sqrt3, sqrt3))
	assert.Equal(t, NewVector(sqrt3, sqrt3, sqrt3), n)
}

func TestNormalIsNormalizedVector(t *testing.T) {
	s := NewSphere()
	sqrt3 := math.Sqrt(3) / 3
	n := s.NormalAt(NewPoint(sqrt3, sqrt3, sqrt3))
	assert.Equal(t, n.Normalize(), n)
}

func TestNormalOnTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.Transform = Translation(0, 5, 0)
	n := s.NormalAt(NewPoint(1, 5, 0))
	assert.Equal(t, NewVector(1, 0, 0), n)
}
