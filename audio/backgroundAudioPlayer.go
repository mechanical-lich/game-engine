package audio

import (
	"errors"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

type BackgroundAudioPlayer struct {
	audioPlayers []*audio.Player
	activePlayer int
	volume       float64
}

func NewBackgroundAudioPlayer(resources []*AudioResource) (*BackgroundAudioPlayer, error) {
	b := &BackgroundAudioPlayer{audioPlayers: make([]*audio.Player, 0), volume: 1}

	if audioContext == nil {
		audioContext = audio.NewContext(sampleRate)
	}

	for i := range resources {
		p, err := audio.NewPlayer(audioContext, resources[i].Source)
		if err != nil {
			return nil, err
		}

		b.audioPlayers = append(b.audioPlayers, p)
	}

	return b, nil
}

func (b *BackgroundAudioPlayer) SetActiveSong(i int) error {
	if i < 0 || i >= len(b.audioPlayers) {
		return errors.New("invalid index")
	}
	b.audioPlayers[b.activePlayer].Pause()
	b.audioPlayers[b.activePlayer].Rewind()
	b.activePlayer = i
	b.audioPlayers[b.activePlayer].Rewind()
	b.audioPlayers[b.activePlayer].SetVolume(b.volume)
	b.audioPlayers[b.activePlayer].Play()
	return nil
}

func (b *BackgroundAudioPlayer) Update() {
	if len(b.audioPlayers) > 0 {
		if !b.audioPlayers[b.activePlayer].IsPlaying() {

			next := b.activePlayer + 1
			if next >= len(b.audioPlayers) {
				next = 0
			}
			fmt.Printf("Switch songs %v \n", next)
			b.SetActiveSong(next)
		}
	}
}

func (b *BackgroundAudioPlayer) SetVolume(v float64) {
	b.volume = v

	if b.activePlayer > 0 && b.activePlayer < len(b.audioPlayers) {
		b.audioPlayers[b.activePlayer].SetVolume(v)
	}
}
