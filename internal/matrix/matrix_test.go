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
	c := `
		| 24 | 49 | 98  | 196 |
		| 31 | 64 | 128 | 256 |
		| 38 | 79 | 158 | 316 |
		| 45 | 94 | 188 | 376 |
	`
	result := Multiply(Parse(a), Parse(b))
	assert.Equal(t, Parse(c), result)
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
