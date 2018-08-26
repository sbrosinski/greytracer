package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritingPixelsToCanvas(t *testing.T) {

	c := NewCanvas(10, 20)
	red := Color{1, 0, 0}

	c.writePixel(2, 3, red)
	assert.Equal(t, red, c.pixelAt(2, 3))

}
