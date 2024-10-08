package main

import (
	"github.com/alikastrati/space-007/game"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
