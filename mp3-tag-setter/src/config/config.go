package config

import (
	"bytes"
	"encoding/json"
	"log"
	"mp3-tag-setter/src/model"
	"os"
	"path/filepath"
)

func OpenMetadataConfigFile(filepath string) *model.Descriptor {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error opening metadata file", err)
	}

	var descriptor model.Descriptor
	if err = json.NewDecoder(f).Decode(&descriptor); err != nil {
		log.Fatal("Error reading metadata file", err)
	}

	return &descriptor
}

func GenerateTemplate(filePath string, askForFields bool) {
	log.Println("Starting template generation...")
	parent, _ := filepath.Split(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error generating template: ", err)
		return
	}
	defer f.Close()

	var descriptor *model.Descriptor
	if askForFields {
		descriptor = generateManualTemplate(parent)
	} else {
		descriptor = generateAutomaticTemplate(parent)
	}

	var buffer *bytes.Buffer = new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(descriptor); err != nil {
		log.Fatal("Error parsing properties as json object: ", err)
	}

	wrote, err := f.Write(buffer.Bytes())
	if err != nil {
		log.Fatal("Error writing template file")
		return
	}

	log.Printf("Wrote %d bytes on file -> %s", wrote, filePath)
}
