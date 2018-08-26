package trace

// Color defines a color
type Color struct {
	Red, Green, Blue float64
}

// Add adds a tuple to this tuple
func (c Color) Add(a Color) Color {
	return Color{
		a.Red + c.Red,
		a.Green + c.Green,
		a.Blue + c.Blue}
}
