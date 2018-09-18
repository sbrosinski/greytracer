package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"time"

	"github.com/sbrosinski/greytracer/trace"
)

func main() {
	floor := trace.NewSphere()
	floor.Shape.Transform = trace.Scaling(10, 0.01, 10)
	floor.Shape.Material = trace.NewMaterial()
	floor.Shape.Material.Color = trace.Color{1, 0.9, 0.9}
	floor.Shape.Material.Specular = 0.0

	leftWall := trace.NewSphere()
	leftWall.Shape.Transform = trace.Translation(0, 0, 5.).Multiply(trace.RotationY(-math.Pi / 4)).Multiply(trace.RotationX(math.Pi / 2)).Multiply(trace.Scaling(10.0, 0.01, 10.0))
	leftWall.Shape.Material = floor.Shape.Material

	rightWall := trace.NewSphere()
	rightWall.Shape.Transform = trace.Translation(0, 0, 5.).Multiply(trace.RotationY(math.Pi / 4)).Multiply(trace.RotationX(math.Pi / 2)).Multiply(trace.Scaling(10.0, 0.01, 10.0))
	rightWall.Shape.Material = floor.Shape.Material

	bottom := trace.NewPlane()
	bottom.Shape.Material = trace.NewMaterial()
	bottom.Shape.Material.Color = trace.Color{1, 0.9, 0.9}
	bottom.Shape.Material.Specular = 0.0

	m := trace.NewMaterial()
	m.Color = trace.Color{0.1, 1, 0.5}
	m.Diffuse = 0.7
	m.Specular = 0.2
	m.Shininess = 300.0

	middle := trace.NewSphereWithTrans(trace.Translation(-0.5, 1.0, 0.5))
	middle.Shape.Material = m

	right := trace.NewSphereWithTrans(trace.Translation(1.5, 0.5, -0.5).Multiply(trace.Scaling(0.5, 0.5, 0.5)))
	right.Shape.Material = m

	//left.Transform = trace.Translation(-1.5, 0.33, -0.75).Multiply(trace.Scaling(0.33, 0.33, 0.33))
	left := trace.NewSphereWithTrans(trace.Translation(-1.5, 0.33, -0.75).Multiply(trace.Scaling(0.33, 0.33, 0.33)))
	left.Shape.Material = m

	light := trace.Light{trace.NewPoint(-10, 10, -10), trace.Color{1, 1, 1}}
	world := trace.World{light, []trace.ShapeOps{bottom, left, middle, right}}
	camera := trace.NewCamera(600., 300., math.Pi/3)
	camera.Transform = trace.ViewTransform(trace.NewPoint(0, 1.5, -5),
		trace.NewPoint(0, 1, 0),
		trace.NewVector(0, 1, 0))

	start := time.Now()
	img := camera.RenderParallel(world)
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)
	savePNG(img, "world.png")

}

func savePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}
