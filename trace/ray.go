package trace

import (
	"math"
)

// Ray describes a ray in a scene
type Ray struct {
	Origin, Direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r *Ray) Position(t float64) Tuple {
	distanceTraveled := r.Direction.Multiply(t)
	return r.Origin.Add(distanceTraveled)
}

func (r *Ray) Transform(trans Matrix) Ray {
	return Ray{
		Origin: trans.MultiplyWithTuple(r.Origin), Direction: trans.MultiplyWithTuple(r.Direction)}
}

type SceneObject interface {
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
}

type Intersection struct {
	T      float64
	Object SceneObject
}

type Intersections struct {
	xs []Intersection
}

func (i *Intersections) Hit() (Intersection, bool) {
	lowestNonNegative := Intersection{math.MaxFloat64, nil}
	for _, intersection := range i.xs {
		if intersection.T > 0 && intersection.T < lowestNonNegative.T {
			lowestNonNegative = intersection
		}
	}
	if lowestNonNegative.T < math.MaxFloat64 {
		return lowestNonNegative, true
	} else {
		return lowestNonNegative, false
	}
}

func NewIntersections(i ...Intersection) Intersections {
	return Intersections{xs: i}
}
