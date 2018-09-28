package main

import (
	"log"
	"math"
	"time"

	"github.com/sbrosinski/greytracer/pkg/render"
	"github.com/sbrosinski/greytracer/trace"
)

func main() {
	floor := trace.NewSphere()
	floor.Transform = trace.Scaling(10, 0.01, 10)
	floor.Material = trace.NewMaterial()
	floor.Material.Color = trace.Color{Red: 1, Green: 0.9, Blue: 0.9}
	floor.Material.Specular = 0.0
	floor.Material.Pattern = trace.NewGradientPattern(trace.Red, trace.White)
	//floor.Material.Pattern.Transform = trace.NewTransform().RotateZ(math.Pi/2).Scale(-10, -10, -10).Matrix()

	leftWall := trace.NewSphere()
	//leftWall.Transform = trace.Translation(0, 0, 5.).Multiply(trace.RotationY(-math.Pi / 4)).Multiply(trace.RotationX(math.Pi / 2)).Multiply(trace.Scaling(10.0, 0.01, 10.0))
	leftWall.Transform = trace.NewTransform().Translate(0, 0, 5.).RotateY(-math.Pi/4).RotateX(math.Pi/2).Scale(10.0, 0.01, 10.0).Matrix()
	leftWall.Material = floor.Material

	rightWall := trace.NewSphere()
	rightWall.Shape.Transform = trace.Translation(0, 0, 5.).Multiply(trace.RotationY(math.Pi / 4)).Multiply(trace.RotationX(math.Pi / 2)).Multiply(trace.Scaling(10.0, 0.01, 10.0))
	rightWall.Shape.Material = floor.Material

	bottom := trace.NewPlane()
	//bottom.Transform = trace.NewTransform().Scale(10, 0.01, 10).Matrix()
	bottom.Shape.Material = floor.Material
	bottom.Material.Pattern = trace.NewGradientPattern(trace.Red, trace.White)
	//bottom.Material.Pattern.Transform = trace.NewTransform().Scale(10, 0.01, 10).Matrix()

	m := trace.NewMaterial()
	m.Color = trace.Color{0.1, 1, 0.5}
	m.Diffuse = 0.7
	m.Specular = 0.2
	m.Shininess = 300.0
	m.Pattern = trace.NewGradientPattern(trace.Red, trace.White)

	middle := trace.NewSphereWithTrans(trace.Translation(-0.5, 1.0, 0.5))
	middle.Material = m

	right := trace.NewSphereWithTrans(trace.Translation(1.5, 0.5, -0.5).Multiply(trace.Scaling(0.5, 0.5, 0.5)))
	right.Material = m

	//left := trace.NewSphereWithTrans(trace.Translation(-1.5, 0.33, -0.75).Multiply(trace.Scaling(0.33, 0.33, 0.33)))
	trans := trace.NewTransform().Translate(-1.5, 0.33, -0.75).Scale(0.33, 0.33, 0.33).Matrix()

	left := trace.NewSphereWithTrans(trans)
	left.Material = m

	light := trace.Light{trace.NewPoint(-10, 10, -10), trace.Color{1, 1, 1}}

	world := trace.World{light, []trace.ShapeOps{bottom, left, middle, right}}

	camera := trace.NewCamera(300, 150, math.Pi/3)
	camera.Transform = trace.ViewTransform(trace.NewPoint(0, 1.5, -5),
		trace.NewPoint(0, 1, 0),
		trace.NewVector(0, 1, 0))

	start := time.Now()

	scene := render.Scene{}
	scene.ToFile(camera, world, "world.png")
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

}
