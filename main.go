package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rillmind/nav-game-go/game"
)

func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
