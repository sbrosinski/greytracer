package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalIsConstantEverywhere(t *testing.T) {
	p := NewPlane()
	assert.Equal(t, NewVector(0, 1, 0), p.NormalAt(NewPoint(10, 0, -10)))
}

func TestIntersectParallelToPlane(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))
	xs := p.Intersect(r)
	assert.True(t, xs.Len() == 0)
}
func TestIntersectWithCoplanarRay(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	xs := p.Intersect(r)
	assert.True(t, xs.Len() == 0)
}
func TestIntersectFromAbove(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0))
	xs := p.Intersect(r)
	assert.True(t, xs.Len() == 1)
	assert.Equal(t, 1.0, xs.xs[0].T)
	assert.Equal(t, p, xs.xs[0].Shape)
}
func TestIntersectFromBelow(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0))
	xs := p.Intersect(r)
	assert.True(t, xs.Len() == 1)
	assert.Equal(t, 1.0, xs.xs[0].T)
	assert.Equal(t, p, xs.xs[0].Shape)
}
