package main

import (
	"Game1/scenes" // Ensure your go.mod matches this path
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := &Game{
		state:        NewGlobalState(),
		currentScene: &scenes.MainMenuScene{},
	}

	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Micro Explorer: The War of the Small World")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
