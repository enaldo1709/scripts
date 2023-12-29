package config

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/enaldo1709/scripts/mp3-tag-setter/src/fileutils"
	"github.com/enaldo1709/scripts/mp3-tag-setter/src/model"
)

func generateAutomaticTemplate(dirPath string) *model.Descriptor {
	dirPath = strings.TrimSuffix(dirPath, string(os.PathSeparator))

	tracks := getTracksMetadataFromFolder(dirPath)

	descriptor := &model.Descriptor{
		Metadata: &model.AlbumMetadata{
			Tracks:      tracks,
			TotalTracks: len(tracks),
		},
		TitleSeparator: "",
	}

	year, name := getAlbumNameFromFolder(dirPath)

	descriptor.Metadata.Album = name
	if year > 0 {
		descriptor.Metadata.Year = year
	}

	return descriptor
}

func generateManualTemplate(dirPath string) *model.Descriptor {
	descriptor := generateAutomaticTemplate(dirPath)

	askForAlbum(descriptor)
	askForAlbumArt(descriptor)
	askForArtist(descriptor)
	askForYear(descriptor)
	askForTracks(descriptor)

	return descriptor
}

func getTracksMetadataFromFolder(dirPath string) []model.TrackMetadata {
	files := fileutils.ReadTracks(dirPath)
	tracks := []model.TrackMetadata{}

	for _, file := range files {
		name, _ := fileutils.GetNameAndExtension(file.Name())
		tracks = append(tracks, model.TrackMetadata{Title: name, FileName: file.Name()})
	}

	return tracks
}

func getAlbumNameFromFolder(dirPath string) (int, string) {
	_, dirName := filepath.Split(dirPath)
	splitted := strings.Split(dirName, "-")
	name := ""
	year := 0

	if len(splitted) > 1 {
		yearS := strings.TrimSpace(splitted[0])
		parsed, err := strconv.Atoi(yearS)
		if err != nil {
			name = yearS
		} else {
			name = strings.TrimSpace(splitted[1])
			year = parsed
		}
	} else {
		name = dirName
	}
	return year, name
}
