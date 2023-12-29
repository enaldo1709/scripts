package model

type Descriptor struct {
	Metadata       *AlbumMetadata `json:"meta"`
	TitleSeparator string         `json:"separator"`
}

type AlbumMetadata struct {
	Album        string          `json:"album"`
	AlbumArtPath string          `json:"album_art"`
	Artist       string          `json:"artist"`
	Genre        string          `json:"genre"`
	TotalTracks  int             `json:"total_tracks"`
	Tracks       []TrackMetadata `json:"tracks"`
	Year         int             `json:"year"`
}

type TrackMetadata struct {
	Title       string `json:"title"`
	Composer    string `json:"composer"`
	TrackNumber int    `json:"track_number"`
}
