package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene defines what a game state needs to do
type Scene interface {
	Update(state *GlobalState) Scene // Returns the next scene to switch to, or nil to stay
	Draw(screen *ebiten.Image)
}
