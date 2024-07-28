package game

import (
	"github.com/alikastrati/space-007/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// Set sprite position to the middle of the screen
	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
		rotation: 0,
	}
}

func (p *Player) Update() {

	speed := math.Pi / float64(ebiten.TPS())
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X -= speed
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X += speed
		p.rotation += speed
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
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
