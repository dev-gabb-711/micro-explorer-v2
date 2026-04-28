package scenes

import (
	"os"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// We need a way to reference the GlobalState from another package,
// so we use a placeholder for now or pass it in.
type MainMenuScene struct {
	selectedOption int // 0: Play, 1: How to Play, 2: Exit
}

func (s *MainMenuScene) Update(state interface{}) interface{} {
	// Simple menu navigation
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.selectedOption = (s.selectedOption + 1) % 3
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.selectedOption = (s.selectedOption + 2) % 3
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch s.selectedOption {
		case 0: // Play
				// Future: return &MapScene{}
		case 2: // Exit
			os.Exit(0)
		}
	}
	return nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "MICRO EXPLORER: THE WAR OF THE SMALL WORLD\n\n"+
		"> Play\n"+
		"  How to Play\n"+
		"  Exit")
}