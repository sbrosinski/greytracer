package matrix

import (
	"testing"

	"github.com/sbrosinski/greytracer/internal/trace"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	str := `
	 | 1    | 2    | 3    | 4    |
	 | 5.5  | 6.5  | 7.5  | 8.5  |
	 | 9    | 10   | 11   | 12   |
	 | 13.5 | 14.5 | 15.5 | 16.5 |
	`
	matrix := Parse(str)
	assert.Equal(t, 4, matrix.rows)
	assert.Equal(t, 4, matrix.cols)
	assert.Equal(t, 1.0, matrix.At(0, 0))
	assert.Equal(t, 4.0, matrix.At(0, 3))
	assert.Equal(t, 5.5, matrix.At(1, 0))
	assert.Equal(t, 7.5, matrix.At(1, 2))
	assert.Equal(t, 11.0, matrix.At(2, 2))
	assert.Equal(t, 13.5, matrix.At(3, 0))
	assert.Equal(t, 15.5, matrix.At(3, 2))
}

func TestMultiply(t *testing.T) {
	a := `
		| 1 | 2 | 3 | 4 |
		| 2 | 3 | 4 | 5 |
		| 3 | 4 | 5 | 6 |
		| 4 | 5 | 6 | 7 |
	`
	b := `
		| 0 | 1 | 2  | 4  |
		| 1 | 2 | 4  | 8  |
		| 2 | 4 | 8  | 16 |
		| 4 | 8 | 16 | 32 |
	`
	expected := `
		| 24 | 49 | 98  | 196 |
		| 31 | 64 | 128 | 256 |
		| 38 | 79 | 158 | 316 |
		| 45 | 94 | 188 | 376 |
	`
	result := Multiply(Parse(a), Parse(b))
	assert.Equal(t, Parse(expected), result)
}

func TestMultiplyWithTuple(t *testing.T) {
	a := `
		| 1 | 2 | 3 | 4 |
		| 2 | 4 | 4 | 2 |
		| 8 | 6 | 4 | 1 |
		| 0 | 0 | 0 | 1 |
	`
	tuple := trace.Tuple{X: 1, Y: 2, Z: 3, W: 1}
	result := MultiplyWithTuple(Parse(a), tuple)
	assert.Equal(t, trace.Tuple{X: 18, Y: 24, Z: 33, W: 1}, result)
}

func TestMultiplyingMatrixByIdentity(t *testing.T) {
	a := `
		| 0 | 1 | 2 | 4 |
		| 1 | 2 | 4 | 8 |
		| 2 | 4 | 8 | 16 |
		| 4 | 8 | 16 | 32 |
	`
	result := Multiply(Parse(a), Identidy4x4)
	assert.Equal(t, Parse(a), result)
}

func TestMultiplyingIdentityByTuple(t *testing.T) {
	tuple := trace.Tuple{X: 1, Y: 2, Z: 3, W: 4}
	result := MultiplyWithTuple(Identidy4x4, tuple)
	assert.Equal(t, tuple, result)
}

func TestTransposingMatrix(t *testing.T) {
	a := `
		| 0 | 9 | 3 | 0 |
		| 9 | 8 | 0 | 8 |
		| 1 | 8 | 5 | 3 |
		| 0 | 0 | 5 | 8 |
	`
	expected := `
		| 0 | 9 | 1 | 0 |
		| 9 | 8 | 8 | 0 |
		| 3 | 0 | 5 | 5 |
		| 0 | 8 | 3 | 8 |
	`
	result := Transpose(Parse(a))
	assert.Equal(t, Parse(expected), result)
}

func TestTransposingIdentiyMatrix(t *testing.T) {
	result := Transpose(Identidy4x4)
	assert.Equal(t, Identidy4x4, result)
}

func Test2x2Matrix(t *testing.T) {
	a := `|  1 | 5 |
		  | -2 | 3 |`
	matrix := Parse(a)
	assert.Equal(t, 2, matrix.rows)
	assert.Equal(t, 2, matrix.cols)
	assert.Equal(t, 1.0, matrix.At(0, 0))
	assert.Equal(t, 5.0, matrix.At(0, 1))
	assert.Equal(t, -2.0, matrix.At(1, 0))
	assert.Equal(t, 3.0, matrix.At(1, 1))
}

func TestCalculatingDeterminantOf2x2Matrix(t *testing.T) {
	a := `|  1 | 5 |
		  | -3 | 2 |`
	assert.Equal(t, 17.0, Determinant(Parse(a)))
}

func TestSubmatrixOf3x3MatrixIs2x2Matrix(t *testing.T) {
	a := `
		| 1  | 5 | 0  |
		| -3 | 2 | 7  |
		| 0  | 6 | -3 |`
	expected := `
		| -3 | 2 |
		| 0  | 6 |`
	result := Submatrix(Parse(a), 0, 2)
	assert.Equal(t, Parse(expected), result)
}

func TestSubmatrixOf4x4MatrixIs3x3Matrix(t *testing.T) {
	a := `
		| -6 | 1 | 1  | 6 |
		| -8 | 5 | 8  | 6 |
		| -1 | 0 | 8  | 2 |
		| -7 | 1 | -1 | 1 |`
	expected := `
		| -6 | 1  | 6 |
		| -8 | 8  | 6 |
		| -7 | -1 | 1 |`
	result := Submatrix(Parse(a), 2, 1)
	assert.Equal(t, Parse(expected), result)
}

func TestCalculatingMinorOf3x3Matrix(t *testing.T) {
	a := `
		| 3 | 5  |  0 |
		| 2 | -1 | -7 |
		| 6 | -1 |  5 |`
	b := Submatrix(Parse(a), 1, 0)
	assert.Equal(t, 25.0, Determinant(b))
	assert.Equal(t, 25.0, Minor(Parse(a), 1, 0))
}

func TestCalculatingCofactorOf3x3Matrix(t *testing.T) {
	a := Parse(`
		| 3 |  5 |  0 |
		| 2 | -1 | -7 |
		| 6 | -1 |  5 |`)
	assert.Equal(t, -12.0, Minor(a, 0, 0))
	assert.Equal(t, -12.0, Cofactor(a, 0, 0))
	assert.Equal(t, 25.0, Minor(a, 1, 0))
	assert.Equal(t, -25.0, Cofactor(a, 1, 0))
}
