package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float64
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float64) {
	c.Radius *= factor
}

type Square struct {
	Side float64
}

func (s *Square) Render() string {
	return fmt.Sprintf("Circle of radius %f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

type TransparentShape struct {
	Shape        Shape
	Transparency float64
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has the color %v", t.Shape.Render(), t.Transparency)
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

func main() {
	c := Circle{45}
	s := Square{45}

	fmt.Println(c, s)

	redCircle := ColoredShape{&c, "red"}

	fmt.Println(redCircle.Render())

	tCircle := TransparentShape{&c, 34}

	fmt.Println(tCircle)
}
