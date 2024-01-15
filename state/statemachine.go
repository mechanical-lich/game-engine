package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type StateMachine struct {
	states       []StateInterface
	currentState int
}

func (s *StateMachine) Update() {
	if s.currentState >= 0 && len(s.states) > 0 {
		nextState := s.states[s.currentState].Update()
		if s.states[s.currentState].Done() {
			s.PopCurrentState()
		}
		if nextState != nil {
			s.PushState(nextState)
		}
	}
}

func (s *StateMachine) Draw(screen *ebiten.Image) {
	if s.currentState >= 0 && len(s.states) > 0 {
		s.states[s.currentState].Draw(screen)
	}
}

// TODO it's pretending it's a stack, but it is really a queue.
func (s *StateMachine) PushState(state StateInterface) {
	s.states = append(s.states, state)
	s.currentState = len(s.states) - 1
}

func (s *StateMachine) PopCurrentState() {
	s.states = append(s.states[:s.currentState], s.states[s.currentState+1:]...)
	s.currentState = len(s.states) - 1
}
