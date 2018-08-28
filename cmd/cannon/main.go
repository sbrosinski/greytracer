package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/sbrosinski/greytracer/internal/trace"
)

// Projectile has a position (a point) and a velocity (a vector)
type Projectile struct {
	Position trace.Tuple
	Velocity trace.Tuple
}

// World has gravity (a vector) and wind (a vector).
type World struct {
	Gravity trace.Tuple
	Wind    trace.Tuple
}

func tick(world World, p Projectile) Projectile {
	var pos = p.Position.Add(p.Velocity)
	var velocity = p.Velocity.Add(world.Gravity)
	velocity = velocity.Add(world.Wind)
	return Projectile{pos, velocity}
}

func main() {
	var canvas = trace.NewCanvas(900, 550)
	// projectile starts one unit above the origin.
	// velocity is normalized to 1 unit/tick.
	// p ← projectile(point(0, 1, 0), normalize(vector(1, 1, 0)))
	var vel = trace.NewVector(1, 1.8, 0)
	velNorm := vel.Normalize()
	var p = Projectile{trace.NewPoint(0, 1, 0), velNorm.Multiply(11.25)}

	// world gravity -0.1 unit/tick, and wind is -0.01 unit/tick.
	// w ← world(vector(0, -0.1, 0), vector(-0.01, 0, 0))
	var w = World{trace.NewVector(0, -0.1, 0), trace.NewVector(-0.01, 0, 0)}

	fmt.Println("Starting projectile ...")
	fmt.Printf("x=%f - y=%f", p.Position.X, p.Position.Y)
	var count = 0
	for p.Position.Y >= 0 {
		p = tick(w, p)
		fmt.Printf("%d - x=%f - y=%f\n", count, p.Position.X, p.Position.Y)
		count++

		canvX := int(p.Position.X)
		canvY := canvas.Height - int(p.Position.Y)
		red := trace.Color{Red: 1, Green: 0, Blue: 0}
		canvas.WritePixel(canvX, canvY, red)
		canvas.WritePixel(canvX-1, canvY, red)
		canvas.WritePixel(canvX+1, canvY, red)
		canvas.WritePixel(canvX, canvY-1, red)
		canvas.WritePixel(canvX, canvY+1, red)

	}
	fmt.Println("Done")

	f, err := os.Create("cannon.png")
	check(err)
	defer f.Close()
	err = png.Encode(f, canvas.ToImage())
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
