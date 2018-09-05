package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointFromDistance(t *testing.T) {
	r := Ray{NewPoint(2, 3, 4), NewVector(1, 0, 0)}
	assert.Equal(t, NewPoint(2, 3, 4), r.Position(0))
	assert.Equal(t, NewPoint(3, 3, 4), r.Position(1))
	assert.Equal(t, NewPoint(1, 3, 4), r.Position(-1))
	assert.Equal(t, NewPoint(4.5, 3, 4), r.Position(2.5))
}
