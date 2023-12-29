package ffmpegwrapper

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/enaldo1709/scripts/audio-file-converter/src/fileutils"
)

func ConvertAudioFile(f string) bool {
	parent, file := filepath.Split(f)
	name, ext := fileutils.GetNameAndExtension(file)
	if eq := fileutils.OutputExt == ext; eq {
		return !eq
	}
	outName := fmt.Sprintf("%s.%s", name, fileutils.OutputExt)
	cmd := exec.Command("ffmpeg", "-i", f, "-b:a", "192K", filepath.Join(parent, outName))
	var outErr bytes.Buffer
	cmd.Stderr = &outErr

	if err := cmd.Run(); err != nil {
		log.Panic("error processing file", outErr.String(), err)
		recover()
		return false
	}

	if err := os.Remove(f); err != nil {
		log.Println("error: removing original file ", f, err)
	}

	return true
}

func LoopThroughFolder(d string, transform func(string) bool) {
	dir := fileutils.ReadDir(d)

	log.Printf("working on directory -> '%s'\n", d)
	count := 0
	for _, file := range dir {
		if file.IsDir() {
			LoopThroughFolder(filepath.Join(d, file.Name()), transform)
			continue
		}
		if transform(filepath.Join(d, file.Name())) {
			count++
		}
	}
	log.Printf("processed %d files in path -> %s\n", count, d)
}
