package models

import (
	"encoding/base64"
	"fmt"
)

type Track struct {
	Metadata  TrackMetadata
	Locations []TrackLocation
}

type Playlist interface {
	Title() string
	Tracks() ([]Track, error)
	Location() PlaylistLocation
}

type ResourceLocation struct {
	Namespace string
	String    string
}

func (rl *ResourceLocation) URI() string {
	return fmt.Sprintf("muzk://?urn:%s:%s", rl.Namespace, base64.URLEncoding.EncodeToString([]byte(rl.String)))
}

type TrackLocation struct {
	ResourceLocation
}

type PlaylistLocation struct {
	ResourceLocation
}
