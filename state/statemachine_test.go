package state

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

type TestState struct {
	Value      string
	IsDone     bool
	UpdateFunc func() StateInterface
}

func (s *TestState) Update() StateInterface {
	if s.UpdateFunc != nil {
		return s.UpdateFunc()
	}
	return nil
}

func (s *TestState) Draw(screen *ebiten.Image) {
	s.Value = "Drawn"
}

func (s *TestState) Done() bool {
	return s.IsDone
}

func TestStateUpdates(t *testing.T) {
	sm := StateMachine{}

	s := &TestState{}
	s.UpdateFunc = func() StateInterface {
		s.Value = "Updated"
		return nil
	}
	sm.PushState(s)

	sm.Update()

	if s.Value != "Updated" {
		t.Errorf("Failed to updated state")
	}
}

func TestStateDraw(t *testing.T) {
	sm := StateMachine{}

	s := &TestState{}
	sm.PushState(s)

	sm.Draw(&ebiten.Image{})

	if s.Value != "Drawn" {
		t.Errorf("Failed to draw state")
	}
}

func TestStatePopsWhenDone(t *testing.T) {
	sm := StateMachine{}

	s := &TestState{}
	s.UpdateFunc = func() StateInterface {
		s.IsDone = true
		return nil
	}
	sm.PushState(s)

	sm.Update()

	if len(sm.states) > 0 {
		t.Errorf("Failed to pop finished state")
	}
}

func TestStatePushesNewState(t *testing.T) {
	sm := StateMachine{}

	s := &TestState{}
	s2 := &TestState{}
	s2.UpdateFunc = func() StateInterface {
		s2.Value = "Updated"
		return nil
	}

	s.UpdateFunc = func() StateInterface {
		return s2
	}

	sm.PushState(s)

	sm.Update()

	if len(sm.states) < 2 {
		t.Errorf("Failed to push new state")
	}

	if sm.states[sm.currentState] != s2 {
		t.Error("Current state is not the most recently pushed state")
	}

	if s2.Value != "" {
		t.Error("State data updated before update called")
	}

	sm.Update()
	if s2.Value != "Updated" {
		t.Error("Failed to update state 2 data")
	}
}
