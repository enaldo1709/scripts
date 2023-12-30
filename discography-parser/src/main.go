package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/enaldo1709/scripts/discography-parser/src/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var months = map[string]string{
	"enero":      "01",
	"febrero":    "02",
	"marzo":      "03",
	"abril":      "04",
	"mayo":       "05",
	"junio":      "06",
	"julio":      "07",
	"agosto":     "08",
	"septiembre": "09",
	"octubre":    "10",
	"noviembre":  "11",
	"diciembre":  "12",
}

func main() {
	filename := "diomedes-diaz-local.txt"
	lastDot := strings.LastIndex(filename, ".")
	filetitle := filename[:lastDot]

	discography := ReadFile("diomedes-diaz-local.txt")

	buffer := new(bytes.Buffer)

	json.NewEncoder(buffer).Encode(discography)
	f, err := os.Create(filetitle + ".json")
	if err != nil {
		log.Fatal("Error creating output file: ", err)
	}
	defer f.Close()

	wrote, err := f.Write(buffer.Bytes())
	if err != nil {
		log.Fatal("Error writing file: ", err)
	}
	log.Printf("Wrote %d bytes on file %s", wrote, f.Name())

}

func ReadFile(filePath string) *model.Discography {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	discography := &model.Discography{
		Albums: []model.AlbumMetadata{},
	}

	caser := cases.Title(language.Spanish)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if artist, ok := strings.CutPrefix(line, "Artist: "); ok {
			discography.Artist = artist
			continue
		}

		if albumLine, ok := strings.CutPrefix(line, "- "); ok {
			albumName := caser.String(strings.Split(albumLine, " (")[0])
			scanner.Scan()
			collabLine := strings.TrimSpace(scanner.Text())
			collaborators := strings.Split(strings.TrimSpace(strings.Split(collabLine, "-")[0]), ",")
			collaborators = append([]string{discography.Artist}, collaborators...)
			releaseDate, year := "", 0
			if collabSplit := strings.Split(collabLine, "-"); len(collabSplit) > 1 {
				releaseDate, year = parseDateISODate(strings.TrimSpace(collabSplit[1]))
			}

			album := &model.AlbumMetadata{
				Album:  albumName,
				Tracks: []model.TrackMetadata{},
				Year:   year,
				Date:   releaseDate,
			}

			trackCount := 0
			scanner.Scan() // skip white line
			for scanner.Scan() {
				trackLine := strings.TrimSpace(scanner.Text())
				if trackLine == "" {
					break
				}
				trackCount++

				track := &model.TrackMetadata{TrackNumber: trackCount, Interpreters: []string{}}

				title := strings.Split(strings.Split(strings.Split(trackLine, "{")[0], "(")[0], "-")[0]
				track.Title = caser.String(strings.TrimSpace(title))
				if strings.Contains(trackLine, "{") && strings.Contains(trackLine, "}") {
					track.Genre = strings.TrimSpace(
						strings.Split(
							strings.Split(trackLine, "{")[1], "}",
						)[0],
					)
				}
				if strings.Contains(trackLine, "(") && strings.Contains(trackLine, ")") {
					track.Composer = strings.TrimSpace(
						strings.Split(
							strings.Split(trackLine, "(")[1], ")",
						)[0],
					)
				}
				for _, collab := range collaborators {
					if collab == "" {
						continue
					}
					track.Interpreters = append(track.Interpreters, strings.TrimSpace(collab))
				}
				if strings.Contains(trackLine, "-") {
					otherCollabs := strings.Split(strings.Split(trackLine, "-")[1], ",")
					for _, collab := range otherCollabs {
						if collab == "" {
							continue
						}
						track.Interpreters = append(track.Interpreters, strings.TrimSpace(collab))
					}
				}
				album.Tracks = append(album.Tracks, *track)
			}
			album.TotalTracks = trackCount

			discography.Albums = append(discography.Albums, *album)
		}
	}

	return discography
}

func parseDateISODate(dateS string) (string, int) {
	splitted := strings.Split(dateS, "de")
	for i := 0; i < len(splitted); i++ {
		splitted[i] = strings.TrimSpace(splitted[i])
	}

	if mint, ok := months[splitted[1]]; ok {
		splitted[1] = mint
	}

	year, err := strconv.Atoi(splitted[2])
	if err != nil {
		year = 0
	}

	return strings.Join(splitted, "-"), year
}
