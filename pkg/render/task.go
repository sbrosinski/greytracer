package render

import (
	"fmt"
	"math"

	"github.com/sbrosinski/greytracer/trace"
)

type Task struct {
	startX, endX, startY, endY float64
	canvas                     trace.Canvas
}

func (t Task) Equal(a Task) bool {
	return t.startX == a.startX && t.startY == a.startY && t.endX == a.endX && t.endY == a.endY
}

func (t Task) String() string {
	return fmt.Sprintf("%.0f,%.0f,%.0f,%.0f\n", t.startX, t.startY, t.endX, t.endY)
}

func NewTask(startX, startY, endX, endY float64) Task {
	return Task{startX: startX, startY: startY, endX: endX, endY: endY}
}

func generateTasks(totalWidth, totalHeight, taskCount float64) []Task {
	taskWidth := math.Ceil(totalWidth / taskCount)
	taskHeight := math.Ceil(totalHeight / taskCount)
	tasks := []Task{}
	for y := 0.0; y <= totalHeight-1.0; y += taskHeight {
		for x := 0.0; x <= totalWidth-1.0; x += taskWidth {
			task := Task{
				x,
				math.Min(x+taskWidth-1, totalWidth-1),
				y,
				math.Min(y+taskHeight-1, totalHeight-1),
				trace.NewCanvas(int(taskWidth), int(taskHeight)),
			}
			tasks = append(tasks, task)
		}
	}
	return tasks
}
