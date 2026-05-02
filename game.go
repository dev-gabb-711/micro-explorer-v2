package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	state        *GlobalState
	currentScene Scene
}

func (g *Game) Update() error {
	nextScene := g.currentScene.Update(g.state)
	if nextScene != nil {
		g.currentScene = nextScene
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 1024, 768
}
