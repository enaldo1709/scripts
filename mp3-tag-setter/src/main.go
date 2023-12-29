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
	var generateTemplate bool = false
	var askForTemplate bool = false

	for i, arg := range os.Args {
		if arg == "-d" {
			dirpath = os.Args[i+1]
		}
		if arg == "-m" {
			metafilepath = os.Args[i+1]
		}
		if arg == "--generate-template" {
			for _, arg2 := range os.Args {
				if arg2 == "-y" || arg2 == "-a" || arg2 == "--ask" {
					askForTemplate = true
				}
			}
			generateTemplate = true
		}
	}

	if dirpath == "" {
		dirpath, _ = os.Getwd()
	}
	if metafilepath == "" {
		metafilepath = path.Join(dirpath, "meta.json")
	}

	if generateTemplate {
		config.GenerateTemplate(metafilepath, askForTemplate)
		return
	}

	log.Println(metafilepath, dirpath)

	descriptor := config.OpenMetadataConfigFile(metafilepath)
	fileprocessor.LoopThroughDir(dirpath, descriptor)
}
