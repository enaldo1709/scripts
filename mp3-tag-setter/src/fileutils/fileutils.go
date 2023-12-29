package fileutils

import (
	"os"
	"path/filepath"
	"slices"
)

var defaultAudioFiles = []string{"mp3", "aac", "m4a", "opus"}

func CheckAudioFileFilter(filePath string) bool {
	_, filename := filepath.Split(filePath)
	_, ext := GetNameAndExtension(filename)
	return slices.Contains(defaultAudioFiles, ext)
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func GetAlbumArtPath(dirpath, artpath string) string {
	if FileExists(artpath) {
		return artpath
	}

	_, artFileName := filepath.Split(artpath)
	rootArtFilePath := filepath.Join(dirpath, artFileName)
	if FileExists(rootArtFilePath) {
		return rootArtFilePath
	}

	return ""
}
