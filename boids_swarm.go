package main

import "sync"

type BoidsSwarm struct {
	screenWidth  int
	screenHeight int

	bMap  [][]int
	boids []*Boid

	mu *sync.RWMutex
}

func NewBoidsSwarm(width, height, bQuantity int) *BoidsSwarm {
	bMap := make([][]int, width+1)
	for i := range bMap {
		bMap[i] = make([]int, height+1)
		for j := range bMap[i] {
			bMap[i][j] = -1
		}
	}

	bs := &BoidsSwarm{
		screenWidth:  width,
		screenHeight: height,
		bMap:         bMap,
		boids:        make([]*Boid, 0, bQuantity),
		mu:           &sync.RWMutex{},
	}

	bs.init()

	return bs
}

func (bs *BoidsSwarm) init() {
	for i := 0; i < cap(bs.boids); i++ {
		bs.AddBoid()
	}
}

func (bs *BoidsSwarm) AddBoid() {
	b := NewBoid(bs.screenWidth, bs.screenHeight, len(bs.boids), bs)

	bs.boids = append(bs.boids, b)
	bs.bMap[int(b.position.x)][int(b.position.y)] = b.id

	go b.start()
}
