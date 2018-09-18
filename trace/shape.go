package trace

// Shape describes an object in a rendered world. A shape can e.g. be a sphere, a plane.
type Shape interface {
	Intersect(ray Ray) Intersections
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
}
