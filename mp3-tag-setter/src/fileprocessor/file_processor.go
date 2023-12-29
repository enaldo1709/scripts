package fileprocessor

import (
	"io/fs"
	"log"
	"mp3-tag-setter/src/fileutils"
	"mp3-tag-setter/src/id3utils"
	"mp3-tag-setter/src/model"
	"os"
	"path"
)

func LoopThroughDir(dirpath string, descriptor *model.Descriptor) {
	filesProcessed := 0
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

	log.Printf("working on directory -> '%s'\n", dirpath)

	for _, file := range files {
		filePath := path.Join(dirpath, file.Name())
		if fileutils.CheckAudioFileFilter(filePath) {
			id3utils.WriteMetadata(filePath, descriptor)
			filesProcessed++
		}
	}
	log.Printf("Finish processing with success -> files processed: %d", filesProcessed)
}
