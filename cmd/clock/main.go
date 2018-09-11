package main

import (
	"math"

	"github.com/sbrosinski/greytracer/trace"
)

var canvas = trace.NewCanvas(250, 250)

func main() {

	for hour := 1; hour <= 12; hour++ {
		rotate := trace.RotationY(float64(hour) * math.Pi / 6)
		twelve := trace.NewPoint(0, 0, 1)
		hourPoint := rotate.MultiplyWithTuple(twelve)
		writeThickPixe(3, int(hourPoint.X*3/8*float64(canvas.Width)+125.0), int(hourPoint.Z*3/8*float64(canvas.Width)+125))
	}

	canvas.SavePNG("clock.png")
}

func writeThickPixe(size, canvX, canvY int) {
	for x := canvX - size; x <= canvX+size; x++ {
		for y := canvY - size; y <= canvY+size; y++ {
			canvas.WritePixel(x, y, trace.White)
		}
	}
}
