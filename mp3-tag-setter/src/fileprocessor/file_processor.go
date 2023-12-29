package fileprocessor

import (
	"log"
	"path"

	"github.com/enaldo1709/scripts/mp3-tag-setter/src/fileutils"
	"github.com/enaldo1709/scripts/mp3-tag-setter/src/id3utils"
	"github.com/enaldo1709/scripts/mp3-tag-setter/src/model"
)

func LoopThroughDir(dirpath string, descriptor *model.Descriptor) {
	filesProcessed := 0
	files := fileutils.ReadDir(dirpath)

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
