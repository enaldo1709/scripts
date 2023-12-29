package fileutils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var (
	AudioFiles = []string{}
	OutputExt  = ""
)

func GetNameAndExtension(f string) (string, string) {
	lastDot := strings.LastIndex(f, ".")
	return f[:lastDot], f[lastDot+1:]
}

func CheckAudioFileFilter(transform func(string) bool) func(string) bool {
	return func(s string) bool {
		_, filename := filepath.Split(s)
		_, ext := GetNameAndExtension(filename)
		if cont := slices.Contains(AudioFiles, ext); !cont {
			return cont
		}
		if eq := OutputExt == ext; eq {
			return !eq
		}
		return transform(s)
	}
}

func ReadDir(d string) []fs.DirEntry {
	dfs := os.DirFS(d)
	dinf, err := fs.Stat(dfs, ".")
	if err != nil {
		log.Fatal("Error reading directory info: ", err)
	}
	if !dinf.IsDir() {
		log.Fatalf("Error: given path %s isn't a directory", d)
	}

	dir, err := fs.ReadDir(dfs, dinf.Name())
	if err != nil {
		log.Fatal("Error reading directory content: ", err)
	}
	return dir
}
