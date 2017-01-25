package models

type TrackMetadata interface {
	Title() string
	Artist() *Artist
	Album() *Album
	Year() int
}

type Artist struct {
	title      string
	albumCache []Album
	trackCache []Track
}

func (a *Artist) Title() string {
	return a.title
}

func (a *Artist) Albums() ([]Album, error) {
	panic("not implemented")
}

func (a *Artist) Tracks() ([]Track, error) {
	panic("not implemented")
}

func (a *Artist) Location() PlaylistLocation {
	panic("not implemented")
}

type Album struct {
	title      string
	trackCache []Track
	coverCache string
	Year       int
}

func (a *Album) Title() string {
	return a.title
}

func (a *Album) Tracks() ([]Track, error) {
	panic("not implemented")
}

func (a *Album) Location() PlaylistLocation {
	panic("not implemented")
}

func (a *Album) Cover() string {
	panic("not implemented")
}
