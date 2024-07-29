package game

import (
	"github.com/alikastrati/space-007/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	laserSpeedPerSecond = 350.0
)

type Laser struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewLaser(pos Vector, rotation float64) *Laser {
	sprite := assets.ShootingSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	b := &Laser{
		position: pos,
		rotation: rotation,
		sprite:   sprite,
	}

	return b
}

func (b *Laser) Update() {
	speed := laserSpeedPerSecond / float64(ebiten.TPS())

	b.position.X += math.Sin(b.rotation) * speed
	b.position.Y += math.Cos(b.rotation) * -speed
}

func (b *Laser) Draw(screen *ebiten.Image) {
	bounds := b.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.sprite, op)
}
