package main

import (
	"image/color"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type BoidsSim struct {
	screenWidth  int
	screenHeight int

	bColor    color.RGBA
	bQuantity int

	bSwarm *BoidsSwarm
}

func NewBoidsSim(screenWidth, screenHeight, bQuantity int, bColor color.RGBA) *BoidsSim {
	bs := &BoidsSim{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		bQuantity:    bQuantity,
		bColor:       bColor,
		bSwarm:       NewBoidsSwarm(screenWidth, screenHeight, bQuantity),
	}

	return bs
}

func (bs *BoidsSim) Update() error {
	return nil
}

func (bs *BoidsSim) Draw(screen *ebiten.Image) {
	for _, boid := range bs.bSwarm.boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), bs.bColor)
		screen.Set(int(boid.position.x-1), int(boid.position.y), bs.bColor)
		screen.Set(int(boid.position.x), int(boid.position.y-1), bs.bColor)
		screen.Set(int(boid.position.x), int(boid.position.y+1), bs.bColor)
	}
}

func (bs *BoidsSim) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return bs.screenWidth, bs.screenHeight
}

func main() {
	width := 1280
	height := 720

	title := "BOIDS"

	greenColor := color.RGBA{R: 10, G: 255, B: 50, A: 255}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	bs := NewBoidsSim(width, height, 500, greenColor)

	if err := ebiten.RunGame(bs); err != nil {
		log.Fatal(err)
	}
}
