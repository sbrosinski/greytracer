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
	result := m.Lighting(NewSphere(), light, position, eyev, normalv, false)
	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}

func TestEyeBetweenLightAndSurfaceEyeOffset45(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	result := m.Lighting(NewSphere(), light, position, eyev, normalv, false)
	assert.Equal(t, Color{1.0, 1.0, 1.0}, result)
}

func TestEyeOppositeSurfaceLightOffset45(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 10, -10), Intensity: White}
	result := m.Lighting(NewSphere(), light, position, eyev, normalv, false)
	assert.True(t, Color{0.7364, 0.7364, 0.7364}.Equals(result))
}

func TestEyeInPathOfReflectionVector(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 10, -10), Intensity: White}
	result := m.Lighting(NewSphere(), light, position, eyev, normalv, false)
	assert.True(t, Color{1.6364, 1.6364, 1.6364}.Equals(result))
}

func TestLightBehindSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, 10), Intensity: White}
	result := m.Lighting(NewSphere(), light, position, eyev, normalv, false)
	assert.True(t, Color{0.1, 0.1, 0.1}.Equals(result))
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	result := m.Lighting(NewSphere(), light, position, eye, normal, false)
	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}

func TestLightingWithSurfaceInShadow(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	inShadow := true
	result := m.Lighting(NewSphere(), light, position, eye, normal, inShadow)
	assert.Equal(t, Color{0.1, 0.1, 0.1}, result)
}

func TestStripePattern(t *testing.T) {
	pattern := NewStripePattern(White, Black)
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 1, 0)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 2, 0)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 0, 1)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 2, 0)))

	assert.Equal(t, White, pattern.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(0.9, 0, 0)))
	assert.Equal(t, Black, pattern.ColorAt(NewPoint(1, 0, 0)))
	assert.Equal(t, Black, pattern.ColorAt(NewPoint(-0.1, 0, 0)))
	assert.Equal(t, Black, pattern.ColorAt(NewPoint(-1, 0, 0)))
	assert.Equal(t, White, pattern.ColorAt(NewPoint(-1.1, 0, 0)))
}

func TestLightingWithPattern(t *testing.T) {
	m := Material{
		Pattern:  NewStripePattern(White, Black),
		Ambient:  1,
		Diffuse:  0,
		Specular: 0,
	}
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := Light{Position: NewPoint(0, 0, -10), Intensity: White}
	c1 := m.Lighting(NewSphere(), light, NewPoint(0.9, 0, 0), eyev, normalv, false)
	c2 := m.Lighting(NewSphere(), light, NewPoint(1.1, 0, 0), eyev, normalv, false)
	assert.Equal(t, White, c1)
	assert.Equal(t, Black, c2)
}

func TestStripesWithObjectTransform(t *testing.T) {
	sphere := NewSphereWithTrans(NewTransform().Scale(2, 2, 2).Matrix())
	pattern := NewStripePattern(Black, White)
	color := pattern.ColorAtShape(sphere.Shape, NewPoint(1.5, 0, 0))
	assert.Equal(t, Black, color)
}

func TestStripesWithPatternTransform(t *testing.T) {
	sphere := NewSphere()
	pattern := NewStripePattern(Black, White)
	pattern.Transform = NewTransform().Scale(2, 2, 2).Matrix()
	color := pattern.ColorAtShape(sphere.Shape, NewPoint(1.5, 0, 0))
	assert.Equal(t, Black, color)
}

func TestStripesWithBothPatternAndShapeTransform(t *testing.T) {
	sphere := NewSphereWithTrans(NewTransform().Scale(2, 2, 2).Matrix())
	pattern := NewStripePattern(Black, White)
	pattern.Transform = NewTransform().Scale(0.5, 0, 0).Matrix()
	color := pattern.ColorAtShape(sphere.Shape, NewPoint(2.5, 0, 0))
	assert.Equal(t, Black, color)
}

func TestGradientPattern(t *testing.T) {
	p := NewGradientPattern(Black, White)
	assert.Equal(t, Black, p.ColorAt(NewPoint(0, 0, 0)))
	assert.Equal(t, Color{0.25, 0.25, 0.25}, p.ColorAt(NewPoint(0.25, 0, 0)))
	assert.Equal(t, Color{0.5, 0.5, 0.5}, p.ColorAt(NewPoint(0.5, 0, 0)))
	assert.Equal(t, Color{0.75, 0.75, 0.75}, p.ColorAt(NewPoint(0.75, 0, 0)))
}

func testPattern() Pattern {
	colorFn := func(point Tuple) Color {
		return Color{point.X, point.Y, point.Z}
	}
	return Pattern{
		Transform: Identidy4x4,
		colorFn:   colorFn,
	}
}

func TestPatternWithBothPatternAndShapeTransform(t *testing.T) {
	sphere := NewSphereWithTrans(NewTransform().Scale(2, 2, 2).Matrix())
	pattern := testPattern()
	pattern.Transform = NewTransform().Translate(0.5, 1, 1.5).Matrix()
	color := pattern.ColorAtShape(sphere.Shape, NewPoint(2.5, 3, 3.5))
	assert.Equal(t, Color{0.75, 0.5, 0.25}, color)
}
