package trace

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

// Canvas describes the rendering surface
type Canvas struct {
	Width  int
	Height int
	Data   [][]Color
}

// NewCanvas creates a canvas
func NewCanvas(width int, height int) Canvas {
	data := make([][]Color, height)
	for i := range data {
		data[i] = make([]Color, width)
	}
	return Canvas{width, height, data}
}

// ToPPM converts canvas to PPM format
func (c *Canvas) ToPPM() string {
	var ppm strings.Builder
	ppm.WriteString("P3\n")
	ppm.WriteString(fmt.Sprintf("%d %d\n", c.Width, c.Height))
	ppm.WriteString("255\n")
	ppm.WriteString(`"""`)
	ppm.WriteString("\n")

	for y := range c.Data {
		var line strings.Builder
		leadingSpace := ""
		for x := range c.Data[y] {
			color := c.Data[y][x]
			red, green, blue := color.Normalize()
			line.WriteString(fmt.Sprintf("%s%d %d %d", leadingSpace, red, green, blue))
			leadingSpace = " "
			if line.Len() > 70 {
				leadingSpace = "\n"
				ppm.WriteString(line.String())
				line.Reset()
			}
		}
		ppm.WriteString(line.String())
		ppm.WriteString("\n")
	}
	ppm.WriteString(`"""`)
	ppm.WriteString("\n")
	return ppm.String()
}

// ToImage generates an Image from Canvas
func (c *Canvas) ToImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for y := range c.Data {
		for x := range c.Data[y] {
			canvasColor := c.Data[y][x]
			red, green, blue := canvasColor.Normalize()
			img.Set(x, y, color.RGBA{uint8(red), uint8(green), uint8(blue), 255})
		}
	}
	return img
}

// WritePixel sets a pixel on the canvas to a color
func (c *Canvas) WritePixel(x int, y int, color Color) {
	if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
		c.Data[y][x] = color
	} else {
		//panic(fmt.Errorf("%d, %d is outside canvas %d,%d", x, y, c.Width, c.Height))
	}
}

// PixelAt gets the color of a canvas at x/y
func (c *Canvas) PixelAt(x int, y int) Color {
	return c.Data[x][y]
}

func (c *Canvas) String() string {
	var out strings.Builder
	for y := range c.Data {
		out.WriteString(fmt.Sprintf("row(%v)  -  %v\n", y, c.Data[y]))
	}
	return out.String()
}
