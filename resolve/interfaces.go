package resolve

import "gitlab.com/mixedCase/muzk/models"

type Service interface {
	Namespace() string
}

type TrackLocationResolver interface {
	Service
	Resolve(location models.TrackLocation) (string, error)
}

type PlaylistResolver interface {
	Service
	Resolve(location models.PlaylistLocation) (*models.Playlist, error)
}

type TrackSearcher interface {
	Service
	SearchTracks(terms string) ([]models.Track, error)
}

type PlaylistSearcher interface {
	Service
	SearchPlaylists(terms string) ([]models.Playlist, error)
}
