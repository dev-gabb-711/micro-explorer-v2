package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	TitleFont font.Face
	MenuFont  font.Face
	GridFont  font.Face
)

func init() {
	// Parse the built-in Arcade pixel font
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	// Size 12 ensures that 5 columns fit perfectly inside a 1024px width
	MenuFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	// Larger size for Headers and Titles
	TitleFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    28,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	GridFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     72,
		Hinting: font.HintingNone,
	})
}
