package trace

import (
	"math"
)

type sphere struct {
	Transform Matrix
}

func NewSphere() sphere {
	return sphere{Transform: Identidy4x4}
}

func (s *sphere) Intersect(ray Ray) Intersections {
	transRay := ray.Transform(s.Transform.Inverse())

	sphereToRay := transRay.origin.Subtract(NewPoint(0, 0, 0))
	a := transRay.direction.Dot(transRay.direction)
	b := 2 * transRay.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return NewIntersections()
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	if t1 > t2 {
		t1, t2 = t2, t1
	}
	return NewIntersections(Intersection{object: s, t: t1}, Intersection{object: s, t: t2})
}

func (s *sphere) NormalAt(worldPoint Tuple) {

}
