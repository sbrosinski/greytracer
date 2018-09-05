package trace

import (
	"math"
)

type Sphere struct {
}

func (s *Sphere) intersect(ray Ray) []float64 {
	sphereToRay := ray.origin.Subtract(NewPoint(0, 0, 0))
	a := Dot(ray.direction, ray.direction)
	b := 2 * Dot(ray.direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []float64{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	if t1 > t2 {
		t1, t2 = t2, t1
	}
	return []float64{t1, t2}
}
