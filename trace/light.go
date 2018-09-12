package trace

import (
	"math"
)

type Light struct {
	Position  Tuple
	Intensity Color
}

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() Material {
	return Material{Color{1, 1, 1}, 0.1, 0.9, 0.9, 200.0}
}

func (m Material) Lighting(light Light, point, eye, normal Tuple) Color {
	var ambient, diffuse, specular Color
	effectiveColor := m.Color.Multiply(light.Intensity)
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
	return ambient.Add(diffuse).Add(specular)
}
