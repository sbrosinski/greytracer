package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectWorldWithRay(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := world.Intersect(ray)
	assert.Equal(t, 4, xs.Len())
	assert.Equal(t, 4., xs.xs[0].T)
	assert.Equal(t, 4.5, xs.xs[1].T)
	assert.Equal(t, 5.5, xs.xs[2].T)
	assert.Equal(t, 6., xs.xs[3].T)
}

func TestShadeIntersection(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := world.Objects[0]
	hit := Intersection{T: 4, Object: shape}
	hit.PrepareHit(ray)
	c := world.ShadeHit(hit)
	assert.True(t, Color{0.38066, 0.47583, 0.2855}.Equals(c))
}

func TestShadeIntersectionFromInside(t *testing.T) {
	world := NewDefaultWorld()
	world.Light = Light{Position: NewPoint(0, 0.25, 0), Intensity: Color{1, 1, 1}}
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := world.Objects[1]
	hit := Intersection{T: 0.5, Object: shape}
	hit.PrepareHit(ray)
	c := world.ShadeHit(hit)
	assert.True(t, Color{0.90498, 0.90498, 0.90498}.Equals(c))
}

func TestColorWhenRayMisses(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	assert.True(t, Black.Equals(world.ColorAt(ray)))
}

func TestColorWhenRayHits(t *testing.T) {
	world := NewDefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	assert.True(t, Color{0.38066, 0.47583, 0.2855}.Equals(world.ColorAt(ray)))
}
