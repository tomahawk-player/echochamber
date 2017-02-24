package echochamber

import (
	"io"
)

// A Service is any data provider tied to a particular identifiable service, such as YouTube, SoundCloud or your own machine's Filesystem.
type Service interface {
	Namespace() string
}

// ConcurrentlyStreamable indicates wether a service supports streaming more than one file simultaneously, allowing the Player to take advantage of it with functionality like caching, if it implements it.
type ConcurrentlyStreamable interface {
	MaxAllowedStreams() int
}

// A TrackLocationResolver returns file streams for a particular TrackLocation, if the latter belongs to the resolver's associated service.
type TrackLocationResolver interface {
	Service
	Resolve(location TrackLocation) (io.ReadCloser, error)
}

// A PlaylistResolver returns a Playlist if a valid location associated with the resolver's service is given.
type PlaylistResolver interface {
	Service
	Resolve(location PlaylistLocation) (*Playlist, error)
}

// A TrackSplitter helps break down sources into many ones, like YouTube albums uploaded as a single video or SoundCloud live shows/mixtapes uploaded as a single track.
type TrackSplitter interface {
	Service
	SplitTracks(location TrackLocation) (TrackSplitResult, error)
}

// A TrackSplitResult contains a Playlist product of splitting a Track into many, and information from the procedure that may aid in using the splitted tracks; e.g.: YouTube description or comment including timestamps for each song in a single-video album.
type TrackSplitResult interface {
	Playlist
	Info() string
}

// An OmniSearcher implementor uses the given search terms to scour for Tracks, Artists and Albums on its associated service, and returns them in a categorized fashion.
type OmniSearcher interface {
	Service
	OmniSearch(terms string) (OmniSearchResult, error)
}

// An OmniSearchResult returns the results from a succesful OmniSearch
type OmniSearchResult interface {
	Tracks() []Track
	Artists() []Artist
	Albums() []Album
}

// A TrackSearcher performs a search on its correspondant service with the given terms and returns the resulting Tracks, if any.
type TrackSearcher interface {
	Service
	SearchTracks(terms string) ([]Track, error)
}

// An ArtistSearcher performs a search on its correspondant service with the given terms and returns the resulting Artists, if any.
type ArtistSearcher interface {
	Service
	SearchArtists(terms string) ([]Artist, error)
}

// An AlbumSearcher performs a search on its correspondant service with the given terms and returns the resulting Albums, if any.
type AlbumSearcher interface {
	Service
	SearchAlbums(terms string) ([]Album, error)
}

// A PlaylistSearcher searches its service for playlists associated with the given terms and returns a list of the Playlists it found, if any.
type PlaylistSearcher interface {
	Service
	SearchPlaylists(terms string) ([]Playlist, error)
}
