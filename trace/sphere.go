package trace

import (
	"math"
)

type sphere struct {
	Transform Matrix
	Material  Material
}

func NewSphere() sphere {
	return sphere{Transform: Identidy4x4, Material: NewMaterial()}
}

func (s sphere) Intersect(ray Ray) Intersections {
	transRay := ray.Transform(s.Transform.Inverse())

	sphereToRay := transRay.Origin.Subtract(NewPoint(0, 0, 0))
	a := transRay.Direction.Dot(transRay.Direction)
	b := 2 * transRay.Direction.Dot(sphereToRay)
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
	return NewIntersections(Intersection{Object: s, T: t1}, Intersection{Object: s, T: t2})
}

func (s sphere) NormalAt(worldPoint Tuple) Tuple {
	invTransf := s.Transform.Inverse()
	objectPoint := invTransf.MultiplyWithTuple(worldPoint)
	objectNormal := objectPoint.Subtract(NewPoint(0, 0, 0))
	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.MultiplyWithTuple(objectNormal)
	return worldNormal.Normalize()
}

func (s sphere) GetMaterial() Material {
	return s.Material
}
