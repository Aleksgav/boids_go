package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
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

func (bs *BoidsSim) Update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range bs.bSwarm.boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), bs.bColor)
			screen.Set(int(boid.position.x-1), int(boid.position.y), bs.bColor)
			screen.Set(int(boid.position.x), int(boid.position.y-1), bs.bColor)
			screen.Set(int(boid.position.x), int(boid.position.y+1), bs.bColor)
		}
	}

	return nil
}

func (bs *BoidsSim) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return bs.screenWidth, bs.screenHeight
}

func main() {
	width := 640
	height := 360

	title := "BOIDS"

	greenColor := color.RGBA{R: 10, G: 255, B: 50, A: 255}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	bs := NewBoidsSim(640, 360, 500, greenColor)

	if err := ebiten.RunGame(bs); err != nil {
		log.Fatal(err)
	}
}
