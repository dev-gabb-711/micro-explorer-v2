package main

import (
	"Game1/core"
	"Game1/scenes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &Game{
		state:        core.NewGlobalState(),
		currentScene: &scenes.MainMenuScene{},
	}

	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Micro Explorer: The War of the Small World")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
