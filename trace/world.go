package trace

import "sort"

// World defines a scene to be rendered, containing one light and a list of objects
type World struct {
	Light   Light
	Objects []SceneObject
}

// NewDefaultWorld constructs a default world with a light and two spheres
func NewDefaultWorld() World {
	light := Light{NewPoint(-10, 10, -10), Color{1, 1, 1}}

	sphere1 := NewSphere()
	sphere1.Material.Color = Color{0.8, 1.0, 0.6}
	sphere1.Material.Diffuse = 0.7
	sphere1.Material.Specular = 0.2

	sphere2 := NewSphere()
	sphere2.Transform = Scaling(0.5, 0.5, 0.5)

	return World{light, []SceneObject{sphere1, sphere2}}
}

// Intersect calculates intersections of a ray across all objects in this world
func (w World) Intersect(ray Ray) Intersections {
	xs := Intersections{}
	for _, object := range w.Objects {
		objectXs := object.Intersect(ray)
		xs.Append(objectXs.xs)
	}
	sort.Sort(xs)
	return xs
}
