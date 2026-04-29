package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &Game{
		state:        NewGlobalState(),
		currentScene: &MainMenuScene{},
	}

	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Micro Explorer: The War of the Small World")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
