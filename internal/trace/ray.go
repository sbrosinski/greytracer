package trace

// Ray describes a ray in a scene
type Ray struct {
	origin, direction Tuple
}

func (r *Ray) Position(t float64) Tuple {
	distanceTraveled := r.direction.Multiply(t)
	return r.origin.Add(distanceTraveled)
}
