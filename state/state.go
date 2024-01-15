package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// StateInterface - Interface representing a state to be used by the game.
type StateInterface interface {
	Update() StateInterface
	Draw(screen *ebiten.Image)
	Done() bool
}
