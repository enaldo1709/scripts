package id3utils

import (
	"mp3-tag-setter/src/fileutils"
	"mp3-tag-setter/src/model"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetTrackMetadata(filePath, separator, tagTitle string, tracks []model.TrackMetadata) model.TrackMetadata {
	_, filename := filepath.Split(filePath)
	for _, track := range tracks {
		if track.FileName == filename {
			return track
		}
	}

	filename, _ = fileutils.GetNameAndExtension(filename)
	caser := cases.Title(language.Spanish)
	filename = caser.String(strings.ToLower(strings.Split(filename, separator)[0]))
	meta := model.TrackMetadata{
		Title: filename,
	}

	for _, track := range tracks {
		if fileutils.LooksLike(filename, track.Title) {
			meta = track
			break
		}
	}

	return meta
}
