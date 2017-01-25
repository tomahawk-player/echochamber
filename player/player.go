package player

import "gitlab.com/mixedCase/muzk/models"

const ( // Playback status
	StatusPlaying = iota
	StatusStopped
	StatusPaused
)

const ( // Loop status
	NoLooping = iota
	LoopTrack
	LoopTracklist
)

const ( // Signals
	Opening = iota
	Buffering
	Playing
	Paused
	Stopped
	Forward
	Backward
	EndReached
	EncounteredError
	LoopChanged
	ShuffleChanged
	RateChanged
	CurrentTrackMetadataChanged
	CanGoNextChanged
	CanGoPreviousChanged
	CanPlayChanged
	CanPauseChanged
	CanSeekChanged
)

type Player interface {
	Status() uint
	Play()
	Pause()
	PlayPause()
	Stop()
	Next()
	Previous()
	CurrentTrack() *models.Track
	Tracklist() models.Playlist
	SetTracklist(models.Playlist) error
	Enqueue(...models.Track) error
	Position() int
	SetPosition(msPosition int) error
	Seek(msOffset int) error
	Loop() uint
	SetLoop(loopType uint) error
	Shuffle() bool
	SetShuffle(enable bool) error
	Rate() float32
	SetRate(float32) error
	CanGoNext() bool
	CanGoPrevious() bool
	CanPlay() bool
	CanPause() bool
	CanSeek() bool
	Subscribe(ch chan<- uint)
}
