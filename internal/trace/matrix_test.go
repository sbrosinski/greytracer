package trace

import (
	"testing"

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
	a := Parse(`
		| 1 | 2 | 3 | 4 |
		| 2 | 3 | 4 | 5 |
		| 3 | 4 | 5 | 6 |
		| 4 | 5 | 6 | 7 |
	`)
	b := Parse(`
		| 0 | 1 | 2  | 4  |
		| 1 | 2 | 4  | 8  |
		| 2 | 4 | 8  | 16 |
		| 4 | 8 | 16 | 32 |
	`)
	expected := Parse(`
		| 24 | 49 | 98  | 196 |
		| 31 | 64 | 128 | 256 |
		| 38 | 79 | 158 | 316 |
		| 45 | 94 | 188 | 376 |
	`)
	result := a.Multiply(b)
	assert.Equal(t, expected, result)
}

func TestMultiplyWithTuple(t *testing.T) {
	a := Parse(`
		| 1 | 2 | 3 | 4 |
		| 2 | 4 | 4 | 2 |
		| 8 | 6 | 4 | 1 |
		| 0 | 0 | 0 | 1 |
	`)
	tuple := Tuple{X: 1, Y: 2, Z: 3, W: 1}
	result := a.MultiplyWithTuple(tuple)
	assert.Equal(t, Tuple{X: 18, Y: 24, Z: 33, W: 1}, result)
}

func TestMultiplyingMatrixByIdentity(t *testing.T) {
	a := Parse(`
		| 0 | 1 | 2 | 4 |
		| 1 | 2 | 4 | 8 |
		| 2 | 4 | 8 | 16 |
		| 4 | 8 | 16 | 32 |
	`)
	result := a.Multiply(Identidy4x4)
	assert.Equal(t, a, result)
}

func TestMultiplyingIdentityByTuple(t *testing.T) {
	tuple := Tuple{X: 1, Y: 2, Z: 3, W: 4}
	result := Identidy4x4.MultiplyWithTuple(tuple)
	assert.Equal(t, tuple, result)
}

func TestTransposingMatrix(t *testing.T) {
	a := Parse(`
		| 0 | 9 | 3 | 0 |
		| 9 | 8 | 0 | 8 |
		| 1 | 8 | 5 | 3 |
		| 0 | 0 | 5 | 8 |
	`)
	expected := Parse(`
		| 0 | 9 | 1 | 0 |
		| 9 | 8 | 8 | 0 |
		| 3 | 0 | 5 | 5 |
		| 0 | 8 | 3 | 8 |
	`)
	result := a.Transpose()
	assert.Equal(t, expected, result)
}

func TestTransposingIdentiyMatrix(t *testing.T) {
	result := Identidy4x4.Transpose()
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
	a := Parse(`|  1 | 5 |
		  		| -3 | 2 |`)
	assert.Equal(t, 17.0, a.Determinant())
}

func TestSubmatrixOf3x3MatrixIs2x2Matrix(t *testing.T) {
	a := Parse(`
		| 1  | 5 | 0  |
		| -3 | 2 | 7  |
		| 0  | 6 | -3 |`)
	expected := Parse(`
		| -3 | 2 |
		| 0  | 6 |`)
	result := a.Submatrix(0, 2)
	assert.Equal(t, expected, result)
}

func TestSubmatrixOf4x4MatrixIs3x3Matrix(t *testing.T) {
	a := Parse(`
		| -6 | 1 | 1  | 6 |
		| -8 | 5 | 8  | 6 |
		| -1 | 0 | 8  | 2 |
		| -7 | 1 | -1 | 1 |`)
	expected := Parse(`
		| -6 | 1  | 6 |
		| -8 | 8  | 6 |
		| -7 | -1 | 1 |`)
	result := a.Submatrix(2, 1)
	assert.Equal(t, expected, result)
}

func TestCalculatingMinorOf3x3Matrix(t *testing.T) {
	a := Parse(`
		| 3 | 5  |  0 |
		| 2 | -1 | -7 |
		| 6 | -1 |  5 |`)
	b := a.Submatrix(1, 0)
	assert.Equal(t, 25.0, b.Determinant())
	assert.Equal(t, 25.0, a.Minor(1, 0))
}

func TestCalculatingCofactorOf3x3Matrix(t *testing.T) {
	a := Parse(`
		| 3 |  5 |  0 |
		| 2 | -1 | -7 |
		| 6 | -1 |  5 |`)
	assert.Equal(t, -12.0, a.Minor(0, 0))
	assert.Equal(t, -12.0, a.Cofactor(0, 0))
	assert.Equal(t, 25.0, a.Minor(1, 0))
	assert.Equal(t, -25.0, a.Cofactor(1, 0))
	assert.Equal(t, 15.0, a.Minor(1, 1))
	assert.Equal(t, 15.0, a.Cofactor(1, 1))
}

