package id3utils

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/enaldo1709/scripts/mp3-tag-setter/src/fileutils"
	"github.com/enaldo1709/scripts/mp3-tag-setter/src/model"

	"github.com/bogem/id3v2"
	"github.com/gabriel-vasile/mimetype"
)

func WriteMetadata(filePath string, descriptor *model.Descriptor) {
	log.Printf("Writting Metadata to file -> %s", filePath)
	tag, err := id3v2.Open(filePath, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error reading file metadata: ", err)
	}

	setAlbum(tag, descriptor.Metadata.Album)
	setArtist(tag, descriptor.Metadata.Artist)
	setYear(tag, descriptor.Metadata.Year)
	setTracksMetadata(tag, descriptor.Metadata.Tracks, filePath, descriptor.TitleSeparator)
	setArtWork(tag, filePath, descriptor.Metadata.AlbumArtPath)
	setAlbumTotalTracks(tag, descriptor.Metadata.TotalTracks)

	if err = tag.Save(); err != nil {
		log.Fatal("Error writing file metadata: ", err)
	}
	log.Println("Metadata wrote with success... ")
}

func setAlbum(tag *id3v2.Tag, album string) {
	if tag == nil {
		return
	}

	if album != "" {
		tag.SetAlbum(album)
	}
}

func setArtist(tag *id3v2.Tag, artist string) {
	if tag == nil {
		return
	}

	if artist != "" {
		tag.SetArtist(artist)
	}
}

func setGenre(tag *id3v2.Tag, genre string) {
	if tag == nil {
		return
	}

	if genre != "" {
		tag.SetGenre(genre)
	}
}

func setYear(tag *id3v2.Tag, year int) {
	if tag == nil {
		return
	}

	if year > 0 {
		tag.SetYear(strconv.Itoa(year))
	}
}

func setTracksMetadata(tag *id3v2.Tag, tracks []model.TrackMetadata, filePath, separator string) {
	if len(tracks) > 0 {

		meta := GetTrackMetadata(filePath, separator, tag.Title(), tracks)
		tag.SetTitle(meta.Title)

		if meta.Genre != "" {
			setGenre(tag, meta.Genre)
		}

		if meta.TrackNumber != 0 {
			tag.AddTextFrame(id3v2.V24CommonIDs["Track number/Position in set"], tag.DefaultEncoding(), strconv.Itoa(meta.TrackNumber))
		}
		if meta.Composer != "" {
			tag.AddTextFrame(id3v2.V24CommonIDs["Composer"], tag.DefaultEncoding(), meta.Composer)
		}
	}
}

func setArtWork(tag *id3v2.Tag, filePath, artWorkPath string) {
	dirPath, _ := filepath.Split(filePath)
	artPath := fileutils.GetAlbumArtPath(dirPath, artWorkPath)
	if artPath != "" {
		artwork, err := os.ReadFile(artPath)
		if err != nil {
			log.Fatal("Error while reading artwork file", err)
		}

		mt, err := mimetype.DetectFile(artPath)
		if err != nil {
			log.Fatal("Error reading mimetype from artwork file", err)
		}

		pic := id3v2.PictureFrame{
			Encoding:    id3v2.EncodingUTF8,
			MimeType:    mt.String(),
			PictureType: id3v2.PTFrontCover,
			Description: "Front cover",
			Picture:     artwork,
		}
		tag.AddAttachedPicture(pic)
	}
}

func setAlbumTotalTracks(tag *id3v2.Tag, totalTracks int) {
	if totalTracks > 0 {
		udtf := id3v2.UserDefinedTextFrame{
			Encoding:    tag.DefaultEncoding(),
			Description: "TRACKTOTAL",
			Value:       strconv.Itoa(totalTracks),
		}

		tag.AddUserDefinedTextFrame(udtf)
	}
}
