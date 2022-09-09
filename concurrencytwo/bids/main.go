package main

import (
	"github.com/guillospy92/utilsgolang/concurrencytwo/bids/bid"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

type Game struct{}

const screenWidth, screenHeight = 640, 360

func (g *Game) Update(screen *ebiten.Image) error {

	for _, boid := range bid.Boids {
		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), green)
	}

	return nil
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
)

// https://github.com/cutajarj/multithreadingingo/blob/master/boids/main.go
func main() {

	for i, row := range bid.BitMap {
		for j := range row {
			bid.BitMap[i][j] = -1
		}
	}

	for i := 0; i < bid.MaxFormCountBid; i++ {
		bid.CreateBid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
