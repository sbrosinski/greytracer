package trace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEyeBetweenLightAndSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	result := m.Lighting(light, position, eyev, normalv, false)
	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}

func TestEyeBetweenLightAndSurfaceEyeOffset45(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	result := m.Lighting(light, position, eyev, normalv, false)
	assert.Equal(t, Color{1.0, 1.0, 1.0}, result)
}

func TestEyeOppositeSurfaceLightOffset45(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 10, -10), Intensity: White}
	result := m.Lighting(light, position, eyev, normalv, false)
	assert.True(t, Color{0.7364, 0.7364, 0.7364}.Equals(result))
}

func TestEyeInPathOfReflectionVector(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 10, -10), Intensity: White}
	result := m.Lighting(light, position, eyev, normalv, false)
	assert.True(t, Color{1.6364, 1.6364, 1.6364}.Equals(result))
}

func TestLightBehindSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, 10), Intensity: White}
	result := m.Lighting(light, position, eyev, normalv, false)
	assert.True(t, Color{0.1, 0.1, 0.1}.Equals(result))
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	result := m.Lighting(light, position, eye, normal, false)
	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}

func TestLightingWithSurfaceInShadow(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	inShadow := true
	result := m.Lighting(light, position, eye, normal, inShadow)
	assert.Equal(t, Color{0.1, 0.1, 0.1}, result)
}
