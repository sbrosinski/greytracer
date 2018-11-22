package trace

import (
	"math"
)

type Camera struct {
	HSize      float64
	VSize      float64
	FOV        float64
	Transform  Matrix
	HalfWidth  float64
	HalfHeight float64
	PixelSize  float64
}

func NewCamera(width, height, fov float64) Camera {
	c := Camera{HSize: width, VSize: height, FOV: fov, Transform: Identidy4x4}
	c.updatePixelSize()
	return c
}

func (c *Camera) updatePixelSize() {
	halfView := math.Tan(c.FOV / 2)
	aspect := c.HSize / c.VSize
	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2) / c.HSize
}

func (c *Camera) RayForPixel(px, py float64) Ray {
	xoffset := (px + 0.5) * c.PixelSize
	yoffset := (py + 0.5) * c.PixelSize
	worldx := c.HalfWidth - xoffset
	worldy := c.HalfHeight - yoffset
	pixel := c.Transform.Inverse().MultiplyWithTuple(NewPoint(worldx, worldy, -1))
	origin := c.Transform.Inverse().MultiplyWithTuple(NewPoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()
	return Ray{Origin: origin, Direction: direction}
}

func (c *Camera) Render(world World) Canvas {
	canvas := NewCanvas(int(c.HSize), int(c.VSize))
	for y := 0; y <= canvas.Height-1; y++ {
		for x := 0; x <= canvas.Width-1; x++ {
			ray := c.RayForPixel(float64(x), float64(y))
			color := world.ColorAt(ray)
			canvas.WritePixel(x, y, color)

		}
	}
	return canvas
}
