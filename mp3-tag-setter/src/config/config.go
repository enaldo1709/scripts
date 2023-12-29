package config

import (
	"encoding/json"
	"log"
	"mp3-tag-setter/src/model"
	"os"
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
