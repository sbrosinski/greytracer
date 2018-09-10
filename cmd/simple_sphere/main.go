package main

import (
	"math"

	"github.com/sbrosinski/greytracer/internal/trace"
)

var canvas = trace.NewCanvas(100, 100)

func main() {
	sphere := trace.NewSphere()
	scale := trace.Scaling(1, 0.5, 1)
	rotate := trace.RotationZ(math.Pi / 4)
	sphere.Transform = trace.Multiply(rotate, scale)

	rayOrigin := trace.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	pixelSize := wallSize / float64(canvas.Width)
	half := wallSize / 2

	// for each row of pixels in the canvas
	for y := 0; y <= canvas.Height-1; y++ {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := half - pixelSize*float64(y)
		// for each pixel in the row
		for x := 0; x <= canvas.Width-1; x++ {
			//  compute the world x coordinate (left = -half, right = half)
			worldX := -half + pixelSize*float64(x)
			// describe the point on the wall that the ray will target
			position := trace.NewPoint(worldX, worldY, wallZ)
			substractedPosition := position.Subtract(rayOrigin)
			normalizedPosition := substractedPosition.Normalize()
			r := trace.NewRay(rayOrigin, normalizedPosition)
			intersections := sphere.Intersect(r)
			_, hasHit := intersections.Hit()
			if hasHit {
				canvas.WritePixel(x, y, trace.Red)
			}
		}
	}

	canvas.SavePNG("simple_sphere.png")
}
