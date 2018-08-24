package main

import (
	"fmt"

	"github.com/sbrosinski/graytracer/internal/trace"
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
	// projectile starts one unit above the origin.
	// velocity is normalized to 1 unit/tick.
	// p ← projectile(point(0, 1, 0), normalize(vector(1, 1, 0)))
	var vel = trace.NewVector(1, 1, 0)
	var p = Projectile{trace.NewPoint(0, 1, 0), vel.Normalize()}

	// world gravity -0.1 unit/tick, and wind is -0.01 unit/tick.
	// w ← world(vector(0, -0.1, 0), vector(-0.01, 0, 0))
	var w = World{trace.NewVector(0, -0.1, 0), trace.NewVector(-0.01, 0, 0)}

	fmt.Println("Starting projectile ...")
	fmt.Printf("x=%f - y=%f", p.Position.X, p.Position.Y)
	for p.Position.Y <= 0 {
		p = tick(w, p)
		fmt.Printf("x=%f - y=%f\n", p.Position.X, p.Position.Y)
	}
	fmt.Println("Done")
}
