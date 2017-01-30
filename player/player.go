package player

import "github.com/tomahawk-player/echochamber/models"

const ( // Playback status
	// StatusPlaying indicates music is being played.
	StatusPlaying = iota
	// StatusStopped indicates reproduction has stopped.
	StatusStopped
	// StatusPaused indicates there's a track loaded but reproduction has been halted.
	StatusPaused
)

const ( // Loop status
	// NoLooping indicates the Player will not restart reproduction.
	NoLooping = iota
	// LoopTrack indicates any Track that plays will be played again every time it is played to its end.
	LoopTrack
	// LoopTracklist means the first item in the Tracklist will be played once the end of its last item is played.
	LoopTracklist
)

const ( // Signals
	// Opening is emitted when the player is acquiring the resource.
	Opening = iota
	// Buffering is emitted when the player is gathering data to play.
	Buffering
	// Playing is emitted when reproduction has started.
	Playing
	// Paused is emitted when reproduction of a track is (temporarily) halted.
	Paused
	// Stopped means reproduction has just stopped.
	Stopped
	// Forward is emitted when the player starts playing the next song of the tracklist.
	Forward
	// Backward is emitted when the player starts playing the previous song of the tracklist.
	Backward
	// EndReached is emitted when there are no more songs to play in the tracklist.
	EndReached
	// EncounteredError is emitted when an error occurs while trying to play a track.
	EncounteredError
	// LoopChanged is emitted when the looping logic was updated.
	LoopChanged
	// ShuffleChanged is emitted when "Shuffle" was enabled or disabled.
	ShuffleChanged
	// RateChanged is emitted when the speed of reproduction changes.
	RateChanged
	// CurrentTrackMetadataChanged is emitted when the track being currently played has its metadata updated.
	CurrentTrackMetadataChanged
	// CanGoNextChanged is emitted when there are no songs in the tracklist after the current one, or if that is no longer the case.
	CanGoNextChanged
	// CanGoPreviousChanged is emitted when there are no songs in the tracklist before the current one, or if that is no longer the case.
	CanGoPreviousChanged
	// CanPlayChanged is emitted when there are no songs in the tracklist to play or if that's no longer the case.
	CanPlayChanged
	// CanPauseChanged is emitted when reproduction can no longer b e paused, or if it can now be paused.
	CanPauseChanged
	// CanSeekChanged is emitted when the current track no longer supports seeking, or if it starts supporting it.
	CanSeekChanged
)

// The Player is the entry point of the music player application, from which reproduction is controlled and information about it is gathered.
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
	Position() (int, error)
	SetPosition(msPosition int) error
	Seek(msOffset int) error
	Loop() uint
	SetLoop(loopType uint) error
	Shuffle() bool
	SetShuffle(enable bool)
	Rate() float32
	SetRate(float32) error
	CanGoNext() bool
	CanGoPrevious() bool
	CanPlay() bool
	CanPause() bool
	CanSeek() bool
	Subscribe(ch chan<- uint)
}
