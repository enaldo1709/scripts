package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/enaldo1709/scripts/audio-file-converter/src/ffmpegwrapper"
	"github.com/enaldo1709/scripts/audio-file-converter/src/fileutils"
)

var (
	defaultAudioFiles = []string{"mp3", "aac", "m4a", "opus"}
	defaultOutputExt  = "mp3"
)

func main() {
	fmt.Println(os.Args)

	cwd, err := os.Getwd()
	if err != nil {
		log.Panicln("error getting current directory", err)
	}
	if len(os.Args) > 1 {
		cwd = os.Args[1]
	}

	_, err = os.ReadDir(cwd)
	if err != nil {
		log.Panicln(err)
	}

	fileutils.OutputExt = defaultOutputExt
	if len(os.Args) > 2 {
		fileutils.OutputExt = os.Args[2]
	}

	fileutils.AudioFiles = defaultAudioFiles
	if len(os.Args) > 3 {
		fileutils.AudioFiles = strings.Split(os.Args[3], ",")
	}

	ffmpegwrapper.LoopThroughFolder(cwd,
		fileutils.CheckAudioFileFilter(ffmpegwrapper.ConvertAudioFile))
	log.Printf("¡done!")
}
