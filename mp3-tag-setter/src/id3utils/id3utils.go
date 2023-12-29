package id3utils

import (
	"path/filepath"
	"strings"

	"github.com/enaldo1709/scripts/mp3-tag-setter/src/fileutils"
	"github.com/enaldo1709/scripts/mp3-tag-setter/src/model"

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
