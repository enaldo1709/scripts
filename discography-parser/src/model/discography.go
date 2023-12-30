package model

type Discography struct {
	Albums []AlbumMetadata `json:"albums"`
	Artist string          `json:"artist"`
}

type AlbumMetadata struct {
	Album       string          `json:"album"`
	TotalTracks int             `json:"total_tracks"`
	Tracks      []TrackMetadata `json:"tracks"`
	Year        int             `json:"year"`
	Date        string          `json:"date"`
}

type TrackMetadata struct {
	Title        string   `json:"title"`
	Interpreters []string `json:"interpreters"`
	Composer     string   `json:"composer"`
	TrackNumber  int      `json:"track_number"`
	Genre        string   `json:"genre"`
}
