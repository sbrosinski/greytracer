package matrix

// Matrix describes a 2d matrix
type Matrix struct {
	rows int
	cols int
	// flattened matrix data. elements[i*step+j] is row i, col j
	elements []float64
	// actual offset between rows
	step int
}
