package trace

import (
	"math"
)

type Sphere struct {
}

func (s *Sphere) intersect(ray Ray) Intersections {
	sphereToRay := ray.origin.Subtract(NewPoint(0, 0, 0))
	a := ray.direction.Dot(ray.direction)
	b := 2 * ray.direction.Dot(sphereToRay)
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
