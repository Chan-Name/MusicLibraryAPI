package storage

import (
	"fmt"
	"log/slog"
	"os"
)

type Library struct {
	SongId      []int
	GroupName   []string
	SongName    []string
	SongText    []string
	Link        []string
	ReleaseDate []string
}

type SongText struct {
	SongText []string
}

func (s *Storage) ReturnLibraryDate() (*Library, error) {

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "give_library.sql")

	req, err := s.PrepareSQL(path)
	if err != nil {
		return nil, err
	}

	rows, err := req.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libraryDate Library
	for rows.Next() {
		var songId int
		var groupName string
		var songName string
		var songText string
		var link string
		var releaseDate string

		if err := rows.Scan(&songId, &groupName, &songName, &songText, &link, &releaseDate); err != nil {
			return nil, err
		}
		libraryDate.SongId = append(libraryDate.SongId, songId)
		libraryDate.GroupName = append(libraryDate.GroupName, groupName)
		libraryDate.SongName = append(libraryDate.SongName, songName)
		libraryDate.SongText = append(libraryDate.SongText, songText)
		libraryDate.Link = append(libraryDate.Link, link)
		libraryDate.ReleaseDate = append(libraryDate.ReleaseDate, releaseDate)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	slog.Info("Output Song with Verses")
	return &libraryDate, nil
}

func (s *Storage) ReturnSongDateWithName(group, song string) (*SongText, error) {

	path := fmt.Sprintf("%s%s", os.Getenv("SQL_PATH"), "give_song_from_name.sql")

	req, err := s.PrepareSQL(path)
	if err != nil {
		return nil, err
	}

	rows, err := req.Query(group, song)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songData SongText
	for rows.Next() {
		var songText string
		if err := rows.Scan(&songText); err != nil {
			return nil, err
		}
		songData.SongText = append(songData.SongText, songText)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	slog.Info("Output Song with Verses")
	return &songData, nil
}
