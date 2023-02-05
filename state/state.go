package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// StateInterface - Interface representing a state to be used by the game.
type StateInterface interface {
	Update()
	Draw(screen *ebiten.Image)
}
