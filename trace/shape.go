package trace

type Shape struct {
	Transform Matrix
	Material  Material
}

func (s *Shape) ApplyTransformation(tansformation Matrix) {
	s.Transform = tansformation
}

func NewDefaultShape() Shape {
	return Shape{
		Transform: Identidy4x4,
		Material:  NewMaterial(),
	}
}

// ShapeOps describes an object in a rendered world. A shape can e.g. be a sphere, a plane.
type ShapeOps interface {
	Intersect(ray Ray) Intersections
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
	//Lighting(light Light, point, eye, normal Tuple, inShadow bool)
	GetShape() Shape
}
