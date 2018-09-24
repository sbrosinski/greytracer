package trace

import (
	"math"
)

type Light struct {
	Position  Tuple
	Intensity Color
}

type PatternColorFn func(point Tuple) Color

type Pattern struct {
	Transform Matrix
	colorFn   PatternColorFn
}

func (p Pattern) ColorAt(point Tuple) Color {
	return p.colorFn(point)
}

func (p Pattern) ColorAtShape(shape Shape, worldPoint Tuple) Color {
	objectPoint := shape.Transform.Inverse().MultiplyWithTuple(worldPoint)
	patternPoint := p.Transform.Inverse().MultiplyWithTuple(objectPoint)
	return p.ColorAt(patternPoint)
}

func NewStripePattern(color1, color2 Color) Pattern {
	colorFn := func(point Tuple) Color {
		if int(math.Floor(point.X))%2 == 0 {
			return color1
		}
		return color2
	}

	return Pattern{
		Transform: Identidy4x4,
		colorFn:   colorFn,
	}
}

func NewGradientPattern(color1, color2 Color) Pattern {
	colorFn := func(point Tuple) Color {
		distance := color2.Substract(color1)
		fraction := point.X - math.Floor(point.X)
		return color1.Add(distance.MultiplyByScalar(fraction))
	}

	return Pattern{
		Transform: Identidy4x4,
		colorFn:   colorFn,
	}
}

type Material struct {
	Pattern   Pattern
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() Material {
	return Material{Pattern{}, Color{1, 1, 1}, 0.1, 0.9, 0.9, 200.0}
}

func (m Material) Lighting(shape ShapeOps, light Light, point, eye, normal Tuple, inShadow bool) Color {
	materialColor := m.Color
	if m.Pattern.colorFn != nil {
		materialColor = m.Pattern.ColorAtShape(shape.GetShape(), point)
	}

	var ambient, diffuse, specular Color
	effectiveColor := materialColor.Multiply(light.Intensity)
	lightv := light.Position.Subtract(point).Normalize()
	ambient = effectiveColor.MultiplyByScalar(m.Ambient)
	lightDotNormal := lightv.Dot(normal)
	if lightDotNormal < 0 {
		diffuse = Black
		specular = Black
	} else {
		diffuse = effectiveColor.MultiplyByScalar(m.Diffuse).MultiplyByScalar(lightDotNormal)
		reflectv := lightv.Multiply(-1).Reflect(normal)
		reflectDotEye := math.Pow(reflectv.Dot(eye), m.Shininess)
		if reflectDotEye <= 0 {
			specular = Black
		} else {
			specular = light.Intensity.MultiplyByScalar(m.Specular).MultiplyByScalar(reflectDotEye)
		}
	}

	if inShadow {
		return ambient
	} else {
		return ambient.Add(diffuse).Add(specular)
	}
}
