package main

import (
	"github.com/sbrosinski/greytracer/trace"
)

var canvas = trace.NewCanvas(400, 400)

func main() {
	sphere := trace.NewSphere()
	//scale := trace.Scaling(1, 0.5, 1)
	//rotate := trace.RotationZ(math.Pi / 4)
	//sphere.Transform = rotate.Multiply(scale)
	color := trace.Color{0.3, 0.3, 1}
	sphere.Material = trace.NewMaterial()
	sphere.Material.Color = color

	lightPosition := trace.NewPoint(-10, 10, -10)
	lightColor := trace.Color{1, 1, 1}
	light := trace.Light{lightPosition, lightColor}

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
			intersection, hasHit := intersections.Hit()
			if hasHit {
				point := r.Position(intersection.T)
				normal := intersection.Object.NormalAt(point)
				eye := r.Direction
				color := intersection.Object.GetMaterial().Lighting(light, point, eye, normal)

				canvas.WritePixel(x, y, color)
			}
		}
	}

	canvas.SavePNG("simple_sphere.png")
}
