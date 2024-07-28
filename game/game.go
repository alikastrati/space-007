package game

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	// How we use our timer
	// g.attackTimer.Update()
	// if g.attackTimer.IsReady() {
	// 	g.attackTimer.Reset()
	//
	// 	// Execute something (an attack for example)
	// }
	//
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(3)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	return nil

}

// This function is used to draw an image on our board/canvas?
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

}

// This creates the Layout (board/canvas?) where our objects/images will be shown
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
