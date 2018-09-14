package trace

import "math"

// Color defines a color
type Color struct {
	Red, Green, Blue float64
}

// Red Color
var Red = Color{1, 0, 0}
var White = Color{1, 1, 1}
var Black = Color{0, 0, 0}

// Add adds a tuple to this tuple
func (c Color) Add(a Color) Color {
	return Color{
		a.Red + c.Red,
		a.Green + c.Green,
		a.Blue + c.Blue}
}

// Multiply multiplies a tuple from this tuple
func (c Color) Multiply(a Color) Color {
	r := c.Red * a.Red
	g := c.Green * a.Green
	b := c.Blue * a.Blue
	return Color{r, g, b}
}

// Multiply multiplies a tuple from this tuple
func (c Color) MultiplyByScalar(a float64) Color {
	r := c.Red * a
	g := c.Green * a
	b := c.Blue * a
	return Color{r, g, b}
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

func (c Color) Equals(a Color) bool {
	return floatEquals(c.Red, a.Red) &&
		floatEquals(c.Green, a.Green) &&
		floatEquals(c.Blue, a.Blue)
}
