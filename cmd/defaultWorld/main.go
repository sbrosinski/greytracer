package main

import (
	"log"
	"math"
	"time"

	"github.com/sbrosinski/greytracer/pkg/render"
	"github.com/sbrosinski/greytracer/trace"
)

func main() {
	world := trace.NewDefaultWorld()

	camera := trace.NewCamera(200, 150, math.Pi/3)
	camera.Transform = trace.ViewTransform(trace.NewPoint(0, 1.5, -5),
		trace.NewPoint(0, 0, 0),
		trace.NewVector(0, 1, 0))

	start := time.Now()

	scene := render.Scene{}
	scene.ToFile(camera, world, "defaultWorld.png")
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

}
