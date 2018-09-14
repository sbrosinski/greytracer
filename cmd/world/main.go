package main

import (
	"image"
	"image/png"
	"math"
	"os"

	"github.com/sbrosinski/greytracer/trace"
)

func main() {
	floor := trace.NewSphere()
	floor.Transform = trace.Scaling(10, 0.01, 10)
	floor.Material = trace.NewMaterial()
	floor.Material.Color = trace.Color{1, 0.9, 0.9}
	floor.Material.Specular = 0.0

	leftWall := trace.NewSphere()
	leftWall.Transform = trace.Translation(0, 0, 5.).Multiply(trace.RotationY(-math.Pi / 4)).Multiply(trace.RotationX(math.Pi / 2)).Multiply(trace.Scaling(10.0, 0.01, 10.0))
	leftWall.Material = floor.Material

	rightWall := trace.NewSphere()
	rightWall.Transform = trace.Translation(0, 0, 5.).Multiply(trace.RotationY(math.Pi / 4)).Multiply(trace.RotationX(math.Pi / 2)).Multiply(trace.Scaling(10.0, 0.01, 10.0))
	rightWall.Material = floor.Material

	middle := trace.NewSphere()
	middle.Transform = trace.Translation(-0.5, 1.0, 0.5)
	middle.Material = trace.NewMaterial()
	middle.Material.Color = trace.Color{0.1, 1, 0.5}
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := trace.NewSphere()
	right.Transform = trace.Translation(1.5, 0.5, -0.5).Multiply(trace.Scaling(0.5, 0.5, 0.5))
	right.Material = trace.NewMaterial()
	right.Material.Color = trace.Color{1., 0.8, 0.1}
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := trace.NewSphere()
	left.Transform = trace.Translation(-1.5, 0.33, -0.75).Multiply(trace.Scaling(0.33, 0.33, 0.33))
	left.Material = trace.NewMaterial()
	left.Material.Color = trace.Color{1., 0.8, 0.1}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	light := trace.Light{trace.NewPoint(-10, 10, -10), trace.Color{1, 1, 1}}
	world := trace.World{light, []trace.SceneObject{floor, leftWall, rightWall, middle, right, left}}
	camera := trace.NewCamera(600., 300., math.Pi/3)
	camera.Transform = trace.ViewTransform(trace.NewPoint(0, 1.5, -5),
		trace.NewPoint(0, 1, 0),
		trace.NewVector(0, 1, 0))

	img := camera.Render(world)
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
