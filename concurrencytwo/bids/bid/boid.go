package bid

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const screenWidth, screenHeight = 640, 360

const MaxFormCountBid = 500

const viewRadius = 13

const adjRate = 0.015

var BitMap [screenWidth + 1][screenHeight + 1]int

var Boids [MaxFormCountBid]*Bid

type Bid struct {
	Position vector2D
	Velocity vector2D
	id       int
}

func (b *Bid) start() {
	for {
		b.moveOne()
		time.Sleep(10 * time.Millisecond)
	}
}

func (b *Bid) calcAcceleration() vector2D {
	upper, lower := b.Position.AddValue(viewRadius), b.Position.AddValue(-viewRadius)

	avgVelocity := vector2D{0, 0}
	count := 0.0

	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.X, 0); j <= math.Min(upper.Y, screenWidth); j++ {
			if otherBidID := BitMap[int(i)][int(j)]; otherBidID != -1 && otherBidID != b.id {
				if distance := Boids[otherBidID].Position.Distance(b.Position); distance < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(Boids[otherBidID].Velocity)
				}
			}
		}
	}

	accel := vector2D{
		0, 0,
	}

	if count > 0 {
		avgVelocity = avgVelocity.DivisionV(count)
		subtract := avgVelocity.Subtract(b.Velocity)
		accel = subtract.MultipleValue(adjRate)
	}

	return accel
}

func (b *Bid) moveOne() {
	add := b.Velocity.Add(b.calcAcceleration())
	b.Velocity = add.Limit(-1, 1)
	BitMap[int(b.Position.X)][int(b.Position.Y)] = -1

	b.Position = b.Position.Add(b.Velocity)

	BitMap[int(b.Position.X)][int(b.Position.Y)] = b.id

	next := b.Position.Add(b.Velocity)

	if next.X >= screenWidth || next.X < 0 {
		b.Velocity = vector2D{
			X: -b.Velocity.X,
			Y: b.Velocity.Y,
		}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		fmt.Println("velocity", -b.Velocity.Y, b.Velocity)
		b.Velocity = vector2D{
			X: b.Velocity.X,
			Y: -b.Velocity.Y,
		}
	}
}

func CreateBid(bit int) {
	b := Bid{
		Position: vector2D{
			X: rand.Float64() * screenWidth,
			Y: rand.Float64() * screenHeight,
		},
		Velocity: vector2D{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
		id: bit,
	}

	fmt.Println(b)

	Boids[bit] = &b

	BitMap[int(b.Position.X)][int(b.Position.Y)] = b.id

	go b.start()
}