func TestCalculatingDeterminantOf3x3Matrix(t *testing.T) {
	a := Parse(`
		|  1 | 2 |  6|
		| -5 | 8 | -4|
		|  2 | 6 |  4|`)
	assert.Equal(t, 56.0, a.Cofactor(0, 0))
	assert.Equal(t, 12.0, a.Cofactor(0, 1))
	assert.Equal(t, -46.0, a.Cofactor(0, 2))
	assert.Equal(t, -196.0, a.Determinant())
}

func TestCalculatingDeterminantOf4x4Matrix(t *testing.T) {
	a := Parse(`
		|-2|-8| 3| 5|
		|-3| 1| 7| 3|
		| 1| 2|-9| 6|
		|-6| 7| 7|-9|
	`)
	assert.Equal(t, 690.0, a.Cofactor(0, 0))
	assert.Equal(t, 447.0, a.Cofactor(0, 1))
	assert.Equal(t, 210.0, a.Cofactor(0, 2))
	assert.Equal(t, 51.0, a.Cofactor(0, 3))
	assert.Equal(t, -4071.0, a.Determinant())
}

func TestInvertibleMatrix(t *testing.T) {
	a := Parse(`
    	|  6 |  4 |  4 |  4 |
    	|  5 |  5 |  7 |  6 |
    	|  4 | -9 |  3 | -7 |
		|  9 |  1 |  7 | -6 |`)
	assert.Equal(t, -2120.0, a.Determinant())
	assert.True(t, a.IsInvertible())
}

func TestNonInvertibleMatrix(t *testing.T) {
	a := Parse(`
    | -4 |  2 | -2 | -3 |
    |  9 |  6 |  2 |  6 |
    |  0 | -5 |  1 | -5 |
    |  0 |  0 |  0 |  0 |`)
	assert.Equal(t, 0.0, a.Determinant())
	assert.False(t, a.IsInvertible())
}

func TestCalculatingInverseOfMatrix(t *testing.T) {
	a := Parse(`      
		| -5 |  2 |  6 | -8 |
		|  1 | -5 |  1 |  8 |
		|  7 |  7 | -6 | -7 |
		|  1 | -3 |  7 |  4 |`)
	b := a.Inverse()
	expected := Parse(`
		|  0.21805 |  0.45113 |  0.24060 | -0.04511 |
		| -0.80827 | -1.45677 | -0.44361 |  0.52068 |
		| -0.07895 | -0.22368 | -0.05263 |  0.19737 |
		| -0.52256 | -0.81391 | -0.30075 |  0.30639 |
	`)
	assert.Equal(t, 532.0, a.Determinant())
	assert.Equal(t, -160.0, a.Cofactor(2, 3))
	assert.Equal(t, -160.0/532.0, b.At(3, 2))
	assert.Equal(t, 105.0, a.Cofactor(3, 2))
	assert.Equal(t, 105.0/532.0, b.At(2, 3))
	assert.True(t, expected.Equal(b))
}

func TestCalculatingInverseOfMatrix2(t *testing.T) {
	a := Parse(`      
		|  8 | -5 |  9 |  2 |
		|  7 |  5 |  6 |  1 |
		| -6 |  0 |  9 |  6 |
		| -3 |  0 | -9 | -4 |`)
	b := a.Inverse()
	expected := Parse(`
		| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
		| -0.07692 |  0.12308 |  0.02564 |  0.03077 |
		|  0.35897 |  0.35897 |  0.43590 |  0.92308 |
		| -0.69231 | -0.69231 | -0.76923 | -1.92308 |
	`)
	assert.True(t, expected.Equal(b))
}

func TestCalculatingInverseOfMatrix3(t *testing.T) {
	a := Parse(`      
		|  9 |  3 |  0 |  9 |
		| -5 | -2 | -6 | -3 |
		| -4 |  9 |  6 |  4 |
		| -7 |  6 |  6 |  2 |`)
	b := a.Inverse()
	expected := Parse(`
		| -0.04074 | -0.07778 |  0.14444 | -0.22222 |
		| -0.07778 |  0.03333 |  0.36667 | -0.33333 |
		| -0.02901 | -0.14630 | -0.10926 |  0.12963 |
		|  0.17778 |  0.06667 | -0.26667 |  0.33333 |
	`)
	assert.True(t, expected.Equal(b))
}

func TestMultiplyingAProductByItsInverse(t *testing.T) {
	a := Parse(`
		|  3 | -9 |  7 |  3 |
		|  3 | -8 |  2 | -9 |
		| -4 |  4 |  4 |  1 |
		| -6 |  5 | -1 |  1 |`)
	b := Parse(`
		|  8 |  2 |  2 |  2 |
		|  3 | -1 |  7 |  0 |
		|  7 |  0 |  5 |  4 |
		|  6 | -2 |  0 |  5 |`)
	c := a.Multiply(b)
	assert.True(t, a.Equal(c.Multiply(b.Inverse())))
}
