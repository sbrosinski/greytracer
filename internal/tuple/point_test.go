package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingTwoPoints(t *testing.T) {
	var p1 = NewPoint(3, -2, 5)
	var p2 = NewPoint(-2, 3, 1)

	var result = Add(p1, p2)

	var expected = NewPoint(1, 1, 6)
	assert.Equal(t, expected, result)
}
