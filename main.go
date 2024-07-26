package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"math"
	// "time"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
	// "io/fs"
)

//go:embed assets/*
var assets embed.FS

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

var PlayerSprite = mustLoadImage("assets/playerShip1_blue.png")

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector
}

func (g *Game) Update() error {
	speed := 5.0
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y = -speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X = -speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X = +speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y = +speed
	}

	// This checks for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	g.playerPosition.X += delta.X
	g.playerPosition.Y += delta.Y

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(PlayerSprite, op)

	// ops := &colorm.DrawImageOptions{}
	// cm := colorm.ColorM{}
	//
	// cm.Translate(1.0, 1.0, 1.0, 0.0)
	// colorm.DrawImage(screen, PlayerSprite, cm, ops)

	// width := PlayerSprite.Bounds().Dx()
	// height := PlayerSprite.Bounds().Dy()
	//
	// halfW := float64(width / 2)
	// halfH := float64(height / 2)
	//
	// op.GeoM.Translate(-halfW, -halfH)
	// op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	// op.GeoM.Translate(halfW, halfH)
	//
	// screen.DrawImage(PlayerSprite, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerPosition: Vector{X: 100, Y: 10},
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
