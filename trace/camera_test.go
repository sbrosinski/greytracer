package trace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPixelSizeHorizontal(t *testing.T) {
	assert.InDelta(t, 0.01, NewCamera(200, 125, math.Pi/2).PixelSize, 0.00001)
}

func TestPixelSizeVertical(t *testing.T) {
	assert.InDelta(t, 0.01, NewCamera(125, 200, math.Pi/2).PixelSize, 0.00001)
}

func TestRayThroughCenterOfCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	assert.Equal(t, NewPoint(0, 0, 0), r.Origin)
	assert.Equal(t, NewVector(0, 0, -1), r.Direction)
}

func TestRayThroughCornerOfCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)
	assert.Equal(t, NewPoint(0, 0, 0), r.Origin)
	assert.True(t, NewVector(0.66519, 0.33259, -0.66851).Equals(r.Direction))
}

func TestRayWhenCameraIsTransformed(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = NewTransform().RotateY(math.Pi/4).Translate(0, -2, 5).Matrix()
	r := c.RayForPixel(100, 50)
	assert.Equal(t, NewPoint(0, 2, -5), r.Origin)
	assert.True(t, NewVector(math.Sqrt2/2, 0, -math.Sqrt2/2).Equals(r.Direction))
}
