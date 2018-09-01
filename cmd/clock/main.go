package main

import (
	"math"

	"github.com/sbrosinski/greytracer/internal/matrix"
	"github.com/sbrosinski/greytracer/internal/trace"
	"github.com/sbrosinski/greytracer/internal/transformation"
)

var canvas = trace.NewCanvas(250, 250)

func main() {

	point := trace.NewPoint(125, 125, 0)
	rotate := transformation.RotationZ(math.Pi / 6)
	point = matrix.MultiplyWithTuple(rotate, point)
	moveUp := transformation.Translation(0, 100, 0)
	point = matrix.MultiplyWithTuple(moveUp, point)
	writeThickPixel(5, point)

	point2 := trace.NewPoint(125, 125, 0)
	rotate2 := transformation.RotationZ(2 * math.Pi / 6)
	point2 = matrix.MultiplyWithTuple(rotate2, point2)
	moveUp2 := transformation.Translation(0, 100, 0)
	point2 = matrix.MultiplyWithTuple(moveUp2, point2)
	writeThickPixel(5, point2)

	canvas.SavePNG("clock.png")
}

func writeThickPixel(size int, point trace.Tuple) {
	canvX := int(point.X)
	canvY := canvas.Height - int(point.Y)
	for x := canvX - size; x <= canvX+size; x++ {
		for y := canvY - size; y <= canvY+size; y++ {
			canvas.WritePixel(x, y, trace.White)
		}
	}
}
