package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayIntersectsSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := Sphere{}
	intersections := s.intersect(r)
	assert.Equal(t, 4., intersections.xs[0].t)
	assert.Equal(t, 6., intersections.xs[1].t)
	assert.Equal(t, &s, intersections.xs[0].object)
	assert.Equal(t, &s, intersections.xs[1].object)
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	r := Ray{NewPoint(0, 1, -5), NewVector(0, 0, 1)}
	s := Sphere{}
	intersections := s.intersect(r)
	assert.Equal(t, 5., intersections.xs[0].t)
	assert.Equal(t, 5., intersections.xs[1].t)
}

func TestRayMissesSphere(t *testing.T) {
	r := Ray{NewPoint(0, 2, -5), NewVector(0, 0, 1)}
	s := Sphere{}
	intersections := s.intersect(r)
	assert.Equal(t, 0, len(intersections.xs))
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	s := Sphere{}
	intersections := s.intersect(r)
	assert.Equal(t, -1., intersections.xs[0].t)
	assert.Equal(t, 1., intersections.xs[1].t)
}

func TestSphereIsBehindRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, 5), NewVector(0, 0, 1)}
	s := Sphere{}
	intersections := s.intersect(r)
	assert.Equal(t, -6., intersections.xs[0].t)
	assert.Equal(t, -4., intersections.xs[1].t)
}
