package echochamber

import (
	"encoding/base64"
	"fmt"
)

// A Track is the most fundamental unit of the application, containing some Metadata about the song it represents and locations where it may be found.
type Track struct {
	Metadata  TrackMetadata
	Locations []TrackLocation
}

// A Playlist is a set of Tracks, and additionally may have a Title to identify it and a canonical location that may be used for reconstructing or updating it.
type Playlist interface {
	Title() string
	Tracks() ([]Track, error)
	Location() PlaylistLocation
}

// TrackMetadata provides useful information about a particular Track.
type TrackMetadata interface {
	Title() string
	Artist() Artist
	Album() Album
	Year() int
	Cover() string
}

// An Artist is a Playlist composed of an musician's works, coupled with some information about him/her.
type Artist interface {
	Playlist
	Name() string
	Albums() ([]Album, error)
}

// An Album is a Playlist composed of the Tracks from an actual album, plus some information about it.
type Album interface {
	Playlist
	Name() string
	Year() int
	Cover() string
}

// A ResourceLocation contains the Namespace where a resource can be found, usually the name of the transport protocol (e.g. "rtmp") or the name of a service (e.g. "youtube"), and a unique identifier for the resource inside that namespace. It provides a utility function to represent itself with a magnet-inspired URI scheme.
type ResourceLocation struct {
	Namespace  string
	Identifier string
}

// URI returns a magnet-inspired representation of a ResourceLocation.
func (rl *ResourceLocation) URI() string {
	return fmt.Sprintf("muzk://?urn:%s:%s", rl.Namespace, base64.URLEncoding.EncodeToString([]byte(rl.Identifier)))
}

// TrackLocation is a Track-specific ResourceLocation, mostly for facilitating type bureaucracy.
type TrackLocation struct {
	ResourceLocation
}

// PlaylistLocation is a Playlist-specific ResourceLocation, mostly for facilitating type bureaucracy.
type PlaylistLocation struct {
	ResourceLocation
}

// NewTrackLocation creates a new TrackLocation
func NewTrackLocation(namespace, identifier string) TrackLocation {
	return TrackLocation{
		ResourceLocation{
			Namespace:  namespace,
			Identifier: identifier,
		},
	}
}

// NewPlaylistLocation creates a new PlaylistLocation
func NewPlaylistLocation(namespace, identifier string) PlaylistLocation {
	return PlaylistLocation{
		ResourceLocation{
			Namespace:  namespace,
			Identifier: identifier,
		},
	}
}

// DynamicPlaylist is a very special type of Playlist that may not be considered "complete", and may add songs in different ways according to implementation logic.
type DynamicPlaylist interface {
	Playlist
	// Finite indicates if the Playlist is expected to no longer add songs at some point under normal conditions or not.
	Finite() bool
	// Sequential indicates if the Playlist will only append songs at the end or at any position.
	Sequential() bool
	// Subscribe allows the user to provide a channel to receive signals whenever the Playlist is updated in some way.
	Subscribe(chan<- uint)
}

// TODO: DynamicPlaylist signals
