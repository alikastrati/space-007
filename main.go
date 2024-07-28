package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"math"
	"time"
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

// Assets
var PlayerSprite = mustLoadImage("assets/playerShip1_blue.png")

// Our X, Y coordinates
type Vector struct {
	X float64
	Y float64
}

type Game struct {
	player *Player
}

// Timer that we can use for time based actions
type Timer struct {
	currentTicks int
	targetTicks  int
}

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	return &Player{
		position: Vector{X: 100, Y: 100},
		sprite:   PlayerSprite,
	}
}

func (p *Player) Update() {

	speed := 5.0
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y -= speed
	}

	// Makes diagonal movement slower and the same as the other inputs
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.position.X += delta.X
	p.position.Y -= delta.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

// Creates a NewTimer where the targetTicks is given as a parameter
// and the rest of it's methods update the currentTicks to match
// that parameter. Once it does, it resets the currentTicks
// to the defualt value (0)
func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}

// This function runs at 60 Ticks per Second (TPS)
func (g *Game) Update() error {
	// How we use our timer
	// g.attackTimer.Update()
	// if g.attackTimer.IsReady() {
	// 	g.attackTimer.Reset()
	//
	// 	// Execute something (an attack for example)
	// }
	//
	g.player.Update()
	return nil

}

// This function is used to draw an image on our board/canvas?
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

}

// This creates the Layout (board/canvas?) where our objects/images will be shown
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		player: NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
