package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Timer struct {
	currentTicks int
	targetTicks  int
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
