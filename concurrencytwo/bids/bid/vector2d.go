package bid

import "math"

type vector2D struct {
	X float64
	Y float64
}

// Add new add Position vector new
func (v1 *vector2D) Add(v2 vector2D) vector2D {
	return vector2D{v1.X + v2.X, v1.Y + v2.Y}
}

// Subtract new subtract Position vector new
func (v1 *vector2D) Subtract(v2 vector2D) vector2D {
	return vector2D{v1.X - v2.X, v1.Y - v2.Y}
}

// Multiply new multiply Position vector new
func (v1 *vector2D) Multiply(v2 vector2D) vector2D {
	return vector2D{v1.X * v2.X, v1.Y * v2.Y}
}

// AddValue new sum Position vector new
func (v1 *vector2D) AddValue(value float64) vector2D {
	return vector2D{v1.X + value, v1.Y + value}
}

// MultipleValue new multiply Position vector new
func (v1 *vector2D) MultipleValue(value float64) vector2D {
	return vector2D{v1.X + value, v1.Y + value}
}

func (v1 *vector2D) DivisionV(d float64) vector2D {
	return vector2D{v1.X / d, v1.Y / d}
}

// Limit new multiply Position vector new
func (v1 *vector2D) Limit(lower, upper float64) vector2D {
	x := math.Min(math.Max(v1.X, lower), upper)
	y := math.Min(math.Max(v1.Y, lower), upper)
	return vector2D{x, y}
}

func (v1 *vector2D) Distance(v2 vector2D) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}
