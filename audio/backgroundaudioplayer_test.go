package audio

import (
	"testing"
)

func TestNewBackgroundAudioPlayer(t *testing.T) {
	Init()
	a, err := LoadAudioFromFile("../assets/audio/ding.mp3", TypeMP3)
	if err != nil {
		t.Errorf("Error loading audio frome file. %s", err.Error())
	}

	b, err := NewBackgroundAudioPlayer([]*AudioResource{a})

	if b == nil {
		t.Errorf("No background player returned")
	}

	if err != nil {
		t.Errorf("Error creating background player %s", err.Error())
	}
}
