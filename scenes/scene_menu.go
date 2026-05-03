package scenes

import (
	"Game1/assets"
	"Game1/core"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type MainMenuScene struct {
	selectedOption int // 0: Play, 1: How to Play, 2: Exit
}

func (s *MainMenuScene) Update(state *core.GlobalState) core.Scene {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.selectedOption = (s.selectedOption + 1) % 3
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.selectedOption = (s.selectedOption + 2) % 3
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch s.selectedOption {
		case 0: // Play
			return NewMapScene()
		case 1: // How to Play
			return nil
		case 2: // Exit
			os.Exit(0)
		}
	}
	return nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()

	// 1. Draw Title
	title := "MICRO EXPLORER"
	subtitle := "The War of the Small World"

	titleBounds := text.BoundString(assets.TitleFont, title)
	titleX := (screenWidth - titleBounds.Dx()) / 2
	text.Draw(screen, title, assets.TitleFont, titleX, screenHeight/3, color.White)

	subBounds := text.BoundString(assets.MenuFont, subtitle)
	subX := (screenWidth - subBounds.Dx()) / 2
	text.Draw(screen, subtitle, assets.MenuFont, subX, (screenHeight/3)+40, color.RGBA{147, 250, 165, 1}) // Light blue

	// 2. Draw Options
	options := []string{"Play", "How to Play", "Exit"}
	startY := (screenHeight / 2) + 80
	rowSpacing := 60

	for i, opt := range options {
		c := color.RGBA{150, 150, 150, 255} // Unselected color

		bounds := text.BoundString(assets.MenuFont, opt)
		x := (screenWidth - bounds.Dx()) / 2
		y := startY + (i * rowSpacing)

		if i == s.selectedOption {
			c = color.RGBA{255, 255, 255, 255}
			text.Draw(screen, ">", assets.MenuFont, x-24, y, c)
		}
		text.Draw(screen, opt, assets.MenuFont, x, y, c)
	}
}
