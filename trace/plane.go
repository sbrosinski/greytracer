package trace

import (
	"math"
)

type Plane struct {
	Shape
}

func NewPlane() Plane {
	return Plane{Shape: NewDefaultShape()}
}

func (p Plane) NormalAt(worldPoint Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (p Plane) Intersect(ray Ray) Intersections {
	if math.Abs(ray.Direction.Y) < 0.0001 {
		return NewIntersections()
	}
	t := -ray.Origin.Y / ray.Direction.Y

	return NewIntersections(Intersection{Shape: p, T: t})
}

func (p Plane) GetMaterial() Material {
	return p.Material
}
