package trace

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanvasCreation(t *testing.T) {
	t.Log("test")
	data := NewCanvas(5, 3).Data
	for y := range data {
		t.Logf("row(%v)  -  %v\n", y, data[y])
	}
}

func TestWritingPixelsToCanvas(t *testing.T) {

	c := NewCanvas(10, 20)
	red := Color{1, 0, 0}

	c.WritePixel(2, 3, red)
	assert.Equal(t, red, c.PixelAt(2, 3))

}

func TestConstructingPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := c.ToPPM()
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, "P3", lines[0])
	assert.Equal(t, "5 3", lines[1])
	assert.Equal(t, "255", lines[2])
	assert.Equal(t, `"""`, lines[3])
}

func TestConstructingPPMPixelData(t *testing.T) {
	c := NewCanvas(5, 3)
	c.WritePixel(0, 0, Color{1.5, 0, 0})
	c.WritePixel(2, 1, Color{0, 0.5, 0})
	c.WritePixel(4, 2, Color{-0.5, 0, 1})
	ppm := c.ToPPM()
	t.Log(ppm)
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, `"""`, lines[3])
	assert.Equal(t, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", lines[4])
	assert.Equal(t, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", lines[5])
	assert.Equal(t, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", lines[6])
	assert.Equal(t, `"""`, lines[7])
}

func TestSplittingLongLinesInPPMFiles(t *testing.T) {
	c := NewCanvas(10, 2)

	for x := 0; x < 10; x++ {
		for y := 0; y < 2; y++ {
			c.WritePixel(x, y, Color{1, 0.8, 0.6})
		}
	}
	ppm := c.ToPPM()
	t.Log(ppm)
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, `"""`, lines[3])
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153", lines[4])
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153", lines[5])
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153", lines[6])
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153", lines[7])
	assert.Equal(t, `"""`, lines[8])

}
