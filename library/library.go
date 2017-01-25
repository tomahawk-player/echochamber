package library

import "gitlab.com/mixedCase/muzk/resolve"

type Library interface {
	resolve.TrackSearcher
	resolve.PlaylistSearcher
}
