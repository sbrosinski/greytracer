package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	var c = Color{-0.5, 0.4, 1.7}
	assert.Equal(t, -0.5, c.Red)
	assert.Equal(t, 0.4, c.Green)
	assert.Equal(t, 1.7, c.Blue)
}
