package matrix

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/sbrosinski/greytracer/internal/trace"
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

// Identidy4x4 defines a 4x4 identidy matrix
var Identidy4x4 = Matrix{rows: 4, cols: 4, elements: []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, step: 4}

// At returns the value of the matrix at this row and column
func (m Matrix) At(row, col int) float64 {
	return m.elements[row*m.step+col]
}

// Multiply multiplies matrix a with b, need to be the same size
func Multiply(a, b Matrix) Matrix {
	var result []float64
	for row := 0; row < a.rows; row++ {
		for col := 0; col < a.cols; col++ {
			product := a.At(row, 0)*b.At(0, col) +
				a.At(row, 1)*b.At(1, col) +
				a.At(row, 2)*b.At(2, col) +
				a.At(row, 3)*b.At(3, col)
			result = append(result, product)
		}
	}
	return Matrix{rows: a.rows, cols: a.cols, elements: result, step: a.cols}
}

// MultiplyWithTuple multiplies a matrix with a tuple
func MultiplyWithTuple(a Matrix, t trace.Tuple) trace.Tuple {
	var result []float64
	for row := 0; row < a.rows; row++ {
		rowProduct := a.At(row, 0)*t.X +
			a.At(row, 1)*t.Y +
			a.At(row, 2)*t.Z +
			a.At(row, 3)*t.W
		result = append(result, rowProduct)
	}
	return trace.Tuple{X: result[0], Y: result[1], Z: result[2], W: result[3]}
}

// Transpose switches rows and columns of a matrix
func Transpose(a Matrix) Matrix {
	var result []float64
	for col := 0; col < a.cols; col++ {
		for row := 0; row < a.rows; row++ {
			result = append(result, a.At(row, col))
		}
	}
	return Matrix{rows: a.rows, cols: a.cols, elements: result, step: a.cols}
}

// Determinant calculates the determinant of a 2x2 matrix
func Determinant(a Matrix) float64 {
	return a.At(0, 0)*a.At(1, 1) - a.At(0, 1)*a.At(1, 0)
}

// Submatrix returns a submatrix by removing row and col
func Submatrix(a Matrix, removeRow, removeCol int) Matrix {
	var result []float64
	for row := 0; row < a.rows; row++ {
		for col := 0; col < a.cols; col++ {
			if row != removeRow && col != removeCol {
				result = append(result, a.At(row, col))
			}
		}
	}
	return Matrix{rows: a.rows - 1, cols: a.cols - 1, elements: result, step: a.cols - 1}
}

// Minor calculates the minor of a matrix at a row, col.
// The minor of an element at row i and column j is the determinant of the submatrix at (i,j).
func Minor(a Matrix, atRow, atCol int) float64 {
	return Determinant(Submatrix(a, atRow, atCol))
}

// Cofactor calculates the cofactor of a matrix at row, col.
func Cofactor(a Matrix, atRow, atCol int) float64 {
	minor := Minor(a, atRow, atCol)
	if atRow+atCol%2 != 0 {
		return minor * -1.0
	} else {
		return minor
	}
}

// Equal checks if to matrix are equal, slow for now since it's using reflection
func Equal(a, b Matrix) bool {
	return reflect.DeepEqual(a, b)
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

	fmt.Printf("%d-%d %v\n", rowCount, colCount, elements)
	return Matrix{rows: rowCount, cols: colCount, elements: elements, step: colCount}
}
