package event

import (
	"fmt"
	"testing"
)

const (
	Test EventType = iota
)

type TestEventData struct {
}

func (t TestEventData) GetType() EventType {
	return Test
}

type TestListener struct {
}

func (t *TestListener) HandleEvent(data EventData) error {
	fmt.Println("Handling event: ", data)
	return nil
}

func TestRegisterListener(t *testing.T) {
	m := &QueuedEventManager{}

	testListener := &TestListener{}

	m.RegisterListener(testListener, Test)
	if len(m.listeners[Test]) != 1 {
		t.Errorf("Output %q not equal to expected %q", len(m.listeners[Test]), 1)
	}
}

func TestUnregisterListener(t *testing.T) {
	m := &QueuedEventManager{}

	testListener := &TestListener{}

	m.RegisterListener(testListener, Test)
	if len(m.listeners[Test]) != 1 {
		t.Errorf("Output %q not equal to expected %q", len(m.listeners[Test]), 1)
	}

	m.UnregisterListener(testListener, Test)
	if len(m.listeners[Test]) != 0 {
		t.Errorf("Output %q not equal to expected %q", len(m.listeners[Test]), 0)
	}
}

func TestUnregisterListenerFromAll(t *testing.T) {
	m := &QueuedEventManager{}

	testListener := &TestListener{}

	m.RegisterListener(testListener, Test)
	if len(m.listeners[Test]) != 1 {
		t.Errorf("Output %q not equal to expected %q", len(m.listeners[Test]), 1)
	}

	m.UnregisterListenerFromAll(testListener)
	if len(m.listeners[Test]) != 0 {
		t.Errorf("Output %q not equal to expected %q", len(m.listeners[Test]), 0)
	}
}
