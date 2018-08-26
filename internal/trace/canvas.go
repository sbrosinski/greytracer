package trace

// Canvas describes the rendering surface
type Canvas struct {
	Width int
	Hight int
	Data  [][]Color
}

func (c *Canvas) writePixel(x int, y int, color Color) {
	c.Data[x][y] = color
}

func (c *Canvas) pixelAt(x int, y int) Color {
	return c.Data[x][y]
}

// NewCanvas creates a canvas
func NewCanvas(width int, height int) Canvas {
	data := make([][]Color, height)
	for i := range data {
		data[i] = make([]Color, width)
	}
	return Canvas{width, height, data}
}
