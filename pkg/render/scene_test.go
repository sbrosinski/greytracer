package render

import (
	"math"
	"testing"

	"github.com/sbrosinski/greytracer/trace"
)

func TestRenderingToFile(t *testing.T) {
	world := trace.NewDefaultWorld()
	camera := trace.NewCamera(75, 100, math.Pi/3)
	camera.Transform = trace.ViewTransform(trace.NewPoint(0, 1.5, -5),
		trace.NewPoint(0, 1, 0),
		trace.NewVector(0, 1, 0))

	scene := Scene{}
	scene.ToFile(camera, world, "test.png")

}
