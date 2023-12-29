package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/enaldo1709/scripts/mp3-tag-setter/src/model"
)

func askFor(message, defaultOpt string) string {
	answer := ""

	if defaultOpt != "" {
		fmt.Printf("%s (%s):", message, defaultOpt)
	} else {
		fmt.Printf("%s: ", message)
	}
	fmt.Scanf("%s", &answer)

	if answer == "" {
		answer = defaultOpt
	}

	return answer
}

func askForAlbum(descriptor *model.Descriptor) {
	descriptor.Metadata.Album = askFor("Type the album name", descriptor.Metadata.Album)
}

func askForAlbumArt(descriptor *model.Descriptor) {
	descriptor.Metadata.AlbumArtPath = askFor("Type the path to album artwork image", descriptor.Metadata.AlbumArtPath)
}

func askForArtist(descriptor *model.Descriptor) {
	descriptor.Metadata.Artist = askFor("Type the artist/interpreter name", descriptor.Metadata.Artist)
}

func askForYear(descriptor *model.Descriptor) {
	sYear := askFor("Type the year of the album launch", strconv.Itoa(descriptor.Metadata.Year))
	year, err := strconv.Atoi(sYear)
	if err != nil {
		log.Println("Error reading year....")
		return
	}

	descriptor.Metadata.Year = year
}

func askForTracks(descriptor *model.Descriptor) {
	createFromFolder := askFor("Create track list from folder? ", "y or n")
	switch createFromFolder {
	case "y or n":
		break
	case "y":
		break
	case "n":
		descriptor.Metadata.Tracks = []model.TrackMetadata{}
		for {
			addNewTrack := askFor("Add new track to list? ", "y or n")
			if addNewTrack == "n" {
				break
			}

			if addNewTrack != "y" {
				log.Printf("Error: invalid option %s", addNewTrack)
				continue
			}
			meta := &model.TrackMetadata{}
			meta.Title = askFor("Type the track name", "")
			meta.Genre = askFor("Type the track genre", "")
			meta.Composer = askFor("Type the track composer", "")
			meta.FileName = askFor("Type the file name", "")

			trackNumber, err := strconv.Atoi(askFor("Type the track number", strconv.Itoa(len(descriptor.Metadata.Tracks)+1)))
			if err != nil {
				trackNumber = len(descriptor.Metadata.Tracks) + 1
				log.Printf("Warning track number is invalid... using -> %d", trackNumber)
			}
			meta.TrackNumber = trackNumber

			descriptor.Metadata.Tracks = append(descriptor.Metadata.Tracks, *meta)
		}

		descriptor.Metadata.TotalTracks = len(descriptor.Metadata.Tracks)
	default:
		log.Println("Skipping... Creating track list from folder...")
	}
}
