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
