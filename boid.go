package main

import (
	"math"
	"math/rand"
	"time"
)

const (
	adjustmentRate    = 0.015
	interactionRadius = 13
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int

	swarm *BoidsSwarm
}

func NewBoid(width, height, id int, swarm *BoidsSwarm) *Boid {
	return &Boid{
		position: Vector2d{
			x: rand.Float64() * float64(width),
			y: rand.Float64() * float64(height),
		},
		velocity: Vector2d{
			x: (rand.Float64() * 2) - 1.0,
			y: (rand.Float64() * 2) - 1.0,
		},
		id:    id,
		swarm: swarm,
	}
}

func (b *Boid) calcAcceleration() Vector2d {
	upper := b.position.AddVal(interactionRadius)
	lower := b.position.AddVal(-interactionRadius)

	avgVelocity := Vector2d{0, 0}
	avgPosition := Vector2d{0, 0}
	separation := Vector2d{0, 0}
	count := 0.0

	width := float64(b.swarm.screenWidth)
	height := float64(b.swarm.screenHeight)

	b.swarm.mu.RLock()

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, width); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, height); j++ {
			if otherBoidID := b.swarm.bMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != b.id {
				if dist := b.swarm.boids[otherBoidID].position.Dist(b.position); dist < interactionRadius {
					count++
					avgVelocity = avgVelocity.Add(b.swarm.boids[otherBoidID].velocity)
					avgPosition = avgPosition.Add(b.swarm.boids[otherBoidID].position)
					separation = separation.Add(b.position.Sub(b.swarm.boids[otherBoidID].position).DivVal(dist))
				}
			}
		}
	}

	b.swarm.mu.RUnlock()

	acc := Vector2d{
		x: b.borderBounce(b.position.x, width),
		y: b.borderBounce(b.position.y, height),
	}

	if count > 0 {
		avgVelocity = avgVelocity.DivVal(count)
		avgPosition = avgPosition.DivVal(count)

		accAlignment := avgVelocity.Sub(b.velocity).MulVal(adjustmentRate)
		accCohesion := avgPosition.Sub(b.position).MulVal(adjustmentRate)
		accSeparation := separation.MulVal(adjustmentRate)
		acc = acc.Add(accAlignment).Add(accCohesion).Add(accSeparation)
	}

	return acc
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < interactionRadius {
		return 1 / pos
	} else if pos > maxBorderPos-interactionRadius {
		return 1 / (pos - maxBorderPos)
	}

	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()

	b.swarm.mu.Lock()

	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	b.swarm.bMap[int(b.position.x)][int(b.position.y)] = -1

	b.position = b.position.Add(b.velocity)
	b.swarm.bMap[int(b.position.x)][int(b.position.y)] = b.id

	b.swarm.mu.Unlock()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}
