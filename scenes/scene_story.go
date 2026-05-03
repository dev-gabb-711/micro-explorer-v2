package scenes

import (
	"Game1/assets"
	"Game1/core"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type StoryScene struct {
	level        int
	title        string
	lines        []string
	currentLine  int
	visibleChars int
	ticks        int
	state        int // 0: Title Delay, 1: Typing, 2: Waiting for Enter
}

type MissionData struct {
	Title string
	Lines []string
}

func NewStoryScene(level int) *StoryScene {
	// Map to hold all the level descriptions
	briefings := map[int]MissionData{
		1: {
			Title: "Level 1: The Pond Microbe - Small Infections first bloom",
			Lines: []string{
				"Objective: Infiltrate the Amoeba",
				"The host is unsuspecting. Evade the lysosomes and swarms.",
				"Reach the nucleus to begin replication",
			},
		},
		2: {
			Title: "Level 2: The Plant Cell - Energy is being stolen",
			Lines: []string{
				"Objective: Breach the cell wall.",
				"Cellulose levels are high today. Use your enzymes and energy.",
				"This is the best way to bypass the outer membrane.",
			},
		},
		// Adding more here in the future
		10: {
			Title: "Level 10: The Source - The supervirus reveals itself",
			Lines: []string{"Objective: Unknown.", "Proceed with extreme caution.", "Get out of there alive BUT", "Save.", "The.", "World."},
		},
	}

	data, ok := briefings[level]
	if !ok {
		data = MissionData{
			Title: "MISSION START",
			Lines: []string{"Objective: Unknown.", "Proceed with extreme caution.", "Get out of there ALIVE"},
		}
	}

	return &StoryScene{
		level: level,
		title: data.Title,
		lines: data.Lines,
		state: 0,
	}
}

func (s *StoryScene) Update(state *core.GlobalState) core.Scene {
	s.ticks++

	switch s.state {
	case 0: // Title Delay
		if s.ticks > 60 {
			s.state = 1
			s.ticks = 0
		}

	case 1: // Typing Mode
		if s.ticks%3 == 0 && s.visibleChars < len(s.lines[s.currentLine]) {
			s.visibleChars++
		}

		// Allow skipping the typing animation for the current line
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			s.visibleChars = len(s.lines[s.currentLine])
		}

		// When the line finishes typing, move to the Waiting state
		if s.visibleChars >= len(s.lines[s.currentLine]) {
			s.state = 2
			s.ticks = 0 // Reset ticks for the blinking cursor
		}

	case 2: // Waiting for Enter
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			s.currentLine++
			s.visibleChars = 0

			// If we've shown all lines, exit the scene
			if s.currentLine >= len(s.lines) {
				return NewMapScene() // This is where we actually transition into gameplay
			}

			// Otherwise, go back to typing the next line
			s.state = 1
		}
	}
	return nil
}

func (s *StoryScene) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Dx()

	// Draw UI Box Header
	header := "--- INCOMING TRANSMISSION ---"
	hBounds := text.BoundString(assets.MenuFont, header)
	text.Draw(screen, header, assets.MenuFont, (screenWidth-hBounds.Dx())/2, 100, color.RGBA{0, 255, 0, 255})

	// Draw Title
	tBounds := text.BoundString(assets.MenuFont, s.title)
	text.Draw(screen, s.title, assets.MenuFont, (screenWidth-tBounds.Dx())/2, 160, color.RGBA{255, 200, 50, 255})

	// Draw Briefing Lines
	startY := 220
	lineSpacing := 35

	// Only draw lines if they are past initial delay state
	if s.state > 0 {
		for i := 0; i < s.currentLine; i++ {
			// Draw all PREVIOUS, fully completed lines
			y := startY + (i * lineSpacing)
			text.Draw(screen, s.lines[i], assets.MenuFont, 100, y, color.White)
		}

		// Draw the CURRENT line being typed
		if s.currentLine < len(s.lines) {
			y := startY + (s.currentLine * lineSpacing)
			displayText := s.lines[s.currentLine][:s.visibleChars]
			text.Draw(screen, displayText, assets.MenuFont, 100, y, color.White)
		}
	}

	if s.state == 2 {
		prompt := "[ PRESS ENTER ]"
		pBounds := text.BoundString(assets.GridFont, prompt)

		alpha := uint8(200)
		if (s.ticks/30)%2 == 0 {
			alpha = 50
		}
		text.Draw(screen, prompt, assets.GridFont, (screenWidth-pBounds.Dx())/2, 500, color.RGBA{255, 255, 255, alpha})
	}
}
