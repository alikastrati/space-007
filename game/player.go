package game

import (
	"github.com/alikastrati/space-007/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

const (
	shootCooldown     = time.Millisecond * 500
	rotationPerSecond = math.Pi + 2

	laserSpawnOffset = 50.0
)

type Player struct {
	game     *Game
	position Vector
	sprite   *ebiten.Image
	rotation float64

	shootCooldown *Timer
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
		game:          game,
		position:      pos,
		sprite:        sprite,
		rotation:      0,
		shootCooldown: NewTimer(shootCooldown),
	}
}

func (p *Player) Update() {

	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()

		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW + math.Sin(p.rotation)*laserSpawnOffset,
			p.position.Y + halfH + math.Cos(p.rotation)*-laserSpawnOffset,
		}

		laser := NewLaser(spawnPos, p.rotation)
		p.game.AddLaser(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
