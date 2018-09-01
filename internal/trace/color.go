package trace

import "math"

// Color defines a color
type Color struct {
	Red, Green, Blue float64
}

// Red Color
var Red = Color{1, 0, 0}
var White = Color{1, 1, 1}

// Add adds a tuple to this tuple
func (c Color) Add(a Color) Color {
	return Color{
		a.Red + c.Red,
		a.Green + c.Green,
		a.Blue + c.Blue}
}

// Normalize scales a color to a range of 0 to 255
func (c Color) Normalize() (int, int, int) {
	cap := func(scaledValue int) int {
		if scaledValue < 0 {
			return 0
		}
		if scaledValue > 255 {
			return 255
		}
		return scaledValue
	}
	red := cap(int(math.Round(c.Red * 255)))
	green := cap(int(math.Round(c.Green * 255)))
	blue := cap(int(math.Round(c.Blue * 255)))
	return red, green, blue
}
