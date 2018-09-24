package trace

import (
	"math"
)

type sphere struct {
	Shape
}

func NewSphere() sphere {
	return sphere{
		Shape: NewDefaultShape()}
}

func NewSphereWithTrans(trans Matrix) sphere {
	return sphere{
		Shape: Shape{Transform: trans, Material: NewMaterial()}}
}

func (s sphere) Intersect(ray Ray) Intersections {
	transRay := ray.Transform(s.Shape.Transform.Inverse())

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
	return NewIntersections(Intersection{Shape: s, T: t1}, Intersection{Shape: s, T: t2})
}

func (s sphere) NormalAt(worldPoint Tuple) Tuple {
	invTransf := s.Shape.Transform.Inverse()
	objectPoint := invTransf.MultiplyWithTuple(worldPoint)
	objectNormal := objectPoint.Subtract(NewPoint(0, 0, 0))
	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.MultiplyWithTuple(objectNormal)
	return worldNormal.Normalize()
}

func (s sphere) GetMaterial() Material {
	return s.Shape.Material
}

func (s sphere) GetShape() Shape {
	return s.Shape
}
