package main

import (
	"log"
	"mp3-tag-setter/src/config"
	"mp3-tag-setter/src/fileprocessor"
	"os"
	"path"
)

var ()

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file especified...")
	}
	var metafilepath string
	var dirpath string

	for i, arg := range os.Args {
		if arg == "-d" {
			dirpath = os.Args[i+1]
		}
		if arg == "-m" {
			metafilepath = os.Args[i+1]
		}
	}

	if dirpath == "" {
		dirpath, _ = os.Getwd()
	}
	if metafilepath == "" {
		metafilepath = path.Join(dirpath, "meta.json")
	}

	log.Println(metafilepath, dirpath)

	descriptor := config.OpenMetadataConfigFile(metafilepath)
	fileprocessor.LoopThroughDir(dirpath, descriptor)
}
