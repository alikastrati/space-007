package game

import (
	"fmt"
	"github.com/alikastrati/space-007/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	meteorSpawnTime = 1 * time.Second

	baseMeteorVelocity  = 0.25
	meteorSpeedUpAmount = 0.1
	meteorSpeedUpTime   = 5 * time.Second
)

type Game struct {
	player           *Player
	meteorSpawnTimer *Timer
	meteors          []*Meteor
	baseVelocity     float64
	velocityTimer    *Timer
	lasers           []*Laser
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity:     baseMeteorVelocity,
		velocityTimer:    NewTimer(meteorSpeedUpTime),
	}

	g.player = NewPlayer(g)

	return g
}

// This function runs at 60 Ticks per Second (TPS)
func (g *Game) Update() error {
	// g.velocityTimer.Update()
	// if g.velocityTimer.IsReady() {
	// 	g.velocityTimer.Reset()
	// 	g.baseVelocity += meteorSpeedUpAmount
	// }
	//
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, l := range g.lasers {
		l.Update()
	}

	// Check for collision (meteor/bullet collision)
	for i, m := range g.meteors {
		for j, b := range g.lasers {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score++

			}
		}
	}

	// Meteor/Player collision
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
	}

	return nil

}

// This function is used to draw an image on our board/canvas?
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 30, color.White)

}

// This creates the Layout (board/canvas?) where our objects/images will be shown
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.lasers = nil
	g.score = 0
	g.meteors = nil
	g.baseVelocity = baseMeteorVelocity
	g.meteorSpawnTimer.Reset()
	g.velocityTimer.Reset()

}

func (g *Game) AddLaser(l *Laser) {
	g.lasers = append(g.lasers, l)
}
