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
	Intersect(ray Ray) Intersections
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
}

type Intersection struct {
	T       float64
	Object  SceneObject
	Point   Tuple
	EyeV    Tuple
	NormalV Tuple
	Inside  bool
}

func (i *Intersection) PrepareHit(ray Ray) {
	i.Point = ray.Position(i.T)
	i.EyeV = ray.Direction.Negate()
	i.NormalV = i.Object.NormalAt(i.Point)
	i.Inside = i.NormalV.Dot(i.EyeV) < 0
	if i.Inside {
		i.NormalV = i.NormalV.Negate()
	}

	// move point slightly above surface, into direction of normal, to make shadows render properly
	i.Point = i.Point.Add(i.NormalV.Multiply(0.0001))
}

type Intersections struct {
	xs []Intersection
}

// Append adds an intersection to the intersecion collection
func (i *Intersections) Append(inter []Intersection) {
	i.xs = append(i.xs, inter...)
}

func (i Intersections) Len() int {
	return len(i.xs)
}

func (i Intersections) Swap(a, b int) {
	i.xs[a], i.xs[b] = i.xs[b], i.xs[a]
}

func (i Intersections) Less(a, b int) bool {
	return i.xs[a].T < i.xs[b].T
}

func (i Intersections) Hit() (Intersection, bool) {
	lowestNonNegative := Intersection{T: math.MaxFloat64, Object: nil}
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
