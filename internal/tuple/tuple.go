package tuple

// A Tuple of variable length
type Tuple struct {
	elements []float64
}

// Parts gets parts of a tuple
type Parts = interface {
	Elements() []float64
}

// NewTuple creates a new tuple
func NewTuple(elements ...float64) Tuple {
	return Tuple{elements}
}

// Elements returns this touple elements
func (t Tuple) Elements() []float64 {
	return t.elements
}

// Add adds a tuple to this tuple
func Add(a Parts, b Parts) Tuple {
	var result []float64
	for i := range a.Elements() {
		result = append(result, a.Elements()[i]+b.Elements()[i])
	}
	return Tuple{result}
}
