package sense

import (
	"errors"
)

type StimulusType int

const (
	SoundStimuli     StimulusType = iota // Audible stimulus, impacted by the resonance of the tile.  Clears out sooner that the other stimuli.
	PheremoneStimuli                     // Stays where put, doesn't spread as far, draft doesn't get rid of it.
	ScentStimuli                         // Wafts outward, moved or cleared by draft
)

type DraftDirection int

const (
	NorthernDraft DraftDirection = iota
	WesternDraft
	EasternDraft
	SouthernDraft
)

type Stimulus struct {
	Type        StimulusType
	Intensity   int
	Decay       int
	ID          string //Custom identify to "name" the stimuli
	tickUpdated int
}

type StimuliTile struct {
	Stimuli   []Stimulus
	Resonance float32 //How well a sound can echo in this tile.
	Draft     float32 //How drafty this tile.
	DraftDir  DraftDirection
	Solid     bool
}

type SenseScape struct {
	width  int
	height int
	data   [][]StimuliTile
	tick   int
}

func NewSenseScape(width int, height int) *SenseScape {
	senseScape := &SenseScape{width: width, height: height}
	data := make([][]StimuliTile, width, height)
	for x := 0; x < width; x++ {
		col := []StimuliTile{}
		for y := 0; y < height; y++ {
			col = append(col, StimuliTile{})
		}
		data[x] = append(data[x], col...)
	}

	senseScape.data = data
	return senseScape
}

func (s *SenseScape) GetStimuliAt(x int, y int) ([]Stimulus, error) {
	if x < s.width && y < s.height && x >= 0 && y >= 0 {
		return s.data[x][y].Stimuli, nil
	}

	return nil, errors.New("outside bounds")
}

func (s *SenseScape) ApplyStimulus(x int, y int, stimulus Stimulus) error {
	stimulus.tickUpdated = s.tick
	if x < s.width && y < s.height && x >= 0 && y >= 0 {
		for i := range s.data[x][y].Stimuli {
			currentStimulus := s.data[x][y].Stimuli[i]
			//Merge with an existing stimulus.
			if currentStimulus.Type == stimulus.Type && currentStimulus.ID == stimulus.ID {
				s.data[x][y].Stimuli[i].Intensity += stimulus.Intensity
				return nil
			}
		}

		s.data[x][y].Stimuli = append(s.data[x][y].Stimuli, stimulus)
		return nil
	}

	return errors.New("outside bounds")
}

func (s *SenseScape) safeApplyStimulus(x int, y int, stimulus Stimulus) error {
	stimulus.tickUpdated = s.tick
	if x < s.width && y < s.height && x >= 0 && y >= 0 {
		for i := range s.data[x][y].Stimuli {
			currentStimulus := s.data[x][y].Stimuli[i]
			//Merge with an existing stimulus.
			if currentStimulus.Type == stimulus.Type && currentStimulus.ID == stimulus.ID {
				if currentStimulus.tickUpdated == s.tick {
					return errors.New("already updated this tick")
				}
				s.data[x][y].Stimuli[i].Intensity += stimulus.Intensity
				return nil
			}
		}

		s.data[x][y].Stimuli = append(s.data[x][y].Stimuli, stimulus)
		return nil
	}

	return errors.New("outside bounds")
}

func (s *SenseScape) MakeSound(x int, y int, ID string, intensity int) {
	for sX := x - intensity/2; sX < x+intensity/2; sX++ {
		for sY := y - intensity/2; sY < y+intensity/2; sY++ {
			if sX < s.width && sY < s.height && sX >= 0 && sY >= 0 {
				if s.data[sX][sY].Solid {
					continue
				}
				s.ApplyStimulus(sX, sY, Stimulus{Type: SoundStimuli, Intensity: intensity, ID: ID})
			}
		}
	}
}

func (s *SenseScape) Update() error {
	s.tick++
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			tile := s.data[x][y]
			if tile.Solid {
				continue
			}
			for i := range tile.Stimuli {
				currentStimulus := &tile.Stimuli[i]
				//Nothing to handle
				if currentStimulus.Intensity == 0.0 {
					continue
				}

				//Disperse the stimuli based on its type
				switch currentStimulus.Type {
				case SoundStimuli:
					currentStimulus.Intensity = 0

				case ScentStimuli:
					if currentStimulus.tickUpdated != s.tick {
						currentStimulus.Intensity -= currentStimulus.Decay
					}

					newStimulus := Stimulus{Intensity: currentStimulus.Intensity - currentStimulus.Decay, Type: currentStimulus.Type, ID: currentStimulus.ID, Decay: currentStimulus.Decay}

					s.safeApplyStimulus(x+1, y, newStimulus)
					s.safeApplyStimulus(x-1, y, newStimulus)
					s.safeApplyStimulus(x, y+1, newStimulus)
					s.safeApplyStimulus(x, y-1, newStimulus)
					s.safeApplyStimulus(x+1, y+1, newStimulus)
					s.safeApplyStimulus(x+1, y-1, newStimulus)
					s.safeApplyStimulus(x-1, y+1, newStimulus)
					s.safeApplyStimulus(x-1, y-1, newStimulus)
				case PheremoneStimuli:
					currentStimulus.Intensity -= currentStimulus.Decay
				}
				currentStimulus.tickUpdated = s.tick
			}
		}
	}

	return nil
}
