package sense

import (
	"fmt"
	"testing"
)

const TestWidth = 10
const TestHeight = 10

func TestNewSenseScapeCreatesAValidSoundScape(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)

	if s.width != TestWidth {
		t.Errorf("Sensescape width %d not equal to expected %d", s.width, TestWidth)
	}

	if s.height != TestHeight {
		t.Errorf("Sensescape height %d not equal to expected %d", s.height, TestHeight)
	}

	if len(s.data) != TestWidth {
		t.Errorf("Sensescape data width %d not equal to expected %d", len(s.data), TestWidth)
	}
	if len(s.data[0]) != TestHeight {
		t.Errorf("Sensescape data height %d not equal to expected %d", len(s.data[0]), TestHeight)
	}
}

func TestGetStimuliAtWithNothingThere(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)

	stimuli, _ := s.GetStimuliAt(0, 0)

	if stimuli != nil {
		t.Errorf("Stimuli did not come back nil")
	}
}

func TestGetStimuliAtOutOfBounds(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)

	_, err := s.GetStimuliAt(TestWidth+1, 0)

	if err == nil {
		t.Errorf("Get stimuli at did not error with out of bounds")
	}
}

func TestAddStimulus(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)
	//Test applying a brand new stimulus
	err := s.ApplyStimulus(0, 0, Stimulus{Type: SoundStimuli, Intensity: 1})
	if err != nil {
		t.Errorf("Got error applying stimulus: %s", err)
	}
	stimuli, _ := s.GetStimuliAt(0, 0)

	if stimuli == nil {
		t.Errorf("Stimuli came back nil")
	}

	if stimuli[0].Intensity != 1 {
		t.Errorf("Stimuli intensity %d not equal to expected %d", stimuli[0].Intensity, 1)
	}

	//Test applying the same stimulus again
	err = s.ApplyStimulus(0, 0, Stimulus{Type: SoundStimuli, Intensity: 1})
	if err != nil {
		t.Errorf("Got error applying stimulus the second time: %s", err)
	}

	stimuli, _ = s.GetStimuliAt(0, 0)
	if stimuli == nil {
		t.Errorf("Stimuli came back nil")
	}
	if stimuli[0].Intensity != 2 {
		t.Errorf("Stimuli intensity %d not equal to expected %d", stimuli[0].Intensity, 2)
	}
}

func TestMakeSound(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)
	//Test applying a brand new stimulus
	s.MakeSound(5, 5, "TEST", 5)

	for x := 0; x < TestWidth; x++ {
		for y := 0; y < TestHeight; y++ {
			if len(s.data[x][y].Stimuli) > 0 {
				fmt.Print(s.data[x][y].Stimuli[0].Intensity)
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println("")
	}

	stimuli, _ := s.GetStimuliAt(5, 5)

	if stimuli == nil {
		t.Errorf("Stimuli came back nil")
	}

	if stimuli[0].Intensity != 5 {
		t.Errorf("Stimuli intensity %d not equal to expected %d", stimuli[0].Intensity, 5)
	}

	stimuli, _ = s.GetStimuliAt(4, 4)
	if stimuli == nil {
		t.Errorf("Stimuli came back nil")
	}
	if stimuli[0].Intensity != 5 {
		t.Errorf("Stimuli intensity %d not equal to expected %d", stimuli[0].Intensity, 5)
	}
}

func TestAddStimulusOutOfBounds(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)

	err := s.ApplyStimulus(TestWidth+1, 0, Stimulus{Type: SoundStimuli, Intensity: 1})

	if err == nil {
		t.Errorf("Get stimuli at did not error with out of bounds")
	}
}

func TestUpdateDoesNotError(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)
	err := s.Update()

	if err != nil {
		t.Errorf("SenseScape update errored: %s", err)
	}
}

func TestUpdateDispersesScentCorrectly(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)

	err := s.ApplyStimulus(0, 0, Stimulus{Type: ScentStimuli, Intensity: 4, Decay: 1})
	if err != nil {
		t.Errorf("Got error applying stimulus during update test: %s", err)
	}

	err = s.Update()

	if err != nil {
		t.Errorf("SenseScape update errored: %s", err)
	}

	stimuli, _ := s.GetStimuliAt(0, 0)
	if stimuli == nil {
		t.Errorf("Stimuli came back nil")
	}
	if stimuli[0].Intensity != 3 {
		t.Errorf("Stimuli intensity %d not equal to expected %d", stimuli[0].Intensity, 3)
	}

	stimuli, _ = s.GetStimuliAt(1, 0)
	if stimuli == nil {
		t.Errorf("Stimuli 2 came back nil")
	}
	if stimuli[0].Intensity != 2 {
		t.Errorf("Stimuli 2 intensity %d not equal to expected %d", stimuli[0].Intensity, 2)
	}

	stimuli, _ = s.GetStimuliAt(2, 0)
	if stimuli == nil {
		t.Errorf("Stimuli 3 came back nil")
	}
	if stimuli[0].Intensity != 1 {
		t.Errorf("Stimuli 3 intensity %d not equal to expected %d", stimuli[0].Intensity, 1)
	}

}

func TestUpdateDecaysPheremones(t *testing.T) {
	s := NewSenseScape(TestWidth, TestHeight)

	err := s.ApplyStimulus(0, 0, Stimulus{Type: PheremoneStimuli, Intensity: 5, Decay: 1})
	if err != nil {
		t.Errorf("Got error applying stimulus during update test: %s", err)
	}

	err = s.Update()

	if err != nil {
		t.Errorf("SenseScape update errored: %s", err)
	}

	stimuli, _ := s.GetStimuliAt(0, 0)
	if stimuli == nil {
		t.Errorf("Stimuli came back nil")
	}
	if stimuli[0].Intensity != 4 {
		t.Errorf("Stimuli intensity %d not equal to expected %d", stimuli[0].Intensity, 4)
	}
}
