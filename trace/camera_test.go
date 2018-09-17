package trace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPixelSizeHorizontal(t *testing.T) {
	assert.InDelta(t, 0.01, NewCamera(200, 125, math.Pi/2).PixelSize, 0.00001)
}

func TestPixelSizeVertical(t *testing.T) {
	assert.InDelta(t, 0.01, NewCamera(125, 200, math.Pi/2).PixelSize, 0.00001)
}

func TestGenerateRenderJobs(t *testing.T) {

	c := NewCamera(600, 300, math.Pi/2)
	jobs := c.GenerateRenderJobs(200, 100)
	for _, job := range jobs {
		t.Logf("%d %d %d %d\n", job.startX, job.startY, job.endX, job.endY)
	}

	t.Fail()
}
