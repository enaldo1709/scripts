package fileutils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
)

var defaultAudioFiles = []string{"mp3", "aac", "m4a", "opus"}

func ReadDir(dirpath string) []fs.DirEntry {
	dfs := os.DirFS(dirpath)
	dinf, err := fs.Stat(dfs, ".")
	if err != nil {
		log.Fatal("Error reading directory info: ", err)
	}
	if !dinf.IsDir() {
		log.Fatal("Error provided dirpath '", dirpath, "' isn't a directory...")
	}

	files, err := fs.ReadDir(dfs, dinf.Name())
	if err != nil {
		log.Fatal("Error reading directory content: ", err)
	}

	return files
}

func ReadTracks(dirPath string) []fs.DirEntry {
	tracks := []fs.DirEntry{}

	for _, file := range ReadDir(dirPath) {
		filePath := filepath.Join(dirPath, file.Name())
		if CheckAudioFileFilter(filePath) {
			tracks = append(tracks, file)
		}
	}

	return tracks
}

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
