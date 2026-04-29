package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"os"
)

type MainMenuScene struct {
	selectedOption int
}

func (s *MainMenuScene) Update(state *GlobalState) Scene {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.selectedOption = (s.selectedOption + 1) % 3
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.selectedOption = (s.selectedOption + 2) % 3
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if s.selectedOption == 0 {
			fmt.Println("Starting Game...")
			// Placeholder scene for how to play
		} else if s.selectedOption == 2 {
			os.Exit(0)
		}
	}
	return nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	opts := []string{"Play", "How to Play", "Exit"}
	msg := "Micro Explorer: The War of the Small World\n\n"
	for i, opt := range opts {
		cursor := "  "
		if i == s.selectedOption {
			cursor = "> "
		}
		msg += cursor + opt + "\n"
	}
	ebitenutil.DebugPrint(screen, msg)
}
