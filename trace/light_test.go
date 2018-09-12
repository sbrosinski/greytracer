package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	result := m.Lighting(light, position, eye, normal)
	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}