package trace

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

type Camera struct {
	HSize      float64
	VSize      float64
	FOV        float64
	Transform  Matrix
	HalfWidth  float64
	HalfHeight float64
	PixelSize  float64
}

func NewCamera(width, height, fov float64) Camera {
	c := Camera{HSize: width, VSize: height, FOV: fov, Transform: Identidy4x4}
	c.updatePixelSize()
	return c
}

func (c *Camera) updatePixelSize() {
	halfView := math.Tan(c.FOV / 2)
	aspect := c.HSize / c.VSize
	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2) / c.HSize
}

func (c *Camera) RayForPixel(px, py float64) Ray {
	xoffset := (px + 0.5) * c.PixelSize
	yoffset := (py + 0.5) * c.PixelSize
	worldx := c.HalfWidth - xoffset
	worldy := c.HalfHeight - yoffset
	pixel := c.Transform.Inverse().MultiplyWithTuple(NewPoint(worldx, worldy, -1))
	origin := c.Transform.Inverse().MultiplyWithTuple(NewPoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()
	return Ray{Origin: origin, Direction: direction}
}

func (c *Camera) Render(world World) image.Image {
	canvas := NewCanvas(int(c.HSize), int(c.VSize))
	for y := 0; y <= canvas.Height-1; y++ {
		for x := 0; x <= canvas.Width-1; x++ {
			ray := c.RayForPixel(float64(x), float64(y))
			color := world.ColorAt(ray)
			canvas.WritePixel(x, y, color)

		}
	}
	return canvas.ToImage()
}

type RenderJob struct {
	startX, endX, startY, endY int
	canvas                     Canvas
}

func (c *Camera) GenerateRenderJobs(jobWidth, jobHeight int) []RenderJob {
	jobs := []RenderJob{}
	for y := 0; y <= int(c.VSize-1); y += jobHeight {
		for x := 0; x <= int(c.HSize-1); x += jobWidth {
			jobs = append(jobs, RenderJob{x, x + jobWidth - 1, y, y + jobHeight - 1, NewCanvas(jobWidth, jobHeight)})
		}
	}
	return jobs
}

func (c *Camera) RenderJob(world World, job RenderJob, results chan<- RenderJob) {
	for y := job.startY; y <= job.endY; y++ {
		for x := job.startX; x <= job.endX; x++ {
			ray := c.RayForPixel(float64(x), float64(y))
			color := world.ColorAt(ray)
			job.canvas.WritePixel(x-job.startX, y-job.startY, color)

		}
	}
	results <- job
}

func (c *Camera) RenderParallel(world World) image.Image {
	renderTasks := c.GenerateRenderJobs(int(c.HSize/4), int(c.VSize/4))

	results := make(chan RenderJob, 100)

	for _, task := range renderTasks {
		go c.RenderJob(world, task, results)
	}

	// Collect results
	canvas := NewCanvas(int(c.HSize), int(c.VSize))
	for i, _ := range renderTasks {
		fmt.Printf("Collecting %d", i)

		finishedJob := <-results

		for y := finishedJob.startY; y <= finishedJob.endY; y++ {
			for x := finishedJob.startX; x <= finishedJob.endX; x++ {
				canvas.WritePixel(x, y, finishedJob.canvas.PixelAt(x-finishedJob.startX, y-finishedJob.startY))
			}
		}

	}
	return canvas.ToImage()
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
