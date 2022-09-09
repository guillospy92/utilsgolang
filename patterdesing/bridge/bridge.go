package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float64)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (v *RasterRenderer) RenderCircle(radius float64) {
	fmt.Println("Drawing a pixel of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float64
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float64) {
	c.radius *= factor
}

func NewCircle(render Renderer, radius float64) *Circle {
	return &Circle{render, radius}
}

func main() {
	//raster := RasterRenderer{}
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
