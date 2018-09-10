package trace

import (
	"math"
)

// Ray describes a ray in a scene
type Ray struct {
	origin, direction Tuple
}

func (r *Ray) Position(t float64) Tuple {
	distanceTraveled := r.direction.Multiply(t)
	return r.origin.Add(distanceTraveled)
}

func (r *Ray) Transform(trans Matrix) Ray {
	return Ray{
		origin: MultiplyWithTuple(trans, r.origin), direction: MultiplyWithTuple(trans, r.direction)}
}

type Intersection struct {
	t      float64
	object interface{}
}

type Intersections struct {
	xs []Intersection
}

func (i *Intersections) hit() (Intersection, bool) {
	lowestNonNegative := Intersection{math.MaxFloat64, nil}
	for _, intersection := range i.xs {
		if intersection.t > 0 && intersection.t < lowestNonNegative.t {
			lowestNonNegative = intersection
		}
	}
	if lowestNonNegative.t < math.MaxFloat64 {
		return lowestNonNegative, true
	} else {
		return lowestNonNegative, false
	}
}

func NewIntersections(i ...Intersection) Intersections {
	return Intersections{xs: i}
}
