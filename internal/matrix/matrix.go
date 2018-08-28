package matrix

import (
	"fmt"
	"reflect"
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
