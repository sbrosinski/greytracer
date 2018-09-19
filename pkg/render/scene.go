package render

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/sbrosinski/greytracer/trace"
)

type Scene struct {
}

func (s *Scene) ToFile(camera trace.Camera, world trace.World, path string) {
	log.Printf("Rendering camera %+v\n", camera)
	tasks := generateTasks(
		camera.HSize,
		camera.VSize,
		4,
	)
	results := make(chan Task, 100)

	for _, task := range tasks {
		go renderTask(camera, world, task, results)
	}

	// Collect results
	canvas := trace.NewCanvas(int(camera.HSize), int(camera.VSize))
	for i, _ := range tasks {
		fmt.Printf("Collecting %d", i)

		finishedJob := <-results

		for y := finishedJob.startY; y <= finishedJob.endY; y++ {
			for x := finishedJob.startX; x <= finishedJob.endX; x++ {
				canvas.WritePixel(int(x), int(y), finishedJob.canvas.PixelAt(int(x-finishedJob.startX), int(y-finishedJob.startY)))
			}
		}

	}
	img := canvas.ToImage()
	savePNG(img, path)
}

func renderTask(camera trace.Camera, world trace.World, job Task, results chan<- Task) {
	for y := job.startY; y <= job.endY; y++ {
		for x := job.startX; x <= job.endX; x++ {
			ray := camera.RayForPixel(float64(x), float64(y))
			color := world.ColorAt(ray)
			job.canvas.WritePixel(int(x-job.startX), int(y-job.startY), color)

		}
	}
	results <- job
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
