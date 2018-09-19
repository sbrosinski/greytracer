package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTasks(t *testing.T) {
	tasks := generateTasks(300, 150, 2)
	assert.Equal(t, 4, len(tasks))
	assert.True(t, NewTask(0, 0, 149, 74).Equal(tasks[0]))
	assert.True(t, NewTask(150, 0, 299, 74).Equal(tasks[1]))
	assert.True(t, NewTask(0, 75, 149, 149).Equal(tasks[2]))
	assert.True(t, NewTask(150, 75, 299, 149).Equal(tasks[3]))

	tasks2 := generateTasks(300, 150, 4)
	assert.Equal(t, 16, len(tasks2))
	assert.True(t, NewTask(0, 0, 74, 37).Equal(tasks2[0]))
	assert.True(t, NewTask(75, 0, 149, 37).Equal(tasks2[1]))
	assert.True(t, NewTask(150, 0, 224, 37).Equal(tasks2[2]))
	assert.True(t, NewTask(225, 0, 299, 37).Equal(tasks2[3]))
	assert.True(t, NewTask(0, 38, 74, 75).Equal(tasks2[4]))
	assert.True(t, NewTask(75, 38, 149, 75).Equal(tasks2[5]))
	assert.True(t, NewTask(150, 38, 224, 75).Equal(tasks2[6]))
	assert.True(t, NewTask(225, 38, 299, 75).Equal(tasks2[7]))
	assert.True(t, NewTask(0, 76, 74, 113).Equal(tasks2[8]))
	assert.True(t, NewTask(75, 76, 149, 113).Equal(tasks2[9]))
	assert.True(t, NewTask(150, 76, 224, 113).Equal(tasks2[10]))
	assert.True(t, NewTask(225, 76, 299, 113).Equal(tasks2[11]))
	assert.True(t, NewTask(0, 114, 74, 150).Equal(tasks2[12]))
	assert.True(t, NewTask(75, 114, 149, 150).Equal(tasks2[13]))
	assert.True(t, NewTask(150, 114, 224, 150).Equal(tasks2[14]))
	assert.True(t, NewTask(225, 114, 299, 150).Equal(tasks2[15]))
}
