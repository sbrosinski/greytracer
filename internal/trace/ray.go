package trace

import (
	"math"
)

// Ray describes a ray in a scene
type Ray struct {
	origin, direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r *Ray) Position(t float64) Tuple {
	distanceTraveled := r.direction.Multiply(t)
	return r.origin.Add(distanceTraveled)
}

func (r *Ray) Transform(trans Matrix) Ray {
	return Ray{
		origin: trans.MultiplyWithTuple(r.origin), direction: trans.MultiplyWithTuple(r.direction)}
}

type Intersection struct {
	t      float64
	object interface{}
}

type Intersections struct {
	xs []Intersection
}

func (i *Intersections) Hit() (Intersection, bool) {
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
