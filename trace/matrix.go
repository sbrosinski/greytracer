package trace

import (
	"strconv"
	"strings"
)

// Matrix describes a 2d matrix
type Matrix struct {
	rows int
	cols int
	// flattened matrix data. elements[i*step+j] is row i, col j
	elements []float64
	// actual offset between rows
	step int
}

// New4X4 creates a new 4x4 matrix
func New4X4(elements ...float64) Matrix {
	return Matrix{4, 4, elements, 4}
}

// Identidy4x4 defines a 4x4 identidy matrix
var Identidy4x4 = Matrix{rows: 4, cols: 4, elements: []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, step: 4}

// At returns the value of the matrix at this row and column
func (m Matrix) At(row, col int) float64 {
	return m.elements[row*m.step+col]
}

// Multiply multiplies matrix a with b, need to be the same size
func (m Matrix) Multiply(a Matrix) Matrix {
	var result []float64
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			product := m.At(row, 0)*a.At(0, col) +
				m.At(row, 1)*a.At(1, col) +
				m.At(row, 2)*a.At(2, col) +
				m.At(row, 3)*a.At(3, col)
			result = append(result, product)
		}
	}
	return Matrix{rows: a.rows, cols: a.cols, elements: result, step: a.cols}
}

// MultiplyWithTuple multiplies a matrix with a tuple
func (m Matrix) MultiplyWithTuple(t Tuple) Tuple {
	var result []float64
	for row := 0; row < m.rows; row++ {
		rowProduct := m.At(row, 0)*t.X +
			m.At(row, 1)*t.Y +
			m.At(row, 2)*t.Z +
			m.At(row, 3)*t.W
		result = append(result, rowProduct)
	}
	return Tuple{X: result[0], Y: result[1], Z: result[2], W: result[3]}
}

// Transpose switches rows and columns of a matrix
func (m Matrix) Transpose() Matrix {
	var result []float64
	for col := 0; col < m.cols; col++ {
		for row := 0; row < m.rows; row++ {
			result = append(result, m.At(row, col))
		}
	}
	return Matrix{rows: m.rows, cols: m.cols, elements: result, step: m.cols}
}

// Determinant calculates the determinant of amatrix
func (m Matrix) Determinant() float64 {
	if m.rows == 2 && m.cols == 2 {
		return m.At(0, 0)*m.At(1, 1) - m.At(0, 1)*m.At(1, 0)
	}
	var result float64
	for col := 0; col < m.cols; col++ {
		cofactor := m.Cofactor(0, col)
		result += m.At(0, col) * cofactor
	}
	return result
}

// Submatrix returns a submatrix by removing row and col
func (m Matrix) Submatrix(removeRow, removeCol int) Matrix {
	var result []float64
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			if row != removeRow && col != removeCol {
				result = append(result, m.At(row, col))
			}
		}
	}
	return Matrix{rows: m.rows - 1, cols: m.cols - 1, elements: result, step: m.cols - 1}
}

// Minor calculates the minor of a matrix at a row, col.
// The minor of an element at row i and column j is the determinant of the submatrix at (i,j).
func (m Matrix) Minor(atRow, atCol int) float64 {
	sub := m.Submatrix(atRow, atCol)
	return sub.Determinant()
}

// Cofactor calculates the cofactor of a matrix at row, col.
func (m Matrix) Cofactor(atRow, atCol int) float64 {
	minor := m.Minor(atRow, atCol)
	if (atRow+atCol)%2 != 0 {
		return minor * -1.0
	}
	return minor
}

// IsInvertible checks if a matrix is invertible
func (m Matrix) IsInvertible() bool {
	return m.Determinant() != 0.0
}

// Inverse inverts a matrix
func (m Matrix) Inverse() Matrix {
	var cofactorElements []float64
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			cofactor := m.Cofactor(row, col)
			cofactorElements = append(cofactorElements, cofactor)
		}
	}
	cofactorMatrix := Matrix{rows: m.rows, cols: m.cols, elements: cofactorElements, step: m.cols}
	transposedCofactorMatrix := cofactorMatrix.Transpose()
	determinentOfA := m.Determinant()
	var inversedElements []float64
	for row := 0; row < transposedCofactorMatrix.rows; row++ {
		for col := 0; col < transposedCofactorMatrix.cols; col++ {
			inversedElements = append(inversedElements, transposedCofactorMatrix.At(row, col)/determinentOfA)
		}
	}
	return Matrix{rows: m.rows, cols: m.cols, elements: inversedElements, step: m.cols}
}

// Equal checks if to matrix are equal
func (m Matrix) Equal(a Matrix) bool {
	if a.cols != m.cols || a.rows != m.rows {
		return false
	}
	for row := 0; row < a.rows; row++ {
		for col := 0; col < a.cols; col++ {
			if !floatEquals(a.At(row, col), m.At(row, col)) {
				return false
			}
		}
	}
	return true
}

func floatEquals(a, b float64) bool {
	diff := 0.001
	if (a-b) < diff && (b-a) < diff {
		return true
	}
	return false
}

// Parse creates a matrix from a string
// | 1    | 2    | 3    | 4    |
// | 5.5  | 6.5  | 7.5  | 8.5  |
// | 9    | 10   | 11   | 12   |
// | 13.5 | 14.5 | 15.5 | 16.5 |
func Parse(s string) Matrix {
	var rowCount, valueCount int
	var elements []float64
	for _, row := range strings.Split(s, "\n") {
		// skip lines which don't contain matrix rows
		if !strings.Contains(row, "|") {
			continue
		}
		rowCount++
		for _, col := range strings.Split(row, "|") {
			colItem := strings.Trim(col, " \t|")
			if len(colItem) == 0 {
				continue
			}
			valueCount++
			value, _ := strconv.ParseFloat(colItem, 64)
			elements = append(elements, value)
		}
	}
	colCount := valueCount / rowCount
	return Matrix{rows: rowCount, cols: colCount, elements: elements, step: colCount}
}
