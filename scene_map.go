package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type MapScene struct {
	selectedLevel int
	levelNames    []string
}

func NewMapScene() *MapScene {
	return &MapScene{
		selectedLevel: 0,
		levelNames: []string{
			"[1] Pond Microbe", "[2] Plant Cell",
			"[3] Insect Host", "[4] Fish Cell",
			"[5] Bird Host", "[6] Mammal Host",
			"[7] Human Blood", "[8] Human Brain",
			"[9] Human Heart", "[10] The Source",
		},
	}
}

func (s *MapScene) Update(state *GlobalState) Scene {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && s.selectedLevel < 5 {
		s.selectedLevel += 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && s.selectedLevel >= 5 {
		s.selectedLevel -= 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && s.selectedLevel%5 != 0 {
		s.selectedLevel -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && s.selectedLevel%5 != 4 {
		s.selectedLevel += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		state.Level = s.selectedLevel + 1
		fmt.Printf("Transitioning to Level %d: %s\n", state.Level, s.levelNames[s.selectedLevel])
		return NewStoryScene(state.Level)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return &MainMenuScene{}
	}

	return nil
}

func (s *MapScene) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()

	titleStr := "SELECT MISSION"
	titleBounds := text.BoundString(TitleFont, titleStr)
	titleX := (screenWidth - titleBounds.Dx()) / 2
	text.Draw(screen, titleStr, TitleFont, titleX, 100, color.White)

	subStr := "Arrows to Navigate, ESC to go back, ENTER to Select"
	subBounds := text.BoundString(MenuFont, subStr)
	subX := (screenWidth - subBounds.Dx()) / 2
	headerBottomY := 160
	text.Draw(screen, subStr, MenuFont, subX, headerBottomY, color.RGBA{150, 150, 150, 255}) // Gray subtitle

	columnSpacing := 205
	rowSpacing := 150

	longestBounds := text.BoundString(GridFont, "[10] The Source")
	gridTotalWidth := (columnSpacing * 4) + longestBounds.Dx()
	startX := (screenWidth - gridTotalWidth) / 2
	availableHeight := screenHeight - headerBottomY
	startY := headerBottomY + ((availableHeight - rowSpacing) / 2)

	for i, name := range s.levelNames {
		col := i % 5
		row := i / 5

		x := startX + (col * columnSpacing)
		y := startY + (row * rowSpacing)

		c := color.RGBA{150, 150, 150, 255} // Default unselected text color

		if i == s.selectedLevel {
			c = color.RGBA{255, 255, 255, 255} // Highlight selected text in bright white
			text.Draw(screen, ">", GridFont, x-20, y, c)
		}

		text.Draw(screen, name, GridFont, x, y, c)
	}
}
