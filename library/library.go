package library

import "github.com/tomahawk-player/echochamber/services"

// A Library is a searchable collection of music metadata.
type Library interface {
	services.OmniSearcher
	services.TrackSearcher
	services.ArtistSearcher
	services.AlbumSearcher
	services.PlaylistSearcher
}
